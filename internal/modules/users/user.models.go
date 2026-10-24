package users

import "time"

type UserDTO struct {
	Pid       string     `json:"pid"`
	Name      string     `json:"name"`
	Username  string     `json:"username"`
	Email     string     `json:"email"`
	Type      string     `json:"type"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
}

type UserRequest struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Type     string `json:"type"`
}

type UserDB struct {
	Id        int        `db:"id"`
	Pid       string     `db:"pid"`
	Active    bool       `db:"active"`
	Name      string     `db:"name"`
	Username  string     `db:"username"`
	Email     string     `db:"email"`
	Token     *string    `db:"token"`
	Password  string     `db:"password"`
	Type      string     `db:"type"`
	CreatedAt time.Time  `db:"created_at"`
	UpdatedAt *time.Time `db:"updated_at"`
	DeletedAt *time.Time `db:"deleted_at"`
}
