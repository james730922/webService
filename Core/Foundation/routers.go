package Foundation

import (
	"github.com/kataras/iris/v12"
)

func InitCoreRouters(app *iris.Application) {
	base := app.Party("@")
	base.Get("/api/:cmdId/:cmdName", RouteApis)
	base.Post("/api/:cmdId/:cmdName", RouteApis)
}


