package crud

import (
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
	return model.User{}, nil
}
