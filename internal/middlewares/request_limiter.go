package middlewares

import (
	"fmt"
	"log/slog"
	"math"
	"net"
	"net/http"
	"strings"
	"sync"
	"time"

	"github.com/danielRamosMencia/api-my-finances/internal/configurations"
	"github.com/danielRamosMencia/api-my-finances/internal/responses"
	"github.com/danielRamosMencia/api-my-finances/internal/sessions"
	"github.com/danielRamosMencia/api-my-finances/internal/utils"
)

const (
	DEVICE_ID_HEADER     = "Device-Id"
	DEVICE_ID_MAX_LENGTH = 128
	RATE_LIMIT_PREFIX    = "rrl-"
	IP_LIMIT_PREFIX      = "rrl-ip-"
)

type limitRule struct {
	maxRequests int
	window      time.Duration
	blockFor    time.Duration
}

type limitRules struct {
	public  limitRule
	private limitRule
}

var (
	deviceRules = limitRules{
		public:  limitRule{maxRequests: 5, window: time.Minute, blockFor: 20 * time.Minute},
		private: limitRule{maxRequests: 30, window: time.Minute, blockFor: 3 * time.Minute},
	}
	ipRules = limitRules{
		public:  limitRule{maxRequests: 15, window: time.Minute, blockFor: 20 * time.Minute},
		private: limitRule{maxRequests: 50, window: time.Minute, blockFor: 3 * time.Minute},
	}
)

type LimitBucket struct {
	Count        int       `json:"count"`
	WindowStart  time.Time `json:"window_start"`
	BlockedUntil time.Time `json:"blocked_until"`
}

type ClientLimit struct {
	Public  LimitBucket `json:"public"`
	Private LimitBucket `json:"private"`
}

type limitStore struct {
	rules  limitRules
	limits map[string]*ClientLimit
}

var (
	limiterMu sync.Mutex
	devices   = limitStore{rules: deviceRules, limits: make(map[string]*ClientLimit)}
	ips       = limitStore{rules: ipRules, limits: make(map[string]*ClientLimit)}
)

func RequestLimiter(w http.ResponseWriter, r *http.Request, private bool) bool {
	deviceId := r.Header.Get(DEVICE_ID_HEADER)
	if deviceId == "" || len(deviceId) > DEVICE_ID_MAX_LENGTH {
		responses.Error(w, http.StatusBadRequest, responses.MISSING_DEVICE_ID)
		return false
	}

	ip := clientIp(r)

	limiterMu.Lock()
	deviceWait := devices.register(RATE_LIMIT_PREFIX+deviceId, private)
	ipWait := ips.register(IP_LIMIT_PREFIX+ip, private)
	limiterMu.Unlock()

	wait := max(deviceWait, ipWait)
	if wait == 0 {
		return true
	}

	seconds := int(math.Ceil(wait.Seconds()))
	w.Header().Set("Retry-After", fmt.Sprint(seconds))

	if private {
		deleteRequestSession(r)
		slog.Warn("[RequestLimiter] private limit exceeded, session closed", "device_id", deviceId, "ip", ip)
		responses.Error(w, http.StatusTooManyRequests, responses.TOO_MANY_REQUESTS_SESSION_CLOSED)
		return false
	}

	slog.Warn("[RequestLimiter] public limit exceeded", "device_id", deviceId, "ip", ip)
	responses.Error(w, http.StatusTooManyRequests, responses.TOO_MANY_REQUESTS)
	return false
}

var proxyIpHeaders = []string{"True-Client-IP", "CF-Connecting-IP"}

func clientIp(r *http.Request) string {
	if configurations.Envs.TRUST_PROXY_HEADERS {
		for _, header := range proxyIpHeaders {
			if ip := net.ParseIP(strings.TrimSpace(r.Header.Get(header))); ip != nil {
				return ip.String()
			}
		}
	}

	host, _, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		return r.RemoteAddr
	}

	return host
}

func (store *limitStore) register(key string, private bool) time.Duration {
	now := time.Now()

	client, ok := store.limits[key]
	if !ok {
		client = &ClientLimit{}
		store.limits[key] = client
	}

	bucket, rule := &client.Public, store.rules.public
	if private {
		bucket, rule = &client.Private, store.rules.private
	}

	if now.Before(bucket.BlockedUntil) {
		return bucket.BlockedUntil.Sub(now)
	}

	if now.Sub(bucket.WindowStart) >= rule.window {
		bucket.WindowStart = now
		bucket.Count = 0
	}

	bucket.Count++

	if bucket.Count > rule.maxRequests {
		bucket.BlockedUntil = now.Add(rule.blockFor)
		return rule.blockFor
	}

	return 0
}

func deleteRequestSession(r *http.Request) {
	tokenString, found := strings.CutPrefix(r.Header.Get("Authorization"), "Bearer ")
	if !found {
		return
	}

	claims, err := utils.VerifyToken(tokenString)
	if err != nil {
		return
	}

	session, _ := claims["session"].(string)
	sessions.Delete(session)
}

func isIdle(bucket LimitBucket, rule limitRule, now time.Time) bool {
	return !now.Before(bucket.BlockedUntil) && now.Sub(bucket.WindowStart) >= rule.window
}

func (store *limitStore) deleteIdle(now time.Time) {
	for key, client := range store.limits {
		if isIdle(client.Public, store.rules.public, now) && isIdle(client.Private, store.rules.private, now) {
			delete(store.limits, key)
		}
	}
}

func StartLimiterCleanup(interval time.Duration) {
	go func() {
		ticker := time.NewTicker(interval)
		defer ticker.Stop()

		for now := range ticker.C {
			limiterMu.Lock()
			devices.deleteIdle(now)
			ips.deleteIdle(now)
			limiterMu.Unlock()
		}
	}()

	slog.Info("Request limiter cleanup started")
}

type LimitsInfo struct {
	Devices map[string]ClientLimit `json:"devices"`
	Ips     map[string]ClientLimit `json:"ips"`
}

func (store *limitStore) snapshot() map[string]ClientLimit {
	snapshot := make(map[string]ClientLimit, len(store.limits))
	for key, client := range store.limits {
		snapshot[key] = *client
	}

	return snapshot
}

func LimitsSnapshot() LimitsInfo {
	limiterMu.Lock()
	defer limiterMu.Unlock()

	return LimitsInfo{
		Devices: devices.snapshot(),
		Ips:     ips.snapshot(),
	}
}
