package models

import "time"

type User struct {
	ID        uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	Username  string    `gorm:"type:varchar(255);not null" json:"username"`
	Password  string    `gorm:"type:text;not null" json:"password"`
	Role      string    `gorm:"type:varchar(255)" json:"role"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	// Relation
	Teacher Teacher `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL" json:"teacher"`
}

func (User) TableName() string {
	return "user"
}
