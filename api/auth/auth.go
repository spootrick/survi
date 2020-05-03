package auth

import (
	"github.com/jinzhu/gorm"
	"github.com/spootrick/survi/api/database"
	"github.com/spootrick/survi/api/model"
	"github.com/spootrick/survi/api/security"
	"github.com/spootrick/survi/api/util/channel"
)

func SignIn(email string, password string) (string, error) {
	user := model.User{}
	var err error
	var db *gorm.DB
	done := make(chan bool)

	go func(ch chan<- bool) {
		defer close(ch)
		db, err = database.Connect()
		if err != nil {
			ch <- false
			return
		}
		defer db.Close()

		err = db.Debug().Model(model.User{}).Where("email = ?", email).Take(&user).Error
		if err != nil {
			ch <- false
			return
		}

		err = security.VerifyPassword(user.Password, password)
		if err != nil {
			ch <- false
			return
		}
		ch <- true
	}(done)

	if channel.Ok(done) {
		user.Password = ""
		return GenerateJWT(user)
	}

	return "", err

}
