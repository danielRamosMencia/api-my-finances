package movementhistory

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/danielRamosMencia/api-my-finances/internal/responses"
)

type dateFilter struct {
	year     int
	month    int
	location *time.Location
}

// getDateFilter reads the year, month and timezone (IANA name, e.g. America/Tegucigalpa) query params.
// Defaults: UTC and the current year/month in that timezone.
func getDateFilter(r *http.Request) (dateFilter, error) {
	query := r.URL.Query()

	location := time.UTC
	if value := query.Get("timezone"); value != "" {
		loaded, err := time.LoadLocation(value)
		// "Local" would be the server's timezone
		if err != nil || value == "Local" {
			return dateFilter{}, responses.NewError(http.StatusBadRequest, responses.INVALID_TIMEZONE)
		}
		location = loaded
	}

	now := time.Now().In(location)
	filter := dateFilter{year: now.Year(), month: int(now.Month()), location: location}

	if value := query.Get("year"); value != "" {
		parsed, err := strconv.Atoi(value)
		if err != nil {
			return dateFilter{}, responses.NewError(http.StatusBadRequest, responses.INVALID_DATE_FILTER)
		}
		filter.year = parsed
	}

	if value := query.Get("month"); value != "" {
		parsed, err := strconv.Atoi(value)
		if err != nil {
			return dateFilter{}, responses.NewError(http.StatusBadRequest, responses.INVALID_DATE_FILTER)
		}
		filter.month = parsed
	}

	return filter, nil
}

func HldrGetUserHistory(w http.ResponseWriter, r *http.Request) {
	filter, err := getDateFilter(r)
	if err != nil {
		responses.HandleError(w, err)
		return
	}

	history, err := ServGetUserHistory(r.Context(), filter.year, filter.month, filter.location)
	if err != nil {
		responses.HandleError(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(history)
}

func HldrGetUserMetrics(w http.ResponseWriter, r *http.Request) {
	filter, err := getDateFilter(r)
	if err != nil {
		responses.HandleError(w, err)
		return
	}

	metrics, err := ServGetUserMetrics(r.Context(), filter.year, filter.month, filter.location)
	if err != nil {
		responses.HandleError(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(metrics)
}
