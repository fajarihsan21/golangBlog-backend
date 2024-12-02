package model

import "time"

// create model based on database
type User struct {
	Id         string    `json:"id" db:"id"`
	UserName   string    `json:"username" db:"username"`
	Password   string    `json:"password" db:"password"`
	Created_at time.Time `json:"created_at" db:"created_at"`
	Updated_at time.Time `json:"updated_at" db:"updated_at"`
}

type UserProfile struct {
	Id         string    `json:"id" db:"id"`
	UserId     string    `json:"user_id" db:"user_id"`
	Created_at time.Time `json:"created_at" db:"created_at"`
	Updated_at time.Time `json:"updated_at" db:"updated_at"`
}
