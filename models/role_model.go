package model

import "gorm.io/gorm"

type Role struct {
    gorm.Model
    RoleName string `gorm:"not null;unique" json:"role_name"`
    Users    []User // 一個角色對應多個使用者
}