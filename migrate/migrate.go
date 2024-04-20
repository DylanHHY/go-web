package main

import (
	"fmt"
	"go-side-project/initializers"
	model "go-side-project/models"
	"math/rand"
	"time"

	"github.com/bxcodec/faker/v3"
)

func init() {
	initializers.ConnectToDB()
}

func main() {

	initializers.DB.AutoMigrate(&model.User{})
	initializers.DB.AutoMigrate(&model.Post{})
	initializers.DB.AutoMigrate(&model.Role{})

}

func createRole(){
	roleNames := []string{"super admin", "admin", "regular user", "paid user", "guest"}
	for _, name := range roleNames {
		role := model.Role{RoleName: name}
		if err := initializers.DB.Create(&role).Error; err != nil {
			fmt.Printf("Failed to create role %s: %v\n", name, err)
		} else {
			fmt.Printf("Role %s created successfully\n", name)
		}
	}
}

func fakerContent(){
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	// 生成假数据
	for i := 0; i < 10; i++ {
		title := fmt.Sprintf("Title %d", i+1)
		content := faker.Paragraph()
		userID := uint(r.Intn(4) + 2) // 随机生成2~5之间的userID

		post := model.Post{
			Title:   title,
			Content: content,
			UserID:  userID,
		}

		initializers.DB.Create(&post)
	}

	fmt.Println("Fake posts generated successfully.")

}