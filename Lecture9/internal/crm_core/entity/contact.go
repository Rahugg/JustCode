package entity

import "gorm.io/gorm"

//Clients

type Contact struct {
	gorm.Model
	FirstName string `gorm:"varchar(255);not null" json:"first_name"`
	LastName  string `gorm:"varchar(255);not null" json:"last_name"`
	Email     string `gorm:"varchar(255);not null" json:"email"`
	Phone     string `gorm:"varchar(255);not null" json:"phone"`
	CompanyID uint   `json:"company_id"`
}
