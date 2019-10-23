package router

import (
	"dtyTrade/models"
	"dtyTrade/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

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
	var body SignInRequest
	if err := ctx.BindJSON(body); err != nil {
		ctx.JSON(http.StatusBadRequest, "参数不正确")
		return
	}

	//hashPW, err := bcrypt.GenerateFromPassword([]byte(body.PassWord), bcrypt.DefaultCost)
	//if err != nil {
	//	ctx.JSON(http.StatusBadRequest, "参数不正确")
	//	return
	//}
	
	//user := models.User{
	//	Email: body.Email,
	//	PassWord: string(hashPW),
	//}

	//user.FindById()

	//secret := utils.Sha256(body.Email)

}

func GetUser(ctx *gin.Context) {
	fmt.Println("call getUser")
	id := ctx.Param("id")
	var userId uint
	if err := utils.StrToUint(id, &userId); err != nil {
		ctx.JSON(http.StatusBadRequest, "参数不正确")
		return
	}

	user := models.User{}
	user.Model.ID = uint(userId)

	user.FindById()

	ctx.JSON(http.StatusOK, nil)
}
