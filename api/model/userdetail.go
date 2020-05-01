package model

import (
	"errors"
	"github.com/spootrick/survi/api/util"
	"time"
)

type UserDetail struct {
	ID         uint      `gorm:"PRIMARY_KEY;AUTO_INCREMENT" json:"id"`
	User       User      `json:"user"` // gorm finds foreign key as UserID
	UserID     uint      `gorm:"NOT NULL" json:"user_id"`
	BirthDate  time.Time `gorm:"type:DATE NOT NULL" json:"birth_date"`
	Gender     string    `gorm:"size:6;NOT NULL" json:"gender"`
	Profession string    `gorm:"size:100;" json:"profession;omitempty"`
	Location   string    `gorm:"size:100;NOT NULL" json:"location"`
	Height     uint      `gorm:"NOT NULL" json:"height"`
	Weight     uint      `gorm:"NOT NULL" json:"weight"`
	Phone      uint64    `gorm:"NOT NULL" json:"phone"`
	Instagram  string    `gorm:"size:30" json:"instagram;omitempty"`
	IsPregnant bool      `gorm:"NOT NULL"`
	CreatedAt  time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt  time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

func (d *UserDetail) Prepare() {
	d.ID = 0
	d.User = User{}
	d.Gender = util.EscapeHTMLAndTrimString(d.Gender)
	d.Profession = util.EscapeHTMLAndTrimString(d.Profession)
	d.Location = util.EscapeHTMLAndTrimString(d.Location)
	d.Instagram = util.EscapeHTMLAndTrimString(d.Instagram)
	d.CreatedAt = time.Now()
	d.UpdatedAt = time.Now()
}

func (d *UserDetail) Verify() error {
	if d.UserID < 1 {
		return errors.New("user is required")
	}

	// TODO: Find a way to check birth date

	if d.Gender == "" {
		return errors.New("gender is required")
	}

	if d.Location == "" {
		return errors.New("location is required")
	}

	if d.Height <= 0 {
		return errors.New("height must be positive")
	}

	if d.Weight <= 0 {
		return errors.New("weight must be positive")
	}

	if d.Phone <= 5000000000 || d.Phone > 5999999999 {
		return errors.New("invalid phone number")
	}

	return nil
}
