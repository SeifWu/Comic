package model

import "github.com/jinzhu/gorm"

// ComicChapter 漫画章节
type ComicChapter struct {
	gorm.Model
	ChapterNumber string `json:"chapterNumber" gorm:"comment:'章节'"`
	Title         string `json:"title" gorm:"comment:'标题'"`
}
