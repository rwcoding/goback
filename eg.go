package goback

import (
	"github.com/gin-gonic/gin"
	"github.com/rwcoding/goback/models"
	"github.com/rwcoding/goback/pkg/boot"
	"github.com/rwcoding/goback/pkg/config"
	"gorm.io/gorm"
)

// App 获取App对象
func App() *boot.Application {
	return boot.App()
}

// SetDev 设置是否开发环境
func SetDev(is bool) {
	config.CC.IsDev = is
}

// SetWriteHeader 设置是否需要写入跨域Header
func SetWriteHeader(is bool) {
	config.CC.Header = is
}

// SetOnlyGP 设置是否仅支持Get|Post方法
func SetOnlyGP(is bool) {
	config.CC.OnlyGP = is
}

// SetLogFile 日志文件，如 /log/back.log
func SetLogFile(log string) {
	config.CC.Log = log
}

// SetLang 语言，验证器会载入语言包，默认zh
func SetLang(lang string) {
	config.CC.Lang = lang
}

// SetDb 设置gorm实例对象
func SetDb(db *gorm.DB) {
	models.SetDb(db)
}

// Route 添加路由
func Route(r string, handler boot.Handler, name string) {
	App().Router.Add(r, handler)
	boot.AddAuthority(r, name)
}

// Run 运行一个gin上下文
func Run(g *gin.Context) {
	App().RunGinCtx(g)
}
