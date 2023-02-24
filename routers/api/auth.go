package api

import (
	"fmt"
	"net/http"

	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"

	"github.com/EDDYCJY/go-gin-example/pkg/app"
	"github.com/EDDYCJY/go-gin-example/pkg/e"
	"github.com/EDDYCJY/go-gin-example/pkg/logging"
	"github.com/EDDYCJY/go-gin-example/pkg/util"
	"github.com/EDDYCJY/go-gin-example/service/auth_service"
)

type auth struct {
	Username string `valid:"Required; MaxSize(50)"`
	Password string `valid:"Required; MaxSize(50)"`
}

// @Summary Get Auth
// @Produce  json
// @Param username query string true "userName"
// @Param password query string true "password"
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /auth [get]
func GetAuth(c *gin.Context) {
	appG := app.Gin{C: c}            //初始化请求类
	valid := validation.Validation{} //初始化验证类

	username := appG.C.PostForm("username")
	password := appG.C.PostForm("password")

	a := auth{Username: username, Password: password} //参数校验赋值
	ok, _ := valid.Valid(&a)                          //参数校验

	if !ok {
		app.MarkErrors(valid.Errors)
		appG.Response(http.StatusBadRequest, e.INVALID_PARAMS, nil)
		return
	}

	authService := auth_service.Auth{Username: username, Password: password} //参数赋值
	isExist, err := authService.Check()                                      //与数据库校对
	logging.Info(a)
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_AUTH_CHECK_TOKEN_FAIL, nil)
		return
	}

	if !isExist {
		appG.Response(http.StatusUnauthorized, e.ERROR_AUTH, nil)
		return
	}

	token, err := util.GenerateToken(username, password)
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_AUTH_TOKEN, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, map[string]string{
		"token": token,
	})
}

func Regist(c *gin.Context) {
	appG := app.Gin{C: c}            //初始化请求类
	valid := validation.Validation{} //初始化验证类

	username := appG.C.PostForm("username")
	password := appG.C.PostForm("password")
	fmt.Printf(username)
	a := auth{Username: username, Password: password} //参数校验赋值
	ok, _ := valid.Valid(&a)                          //参数校验

	if !ok {
		app.MarkErrors(valid.Errors)
		appG.Response(http.StatusBadRequest, e.INVALID_PARAMS, nil)
		return
	}

	authService := auth_service.Auth{Username: username, Password: password} //参数赋值
	isExist := authService.CheckUserName()                                   //与数据库校对

	if isExist {
		appG.Response(http.StatusUnauthorized, e.ERROR_AUTH_CHECK_USER_EXIST, nil)
		return
	}

	id, err_info := authService.AddUser()

	if err_info != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_AUTH_ADD_USER_ERROR, nil)
		return
	}

	if id {
		appG.Response(http.StatusOK, e.SUCCESS, nil)
		return
	} else {
		appG.Response(http.StatusUnauthorized, e.ERROR_AUTH_ADD_USER_ERROR, nil)
		return
	}
}
