package sessions

import (
	"crypto/subtle"
	"log/slog"
	"sync"
	"time"

	"github.com/danielRamosMencia/api-my-finances/internal/configurations"
	"github.com/danielRamosMencia/api-my-finances/internal/utils"
)

type session struct {
	token     string
	expiresAt time.Time
}

var (
	mu       sync.RWMutex
	sessions = make(map[string]session)
)

func key(sessionClaim string) string {
	return utils.HmacSHA256(configurations.Envs.HMAC_KEY, sessionClaim)
}

func Save(sessionClaim, token string, expiresAt time.Time) {
	mu.Lock()
	defer mu.Unlock()

	sessions[key(sessionClaim)] = session{
		token:     token,
		expiresAt: expiresAt,
	}
}

func IsActive(sessionClaim, token string) bool {
	if sessionClaim == "" {
		return false
	}

	mu.RLock()
	stored, ok := sessions[key(sessionClaim)]
	mu.RUnlock()

	if !ok {
		return false
	}

	if time.Now().After(stored.expiresAt) {
		Delete(sessionClaim)
		return false
	}

	return subtle.ConstantTimeCompare([]byte(stored.token), []byte(token)) == 1
}

func Delete(sessionClaim string) {
	mu.Lock()
	defer mu.Unlock()

	delete(sessions, key(sessionClaim))
}

func deleteExpired() {
	now := time.Now()

	mu.Lock()
	defer mu.Unlock()

	for k, s := range sessions {
		if now.After(s.expiresAt) {
			delete(sessions, k)
		}
	}
}

func StartCleanup(interval time.Duration) {
	go func() {
		ticker := time.NewTicker(interval)
		defer ticker.Stop()

		for range ticker.C {
			deleteExpired()
		}
	}()

	slog.Info("Session cleanup started")
}
