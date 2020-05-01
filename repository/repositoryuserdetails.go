package repository

import "github.com/spootrick/survi/api/model"

type UserDetailRepository interface {
	Save(user model.UserDetail) (model.UserDetail, error)
	//FindAll() ([]model.UserDetail, error)
	FindById(uint) (model.UserDetail, error)
	//Update(uint, model.UserDetail) (int64, error)
	//Delete(uint) (int64, error)
}
