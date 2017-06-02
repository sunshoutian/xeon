package main

import (
	"github.com/radicalmind/xeon"
	"github.com/radicalmind/xeon/context"
)

func main() {
	app := xeon.New()

	// This will serve the ./static/favicons/xeon_32_32.ico to: localhost:8080/favicon.ico
	app.Favicon("./static/favicons/xeon_32_32.ico")

	// app.Favicon("./static/favicons/xeon_32_32.ico", "/favicon_48_48.ico")
	// This will serve the ./static/favicons/xeon_32_32.ico to: localhost:8080/favicon_48_48.ico

	app.Get("/", func(ctx context.Context) {
		ctx.HTML(`<a href="/favicon.ico"> press here to see the favicon.ico</a>.
		 At some browsers like chrome, it should be visible at the top-left side of the browser's window,
		 because some browsers make requests to the /favicon.ico automatically,
		  so Xeon serves your favicon in that path too (you can change it).`)
	}) // if favicon doesn't show to you, try to clear your browser's cache.

	app.Run(xeon.Addr(":8080"))
}
