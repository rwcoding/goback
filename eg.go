package goback

import (
	"github.com/gin-gonic/gin"
	"github.com/rwcoding/goback/models"
	"github.com/rwcoding/goback/pkg/boot"
	"github.com/rwcoding/goback/pkg/config"
	"gorm.io/gorm"
)

func App() *boot.Application {
	return boot.App()
}

func SetDev(is bool) {
	config.CC.IsDev = is
}

func SetWriteHeader(is bool) {
	config.CC.Header = is
}

func SetOnlyGP(is bool) {
	config.CC.OnlyGP = is
}

func SetLogFile(log string) {
	config.CC.Log = log
}

func SetLang(lang string) {
	config.CC.Lang = lang
}

func SetDb(db *gorm.DB) {
	models.SetDb(db)
}

func Route(r string, handler boot.Handler, name string) {
	App().Router.Add(r, handler)
	boot.AddAuthority(r, name)
}

func Run(g *gin.Context) {
	App().RunGinCtx(g)
}
