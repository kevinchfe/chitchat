package routes

import (
	"chitchat/handlers"
	"net/http"
)

// 定义一个 WebRoute 结构体用于存放单个路由
type WebRoute struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

// 声明 WebRoutes 切片存放所有 Web 路由
type WebRoutes []WebRoute

// 定义所有 Web 路由
var webRoutes = WebRoutes{
	WebRoute{
		"home",
		"GET",
		"/",
		handlers.Index,
	},
	WebRoute{
		"signup",
		"GET",
		"/signup",
		handlers.Signup,
	},
	WebRoute{
		"signupAccount",
		"POST",
		"/signup_account",
		handlers.SignupAccount,
	},
	WebRoute{
		"login",
		"GET",
		"/login",
		handlers.Login,
	},
	WebRoute{
		"auth",
		"POST",
		"/authenticate",
		handlers.Authenticate,
	},
	WebRoute{
		"logout",
		"GET",
		"/logout",
		handlers.Logout,
	},
}
