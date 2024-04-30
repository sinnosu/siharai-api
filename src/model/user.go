package model

import "time"

type User struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
    CompanyID          int    `json:"company_id"`
	Email     string    `json:"email" gorm:"unique"`
    Name               string `json:"name"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type UserResponse struct {
	ID    uint   `json:"id" gorm:"priaryKey"`
	Email string `json:"email" gorm:"unique"`
}
