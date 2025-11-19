package models

import "time"

type Teacher struct {
	ID        uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	UserID    *uint     `gorm:"index" json:"user_id"` // nullable
	Name      string    `gorm:"type:varchar(255)" json:"name"`
	NIK       string    `gorm:"type:varchar(255)" json:"nik"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	// Relations
	User      *User      `gorm:"foreignKey:UserID;references:ID" json:"user"`
	Schedules []Schedule `gorm:"foreignKey:TeacherID;references:ID" json:"schedules"`
}

func (Teacher) TableName() string {
	return "teacher"
}
