package main

import (
	"github.com/radicalmind/xeon"
)

func main() {
	app := xeon.New()

	// [...]

	// Good when you want to change some of the configuration's field.
	// I use that method :)
	app.Run(xeon.Addr(":8080"), xeon.WithoutBanner, xeon.WithTray, xeon.WithCharset("UTF-8"))

	// or before run:
	// app.Configure(xeon.WithoutBanner, xeon.WithTray, xeon.WithCharset("UTF-8"))
	// app.Run(xeon.Addr(":8080"))
}
