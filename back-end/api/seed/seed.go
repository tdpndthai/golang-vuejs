package seed

import (
	"log"

	"github.com/jinzhu/gorm"
	"github.com/tdpndthai/golang-vuejs/api/models"
)

var users = []models.User{
	models.User{
		Nickname: "thainguyen",
		Email:    "thainguyen@gmail.com",
		Password: "1605",
	},
	models.User{
		Nickname: "danhthai",
		Email:    "danhthai@gmail.com",
		Password: "1605",
	},
}



func Load(db *gorm.DB) {

	err := db.Debug().DropTableIfExists(&models.User{}).Error
	if err != nil {
		log.Fatalf("cannot drop table: %v", err)
	}
	err = db.Debug().AutoMigrate(&models.User{}).Error
	if err != nil {
		log.Fatalf("cannot migrate table: %v", err)
	}

	for i,_ := range users {
		err = db.Debug().Model(&models.User{}).Create(&users[i]).Error
		if err != nil {
			log.Fatalf("cannot seed users table: %v", err)
		}
	}
}