package models

import (
	"fmt"
	"time"

	"github.com/astaxie/beego/orm"
)

type ExchangeAccount struct {
	ID        int       `orm:"column(exchange_account_id);pk;auto"`
	User      *User     `orm:"rel(fk)"`
	Exchange  *Exchange `orm:"rel(fk)"`
	Name      string    `orm:"size(20);null"`
	Key       string    `orm:"size(255)"`
	Secret    string    `orm:"size(255)"`
	CreatedAt time.Time `orm:"auto_now_add;type(datetime)"`
	UpdatedAt time.Time `orm:"auto_now;type(datetime)"`
	IsActive  bool      `orm:"default(true)"`
}

type ExchangeAccountAjaxResponse struct {
	ID           int    `orm:"column(exchange_account_id)"`
	Name         string `orm:"column(name)"`
	Key          string `orm:"column(key)"`
	Secret       string `orm:"column(secret)"`
	CreatedAt    string `orm:"column(created_at)"`
	IsActive     bool   `orm:"column(is_active)"`
	ExchangeName string `orm:"column(exchange_name)"`
}

func (ea *ExchangeAccount) TableName() string {
	return "exchange_accounts"
}

func init() {
	orm.RegisterModel(new(ExchangeAccount))
}
func GetUserExchangeAccounts(id int) []ExchangeAccount {
	o := orm.NewOrm()
	exchangeAccounts := []ExchangeAccount{}
	o.QueryTable("exchange_accounts").RelatedSel().Filter("user_id", id).All(&exchangeAccounts)
	return exchangeAccounts
}

func GetUserActiveExchangeAccounts(id int) []ExchangeAccount {
	o := orm.NewOrm()
	exchangeAccounts := []ExchangeAccount{}
	o.QueryTable("exchange_accounts").RelatedSel().Filter("user_id", id).Filter("is_active", true).All(&exchangeAccounts)
	return exchangeAccounts
}

func GetActiveExchangeAccounts() []ExchangeAccount {
	o := orm.NewOrm()
	exchangeAccounts := []ExchangeAccount{}
	o.QueryTable("exchange_accounts").RelatedSel().Filter("is_active", true).All(&exchangeAccounts)
	return exchangeAccounts
}

func GetUserExchangeAccount(userID, exchangeAccountID int) ExchangeAccountAjaxResponse {
	o := orm.NewOrm()

	exchangeAccountAjaxResponse := ExchangeAccountAjaxResponse{}

	if o.Driver().Type() == orm.DRSqlite {
		o.Raw(`SELECT   exchange_accounts.exchange_account_id,
					exchange_accounts.name,
					exchange_accounts.key,
					exchange_accounts.secret,
					strftime('%d.%m.%Y %H:%M', datetime(exchange_accounts.created_at, '+3 Hour')) as created_at,  
					exchange_accounts.is_active, 
					exchanges.name as exchange_name
			  		FROM exchange_accounts 
			  		JOIN exchanges ON exchanges.exchange_id = exchange_accounts.exchange_id
					WHERE exchange_accounts.user_id = ? AND exchange_accounts.exchange_account_id = ?`, userID, exchangeAccountID).QueryRow(&exchangeAccountAjaxResponse)
	} else {
		fmt.Println("strftime not supported")
	}

	return exchangeAccountAjaxResponse
}

func GetUserExchangeAccountsList(id int) []ExchangeAccountAjaxResponse {
	o := orm.NewOrm()

	exchangeAccountAjaxResponse := []ExchangeAccountAjaxResponse{}

	if o.Driver().Type() == orm.DRSqlite {
		o.Raw(`SELECT   exchange_accounts.exchange_account_id,
					exchange_accounts.name,
					exchange_accounts.key,
					strftime('%d.%m.%Y %H:%M', datetime(exchange_accounts.created_at, '+3 Hour')) as created_at,  
					exchange_accounts.is_active, 
					exchanges.name as exchange_name
			  		FROM exchange_accounts 
			  		JOIN exchanges ON exchanges.exchange_id = exchange_accounts.exchange_id
					WHERE exchange_accounts.user_id = ?`, id).QueryRows(&exchangeAccountAjaxResponse)
	} else {
		fmt.Println("strftime not supported")
	}

	return exchangeAccountAjaxResponse
}

func InsertExchangeAccount(userID, exchangeID int, name, key, secret string) bool {
	o := orm.NewOrm()
	user := User{ID: userID}
	exchange := Exchange{ID: exchangeID}
	exchangeAccount := ExchangeAccount{User: &user, Exchange: &exchange, Name: name, Key: key, Secret: secret, IsActive: true}
	_, err := o.Insert(&exchangeAccount)
	if err != nil {
		return false
	} else {
		return true
	}
}

func UpdateExchangeAccount(userID, exchangeAccountID int, name, key, secret string) bool {
	o := orm.NewOrm()
	_, err := o.Raw("UPDATE exchange_accounts SET name = ?, key = ?, secret = ? WHERE exchange_account_id = ? AND user_id= ?", name, key, secret, exchangeAccountID, userID).Exec()

	if err == nil {
		return true
	} else {
		return false
	}
}

func ToggleExchangeAccount(userID, exchangeAccountID int, isActive bool) bool {
	o := orm.NewOrm()
	_, err := o.Raw("UPDATE exchange_accounts SET is_active = ? WHERE exchange_account_id = ? AND user_id = ?", isActive, exchangeAccountID, userID).Exec()

	if err == nil {
		return true
	} else {
		return false
	}
}

func DeleteExchangeAccount(userID, exchangeAccountID int) bool {
	o := orm.NewOrm()
	_, err := o.Raw("DELETE FROM exchange_accounts WHERE exchange_account_id = ? AND user_id = ?", exchangeAccountID, userID).Exec()

	if err == nil {
		return true
	} else {
		return false
	}
}
