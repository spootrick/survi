package crud

import (
	"errors"
	"github.com/jinzhu/gorm"
	"github.com/spootrick/survi/api/model"
	"github.com/spootrick/survi/api/util/channel"
	"time"
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
		return model.UserDetail{}, errors.New("user detail not found")
	}
	return model.UserDetail{}, err
}

func (d *repositoryUserDetailsCRUD) Update(userId uint, userDetail model.UserDetail) (int64, error) {
	var result *gorm.DB
	done := make(chan bool)
	go func(ch chan<- bool) {
		defer close(ch)
		result = d.db.Debug().Model(&model.UserDetail{}).Where("user_id = ?", userId).Take(&model.UserDetail{}).Updates(model.UserDetail{
			BirthDate:  userDetail.BirthDate,
			Gender:     userDetail.Gender,
			Profession: userDetail.Profession,
			Location:   userDetail.Location,
			Height:     userDetail.Height,
			Weight:     userDetail.Weight,
			Phone:      userDetail.Phone,
			Instagram:  userDetail.Instagram,
			IsPregnant: userDetail.IsPregnant,
			UpdatedAt:  time.Now(),
		})
		ch <- true
	}(done)
	if channel.Ok(done) {
		if result.Error != nil {
			if gorm.IsRecordNotFoundError(result.Error) {
				return 0, errors.New("user detail not found")
			}
			return 0, result.Error
		}
		return result.RowsAffected, nil
	}
	return 0, result.Error
}
