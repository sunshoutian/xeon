package main

import (
	"github.com/radicalmind/xeon"
	"github.com/radicalmind/xeon/context"
	"github.com/radicalmind/xeon/view"
)

func main() {
	app := xeon.New()
	// $ go get -u github.com/jteeuwen/go-bindata/...
	// $ go-bindata ./templates/...
	// $ go build
	// $ ./embedding-templates-into-app
	// html files are not used, you can delete the folder and run the example
	app.AttachView(view.HTML("./templates", ".html").Binary(Asset, AssetNames))
	app.Get("/", hi)

	// http://localhost:8080
	app.Run(xeon.Addr(":8080"))
}

type page struct {
	Title, Name string
}

func hi(ctx context.Context) {
	ctx.ViewData("", page{Title: "Hi Page", Name: "xeon"})
	ctx.View("hi.html")
}
