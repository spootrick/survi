package repository

import "github.com/spootrick/survi/api/model"

type UserRepository interface {
	Save(user model.User) (model.User, error)
	//FindAll([]model.User) error
	//FindById(uint) (model.User, error)
	//Update(uint, model.User) (int64, error)
	//Delete(uint) (int64, error)
}
