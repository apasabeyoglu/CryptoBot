package controllers

import "github.com/astaxie/beego"

type ErrorController struct {
	beego.Controller
}

func (c *ErrorController) Prepare() {
	c.TplName = "error.tpl"
}

func (c *ErrorController) Error403() {
	c.Data["Message"] = "Bu sayfayı görüntüleme yetkiniz yok"
}

func (c *ErrorController) Error404() {
	c.Data["Message"] = "Sayfa bulunamadı"
}

func (c *ErrorController) Error500() {
	c.Data["Message"] = "Beklenmeyen bir hata oluştu"
}

func (c *ErrorController) ErrorDb() {
	c.Data["Message"] = "Veritabanında beklenmeyen bir hata oluştu"
}
