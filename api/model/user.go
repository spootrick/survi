package model

import (
	"errors"
	"github.com/jinzhu/gorm"
	"github.com/spootrick/survi/api/security"
	"github.com/spootrick/survi/api/util"
	"time"
)

type User struct {
	ID         uint      `gorm:"PRIMARY_KEY;AUTO_INCREMENT" json:"id"`
	FirstName  string    `gorm:"size:50;NOT NULL" json:"first_name"`
	LastName   string    `gorm:"size:50;NOT NULL" json:"last_name"`
	Email      string    `gorm:"size:50;NOT NULL;unique;unique_index" json:"email"`
	Password   string    `gorm:"size:100;NOT NULL" json:"password"`
	Role       string    `gorm:"size:50;NOT NULL;default:'ROLE_USER'" json:"role"`
	IsVerified bool      `gorm:"default:false" json:"is_verified"` // e-mail verification
	IsActive   bool      `gorm:"default:false" json:"is_active"`   // account is active or not
	CreatedAt  time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt  time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

// BeforeSave hash the user password before save
func (u *User) BeforeSave(scope *gorm.Scope) error {
	return hashPassword(scope, u)
}

// BeforeUpdate hash the user password before update
func (u *User) BeforeUpdate(scope *gorm.Scope) error {
	return hashPassword(scope, u)
}

func hashPassword(scope *gorm.Scope, user *User) error {
	hashedPassword, err := security.Hash(user.Password)
	if err != nil {
		return err
	}
	return scope.SetColumn("password", hashedPassword)
}

func (u *User) Prepare() {
	u.ID = 0
	u.FirstName = util.EscapeHTMLAndTrimString(u.FirstName)
	u.LastName = util.EscapeHTMLAndTrimString(u.LastName)
	u.Email = util.EscapeHTMLAndTrimString(u.Email)
	u.Role = util.EscapeHTMLAndTrimString(u.Role)
	u.IsVerified = false
	u.IsActive = true
	u.CreatedAt = time.Now()
	u.UpdatedAt = time.Now()
}

type Action string

const (
	Create Action = "create"
	Update Action = "update"
	Login  Action = "login"
)

func (u *User) Verify(action Action) error {
	// TODO: implement password checker 8-20 character allowed chars etc.
	switch action {
	case Update:
		// TODO: make a map for each case with fields and error messages and send this map to a method that verifies
		// given elements of map
		if u.FirstName == "" {
			return errors.New("first name is required")
		}

		if u.LastName == "" {
			return errors.New("last name is required")
		}

		if u.Email == "" {
			return errors.New("e-mail is required")
		}

		if !util.VerifyEmailFormat(u.Email) {
			return errors.New("invalid e-mail")
		}
	case Login:

	default:
		if u.FirstName == "" {
			return errors.New("first name is required")
		}

		if u.LastName == "" {
			return errors.New("last name is required")
		}

		if u.Email == "" {
			return errors.New("e-mail is required")
		}

		if !util.VerifyEmailFormat(u.Email) {
			return errors.New("invalid e-mail")
		}

		if u.Password == "" {
			return errors.New("password is required")
		}
	}
	return nil
}
