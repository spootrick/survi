package model

import "time"

type UserDetail struct {
	ID         uint      `gorm:"PRIMARY_KEY;AUTO_INCREMENT" json:"id"`
	User       User      `json:"user"` // gorm finds foreign key as UserID
	UserID     uint      `gorm:"NOT NULL" json:"user_id"`
	BirthDate  time.Time `gorm:"type:DATE NOT NULL" json:"birth_date"`
	Gender     string    `gorm:"size:6;NOT NULL" json:"gender"`
	Profession string    `gorm:"size:100;NOT NULL" json:"profession"`
	Location   string    `gorm:"size:100;NOT NULL" json:"location"`
	Height     uint      `gorm:"NOT NULL" json:"height"`
	Weight     uint      `gorm:"NOT NULL" json:"weight"`
	Phone      uint64    `gorm:"NOT NULL" json:"phone"`
	Instagram  string    `gorm:"size:30" json:"instagram;omitempty"`
	IsPregnant bool      `gorm:"NOT NULL"`
}
