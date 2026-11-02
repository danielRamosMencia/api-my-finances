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
