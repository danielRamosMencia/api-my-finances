package movementhistory

type MovementHistoryDB struct {
	Pid           string  `db:"pid" json:"pid"`
	HistoryAmount float64 `db:"history_amount" json:"history_amount"`
	Amount        float64 `db:"amount" json:"amount"`
	MovementCode  string  `db:"movement_code" json:"movement_code"`
	Account       string  `db:"account" json:"account"`
	Type          string  `db:"type" json:"type"`
	TypeCode      string  `db:"type_code" json:"type_code"`
	Username      string  `db:"username" json:"username"`
	Email         string  `db:"email" json:"email"`
	Date          string  `db:"date" json:"date"`
}

type UserMetrics struct {
	From           string           `json:"from"`
	To             string           `json:"to"`
	TotalMovements int              `db:"total_movements" json:"total_movements"`
	TotalIncome    float64          `db:"total_income" json:"total_income"`
	TotalExpense   float64          `db:"total_expense" json:"total_expense"`
	Net            float64          `db:"net" json:"net"`
	Accounts       []AccountMetrics `json:"accounts"`
}

type AccountMetrics struct {
	AccountPid     string  `db:"account_pid" json:"account_pid"`
	Account        string  `db:"account" json:"account"`
	TotalMovements int     `db:"total_movements" json:"total_movements"`
	TotalIncome    float64 `db:"total_income" json:"total_income"`
	TotalExpense   float64 `db:"total_expense" json:"total_expense"`
	Net            float64 `db:"net" json:"net"`
}
