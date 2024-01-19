package models

import (
	"time"
)

type Book struct {
	Id            int64     `gorm:"primaryKey" required:"true" json:"id"`
	Title         string    `gorm:"type:varchar(255);index;NOT NULL" required:"true" json:"title"`
	Description   string    `gorm:"type:text;NOT NULL" required:"true" json:"description"`
	Author        string    `gorm:"type:varchar(255);NOT NULL" required:"true" json:"author"`
	PublishedDate string    `gorm:"type:date;NOT NULL" required:"true" json:"published_date"`
	CreatedAt     time.Time `gorm:"type:datetime;default:CURRENT_TIMESTAMP();<-:create" json:"created_at"`
	UpdatedAt     time.Time `gorm:"type:datetime;default:CURRENT_TIMESTAMP()" json:"updated_at"`
	Reviews       []Review  `gorm:"foreign_key:Id" json:"reviews"`
}
