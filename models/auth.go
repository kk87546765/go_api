package models

import (
	"github.com/EDDYCJY/go-gin-example/pkg/logging"
	"github.com/jinzhu/gorm"
)

type Auth struct {
	ID       int    `gorm:"primary_key" json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

// CheckAuth checks if authentication information exists
func CheckAuth(username, password string) (bool, error) {
	var auth Auth
	logging.Info(Auth{Username: username, Password: password})
	err := db.Debug().Select("id").Where(Auth{Username: username, Password: password}).First(&auth).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return false, err
	}

	if auth.ID > 0 {
		return true, nil
	}

	return false, nil
}

func CheckUserName(username string) bool {
	var auth Auth
	err := db.Select("id").Where(Auth{Username: username}).First(&auth).Error

	if err != nil {
		return false
	}

	if auth.ID > 0 {
		return true
	} else {
		return false
	}
}

func AddUser(username, password string) (bool, error) {
	var auth Auth
	auth.Username = username
	auth.Password = password
	logging.Info(auth)

	err := db.Create(&auth).Error
	logging.Info(err)
	logging.Info(auth)
	if err != nil {
		return false, err
	}

	res, err_info := CheckAuth(username, password)

	if err_info != nil {
		return false, err_info
	}

	return res, nil
}
