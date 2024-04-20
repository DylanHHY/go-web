package model

import (
	"gorm.io/gorm"
)

type Post struct {
    gorm.Model
    Title       string `json:"title" gorm:"not null"`
    Content     string `json:"content"`
    UserID      uint   // 外鍵，指向 User 的 ID
}
