package main

import (
	"net/http"

	"github.com/radicalmind/xeon"
	"github.com/radicalmind/xeon/context"
)

func main() {
	app := xeon.New()

	app.Get("/", func(ctx context.Context) {
		ctx.Writef("Hello from the server")
	})

	app.Get("/mypath", func(ctx context.Context) {
		ctx.Writef("Hello from %s", ctx.Path())
	})

	// call .Build before use the 'app' as an http.Handler on a custom http.Server
	if err := app.Build(); err != nil {
		panic(err)
	}

	// create our custom server and assign the Handler/Router
	srv := &http.Server{Handler: app, Addr: ":8080"} // you have to set Handler:app and Addr, see "xeon-way" which does this automatically.
	// http://localhost:8080/
	// http://localhost:8080/mypath
	println("Start a server listening on http://localhost:8080")
	srv.ListenAndServe() // same as app.Run(xeon.Addr(":8080"))

	// Notes:
	// Banne and Tray are not shown at all. Same for the Interrupt Handler, even if app's configuration allows them.
	//
	// `.Run` is the only one function that cares about those three.

	// More:
	// see "multi" if you need to use more than one server at the same app.
	//
	// for a custom listener use: xeon.Listener(net.Listener) or
	// xeon.TLS(cert,key) or xeon.AutoTLS(), see "custom-listener" example for those.
}
