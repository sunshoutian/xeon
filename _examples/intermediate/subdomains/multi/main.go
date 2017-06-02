package main

import (
	"github.com/radicalmind/xeon"
	"github.com/radicalmind/xeon/context"
)

func main() {
	app := xeon.New()

	/*
	 * Setup static files
	 */

	app.StaticWeb("/assets", "./public/assets")
	app.StaticWeb("/upload_resources", "./public/upload_resources")

	dashboard := app.Party("dashboard.")
	{
		dashboard.Get("/", func(ctx context.Context) {
			ctx.Writef("HEY FROM dashboard")
		})
	}
	system := app.Party("system.")
	{
		system.Get("/", func(ctx context.Context) {
			ctx.Writef("HEY FROM system")
		})
	}

	app.Get("/", func(ctx context.Context) {
		ctx.Writef("HEY FROM frontend /")
	})
	// http://domain.local:80
	// http://dashboard.local
	// http://system.local
	// Make sure you prepend the "http" in your browser
	// because .local is a virtual domain we think to show case you
	// that you can declare any syntactical correct name as a subdomain in Xeon.
	app.Run(xeon.Addr("domain.local:80")) // for beginners: look ../hosts file
}
