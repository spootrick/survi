package crud

import (
	"errors"
	"github.com/jinzhu/gorm"
	"github.com/spootrick/survi/api/model"
	"github.com/spootrick/survi/api/util/channel"
)

type repositoryUserDetailsCRUD struct {
	db *gorm.DB
}

func NewRepositoryUserDetailsCRUD(db *gorm.DB) *repositoryUserDetailsCRUD {
	return &repositoryUserDetailsCRUD{db: db}
}

func (d *repositoryUserDetailsCRUD) Save(userDetail model.UserDetail) (model.UserDetail, error) {
	var err error
	done := make(chan bool)
	go func(ch chan<- bool) {
		defer close(ch)
		err = d.db.Debug().Model(&model.UserDetail{}).Create(&userDetail).Error
		if err != nil {
			ch <- false
			return
		}
		ch <- true
	}(done)

	if channel.Ok(done) {
		return userDetail, nil
	}
	return model.UserDetail{}, err
}

func (d *repositoryUserDetailsCRUD) FindById(userId uint) (model.UserDetail, error) {
	var err error
	var userDetail model.UserDetail
	done := make(chan bool)
	go func(ch chan<- bool) {
		defer close(ch)
		err = d.db.Debug().Model(&model.UserDetail{}).Where("user_id = ?", userId).Take(&userDetail).Error
		if err != nil {
			ch <- false
			return
		}
		err = d.db.Debug().Model(&model.User{}).Where("id = ?", userId).Take(&userDetail.User).Error
		if err != nil {
			ch <- false
			return
		}
		ch <- true
	}(done)
	if channel.Ok(done) {
		return userDetail, nil
	}
	if gorm.IsRecordNotFoundError(err) {
		return model.UserDetail{}, errors.New("user not found")
	}
	return model.UserDetail{}, err
}
