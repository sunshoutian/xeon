package main

import (
	"github.com/radicalmind/xeon"
)

func main() {
	app := xeon.New()

	// [...]

	// Good when you have two configurations, one for development and a different one for production use.
	app.Run(xeon.Addr(":8080"), xeon.WithConfiguration(xeon.YAML("./configs/xeon.yml")))

	// or before run:
	// app.Configure(xeon.WithConfiguration(xeon.YAML("./configs/xeon.yml")))
	// app.Run(xeon.Addr(":8080"))
}
