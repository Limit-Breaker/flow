package models

type Settings struct {
	BaseModel
	Classify int `gorm:"column:classify; type:int(11)" json:"classify" form:"classify"` // 设置分类，1 配置信息，2 Ldap配置
	//Content  json.RawMessage `gorm:"column:content; type:json" json:"content" form:"content"`       // 配置内容
}
