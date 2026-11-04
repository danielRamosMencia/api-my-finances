package movements

type MovementDB struct {
	Id               int     `db:"id"`
	Pid              string  `db:"pid"`
	Active           bool    `db:"active"`
	Code             string  `db:"code"`
	Description      string  `db:"description"`
	Amount           float64 `db:"amount"`
	AccountId        int     `db:"account_id"`
	AccountPid       string  `db:"account_pid"`
	AccountAlias     string  `db:"account_alias"`
	MovementTypeId   int     `db:"movement_type_id"`
	MovementTypePid  string  `db:"movement_type_pid"`
	MovementTypeName string  `db:"movement_type_name"`
	UserId           int     `db:"user_id"`
	UserPid          string  `db:"user_pid"`
	Username         string  `db:"username"`
	CreatedAt        string  `db:"created_at"`
	UpdatedAt        *string `db:"updated_at"`
	DeletedAt        *string `db:"deleted_at"`
}

type MovementDTO struct {
	Pid              string  `json:"pid"`
	Active           bool    `json:"active"`
	Code             string  `json:"code"`
	Description      string  `json:"description"`
	Amount           float64 `json:"amount"`
	AccountPid       string  `json:"account_pid"`
	AccountAlias     string  `json:"account_alias"`
	MovementTypePid  string  `json:"movement_type_pid"`
	MovementTypeName string  `json:"movement_type_name"`
	UserPid          string  `json:"user_pid"`
	Username         string  `json:"username"`
	CreatedAt        string  `json:"created_at"`
	UpdatedAt        *string `json:"updated_at"`
}

type MovementRequest struct {
	Description     string  `json:"description"`
	Amount          float64 `json:"amount"`
	MovementTypePid string  `json:"movement_type_pid"`
	AccountPid      string  `json:"account_pid"`
	Income          bool    `json:"income"`
}
