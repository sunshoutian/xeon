package main

import (
	"github.com/radicalmind/xeon"
	"github.com/radicalmind/xeon/context"
)

func main() {
	app := xeon.New()

	app.Get("/", func(ctx context.Context) {
		ctx.Writef("Hello from the SECURE server")
	})

	app.Get("/mypath", func(ctx context.Context) {
		ctx.Writef("Hello from the SECURE server on path /mypath")
	})

	// start the server (HTTPS) on port 443, this is a blocking func
	app.Run(xeon.TLS("127.0.0.1:443", "mycert.cert", "mykey.key"))
}
