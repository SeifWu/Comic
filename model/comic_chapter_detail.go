package model

// ComicChapterDetail 漫画正文
type ComicChapterDetail struct {
	Sort int    `json:"sort" gorm:"comment:'序号'"`
	URL  string `json:"url" gorm:"comment:'图片地址'"`
}
