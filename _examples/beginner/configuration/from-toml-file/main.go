package main

import (
	"github.com/radicalmind/xeon"
)

func main() {
	app := xeon.New()

	// [...]

	// Good when you have two configurations, one for development and a different one for production use.
	app.Run(xeon.Addr(":8080"), xeon.WithConfiguration(xeon.TOML("./configs/xeon.tml")))

	// or before run:
	// app.Configure(xeon.WithConfiguration(xeon.TOML("./configs/xeon.tml")))
	// app.Run(xeon.Addr(":8080"))
}
