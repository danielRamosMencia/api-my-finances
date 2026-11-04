package movementhistory

import (
	"context"
	"net/http"

	"github.com/danielRamosMencia/api-my-finances/internal/responses"
	"github.com/danielRamosMencia/api-my-finances/internal/utils"
)

func ServGetUserHistory(ctx context.Context, year, month int) ([]MovementHistoryDB, error) {
	from, to, err := utils.GetMonthRange(year, month)
	if err != nil {
		return nil, responses.NewError(http.StatusBadRequest, responses.INVALID_DATE_FILTER)
	}

	userPid := utils.GetClaim(ctx, "pid")
	history, err := RepoGetUserHistory(userPid, from, to)
	if err != nil {
		return nil, err
	}

	return history, nil
}

func ServGetUserMetrics(ctx context.Context, year, month int) (UserMetrics, error) {
	from, to, err := utils.GetMonthRange(year, month)
	if err != nil {
		return UserMetrics{}, responses.NewError(http.StatusBadRequest, responses.INVALID_DATE_FILTER)
	}

	userPid := utils.GetClaim(ctx, "pid")
	return RepoGetUserMetrics(userPid, from, to)
}
