package helpers

import (
	con "context"
	"math"
	"math/rand"
	"regexp"
	"time"

	"github.com/adshao/go-binance"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"github.com/astaxie/beego/utils"
)

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
const (
	letterIdxBits = 6                    // 6 bits to represent a letter index
	letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
	letterIdxMax  = 63 / letterIdxBits   // # of letter indices fitting in 63 bits
)

//ConvertCheck - Template function. true => "checked", false => ""
func ConvertCheck(b bool) string {
	if b {
		return "checked"
	} else {
		return ""
	}
}

func CheckLogin(ctx *context.Context) bool {
	u := ctx.Input.Session("uid")
	if u != nil {
		return true
	} else {
		return false
	}
}

func GetCurrentUserID(ctx *context.Context) int {
	return ctx.Input.Session("uid").(int)
}

func GenerateRandomString(n int) string {
	src := rand.NewSource(time.Now().UnixNano())
	b := make([]byte, n)
	// A src.Int63() generates 63 random bits, enough for letterIdxMax characters.
	for i, cache, remain := n-1, src.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = src.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(letterBytes) {
			b[i] = letterBytes[idx]
			i--
		}
		cache >>= letterIdxBits
		remain--
	}

	return string(b)
}

func SetDecimals(number float64, places int) float64 {
	return (math.Round(number*math.Pow10(places)) / math.Pow10(places))
}

func SendMail(to []string, subject, body string) bool {
	username := beego.AppConfig.String("emailusername")
	password := beego.AppConfig.String("emailpassword")
	host := beego.AppConfig.String("emailhost")
	port := beego.AppConfig.String("emailport")

	config := `{"username":"` + username + `","password":"` + password + `","host":"` + host + `","port": ` + port + `}`
	mail := utils.NewEMail(config)
	mail.To = to
	mail.From = username
	mail.Subject = subject
	mail.HTML = body
	err := mail.Send()

	if err == nil {
		return true
	} else {
		return false
	}
}

func IsValidEmail(email string) bool {
	re := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
	return re.MatchString(email)
}

func TestExchangeAccount(exchangeID int, key, secret string) bool {
	if exchangeID == 1 { //Binance
		client := binance.NewClient(key, secret)
		_, err := client.NewGetAccountService().Do(con.Background())
		if err == nil {
			return true
		} else {
			return false
		}
	} else {
		return true
	}
}
