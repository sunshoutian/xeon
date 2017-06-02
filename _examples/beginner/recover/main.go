package main

import (
	"github.com/radicalmind/xeon"
	"github.com/radicalmind/xeon/context"

	"github.com/radicalmind/xeon/middleware/recover"
)

func main() {
	app := xeon.New()
	// use this recover(y) middleware
	app.Use(recover.New())

	i := 0
	// let's simmilate a panic every next request
	app.Get("/", func(ctx context.Context) {
		i++
		if i%2 == 0 {
			panic("a panic here")
		}
		ctx.Writef("Hello, refresh one time more to get panic!")
	})

	// http://localhost:8080, refresh it 5-6 times.
	app.Run(xeon.Addr(":8080"))
}

// Note:
// app := xeon.Default() instead of xeon.New() makes use of the recovery middleware automatically.
