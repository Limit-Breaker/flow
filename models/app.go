package models

type AppModel struct {
	BaseModel
	Name  string
	Group string
	Type  string
}

func NewAppModel() *AppModel {
	return &AppModel{}
}
