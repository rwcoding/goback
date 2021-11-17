package boot

import (
	"github.com/gin-gonic/gin"
	"github.com/rwcoding/goback/pkg/config"
	"github.com/rwcoding/goback/pkg/logger"
	"github.com/rwcoding/goback/pkg/util"
	"net/http"
	"runtime/debug"
	"strconv"
)

var app *Application

type Application struct {
	Router  *Router
	Console *util.Console
}

func App() *Application {
	if app != nil {
		return app
	}
	app = &Application{
		Router:  NewRouter(),
		Console: util.NewConsole(nil),
	}
	return app
}

func (app *Application) Run(ctx *Context) {
	defer func() {
		if r := recover(); r != nil {
			if s, ok := r.(string); ok {
				logger.Error(s)
			}
			logger.Error(string(debug.Stack()))
		}
	}()

	if res := ctx.Verify(); res != nil {
		ctx.g.Data(http.StatusOK, "application/json", res.Json())
		return
	}

	if handler := app.Router.Find(ctx.api); handler != nil {
		req := handler(ctx)
		if err := ctx.NewRequest(req); err != nil {
			ctx.Error(err)
			return
		}
		ctx.Render(req.Run())
	} else {
		ctx.g.String(http.StatusOK, `{"err":19999,"msg":"route error"}`)
	}
}

func (app *Application) RunGinCtx(g *gin.Context) {
	if config.NeedWriteHeader() {
		g.Header("Access-Control-Allow-Origin", "*")
		g.Header("Access-Control-Allow-Headers", "Content-Type,Go-Session,Go-Time,Go-Sign,Go-Api")
		g.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
		g.Header("Access-Control-Max-Age", "86400")

		if g.Request.Method == "OPTIONS" {
			g.AbortWithStatus(http.StatusNoContent)
			return
		}
	}

	cmd := GetParamFromCtx(g, "Go-Api")
	if cmd == "" {
		g.String(http.StatusOK, `{"err":19999,"msg":"api error"}`)
		return
	}

	time, err := strconv.Atoi(GetParamFromCtx(g, "Go-Time"))
	if err != nil {
		g.String(http.StatusOK, `{"err":19999,"msg":"api error"}`)
	}

	app.Run(&Context{
		g:         g,
		api:       cmd,
		sessionId: GetParamFromCtx(g, "Go-Session"),
		sign:      GetParamFromCtx(g, "Go-Sign"),
		time:      time,
	})
}

func GetParamFromCtx(g *gin.Context, name string) string {
	if config.OnlyGetPost() {
		return g.Query(name)
	}
	return g.GetHeader(name)
}
