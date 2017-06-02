package main

import (
	"github.com/radicalmind/xeon"
	"github.com/radicalmind/xeon/context"
	"github.com/radicalmind/xeon/view"
)

func main() {
	app := xeon.New() // defaults to these

	app.AttachView(view.HTML("./templates", ".html"))

	app.Get("/", hi)

	// http://localhost:8080
	app.Run(xeon.Addr(":8080"), xeon.WithCharset("UTF-8")) // defaults to that but you can change it.
}

func hi(ctx context.Context) {
	ctx.ViewData("Title", "Hi Page")
	ctx.ViewData("Name", "Xeon") // {{.Name}} will render: Xeon
	// ctx.ViewData("", myCcustomStruct{})
	ctx.View("hi.html")
}
