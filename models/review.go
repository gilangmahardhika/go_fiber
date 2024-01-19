package models

import "time"

type Review struct {
	Id        int       `gorm:"primaryKey" required:"true" json:"id"`
	Review    string    `gorm:"text;required:true;NOT NULL" json:"review"`
	CreatedAt time.Time `gorm:"type:datetime;default:CURRENT_TIMESTAMP();<-:create" json:"created_at"`
	UpdatedAt time.Time `gorm:"type:datetime;default:CURRENT_TIMESTAMP()" json:"updated_at"`
	BookId    int       `gorm:"foreignKey:books(id)" json:"book"`
	UserId    int       `gorm:"foreignKey:users(id)" json:"user"`
}
