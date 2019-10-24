package router

type SingUpRequest struct {
	Phone    string
	Email    string
	PassWord string
}

type SignInRequest struct {
	Email    string
	PassWord string
}