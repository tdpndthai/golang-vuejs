package models

import (
	"errors"
	"io"
	"net/http"
	"os"

	"github.com/jinzhu/gorm"
)

type Photo struct {
	ID  uint32 `gorm:"primary_key;auto_increment" json:"id"`
	Src string `gorm:"size:45;not null;" json:"src"`
}

// type PhotoCollection struct {
// 	Photos []Photo `json:"items"`
// }

func (photo *Photo) SavePhoto(db *gorm.DB) (*Photo, error) {
	var err error
	err = db.Debug().Create(&photo).Error
	if err != nil {
		return &Photo{}, err
	}
	return photo, nil
}

func FileUpload(r *http.Request) (string, error) {
	//this function returns the filename(to save in database) of the saved file or an error if it occurs
	r.ParseMultipartForm(32 << 20)
	//ParseMultipartForm parses a request body as multipart/form-data
	file, handler, err := r.FormFile("file") //retrieve the file from form data
	//replace file with the key your sent your image with
	if err != nil {
		return "", err
	}
	defer file.Close() //close the file when we finish
	//this is path which  we want to store the file
	f, err := os.OpenFile("../../uploads/"+handler.Filename, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		return "", err
	}
	defer f.Close()
	io.Copy(f, file)
	//here we save our file to our path
	return handler.Filename, nil
}

func (u *Photo) DeletePhoto(db *gorm.DB, uid uint32) (int64, error) {

	db = db.Debug().Model(&Photo{}).Where("id = ?", uid).Take(&Photo{}).Delete(&Photo{})

	if db.Error != nil {
		return 0, db.Error
	}
	return db.RowsAffected, nil
}

func (p *Photo) FindPhotoById(db *gorm.DB, uid uint32) (*Photo, error) {
	var err error
	err = db.Debug().Model(Photo{}).Where("id = ?", uid).Take(&p).Error
	if err != nil {
		return &Photo{}, err
	}
	if gorm.IsRecordNotFoundError(err) {
		return &Photo{}, errors.New("Photo Not Found")
	}
	return p, err
}
