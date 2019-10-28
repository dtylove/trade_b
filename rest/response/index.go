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
		U_EMAIL_NOT_FOUND: "email not found",

		O_ADD_ERR: "create order failed",

		P_NOT_FOUND:         "product not found",
		P_CREATE_FAILED:     "create product failed",
		P_BASE_COUNTER_SAME: "base counter can not same",
	}
}
