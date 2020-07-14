package model

// ComicChapterDetail 漫画正文
type ComicChapterDetail struct {
	Sort           int `json:"sort" gorm:"comment:'序号'"`
	ComicChapterID int
	ComicChapter   ComicChapter
	Attachment     Attachment `gorm:"polymorphic:Owner;"`
}
