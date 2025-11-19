package models

import "time"

type Schedule struct {
	ID          uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	TeacherID   uint      `gorm:"index;not null" json:"teacher_id"`
	ClassCode   string    `gorm:"type:varchar(255);not null" json:"class_code"`
	ClassName   string    `gorm:"type:varchar(255);not null" json:"class_name"`
	SubjectCode string    `gorm:"type:varchar(255);not null" json:"subject_code"`
	Date        time.Time `json:"date"`
	JamKe       int       `json:"jam_ke"`
	TimeStart   string    `gorm:"type:time" json:"time_start"`
	TimeEnd     string    `gorm:"type:time" json:"time_end"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`

	// Relation
	Teacher Teacher `gorm:"foreignKey:TeacherID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"teacher"`
}

func (Schedule) TableName() string {
	return "schedule"
}
