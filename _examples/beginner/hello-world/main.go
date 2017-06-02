package main

import (
	"github.com/radicalmind/xeon"
	"github.com/radicalmind/xeon/context"
)

func main() {
	app := xeon.New()

	// a snapshot you will understand easy:
	// app.Macros().Int.RegisterFunc("even", func() func(paramValue string) bool { // bug
	// 	return func(paramValue string) bool {
	// 		n, err := strconv.Atoi(paramValue)
	// 		if err != nil {
	// 			return false
	// 		}
	// 		return n%2 == 0
	// 	}
	// })

	// app.Handle("GET", "/users/{userid:int even()}", func(ctx context.Context) {
	// 	ctx.Writef("Hello user with id: %s", ctx.Params().Get("userid"))
	// })

	app.Get("/", func(ctx context.Context) {
		ctx.HTML("<h1>Index /</h1>")
	})

	if err := app.Run(xeon.Addr(":8080")); err != nil {
		panic(err)
	}

}
