package models

import (
	"time"

	"github.com/astaxie/beego/orm"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID                      int       `orm:"column(user_id);pk;auto"`
	Email                   string    `orm:"size(255);unique"`
	Password                string    `orm:"size(60)"`
	Username                string    `orm:"size(50)"`
	EmailConfirmationString string    `orm:"size(100);unique"`
	CreatedAt               time.Time `orm:"auto_now_add;type(datetime)"`
	UpdatedAt               time.Time `orm:"auto_now;type(datetime)"`
	IsEmailConfirmed        bool      `orm:"default(false)"`
}

func (u *User) TableName() string {
	return "users"
}

func init() {
	orm.RegisterModel(new(User))
}

func InsertUser(email, username, password, emailConfirmationString string) int64 {
	emailExists, usernameExists := CheckUserEmailUsernameExists(email, username)

	if emailExists {
		return -1
	} else if usernameExists {
		return -2
	}

	hashedPasswordBytes, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	hashedPassword := string(hashedPasswordBytes[:])
	o := orm.NewOrm()
	user := User{Email: email, Username: username, Password: hashedPassword, EmailConfirmationString: emailConfirmationString}
	id, err := o.Insert(&user)
	if err == nil {
		return id
	}

	return 0
}

func CheckUserEmailUsernameExists(email, username string) (bool, bool) {
	o := orm.NewOrm()
	user := User{Email: email, Username: username}

	cond := orm.NewCondition()
	cond1 := cond.And("email", email).Or("username", username)

	err := o.QueryTable("users").SetCond(cond1).One(&user)
	if err == orm.ErrNoRows {
		return false, false
	}

	if user.Email == email {
		return true, false
	} else {
		return false, true
	}

}

func GetUserByEmailConfirmationString(emailConfirmationString string) int {
	o := orm.NewOrm()
	user := User{EmailConfirmationString: emailConfirmationString}
	err := o.Read(&user, "email_confirmation_string")

	if err == orm.ErrNoRows {
		return 0
	}

	return user.ID
}

func ConfirmUserEmail(id int) {
	o := orm.NewOrm()
	user := User{ID: id, IsEmailConfirmed: true}
	o.Update(&user, "is_email_confirmed")
}

func GetUserByEmailPassword(email, password string) int {
	o := orm.NewOrm()
	user := User{}
	err := o.QueryTable("users").Filter("email", email).One(&user)

	if err == orm.ErrNoRows {
		return -1
	} else {
		passwordDB := user.Password
		err = bcrypt.CompareHashAndPassword([]byte(passwordDB), []byte(password))
		if err == nil {
			isEmailConfirmed := user.IsEmailConfirmed

			if isEmailConfirmed {
				return user.ID
			}

			return -2

		}

		return -1
	}
}

func GetUserByEmail(email string) (bool, int) {
	o := orm.NewOrm()
	user := User{}
	err := o.QueryTable("users").Filter("email", email).One(&user)

	if err == orm.ErrNoRows {
		return false, 0
	} else {
		return true, user.ID
	}
}

func UpdateUserPassword(id int, password string) bool {
	o := orm.NewOrm()
	hashedPasswordBytes, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	hashedPassword := string(hashedPasswordBytes[:])

	user := User{ID: id, Password: hashedPassword}
	_, err := o.Update(&user, "password")

	if err == nil {
		return true
	} else {
		return false
	}
}
