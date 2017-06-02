package main

import (
	"github.com/radicalmind/xeon"
)

func main() {
	app := xeon.New()

	// first parameter is the request path
	// second is the operating system directory
	app.StaticWeb("/static", "./assets")

	// http://localhost:8080/static/css/main.css
	app.Run(xeon.Addr(":8080"))
}
