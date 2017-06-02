package main

import (
	"github.com/radicalmind/xeon"
	"github.com/radicalmind/xeon/context"
)

type Company struct {
	Name  string
	City  string
	Other string
}

func MyHandler(ctx context.Context) {
	c := &Company{}
	if err := ctx.ReadJSON(c); err != nil {
		ctx.StatusCode(xeon.StatusBadRequest)
		ctx.WriteString(err.Error())
		return
	}

	ctx.Writef("Received: %#v\n", c)
}

func main() {
	app := xeon.New()

	app.Post("/", MyHandler)

	// use Postman or whatever to do a POST request
	// to the http://localhost:8080 with RAW BODY:
	/*
		{
			"Name": "Xeon-Go",
			"City": "New York",
			"Other": "Something here"
		}
	*/
	// and Content-Type to application/json
	//
	// The response should be:
	// Received: &main.Company{Name:"Xeon-Go", City:"New York", Other:"Something here"}
	app.Run(xeon.Addr(":8080"))
}
