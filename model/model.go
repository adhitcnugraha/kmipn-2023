package model

import "time"

type User struct {
	ID        int       `gorm:"primaryKey" json:"id"`
	Fullname  string    `json:"fullname" gorm:"type:varchar(255);"`
	Email     string    `json:"email" gorm:"type:varchar(255);not null"`
	Password  string    `json:"-" gorm:"type:varchar(255);not null"`
	Address   string    `json:"fullname" gorm:"type:varchar(255);"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type UserLogin struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type UserRegister struct {
	Fullname string `json:"fullname" gorm:"type:varchar(255);"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type Seller struct {
	ID        int    `gorm:"primaryKey" json:"id"`
	Email     string `json:"email" gorm:"type:varchar(255);not null"`
	Password  string `json:"-" gorm:"type:varchar(255);not null"`
	Address   string `json:"fullname" gorm:"type:varchar(255);not null"`
	ShopName  string `json:"shop_name" gorm:"type:varchar(255);not null"`
	Products  []Product
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Admin struct {
	ID       int    `gorm:"primaryKey" json:"id"`
	Email    string `json:"email" gorm:"type:varchar(255);not null"`
	Password string `json:"-" gorm:"type:varchar(255);not null"`
}

type Product struct {
	ID          int       `gorm:"primaryKey" json:"id"`
	SellerID    int       `json:"seller_id" gorm:"not null"`
	ProductName string    `json:"product_name" gorm:"type:varchar(255);not null"`
	Description string    `json:"description" gorm:"type:text;not null"`
	Price       float64   `json:"price"`
	Stock       int       `json:"stock"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type UserProductCategory struct {
	ID          int     `gorm:"primaryKey" json:"id"`
	ProductName string  `json:"product_name" gorm:"type:varchar(255);not null"`
	Description string  `json:"description" gorm:"type:text;not null"`
	Category    string  `json:"category" gorm:"type:varchar(255);not null"`
	Price       float64 `json:"price"`
	Stock       int     `json:"stock"`
}

type Order struct {
	ID        int       `gorm:"primaryKey" json:"id"`
	UserID    int       `json:"user_id" gorm:"not null"`
	ProductID int       `json:"product_id" gorm:"not null"`
	Quantity  int       `json:"quantity" gorm:"not null"`
	Status    string    `json:"status" gorm:"type:varchar(255);not null"`
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
