package entity

import "gorm.io/gorm"

type User struct {
	gorm.Model
	ID        uint64
	FirstName string
	LastName  string
	Age       uint64
	Phone     string
	RoleID    uint `json:"role_id"`
	Role      Role `gorm:"not null" json:"role"`
	Email     string
	Provider  string `gorm:"not null" json:"provider"`
	Password  string
	Verified  bool
}
