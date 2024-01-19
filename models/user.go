package models

import "time"

type User struct {
	Id        int       `gorm:"primary_key" json:"id"`
	Name      string    `gorm:"type:varchar(255);index;NOT NULL" json:"name"`
	Email     string    `gorm:"type:varchar(255);index;NOT NULL" json:"email"`
	Password  string    `gorm:"type:varchar(255);index;NOT NULL" json:"password"`
	Token     string    `gorm:"type:varchar(255);index" json:"token"`
	CreatedAt time.Time `gorm:"type:datetime;default:CURRENT_TIMESTAMP();<-:create" json:"created_at"`
	UpdatedAt time.Time `gorm:"type:datetime;default:CURRENT_TIMESTAMP()" json:"updated_at"`
	Reviews   []Review  `gorm:"foreignKey:Id" json:"reviews"`
}
