package crud

import (
	"errors"
	"github.com/jinzhu/gorm"
	"github.com/spootrick/survi/api/model"
	"github.com/spootrick/survi/api/util/channel"
)

type repositoryUserCRUD struct {
	db *gorm.DB
}

func NewRepositoryUserCRUD(db *gorm.DB) *repositoryUserCRUD {
	return &repositoryUserCRUD{db: db}
}

func (r *repositoryUserCRUD) Save(user model.User) (model.User, error) {
	var err error
	done := make(chan bool)
	go func(ch chan<- bool) {
		err = r.db.Debug().Model(&model.User{}).Create(&user).Error
		if err != nil {
			ch <- false
			return
		}
		ch <- true
	}(done)

	if channel.Ok(done) {
		return user, nil
	}
	return model.User{}, err
}

func (r *repositoryUserCRUD) FindAll() ([]model.User, error) {
	var err error
	var users []model.User
	done := make(chan bool)
	go func(ch chan<- bool) {
		err = r.db.Debug().Model(&model.User{}).Limit(100).Find(&users).Error
		if err != nil {
			ch <- false
			return
		}
		ch <- true
	}(done)

	if channel.Ok(done) {
		return users, nil
	}
	return nil, err
}

func (r *repositoryUserCRUD) FindById(id uint) (model.User, error) {
	var err error
	var user model.User
	done := make(chan bool)
	go func(ch chan<- bool) {
		err = r.db.Debug().Model(&model.User{}).Where("id = ?", id).Take(&user).Error
		if err != nil {
			ch <- false
			return
		}
		ch <- true
	}(done)
	if channel.Ok(done) {
		return user, nil
	}
	if gorm.IsRecordNotFoundError(err) {
		return model.User{}, errors.New("user not found")
	}
	return model.User{}, err
}
