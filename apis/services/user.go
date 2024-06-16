package services

import (
	"flow/common/request"
	"flow/config"
	"flow/models"
	"fmt"
)

type userService struct {
}

var UserService = new(userService)

// Register 注册
func (userService *userService) Register(params *request.Register) (error, *models.UserModel) {
	user := models.UserModel{}
	result := config.DB.Where("mobile = ?", params.Mobile).Select("id").First(&models.UserModel{})
	if result.RowsAffected != 0 {
		return fmt.Errorf("手机号已存在"), nil
	}

	user = models.UserModel{
		Name:     params.Name,
		Mobile:   params.Mobile,
		Password: params.Password,
	}
	if err := config.DB.Create(&user).Error; err != nil {
		return fmt.Errorf("creat user record failed: %s", err), nil
	}
	return nil, &user
}
