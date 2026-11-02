package accounts

import (
	"context"
	"errors"

	"github.com/danielRamosMencia/api-my-finances/internal/utils"
)

func ServGetUserAccounts(ctx context.Context) ([]AccountDTO, error) {
	userPid := utils.GetClaim(ctx, "pid")
	accounts, err := RepoGetUserAccounts(userPid)
	if err != nil {
		return nil, err
	}

	dtos := make([]AccountDTO, len(accounts))
	for i, account := range accounts {
		dtos[i] = toDTO(account)
	}

	return dtos, nil
}

func ServGetByPid(pid string) (AccountDTO, error) {
	account, err := RepoGetByPid(pid)
	if err != nil {
		return AccountDTO{}, err
	}

	if account == nil {
		return AccountDTO{}, errors.New("account with pid " + pid + " not found")
	}

	return toDTO(*account), nil
}

func ServSave(ctx context.Context, request AccountRequest) (AccountDTO, error) {
	instant := utils.GetNowFormat()
	userPid := utils.GetClaim(ctx, "pid")

	account := AccountDB{
		Pid:       utils.RandomString(32),
		Active:    true,
		Alias:     request.Alias,
		amount:    utils.FormatToDigits(request.Amount, 2),
		userPid:   userPid,
		CreatedAt: instant,
	}

	err := RepoInsert(account)
	if err != nil {
		return AccountDTO{}, err
	}

	return toDTO(account), nil
}

func ServUpdate(pid string, request AccountRequest) (AccountDTO, error) {
	account, err := RepoGetByPid(pid)
	if err != nil {
		return AccountDTO{}, err
	}

	if account == nil {
		return AccountDTO{}, errors.New("account with pid " + pid + " not found")
	}

	account.Alias = request.Alias
	account.amount = utils.FormatToDigits(request.Amount, 2)
	account.UpdatedAt = utils.Ptr(utils.GetNowFormat())

	err = RepoUpdate(pid, *account)
	if err != nil {
		return AccountDTO{}, err
	}

	return toDTO(*account), nil
}

func toDTO(account AccountDB) AccountDTO {
	return AccountDTO{
		Pid:       account.Pid,
		Active:    account.Active,
		Alias:     account.Alias,
		Amount:    account.amount,
		UserPid:   account.userPid,
		CreatedAt: account.CreatedAt,
		UpdatedAt: account.UpdatedAt,
	}
}
