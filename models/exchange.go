package models

import (
	"time"

	"github.com/astaxie/beego/orm"
)

type Exchange struct {
	ID            int       `orm:"column(exchange_id);pk;auto"`
	Name          string    `orm:"size(255);unique"`
	LogoImageName string    `orm:"size(255);null"`
	CreatedAt     time.Time `orm:"auto_now_add;type(datetime)"`
	UpdatedAt     time.Time `orm:"auto_now;type(datetime)"`
	IsActive      bool      `orm:"default(true)"`
}

func (e *Exchange) TableName() string {
	return "exchanges"
}

func init() {
	orm.RegisterModel(new(Exchange))
}

func GetActiveExchanges() []Exchange {
	o := orm.NewOrm()
	exchanges := []Exchange{}
	o.QueryTable("exchanges").Filter("is_active", true).All(&exchanges)
	return exchanges
}

func CheckExchangeIsActiveAndExists(id int) bool {
	o := orm.NewOrm()
	exchange := Exchange{ID: id}
	err := o.QueryTable("exchanges").Filter("is_active", true).One(&exchange)

	if err == orm.ErrNoRows || err != nil {
		return false
	} else {
		return true
	}
}
