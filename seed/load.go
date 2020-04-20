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

	err = db.Debug().DropTableIfExists(&model.User{}).Error
	if err != nil {
		log.Fatal("error dropping table in seed.Load():", err)
	}

	err = db.Debug().AutoMigrate(&model.User{}).Error
	if err != nil {
		log.Fatal("error in auto migrate in seed.Load():", err)
	}

	for _, user := range users {
		err = db.Debug().Model(&model.User{}).Create(&user).Error
		if err != nil {
			log.Fatal("error creating dummy user in seed.Load():", err)
		}

		console.Pretty(user)
	}
}
