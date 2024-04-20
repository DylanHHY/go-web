package model

import "gorm.io/gorm"

type User struct {
    gorm.Model
    Name     string `json:"name" gorm:"not null"`
    Email    string `json:"email" gorm:"not null"`
    Password string `gorm:"not null"`
    RoleID   uint   `json:"role_id" gorm:"not null"`
    Role     Role   `gorm:"foreignKey:RoleID"`
    Posts    []Post `gorm:"foreignKey:UserID"`
}