package middlewares

import (
	"strings"

	"github.com/astaxie/beego/context"
)

var LoginRequired = func(ctx *context.Context) {
	if strings.HasPrefix(ctx.Input.URL(), "/kullanici/giris") ||
		strings.HasPrefix(ctx.Input.URL(), "/kullanici/kayit") ||
		strings.HasPrefix(ctx.Input.URL(), "/kullanici/onay") ||
		strings.HasPrefix(ctx.Input.URL(), "/kullanici/sifremi-unuttum") ||
		strings.HasPrefix(ctx.Input.URL(), "/kullanici/yeni-sifre") {
		return
	}

	_, ok := ctx.Input.Session("uid").(int)
	if !ok {
		ctx.Redirect(302, "/kullanici/giris")
	}
}
