package model

import "github.com/jinzhu/gorm"

// Attachment 附件
type Attachment struct {
	gorm.Model

	FileName  string
	FileSize  string
	OwnerID   int
	OwnerType string
}
