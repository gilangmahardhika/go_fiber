package models

type Book struct {
	Id            int64  `gorm:"primaryKey" json:"id"`
	Title         string `gorm:"type:varchar(255)" json:"title"`
	Description   string `gorm:"type:text" json:"description"`
	Author        string `gorm:"type:varchar(255)" json:"author"`
	PublishedDate string `gorm:"type:date" json:"published_date"`
}
