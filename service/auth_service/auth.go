package auth_service

import (
	"github.com/EDDYCJY/go-gin-example/models"
	"github.com/EDDYCJY/go-gin-example/pkg/util"
)

type Auth struct {
	Username string
	Password string
}

func (a *Auth) Check() (bool, error) {
	return models.CheckAuth(a.Username, changePwd(a.Password))
}

func changePwd(pwd string) string {
	return util.EncodeMD5(pwd)
}

func (a *Auth) CheckUserName() bool {
	return models.CheckUserName(a.Username)
}

func (a *Auth) AddUser() (bool, error) {
	return models.AddUser(a.Username, changePwd(a.Password))
}
