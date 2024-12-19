package model

import (
	"time"

	"encoding/json"

	"github.com/google/uuid"
)

type User struct {
	Id         uuid.UUID `json:"id" db:"id"`
	UserName   string    `json:"username" db:"username"`
	Password   string    `json:"password" db:"password"`
	Is_active  string    `json:"is_active" db:"is_active"`
	Created_at time.Time `json:"created_at" db:"created_at"`
	Updated_at time.Time `json:"updated_at" db:"updated_at"`
	
	Profile UserProfile `json:"Profile" db:"-"`
}

func (u *User) Valid() bool {
	return u.Id != uuid.Nil
}

func (u *User) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, u)
}

func (u *User) MarshalBinary(data []byte) ([]byte, error) {
	return json.Marshal(u)
}

type UserProfile struct {
	Id         uuid.UUID `json:"id" db:"id"`
	UserId     string    `json:"user_id" db:"user_id"`
	Name	   string    `json:"name" db:"name"`
	Role	   string    `json:"role" db:"role"`
	Email	   string    `json:"email" db:"email"`
	Phone	   string    `json:"phone" db:"phone"`
	Is_active  string    `json:"is_active" db:"is_active"`
	Created_at time.Time `json:"created_at" db:"created_at"`
	Updated_at time.Time `json:"updated_at" db:"updated_at"`
}

func (up *UserProfile) Valid() bool {
	return up.Id != uuid.Nil
}

func (up *UserProfile) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, up)
}

func (up *UserProfile) MarshalBinary(data []byte) ([]byte, error) {
	return json.Marshal(up)
}