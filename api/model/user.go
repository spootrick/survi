package model

import (
	"github.com/spootrick/survi/api/security"
	"time"
)

type User struct {
	ID        uint      `gorm:"PRIMARY_KEY;AUTO_INCREMENT" json:"id"`
	FirstName string    `gorm:"size:50;NOT NULL" json:"first_name"`
	LastName  string    `gorm:"size:50;NOT NULL" json:"last_name"`
	Email     string    `gorm:"size:50;NOT NULL;unique;unique_index" json:"email"`
	Password  string    `gorm:"size:100;NOT NULL" json:"password"`
	Roles     string    `gorm:"size:50;NOT NULL;default:'ROLE_USER'" json:"roles"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

func (u *User) BeforeSave() error {
	hashedPassword, err := security.Hash(u.Password)
	if err != nil {
		return err
	}
	u.Password = string(hashedPassword)
	return nil
}
