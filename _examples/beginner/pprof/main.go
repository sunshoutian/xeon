package main

import (
	"github.com/radicalmind/xeon"
	"github.com/radicalmind/xeon/context"

	"github.com/radicalmind/xeon/middleware/pprof"
)

func main() {
	app := xeon.New()

	app.Get("/", func(ctx context.Context) {
		ctx.HTML("<h1> Please click <a href='/debug/pprof'>here</a>")
	})

	app.Any("/debug/pprof/{action:path}", pprof.New())
	//                              ___________
	app.Run(xeon.Addr(":8080"))
}
