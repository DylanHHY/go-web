package model

import "gorm.io/gorm"

type User struct {
    gorm.Model
    Name        string `json:"name" gorm:"not null"`
    Email       string `json:"email" gorm:"not null"`
    Password    string `gorm:"not null"`
    Role        string `json:"role" gorm:"not null"`
    Posts       []Post `gorm:"foreignKey:UserID"`
}