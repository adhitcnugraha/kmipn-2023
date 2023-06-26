package model

import "time"

type User struct {
	ID        int       `gorm:"primaryKey" json:"id"`
	Email     string    `json:"email" gorm:"type:varchar(255);not null"`
	Password  string    `json:"-" gorm:"type:varchar(255);not null"`
	Address   string    `json:"fullname" gorm:"type:varchar(255);"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Session struct {
	ID     int       `gorm:"primaryKey" json:"id"`
	Token  string    `json:"token"`
	Email  string    `json:"email"`
	Expiry time.Time `json:"expiry"`
}

type Credential struct {
	Host         string
	Username     string
	Password     string
	DatabaseName string
	Port         int
	Schema       string
}
