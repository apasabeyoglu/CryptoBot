package controllers

import (
	"cryptobot/helpers"
	"cryptobot/models"
	"html/template"

	"github.com/astaxie/beego"
	recaptcha "github.com/dpapathanasiou/go-recaptcha"
)

type UserController struct {
	beego.Controller
}

func (c *UserController) Prepare() {
	c.Data["XsrfData"] = template.HTML(c.XSRFFormHTML())
}

func (c *UserController) LoginForm() {
	beego.ReadFromRequest(&c.Controller)
	if helpers.CheckLogin(c.Ctx) {
		c.Ctx.Redirect(302, "/")
		c.StopRun()
	}

	recaptchaSiteKey := beego.AppConfig.String("sitekey")
	recaptchaSecretKey := beego.AppConfig.String("secretkey")

	recaptcha.Init(recaptchaSecretKey)

	c.Data["Title"] = "Kullanıcı Girişi"
	c.Data["SiteKey"] = recaptchaSiteKey
	c.TplName = "login.tpl"
}

func (c *UserController) RegisterForm() {

	if helpers.CheckLogin(c.Ctx) {
		c.Ctx.Redirect(302, "/")
		c.StopRun()
	}

	c.Data["Title"] = "Yeni Kullanıcı Kaydı"
	c.TplName = "register.tpl"
}

func (c *UserController) ForgotPasswordForm() {
	if helpers.CheckLogin(c.Ctx) {
		c.Ctx.Redirect(302, "/")
		c.StopRun()
	}

	c.Data["Title"] = "Yeni Şifre Talebi"
	c.TplName = "forgotpassword.tpl"
}

func (c *UserController) NewPasswordForm() {
	randomString := c.Ctx.Input.Param(":rs")

	if len(randomString) != 100 || helpers.CheckLogin(c.Ctx) {
		c.Ctx.Redirect(302, "/")
		c.StopRun()
	}

	c.Data["Title"] = "Yeni Şifre Oluştur"
	c.Data["RandomString"] = randomString
	c.TplName = "newpassword.tpl"
}

func (c *UserController) Confirm() {
	emailConfirmationString := c.Ctx.Input.Param(":cs")

	if len(emailConfirmationString) != 100 || helpers.CheckLogin(c.Ctx) {
		c.Ctx.Redirect(302, "/")
		c.StopRun()
	}

	userID := models.GetUserByEmailConfirmationString(emailConfirmationString)

	if userID == 0 {
		c.Ctx.Redirect(302, "/")
		c.StopRun()
	} else {
		models.ConfirmUserEmail(userID)
		flash := beego.NewFlash()
		flash.Success("Email adresiniz başarıyla onaylandı. Giriş yapabilirsiniz.")
		flash.Store(&c.Controller)
		c.Ctx.Redirect(302, "/kullanici/giris")
	}
}

//Ajax
func (c *UserController) Login() {
	if c.IsAjax() {
		email := c.GetString("email")
		password := c.GetString("password")
		recaptchaReponse := c.GetString("captcha")

		response := models.AjaxResponse{}

		if !helpers.IsValidEmail(email) {
			response = models.AjaxResponse{
				Result:  -1,
				Message: "Lütfen geçerli bir email adresi girin.",
			}
		} else if len(password) < 6 || len(password) > 20 {
			response = models.AjaxResponse{
				Result:  -1,
				Message: "Şifre 6-20 karakter aralığında olmalıdır.",
			}
		} else if len(recaptchaReponse) == 0 {
			response = models.AjaxResponse{
				Result:  -1,
				Message: "Hatalı captcha",
			}
		} else {
			result, err := recaptcha.Confirm(c.Ctx.Input.IP(), recaptchaReponse)
			if err != nil || !result {
				response = models.AjaxResponse{
					Result:  -1,
					Message: "Hatalı captcha",
				}
			} else {
				userID := models.GetUserByEmailPassword(email, password)

				if userID > 0 {
					c.SetSession("uid", userID)
					response = models.AjaxResponse{
						Result:  0,
						Message: "Giriş başarılı.",
					}
				} else if userID == -1 {
					response = models.AjaxResponse{
						Result:  -1,
						Message: "Lütfen email adresinizi ve şifrenizi kontrol edin.",
					}
				} else if userID == -2 {
					response = models.AjaxResponse{
						Result:  -1,
						Message: "Giriş yapabilmek için email adresinizi onaylamanız gerekmektedir. Lütfen email adresinize göndermiş olduğumuz linke tıklayarak üyeliğinizi onaylayın.",
					}
				}
			}
		}

		c.Data["json"] = &response
		c.ServeJSON()
	}
}

//Ajax
func (c *UserController) Register() {
	if c.IsAjax() {
		email := c.GetString("email")
		username := c.GetString("username")
		password := c.GetString("password")
		emailConfirmationString := helpers.GenerateRandomString(100)

		response := models.AjaxResponse{}

		if !helpers.IsValidEmail(email) {
			response = models.AjaxResponse{
				Result:  -1,
				Message: "Lütfen geçerli bir email adresi girin.",
			}
		} else if len(username) < 2 || len(username) > 50 {
			response = models.AjaxResponse{
				Result:  -1,
				Message: "Kullanıcı adı 2-50 karakter aralığında olmalıdır.",
			}
		} else if len(password) < 6 || len(password) > 20 {
			response = models.AjaxResponse{
				Result:  -1,
				Message: "Şifre 6-20 karakter aralığında olmalıdır.",
			}
		} else {

			result := models.InsertUser(email, username, password, emailConfirmationString)

			if result > 0 {
				subject := "Üyeliğiniz Hakkında"
				body := "CryptoBot üyeliğinizi aktifleştirmek için lütfen aşağıdaki linke tıklayın.<br><br>"
				body += "<a href='http://localhost:8080/kullanici/onay/" + emailConfirmationString + "'>Aktifleştir</a>"
				helpers.SendMail([]string{email}, subject, body)

				response = models.AjaxResponse{
					Result:  0,
					Message: "Başarıyla kayıt olundu. Giriş yapabilmek için lütfen email adresinize göndermiş olduğumuz linke tıklayarak üyeliğinizi onaylayın.",
				}
			} else if result == 0 {
				response = models.AjaxResponse{
					Result:  -1,
					Message: "Beklenmeyen bir hata oluştu.",
				}
			} else if result == -1 {
				response = models.AjaxResponse{
					Result:  -1,
					Message: "Bu email adresi kullanımda.",
				}
			} else if result == -2 {
				response = models.AjaxResponse{
					Result:  -1,
					Message: "Bu nickname kullanımda.",
				}
			}
		}

		c.Data["json"] = &response
		c.ServeJSON()
	}
}

//Ajax
func (c *UserController) ForgotPassword() {
	if c.IsAjax() {
		email := c.GetString("email")

		response := models.AjaxResponse{}

		if !helpers.IsValidEmail(email) {
			response = models.AjaxResponse{
				Result:  -1,
				Message: "Lütfen geçerli bir email adresi girin.",
			}
		} else {
			userExists, userID := models.GetUserByEmail(email)

			if userExists {
				randomString := helpers.GenerateRandomString(100)
				isSuccess := models.InsertUpdateForgotPassword(userID, randomString)

				if !isSuccess {
					response = models.AjaxResponse{
						Result:  -1,
						Message: "Beklenmeyen bir hata oluştu.",
					}
				} else {
					subject := "Yeni Şifre Talebiniz Hakkında"
					body := "CryptoBot üyeliğinizin şifresini değiştirmek için lütfen aşağıdaki linke tıklayın.<br><br>"
					body += "<a href='http://localhost:8080/kullanici/yeni-sifre/" + randomString + "'>Değiştir</a>"
					isSuccess = helpers.SendMail([]string{email}, subject, body)

					if isSuccess {
						response = models.AjaxResponse{
							Result:  0,
							Message: "Şifrenizi değiştirmeniz için gereken bağlantı email adresinize gönderildi.",
						}
					} else {
						response = models.AjaxResponse{
							Result:  -1,
							Message: "Beklenmeyen bir hata oluştu.",
						}
					}

				}
			} else {
				response = models.AjaxResponse{
					Result:  0,
					Message: "Şifrenizi değiştirmeniz için gereken bağlantı email adresinize gönderildi.",
				}
			}

		}

		c.Data["json"] = &response
		c.ServeJSON()
	}
}

//Ajax
func (c *UserController) NewPassword() {
	if c.IsAjax() {
		password := c.GetString("password")
		randomString := c.GetString("randomString")

		response := models.AjaxResponse{}

		if len(password) < 6 || len(password) > 20 {
			response = models.AjaxResponse{
				Result:  -1,
				Message: "Şifre 6-20 karakter aralığında olmalıdır.",
			}
		} else if len(randomString) != 100 {
			response = models.AjaxResponse{
				Result:  -1,
				Message: "Beklenmeyen bir hata oluştu.",
			}
		} else {
			userID := models.GetUserByForgotPasswordRandomString(randomString)

			if userID > 0 {
				result := models.UpdateUserPassword(userID, password)

				if result {
					models.DeleteForgotPasswordRandomString(userID)
					response = models.AjaxResponse{
						Result:  0,
						Message: "Şifreniz başarıyla değiştirildi.",
					}
				} else {
					response = models.AjaxResponse{
						Result:  -1,
						Message: "Beklenmeyen bir hata oluştu.",
					}
				}
			} else {
				response = models.AjaxResponse{
					Result:  -1,
					Message: "Beklenmeyen bir hata oluştu.",
				}
			}
		}

		c.Data["json"] = &response
		c.ServeJSON()
	}
}

func (c *UserController) Logout() {
	c.DelSession("uid")
	c.DestroySession()
	c.Ctx.Redirect(302, "/")
}
