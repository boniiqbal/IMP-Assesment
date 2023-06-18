package domain

import (
	"time"
)

type Users struct {
	CreatedAt time.Time `json:"created_at" gorm:"created_at"`
	ID        int64     `json:"id" gorm:"id"`
	FullName  string    `json:"full_name" gorm:"full_name"`
	Username  string    `json:"username" gorm:"username"`
	Password  string    `json:"password" gorm:"password"`
}

type UserParams struct {
	BasedFilter
}
