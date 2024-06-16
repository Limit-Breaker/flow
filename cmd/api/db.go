package api

import (
	"flow/models"
	"fmt"
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) {
	//db.SingularTable(true)
	fmt.Println("----db: ", db.Name())
	err := db.AutoMigrate(
		&models.UserModel{},
	)

	if err != nil {
		fmt.Printf("migrate table failed: %s\n", err)
		panic("migrate table failed")
	}
}
