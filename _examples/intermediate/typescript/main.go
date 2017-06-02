package main

import (
	"github.com/radicalmind/xeon"
	"github.com/radicalmind/xeon/typescript"
)

// NOTE: Some machines don't allow to install typescript automatically, so if you don't have typescript installed
// and the typescript adaptor doesn't works for you then follow the below steps:
// 1. close the xeon server
// 2. open your terminal and execute: npm install -g typescript
// 3. start your xeon server, it should be work, as expected, now.
func main() {
	app := xeon.New()

	ts := typescript.New()
	ts.Config.Dir = "./www/scripts"
	ts.Attach(app) // attach the application to the typescript compiler

	app.StaticWeb("/", "./www") // serve the index.html
	app.Run(xeon.Addr(":8080"))
}

// open http://localhost:8080
// go to ./www/scripts/app.ts
// make a change
// reload the http://localhost:8080 and you should see the changes
//
// what it does?
// - compiles the typescript files using default compiler options if not tsconfig found
// - watches for changes on typescript files, if a change then it recompiles the .ts to .js
//
// same as you used to do with gulp-like tools, but here at Xeon I do my bests to help GO developers.
