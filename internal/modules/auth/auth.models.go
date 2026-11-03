package auth

type UserDB struct {
	Name     string `db:"name"`
	Username string `db:"username"`
	Pid      string `db:"pid"`
	Email    string `db:"email"`
	Password string `db:"password"`
}

type LoginRequest struct {
	Username string `db:"username"`
	Password string `db:"password"`
}

type LoginResponse struct {
	Username string `json:"username"`
	Pid      string `json:"pid"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Token    string `json:"token"`
}
