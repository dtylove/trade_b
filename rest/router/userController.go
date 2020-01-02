package router

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"trade_b/rest/models"
	"trade_b/rest/response"
	"trade_b/rest/service"
	"trade_b/utils"
)

const tokenExpired = 7 * 24 * 60 * 60

type signUpRequest struct {
	Phone    string
	Email    string
	PassWord string
}
func SignUp(ctx *gin.Context) {
	var body signUpRequest

	if err := ctx.BindJSON(&body); err != nil {
		// TODO log
		response.Res(ctx, response.C_PARAMS_ERR, nil)
		return
	}

	user := &models.User{
		Email:       body.Email,
		PassWord:    body.PassWord,
		MobilePhone: body.Phone,
	}

	hashPW, err := bcrypt.GenerateFromPassword([]byte(user.PassWord), bcrypt.DefaultCost)
	if err != nil {
		// TODO log
		response.Res(ctx, response.C_PARAMS_ERR, nil)
		return
	}

	user.PassWord = string(hashPW)

	if err = models.Add(user); err != nil {
		// TODO log
		response.Res(ctx, response.U_CREATE_USER_ERR, err)
		return
	}

	response.Res(ctx, response.OK, nil)
}

type signInRequest struct {
	Email    string
	PassWord string
}
func SignIn(ctx *gin.Context) {
	fmt.Println("SignIn")
	var body signInRequest
	if err := ctx.BindJSON(&body); err != nil {
		response.Res(ctx, response.C_PARAMS_ERR, nil)
		return
	}

	user := models.User{
		Email:    body.Email,
	}

	if err := user.FindByEmail(); err != nil {
		response.Res(ctx, response.U_CREATE_USER_ERR, nil)
		return
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.PassWord), []byte(body.PassWord))
	if err != nil {
		response.Res(ctx, response.U_PWD_ERR, nil)
		return
	}

	var token string
	token, err = service.RefreshAccessToken(user.Id, user.Email)
	if err != nil {
		response.Res(ctx, response.C_TOKEN_NOT_FOUND, nil)
		return
	}

	ctx.SetCookie("accessToken", token, tokenExpired, "/", "*", false, false)
	response.Res(ctx, response.OK, token)
}

// 测试接口
func GetUser(ctx *gin.Context) {
	id := ctx.Param("id")
	var userId uint
	if err := utils.StrToUint(id, &userId); err != nil {
		response.Res(ctx, response.C_PARAMS_ERR, nil)
		return
	}

	//requester := ctx.MustGet("user").(models.User)

	user := models.User{Email:"59703122@qq.com"}
	if err := models.FindOne(&user); err != nil {
		response.Res(ctx, response.C_PARAMS_ERR, nil)
		return
	}
	//if err := models.FindOneById(&user, userId); err != nil {
	//	response.Res(ctx, response.C_PARAMS_ERR, nil)
	//	return
	//}

	response.Res(ctx, response.OK, user)
}
