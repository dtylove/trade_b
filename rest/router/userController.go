package router

import (
	"dtyTrade/rest/models"
	"dtyTrade/rest/service"
	"dtyTrade/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

const tokenExpired = 7*24*60*60

func SignUp(ctx *gin.Context) {
	var body SingUpRequest

	if err := ctx.BindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, "参数不正确")
		return
	}

	user := &models.User{
		Email:       body.Email,
		PassWord:    body.PassWord,
		MobilePhone: body.Phone,
	}

	hashPW, err := bcrypt.GenerateFromPassword([]byte(user.PassWord), bcrypt.DefaultCost)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, "参数不正确")
		return
	}

	user.PassWord = string(hashPW)
	err = user.Create()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, "创建失败")
		return
	}

	ctx.JSON(http.StatusOK, nil)
}

func SignIn(ctx *gin.Context) {
	fmt.Println("SignIn")
	var body SignInRequest
	if err := ctx.BindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, "参数不正确1")
		return
	}

	hashPW, err := bcrypt.GenerateFromPassword([]byte(body.PassWord), bcrypt.DefaultCost)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, "参数不正确2")
		return
	}
	
	user := models.User{
		Email: body.Email,
		PassWord: string(hashPW),
	}

	user.FindByEmail()

	err = bcrypt.CompareHashAndPassword([]byte(user.PassWord), []byte(body.PassWord))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, "密码错误")
		return
	}

	var token string
	token, err = service.RefreshAccessToken(user.ID, user.Email)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, "参数不正确2")
		return
	}

	ctx.SetCookie("accessToken", token, tokenExpired, "/", "*", false, false)
	ctx.JSON(http.StatusOK, token)
}

// 测试接口
func GetUser(ctx *gin.Context) {
	id := ctx.Param("id")
	var userId uint
	if err := utils.StrToUint(id, &userId); err != nil {
		ctx.JSON(http.StatusBadRequest, "参数不正确")
		return
	}

	//requester := ctx.MustGet("user").(models.User)

	user := models.User{}
	user.Model.ID = uint(userId)

	user.FindById()

	ctx.JSON(http.StatusOK, nil)
}
