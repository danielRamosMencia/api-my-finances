package movementhistory

import "github.com/danielRamosMencia/api-my-finances/internal/utils"

func ServGetUserHistory() ([]MovementHistoryDB, error) {
	userPid := utils.MockGetClaim("pid")
	history, err := RepoGetUserHistory(userPid)
	if err != nil {
		return nil, err
	}

	return history, nil
}
