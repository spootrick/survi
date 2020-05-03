package seed

import (
	"github.com/spootrick/survi/api/model"
	"github.com/spootrick/survi/api/util"
	"time"
)

var users = []model.User{
	{
		FirstName:  "Furkan",
		LastName:   "Karakoyunlu",
		Email:      "furkan@karakoyunlu.com",
		Password:   "123456",
		Role:       "ROLE_ADMIN,ROLE_USER",
		IsVerified: true,
		IsActive:   true,
	},
	{
		FirstName:  "John",
		LastName:   "Doe",
		Email:      "john@doe.com",
		Password:   "p4ssw0rd",
		IsVerified: false,
		IsActive:   true,
	},
}

var userDetails = []model.UserDetail{
	{
		BirthDate:  time.Now(),
		Gender:     "male",
		Profession: util.StrPtr("Computer Engineer"),
		Location:   "Istanbul",
		Height:     180,
		Weight:     74,
		Phone:      5375600268,
		Instagram:  util.StrPtr("spootrick"),
		IsPregnant: util.BoolPtr(false),
	},
	{
		BirthDate:  time.Now(),
		Gender:     "male",
		Profession: util.StrPtr("Dummy Prof"),
		Location:   "New York",
		Height:     178,
		Weight:     75,
		Phone:      5493857383,
		Instagram:  util.StrPtr("johndoe"),
		IsPregnant: util.BoolPtr(true),
	},
}
