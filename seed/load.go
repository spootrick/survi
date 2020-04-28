package seed

import (
	"github.com/spootrick/survi/api/database"
	"github.com/spootrick/survi/api/model"
	"github.com/spootrick/survi/api/util/console"
	"log"
)

func Load() {
	db, err := database.Connect()
	if err != nil {
		log.Fatal("error in connecting to database in seed.Load():", err)
	}
	defer db.Close()

	err = db.Debug().DropTableIfExists(&model.User{}, &model.UserDetail{}).Error
	if err != nil {
		log.Fatal("error dropping table in seed.Load():", err)
	}

	err = db.Debug().AutoMigrate(&model.User{}, &model.UserDetail{}).Error
	if err != nil {
		log.Fatal("error in auto migrate in seed.Load():", err)
	}

	err = db.Debug().Model(&model.UserDetail{}).AddForeignKey("user_id", "users(id)", "cascade", "cascade").Error
	if err != nil {
		log.Fatal(err)
	}

	for i, _ := range users {
		err = db.Debug().Model(&model.User{}).Create(&users[i]).Error
		if err != nil {
			log.Fatal("error creating dummy user in seed.Load():", err)
		}

		userDetails[i].UserID = users[i].ID

		err = db.Debug().Model(&model.UserDetail{}).Create(&userDetails[i]).Error
		if err != nil {
			log.Fatal("error creating dummy user in seed.Load():", err)
		}

		err = db.Debug().Model(userDetails[i]).Related(&userDetails[i].User).Error
		if err != nil {
			log.Fatal("error creating dummy user in seed.Load():", err)
		}

		console.Pretty(userDetails[i])
	}
}
