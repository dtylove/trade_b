package router

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func SignUp(ctx *gin.Context) {
	var body SingUpJson

	if err := ctx.BindJSON(&body); err != nil {

		//ctx.JSON(http.StatusBadRequest, "参数不正确")
		panic(err)
		return
	}

	data, _ := json.Marshal(body)
	fmt.Println(string(data))

	//user := &models.User{
	//	Email:       body.Email,
	//	PassWord:    body.PassWord,
	//	MobilePhone: body.Phone,
	//}

	//db := models.GetDB()
	//db.Create(user)

	//if err != nil {
	//	panic(err)
	//	//ctx.JSON(http.StatusBadRequest, "创建失败")
	//	return
	//}

	ctx.JSON(http.StatusOK, nil)
}
