package auth

import "time"

type User struct {
	ID        string    `json:"id"`
	FirstName string    `json:"first_name"`
	Username  string    `json:"username"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"created_at"`
}

type Preference struct {
	ID     string `json:"id"`
	UserID string `json:"user_id"`
}
