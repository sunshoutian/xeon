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

	srv := &http.Server{Addr: ":8080" /* Any custom fields here: Handler and ErrorLog are setted to the server automatically */}
	// http://localhost:8080/
	// http://localhost:8080/mypath
	app.Run(xeon.Server(srv)) // same as app.Run(xeon.Addr(":8080"))

	// More:
	// see "multi" if you need to use more than one server at the same app.
	//
	// for a custom listener use: xeon.Listener(net.Listener) or
	// xeon.TLS(cert,key) or xeon.AutoTLS(), see "custom-listener" example for those.
}
