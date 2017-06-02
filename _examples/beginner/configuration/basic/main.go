package main

import (
	"github.com/radicalmind/xeon"
)

func main() {
	app := xeon.New()

	// [...]

	// Good when you want to modify the whole configuration.
	app.Run(xeon.Addr(":8080"), xeon.WithConfiguration(xeon.Configuration{ // default configuration:
		DisableBanner:                     false,
		DisableTray:                       false,
		DisableInterruptHandler:           false,
		DisablePathCorrection:             false,
		EnablePathEscape:                  false,
		FireMethodNotAllowed:              false,
		DisableBodyConsumptionOnUnmarshal: false,
		DisableAutoFireStatusCode:         false,
		TimeFormat:                        "Mon, 02 Jan 2006 15:04:05 GMT",
		Charset:                           "UTF-8",
	}))

	// or before run:
	// app.Configure(xeon.WithConfiguration(...))
	// app.Run(xeon.Addr(":8080"))
}
