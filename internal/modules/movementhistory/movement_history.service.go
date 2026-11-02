package movementhistory

import (
	"context"

	"github.com/danielRamosMencia/api-my-finances/internal/utils"
)

func ServGetUserHistory(ctx context.Context) ([]MovementHistoryDB, error) {
	userPid := utils.GetClaim(ctx, "pid")
	history, err := RepoGetUserHistory(userPid)
	if err != nil {
		return nil, err
	}

	return history, nil
}
