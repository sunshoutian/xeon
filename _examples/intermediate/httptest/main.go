package main

import (
	"github.com/radicalmind/xeon"

	"github.com/radicalmind/xeon/sessions"
)

func newApp() *xeon.Application {
	app := xeon.New()

	app.Adapt(sessions.New(sessions.Config{Cookie: "mysessionid"}))

	app.Get("/hello", func(ctx context.Context) {
		sess := ctx.Session()
		if !sess.HasFlash() /* or sess.GetFlash("name") == "", same thing here */ {
			ctx.HTML(xeon.StatusUnauthorized, "<h1> Unauthorized Page! </h1>")
			return
		}

		ctx.JSON(xeon.StatusOK, context.Map{
			"Message": "Hello",
			"From":    sess.GetFlash("name"),
		})
	})

	app.Post("/login", func(ctx context.Context) {
		sess := ctx.Session()
		if !sess.HasFlash() {
			sess.SetFlash("name", ctx.FormValue("name"))
		}
		// let's no redirect, just set the flash message, nothing more.
	})

	return app
}

func main() {
	app := newApp()
	app.Run(xeon.Addr(":8080"))
}
