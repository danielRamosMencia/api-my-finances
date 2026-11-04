package movementhistory

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/danielRamosMencia/api-my-finances/internal/responses"
)

// getYearMonth reads the year and month query params. A missing param defaults to the current UTC year/month.
func getYearMonth(r *http.Request) (int, int, bool) {
	now := time.Now().UTC()
	year, month := now.Year(), int(now.Month())

	if value := r.URL.Query().Get("year"); value != "" {
		parsed, err := strconv.Atoi(value)
		if err != nil {
			return 0, 0, false
		}
		year = parsed
	}

	if value := r.URL.Query().Get("month"); value != "" {
		parsed, err := strconv.Atoi(value)
		if err != nil {
			return 0, 0, false
		}
		month = parsed
	}

	return year, month, true
}

func HldrGetUserHistory(w http.ResponseWriter, r *http.Request) {
	year, month, ok := getYearMonth(r)
	if !ok {
		responses.Error(w, http.StatusBadRequest, responses.INVALID_DATE_FILTER)
		return
	}

	history, err := ServGetUserHistory(r.Context(), year, month)
	if err != nil {
		responses.HandleError(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(history)
}

func HldrGetUserMetrics(w http.ResponseWriter, r *http.Request) {
	year, month, ok := getYearMonth(r)
	if !ok {
		responses.Error(w, http.StatusBadRequest, responses.INVALID_DATE_FILTER)
		return
	}

	metrics, err := ServGetUserMetrics(r.Context(), year, month)
	if err != nil {
		responses.HandleError(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(metrics)
}
