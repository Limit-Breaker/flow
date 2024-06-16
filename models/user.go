package models

type UserModel struct {
	BaseModel
	Name     string
	Mobile   string
	Password string
}

func (UserModel) TableName() string {
	return "user"
}
