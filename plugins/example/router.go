package example

import (
	"github.com/chenhg5/go-admin/context"
	"github.com/chenhg5/go-admin/modules/auth"
)

func InitRouter(prefix string) *context.App {

	app := context.NewApp()
	app.GET(prefix+"/example", auth.Middleware, TestHandler)
	app.GET(prefix, auth.Middleware, TestHandler)

	return app
}
