package movementtypes

// INCOME_CODE is the movement type code that adds to the account balance. Every other code subtracts.
const INCOME_CODE = "INCOME"

type MovementTypeDB struct {
	Id          int     `db:"id"`
	Pid         string  `db:"pid"`
	Active      bool    `db:"active"`
	Name        string  `db:"name"`
	Description *string `db:"description"`
	Code        string  `db:"code"`
	CreatedAt   string  `db:"created_at"`
	UpdatedAt   *string `db:"updated_at"`
	DeletedAt   *string `db:"deleted_at"`
}

type MovementTypeDTO struct {
	Pid         string  `json:"pid"`
	Active      bool    `json:"active"`
	Name        string  `json:"name"`
	Description *string `json:"description,omitempty"`
	Code        string  `json:"code"`
}
