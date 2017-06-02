package main

import (
	"github.com/radicalmind/xeon"
	"github.com/radicalmind/xeon/context"
)

func main() {
	app := xeon.New()

	app.Get("/", func(ctx context.Context) {
		file := "./files/first.zip"
		ctx.SendFile(file, "c.zip")
	})

	app.Run(xeon.Addr(":8080"))
}
