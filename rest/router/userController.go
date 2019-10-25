package router

import (
	"dtyTrade/rest/models"
	"dtyTrade/rest/response"
	"dtyTrade/rest/service"
	"dtyTrade/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

const tokenExpired = 7*24*60*60

func SignUp(ctx *gin.Context) {
	var body SingUpRequest

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
	err = user.Create()
	if err != nil {
		// TODO log
		response.Res(ctx, response.U_CREATE_USER_ERR, nil)
		return
	}

	response.Res(ctx, response.OK, nil)
}

func SignIn(ctx *gin.Context) {
	fmt.Println("SignIn")
	var body SignInRequest
	if err := ctx.BindJSON(&body); err != nil {
		response.Res(ctx, response.C_PARAMS_ERR, nil)
		return
	}

	hashPW, err := bcrypt.GenerateFromPassword([]byte(body.PassWord), bcrypt.DefaultCost)
	if err != nil {
		response.Res(ctx, response.U_PWD_ERR, nil)
		return
	}
	
	user := models.User{
		Email: body.Email,
		PassWord: string(hashPW),
	}

	user.FindByEmail()

	err = bcrypt.CompareHashAndPassword([]byte(user.PassWord), []byte(body.PassWord))
	if err != nil {
		response.Res(ctx, response.U_PWD_ERR, nil)
		return
	}

	var token string
	token, err = service.RefreshAccessToken(user.ID, user.Email)
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

	user := models.User{}
	user.Model.ID = uint(userId)

	user.FindById()

	response.Res(ctx, response.OK, user)
}
