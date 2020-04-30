package models

import (
	"time"

	"github.com/astaxie/beego/orm"
)

type ForgotPassword struct {
	ID           int       `orm:"column(forgot_password_id);pk;auto"`
	User         *User     `orm:"rel(fk)"`
	RandomString string    `orm:"size(100);unique"`
	CreatedAt    time.Time `orm:"auto_now_add;type(datetime)"`
	UpdatedAt    time.Time `orm:"auto_now;type(datetime)"`
}

func (f *ForgotPassword) TableName() string {
	return "forgot_password"
}

func init() {
	orm.RegisterModel(new(ForgotPassword))
}

func InsertUpdateForgotPassword(id int, randomString string) bool {
	o := orm.NewOrm()
	user := User{ID: id}

	if o.Driver().Type() != orm.DRSqlite {
		forgotPassword := ForgotPassword{User: &user, RandomString: randomString}
		_, err := o.InsertOrUpdate(&forgotPassword, "random_string")

		if err == nil {
			return true
		} else {
			return false
		}
	} else {
		forgotPassword := ForgotPassword{User: &user}
		err := o.Read(&forgotPassword, "user_id")

		forgotPasswordID := forgotPassword.ID

		if err == orm.ErrNoRows {
			forgotPassword = ForgotPassword{User: &user, RandomString: randomString}
			_, err := o.Insert(&forgotPassword)
			if err == nil {
				return true
			} else {
				return false
			}

		} else if err != nil {
			return false
		} else {
			forgotPassword = ForgotPassword{ID: forgotPasswordID, User: &user, RandomString: randomString}
			_, err := o.Update(&forgotPassword, "random_string")

			if err == nil {
				return true
			} else {
				return false
			}
		}
	}

}

func GetUserByForgotPasswordRandomString(randomString string) int {
	o := orm.NewOrm()
	forgotPassword := ForgotPassword{RandomString: randomString}
	err := o.Read(&forgotPassword, "random_string")
	if err != nil {
		return 0
	} else {
		return forgotPassword.User.ID
	}
}

func DeleteForgotPasswordRandomString(id int) bool {
	o := orm.NewOrm()
	user := User{ID: id}
	forgotPassword := ForgotPassword{User: &user}
	_, err := o.Delete(&forgotPassword, "user_id")

	if err == nil {
		return true
	} else {
		return false
	}
}
