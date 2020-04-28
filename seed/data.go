package seed

import (
	"github.com/spootrick/survi/api/model"
	"time"
)

var users = []model.User{
	{
		FirstName: "Furkan",
		LastName:  "Karakoyunlu",
		Email:     "furkan@karakoyunlu.com",
		Password:  "123456",
		Roles:     "ROLE_ADMIN,ROLE_USER",
	},
	{
		FirstName: "John",
		LastName:  "Doe",
		Email:     "john@doe.com",
		Password:  "p4ssw0rd",
	},
}

var userDetails = []model.UserDetail{
	{
		BirthDate:  time.Now(),
		Gender:     "male",
		Profession: "Computer Engineer",
		Location:   "Istanbul",
		Height:     180,
		Weight:     74,
		Phone:      5375600268,
		Instagram:  "spootrick",
		IsPregnant: false,
	},
	{
		BirthDate:  time.Now(),
		Gender:     "male",
		Profession: "Dummy Prof",
		Location:   "New York",
		Height:     178,
		Weight:     75,
		Phone:      5493857383,
		Instagram:  "johndoe",
		IsPregnant: true,
	},
}
