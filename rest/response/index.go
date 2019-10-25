package response

type Response struct {
	Head map[string]string `json:"head"`
	Data interface{}       `json:"body"`
}

var Msg map[string]string

func BuildMsg(code string, data interface{}) Response {
	return Response{Head: map[string]string{"code": code, "msg": Msg[code]}, Data: data}
}

func init() {
	Msg = map[string]string{
		OK:                "success",
		C_PARAMS_ERR:      "params error",
		C_TOKEN_NOT_FOUND: "token not found",
		U_CREATE_USER_ERR: "create user failed",
	}
}
