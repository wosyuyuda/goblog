package model

//友情链接
type Link struct {
	ID   uint   `gorm:"primarykey" form:"id" json:"id"`
	Url  string `gorm:"size:255;" form:"url" json:"url"`
	Name string `gorm:"size:255;" form:"name" json:"name"`
}
