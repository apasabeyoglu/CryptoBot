package controllers

import (
	"cryptobot/helpers"
	"cryptobot/models"
	"html/template"

	"github.com/astaxie/beego"
)

type OrderController struct {
	beego.Controller
}

func (c *OrderController) Prepare() {
	c.Data["XsrfData"] = template.HTML(c.XSRFFormHTML())
}

func (c *OrderController) Orders() {
	c.Data["Title"] = "Emirlerim"
	c.TplName = "orders.tpl"
}

//Ajax
func (c *OrderController) AllOrdersList() {
	if c.IsAjax() {
		userID := helpers.GetCurrentUserID(c.Ctx)
		openOrders := models.GetUserOrders(userID)
		c.Data["json"] = &openOrders
		c.ServeJSON()
	}
}
