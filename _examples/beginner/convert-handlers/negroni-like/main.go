package main

import (
	"net/http"

	"github.com/radicalmind/xeon"
	"github.com/radicalmind/xeon/context"
	"github.com/radicalmind/xeon/core/handlerconv"
)

func main() {
	app := xeon.New()
	xeonMiddleware := handlerconv.FromStdWithNext(negronilikeTestMiddleware)
	app.Use(xeonMiddleware)

	// Method GET: http://localhost:8080/
	app.Get("/", func(ctx context.Context) {
		ctx.HTML("<h1> Home </h1>")
		// this will print an error,
		// this route's handler will never be executed because the middleware's criteria not passed.
	})

	// Method GET: http://localhost:8080/ok
	app.Get("/ok", func(ctx context.Context) {
		ctx.Writef("Hello world!")
		// this will print "OK. Hello world!".
	})

	// http://localhost:8080
	// http://localhost:8080/ok
	app.Run(xeon.Addr(":8080"))
}

func negronilikeTestMiddleware(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	if r.URL.Path == "/ok" && r.Method == "GET" {
		w.Write([]byte("OK. "))
		next(w, r) // go to the next route's handler
		return
	}
	// else print an error and do not forward to the route's handler.
	w.WriteHeader(xeon.StatusBadRequest)
	w.Write([]byte("Bad request"))
}
