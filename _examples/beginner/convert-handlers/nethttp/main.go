package main

import (
	"net/http"

	"github.com/radicalmind/xeon"
	"github.com/radicalmind/xeon/context"
	"github.com/radicalmind/xeon/core/handlerconv"
)

func main() {
	app := xeon.New()
	xeonMiddleware := handlerconv.FromStd(nativeTestMiddleware)
	app.Use(xeonMiddleware)

	// Method GET: http://localhost:8080/
	app.Get("/", func(ctx context.Context) {
		ctx.HTML("Home")
	})

	// Method GET: http://localhost:8080/ok
	app.Get("/ok", func(ctx context.Context) {
		ctx.HTML("<b>Hello world!</b>")
	})

	// http://localhost:8080
	// http://localhost:8080/ok
	app.Run(xeon.Addr(":8080"))
}

func nativeTestMiddleware(w http.ResponseWriter, r *http.Request) {
	println("Request path: " + r.URL.Path)
}
