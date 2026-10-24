package users

type UserDTO struct {
	Pid       string  `json:"pid"`
	Name      string  `json:"name"`
	Active    bool    `json:"active"`
	Username  string  `json:"username"`
	Email     string  `json:"email"`
	Type      string  `json:"type"`
	CreatedAt string  `json:"created_at"`
	UpdatedAt *string `json:"updated_at"`
}

type UserRequest struct {
	Name     string `json:"name"`
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
	Type     string `json:"type"`
}

type UserDB struct {
	Id        int     `db:"id"`
	Pid       string  `db:"pid"`
	Active    bool    `db:"active"`
	Name      string  `db:"name"`
	Username  string  `db:"username"`
	Email     string  `db:"email"`
	Token     *string `db:"token"`
	Password  string  `db:"password"`
	Type      string  `db:"type"`
	CreatedAt string  `db:"created_at"`
	UpdatedAt *string `db:"updated_at"`
	DeletedAt *string `db:"deleted_at"`
}
