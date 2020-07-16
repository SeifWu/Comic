package model

import "github.com/jinzhu/gorm"

// Comic 漫画模型
type Comic struct {
	gorm.Model

	Title    string     `json:"title" gorm:"type:varchar(20); not null; comment:'漫画名称'"`
	Describe string     `json:"describe" gorm:"comment:'描述'"`
	Author   string     `json:"author" gorm:"comment:'作者'"`
	Cover    Attachment `json:"cover" gorm:"polymorphic:Owner;comment:'封面'"`
}
