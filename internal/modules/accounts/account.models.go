package accounts

type AccountDB struct {
	Id        int     `db:"id"`
	Pid       string  `db:"pid"`
	Active    bool    `db:"active"`
	Alias     string  `db:"alias"`
	amount    float64 `db:"amount"`
	userId    int     `db:"user_id"`
	userPid   string  `db:"user_pid"`
	CreatedAt string  `db:"created_at"`
	UpdatedAt *string `db:"updated_at"`
	DeletedAt *string `db:"deleted_at"`
}

type AccountDTO struct {
	Pid       string  `json:"pid"`
	Active    bool    `json:"active"`
	Alias     string  `json:"alias"`
	Amount    float64 `json:"amount"`
	UserPid   string  `json:"user_pid"`
	CreatedAt string  `json:"created_at"`
	UpdatedAt *string `json:"updated_at"`
}

type AccountRequest struct {
	Alias  string  `json:"alias"`
	Amount float64 `json:"amount"`
}
