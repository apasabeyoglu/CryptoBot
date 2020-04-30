package controllers

import (
	"cryptobot/helpers"
	"cryptobot/models"
	"html/template"
	"strconv"

	"github.com/astaxie/beego"
)

type ExchangeAccountController struct {
	beego.Controller
}

func (c *ExchangeAccountController) Prepare() {
	c.Data["XsrfData"] = template.HTML(c.XSRFFormHTML())
}

func (c *ExchangeAccountController) ExchangeAccountForm() {
	c.Data["Title"] = "Hesaplarım"
	c.Data["Exchanges"] = models.GetActiveExchanges()
	c.TplName = "exchangeaccounts.tpl"
}

//Ajax
func (c *ExchangeAccountController) AddExchangeAccount() {
	if c.IsAjax() {
		userID := helpers.GetCurrentUserID(c.Ctx)
		exchangeIDStr := c.GetString("exchangeID")
		name := c.GetString("name")
		apiKey := c.GetString("apiKey")
		apiSecret := c.GetString("apiSecret")

		response := models.AjaxResponse{}

		if exchangeID, err := strconv.Atoi(exchangeIDStr); err != nil {
			response = models.AjaxResponse{
				Result:  -1,
				Message: "Lütfen hesap eklemek isteğiniz borsayı seçin.",
			}
		} else if len(name) > 20 {
			response = models.AjaxResponse{
				Result:  -1,
				Message: "İsim alanı 20 karakterden uzun olamaz.",
			}
		} else if len(apiKey) < 32 || len(apiKey) > 64 {
			response = models.AjaxResponse{
				Result:  -1,
				Message: "Lütfen geçerli bir api key değeri girin.",
			}
		} else if len(apiSecret) < 32 || len(apiSecret) > 64 {
			response = models.AjaxResponse{
				Result:  -1,
				Message: "Lütfen geçerli bir api secret değeri girin.",
			}
		} else {
			exists := models.CheckExchangeIsActiveAndExists(exchangeID)

			if !exists {
				response = models.AjaxResponse{
					Result:  -1,
					Message: "Lütfen hesap eklemek isteğiniz borsayı seçin.",
				}
			} else {
				if !helpers.TestExchangeAccount(exchangeID, apiKey, apiSecret) {
					response = models.AjaxResponse{
						Result:  -1,
						Message: "Lütfen api bilgilerinizi kontrol edin.",
					}
				} else {
					success := models.InsertExchangeAccount(userID, exchangeID, name, apiKey, apiSecret)
					if !success {
						response = models.AjaxResponse{
							Result:  -1,
							Message: "Beklenmeyen bir hata oluştu.",
						}
					} else {
						models.ActiveExchangeAccounts = models.GetActiveExchangeAccounts()
						response = models.AjaxResponse{
							Result:  0,
							Message: "Borsa hesabınız başarıyla eklendi.",
						}
					}
				}
			}
		}

		c.Data["json"] = &response
		c.ServeJSON()
	}
}

//Ajax
func (c *ExchangeAccountController) ExchangeAccountList() {
	if c.IsAjax() {
		userID := helpers.GetCurrentUserID(c.Ctx)
		exchangeAccounts := models.GetUserExchangeAccountsList(userID)
		c.Data["json"] = &exchangeAccounts
		c.ServeJSON()
	}
}

//Ajax
func (c *ExchangeAccountController) ToggleExchangeAccount() {
	if c.IsAjax() {
		userID := helpers.GetCurrentUserID(c.Ctx)
		exchangeAccountIDStr := c.GetString("exchangeAccountID")
		isActive, err := c.GetBool("isActive")

		response := models.AjaxResponse{}

		if exchangeAccountID, err2 := strconv.Atoi(exchangeAccountIDStr); err2 != nil {
			response = models.AjaxResponse{
				Result:  -1,
				Message: "Beklenmeyen bir hata oluştu.",
			}
		} else if err != nil {
			response = models.AjaxResponse{
				Result:  -1,
				Message: "Beklenmeyen bir hata oluştu.",
			}
		} else {
			success := models.ToggleExchangeAccount(userID, exchangeAccountID, isActive)

			if success {
				models.ActiveExchangeAccounts = models.GetActiveExchangeAccounts()
				var message string
				if isActive {
					message = "Hesap başarıyla aktif hale getirildi."
				} else {
					message = "Hesap başarıyla pasif hale getirildi."
				}

				response = models.AjaxResponse{
					Result:  0,
					Message: message,
				}
			}
		}

		c.Data["json"] = &response
		c.ServeJSON()
	}
}

//Ajax
func (c *ExchangeAccountController) GetExchangeAccount() {
	if c.IsAjax() {
		userID := helpers.GetCurrentUserID(c.Ctx)
		exchangeAccountIDStr := c.GetString("exchangeAccountID")

		response := models.AjaxResponse{}

		if exchangeAccountID, err := strconv.Atoi(exchangeAccountIDStr); err != nil {
			response = models.AjaxResponse{
				Result:  -1,
				Message: "Beklenmeyen bir hata oluştu.",
			}
			c.Data["json"] = &response
		} else {
			exchangeAccount := models.GetUserExchangeAccount(userID, exchangeAccountID)
			c.Data["json"] = &exchangeAccount
		}

		c.ServeJSON()

	}
}

//Ajax
func (c *ExchangeAccountController) EditExchangeAccount() {
	if c.IsAjax() {
		userID := helpers.GetCurrentUserID(c.Ctx)
		exchangeAccountIDStr := c.GetString("exchangeAccountID")
		name := c.GetString("name")
		apiKey := c.GetString("apiKey")
		apiSecret := c.GetString("apiSecret")

		response := models.AjaxResponse{}

		if exchangeAccountID, err := strconv.Atoi(exchangeAccountIDStr); err != nil {
			response = models.AjaxResponse{
				Result:  -1,
				Message: "Beklenmeyen bir hata oluştu.",
			}
		} else if len(name) > 20 {
			response = models.AjaxResponse{
				Result:  -1,
				Message: "İsim alanı 20 karakterden uzun olamaz.",
			}
		} else if len(apiKey) < 32 || len(apiKey) > 64 {
			response = models.AjaxResponse{
				Result:  -1,
				Message: "Lütfen geçerli bir api key değeri girin.",
			}
		} else if len(apiSecret) < 32 || len(apiSecret) > 64 {
			response = models.AjaxResponse{
				Result:  -1,
				Message: "Lütfen geçerli bir api secret değeri girin.",
			}
		} else {
			success := models.UpdateExchangeAccount(userID, exchangeAccountID, name, apiKey, apiSecret)

			if !success {
				response = models.AjaxResponse{
					Result:  -1,
					Message: "Beklenmeyen bir hata oluştu.",
				}
			} else {
				models.ActiveExchangeAccounts = models.GetActiveExchangeAccounts()
				response = models.AjaxResponse{
					Result:  0,
					Message: "Borsa hesabınız başarıyla güncellendi.",
				}
			}
		}

		c.Data["json"] = &response
		c.ServeJSON()

	}
}

//Ajax
func (c *ExchangeAccountController) DeleteExchangeAccount() {
	if c.IsAjax() {
		userID := helpers.GetCurrentUserID(c.Ctx)
		exchangeAccountIDStr := c.GetString("exchangeAccountID")

		response := models.AjaxResponse{}

		if exchangeAccountID, err := strconv.Atoi(exchangeAccountIDStr); err != nil {
			response = models.AjaxResponse{
				Result:  -1,
				Message: "Beklenmeyen bir hata oluştu.",
			}
		} else {
			success := models.DeleteExchangeAccount(userID, exchangeAccountID)

			if success {
				models.ActiveExchangeAccounts = models.GetActiveExchangeAccounts()
				response = models.AjaxResponse{
					Result:  0,
					Message: "Borsa hesabınız başarıyla silindi.",
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
