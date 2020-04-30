package routers

import (
	"cryptobot/controllers"
	"cryptobot/middlewares"

	"github.com/astaxie/beego"
)

func init() {

	//User
	beego.Router("/kullanici/giris", &controllers.UserController{}, "get:LoginForm")
	beego.Router("/kullanici/kayit", &controllers.UserController{}, "get:RegisterForm")
	beego.Router("/kullanici/sifremi-unuttum", &controllers.UserController{}, "get:ForgotPasswordForm")
	beego.Router("/kullanici/yeni-sifre/:rs:string", &controllers.UserController{}, "get:NewPasswordForm")
	beego.Router("/kullanici/onay/:cs:string", &controllers.UserController{}, "get:Confirm")
	beego.Router("/kullanici/giris", &controllers.UserController{}, "post:Login")
	beego.Router("/kullanici/kayit", &controllers.UserController{}, "post:Register")
	beego.Router("/kullanici/sifremi-unuttum", &controllers.UserController{}, "post:ForgotPassword")
	beego.Router("/kullanici/yeni-sifre", &controllers.UserController{}, "post:NewPassword")
	beego.Router("/kullanici/cikis", &controllers.UserController{}, "post:Logout")

	//Balances
	beego.Router("/", &controllers.BalanceController{}, "get:Balances")
	beego.Router("/bakiye", &controllers.BalanceController{}, "get:Balances")
	beego.Router("/bakiye/grafik", &controllers.BalanceController{}, "get:Graphic")
	beego.Router("/bakiye/grafik/goruntule", &controllers.BalanceController{}, "get:GraphicData")
	beego.Router("/bakiye/listele", &controllers.BalanceController{}, "get:AllBalancesList")

	//Orders
	beego.Router("/emir", &controllers.OrderController{}, "get:Orders")
	beego.Router("/emir/listele", &controllers.OrderController{}, "get:AllOrdersList")

	//Exchange Account
	beego.Router("/borsa-hesabi", &controllers.ExchangeAccountController{}, "get:ExchangeAccountForm")
	beego.Router("/borsa-hesabi/listele", &controllers.ExchangeAccountController{}, "get:ExchangeAccountList")
	beego.Router("/borsa-hesabi/ekle", &controllers.ExchangeAccountController{}, "post:AddExchangeAccount")
	beego.Router("/borsa-hesabi/aktiflestir", &controllers.ExchangeAccountController{}, "post:ToggleExchangeAccount")
	beego.Router("/borsa-hesabi/duzenle", &controllers.ExchangeAccountController{}, "get:GetExchangeAccount")
	beego.Router("/borsa-hesabi/duzenle", &controllers.ExchangeAccountController{}, "post:EditExchangeAccount")
	beego.Router("/borsa-hesabi/sil", &controllers.ExchangeAccountController{}, "post:DeleteExchangeAccount")

	//Error Handler
	beego.ErrorController(&controllers.ErrorController{})

	//Middlewares
	beego.InsertFilter("/*", beego.BeforeRouter, middlewares.LoginRequired)
}
