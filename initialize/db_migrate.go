package initialize

import (
	"seifwu.com/gin-basic-project/global"
	"seifwu.com/gin-basic-project/model"
)

// DBMigrate Migrate Table
func DBMigrate() {
	db := global.DB

	db.AutoMigrate(
		&model.User{},
		&model.Comic{},
		&model.ComicChapter{},
		&model.ComicChapterDetail{},
		&model.Attachment{},
	)
}
