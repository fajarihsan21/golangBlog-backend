package model

import "time"

type User struct {
	Id         string    `json:"id" db:"id"`
	UserName   string    `json:"username" db:"username"`
	Password   string    `json:"password" db:"password"`
	Is_active  string    `json:"is_active" db:"is_active"`
	Created_at time.Time `json:"created_at" db:"created_at"`
	Updated_at time.Time `json:"updated_at" db:"updated_at"`
	
	Profile UserProfile `json:"Profile" db:"-"`
}

type UserProfile struct {
	Id         string    `json:"id" db:"id"`
	UserId     string    `json:"user_id" db:"user_id"`
	Name	   string    `json:"name" db:"name"`
	Role	   string    `json:"role" db:"role"`
	Email	   string    `json:"email" db:"email"`
	Phone	   string    `json:"phone" db:"phone"`
	Is_active  string    `json:"is_active" db:"is_active"`
	Created_at time.Time `json:"created_at" db:"created_at"`
	Updated_at time.Time `json:"updated_at" db:"updated_at"`
}
