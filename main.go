package main

import (
	"github.com/dixydo/olxmanager-server/controllers"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
)

func main() {
	app := newApp()
	app.Run(iris.Addr(":8080"), iris.WithoutServerError(iris.ErrServerClosed))
}

func newApp() *iris.Application {
	app := iris.New()

	adverts := mvc.New(app.Party("/adverts"))
	adverts.Handle(new(controllers.AdvertController))

	return app
}
