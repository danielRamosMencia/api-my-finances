package movementhistory

import (
	"context"
	"net/http"
	"time"

	"github.com/danielRamosMencia/api-my-finances/internal/responses"
	"github.com/danielRamosMencia/api-my-finances/internal/utils"
)

func ServGetUserHistory(ctx context.Context, year, month int, location *time.Location) ([]MovementHistoryDB, error) {
	monthRange, err := utils.GetMonthRange(year, month, location)
	if err != nil {
		return nil, responses.NewError(http.StatusBadRequest, responses.INVALID_DATE_FILTER)
	}

	userPid := utils.GetClaim(ctx, "pid")
	history, err := RepoGetUserHistory(userPid, monthRange.From, monthRange.To)
	if err != nil {
		return nil, err
	}

	return history, nil
}

func ServGetUserMetrics(ctx context.Context, year, month int, location *time.Location) (UserMetrics, error) {
	monthRange, err := utils.GetMonthRange(year, month, location)
	if err != nil {
		return UserMetrics{}, responses.NewError(http.StatusBadRequest, responses.INVALID_DATE_FILTER)
	}

	userPid := utils.GetClaim(ctx, "pid")
	metrics, err := RepoGetUserMetrics(userPid, monthRange.From, monthRange.To)
	if err != nil {
		return UserMetrics{}, err
	}

	metrics.From = monthRange.FirstDay
	metrics.To = monthRange.LastDay
	metrics.Timezone = location.String()

	return metrics, nil
}
