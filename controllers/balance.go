package controllers

import (
	"cryptobot/helpers"
	"cryptobot/models"
	"html/template"
	"strconv"

	"github.com/astaxie/beego"
)

type BalanceController struct {
	beego.Controller
}

func (c *BalanceController) Prepare() {
	c.Data["XsrfData"] = template.HTML(c.XSRFFormHTML())
}

func (c *BalanceController) Balances() {
	c.Data["Title"] = "Bakiyelerim"
	c.TplName = "balances.tpl"
}

func (c *BalanceController) Graphic() {
	c.Data["Title"] = "Grafik"
	c.TplName = "graphic.tpl"
}

//Ajax
func (c *BalanceController) AllBalancesList() {
	if c.IsAjax() {
		userID := helpers.GetCurrentUserID(c.Ctx)
		allBalances := models.GetUserBalances(userID)
		c.Data["json"] = &allBalances
		c.ServeJSON()
	}
}

//Ajax
func (c *BalanceController) GraphicData() {
	if c.IsAjax() {
		userID := helpers.GetCurrentUserID(c.Ctx)
		exchangeAccountIDStr := c.GetString("exchangeAccountID")

		if exchangeAccountID, err := strconv.Atoi(exchangeAccountIDStr); err == nil {
			balanceGraphicData := models.GetUserBalanceGraphicData(userID, exchangeAccountID)
			c.Data["json"] = &balanceGraphicData
			c.ServeJSON()
		}
	}
}
