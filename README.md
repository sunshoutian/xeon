# Xeon

Fully-featured HTTP/2 backend web framework written entirely in Google’s Go Language.

[![Build Status](https://api.travis-ci.org/radicalmind/xeon.svg?branch=master&style=flat-square)](https://travis-ci.org/radicalmind/xeon)
[![Report Card](https://img.shields.io/badge/report%20card%20-a%2B-006699.svg?style=flat-square)](http://goreportcard.com/report/radicalmind/xeon)
[![Docs](https://img.shields.io/badge/docs-%20reference-5272B4.svg?style=flat-square)](https://godoc.org/github.com/radicalmind/xeon)
[![Chat](https://img.shields.io/badge/community-%20chat-00BCD4.svg?style=flat-square)](https://gitter.im/xeon-go/Lobby)


## Installation

The only requirement is the [Go Programming Language](https://golang.org/dl/), at least 1.8.

```sh
$ go get -u github.com/radicalmind/xeon
```

```go
package main

import (
    "github.com/radicalmind/xeon"
    "github.com/radicalmind/xeon/context"
    "github.com/radicalmind/xeon/view"
)

// User is just a bindable object structure.
type User struct {
    Username  string `json:"username"`
    Firstname string `json:"firstname"`
    Lastname  string `json:"lastname"`
    City      string `json:"city"`
    Age       int    `json:"age"`
}

func main() {
    app := xeon.New()
    
    // Define templates using the std html/template engine.
    // Parse and load all files inside "./views" folder with ".html" file extension.
    // Reload the templates on each request (development mode).
    app.AttachView(view.HTML("./views", ".html").Reload(true))
    
    // Regster custom handler for specific http errors.
    app.OnErrorCode(xeon.StatusInternalServerError, func(ctx context.Context) {
    	// .Values are used to communicate between handlers, middleware.
    	errMessage := ctx.Values().GetString("error")
    	if errMessage != "" {
    		ctx.Writef("Internal server error: %s", errMessage)
    		return
    	}
    
    	ctx.Writef("(Unexpected) internal server error")
    })
    
    app.Use(func(ctx context.Context) {
    	ctx.Application().Log("Begin request for path: %s", ctx.Path())
    	ctx.Next()
    })
    
    // app.Done(func(ctx context.Context) {})
    
    // Method POST: http://localhost:8080/decode
    app.Post("/decode", func(ctx context.Context) {
    	var user User
    	ctx.ReadJSON(&user)
    	ctx.Writef("%s %s is %d years old and comes from %s", user.Firstname, user.Lastname, user.Age, user.City)
    })
    
    // Method GET: http://localhost:8080/encode
    app.Get("/encode", func(ctx context.Context) {
    	doe := User{
    		Username:  "Johndoe",
    		Firstname: "John",
    		Lastname:  "Doe",
    		City:      "Neither FBI knows!!!",
    		Age:       25,
    	}
    
    	ctx.JSON(doe)
    })
    
    // Method GET: http://localhost:8080/profile/anytypeofstring
    app.Get("/profile/{username:string}", profileByUsername)
    
    usersRoutes := app.Party("/users", logThisMiddleware)
    {
    	// Method GET: http://localhost:8080/users/42
    	usersRoutes.Get("/{id:int min(1)}", getUserByID)
    	// Method POST: http://localhost:8080/users/create
    	usersRoutes.Post("/create", createUser)
    }
    
    // Listen for incoming HTTP/1.x & HTTP/2 clients on localhost port 8080.
    app.Run(xeon.Addr(":8080"), xeon.WithCharset("UTF-8"))
}

func logThisMiddleware(ctx context.Context) {
    ctx.Application().Log("Path: %s | IP: %s", ctx.Path(), ctx.RemoteAddr())
    
    // .Next is required to move forward to the chain of handlers,
    // if missing then it stops the execution at this handler.
    ctx.Next()
}

func profileByUsername(ctx context.Context) {
    // .Params are used to get dynamic path parameters.
    username := ctx.Params().Get("username")
    ctx.ViewData("Username", username)
    // renders "./views/users/profile.html"
    // with {{ .Username }} equals to the username dynamic path parameter.
    ctx.View("users/profile.html")
}

func getUserByID(ctx context.Context) {
    userID := ctx.Params().Get("id") // Or convert directly using: .Values().GetInt/GetInt64 etc...
    // your own db fetch here instead of user :=...
    user := User{Username: "username" + userID}
    
    ctx.XML(user)
}

func createUser(ctx context.Context) {
    var user User
    err := ctx.ReadForm(&user)
    if err != nil {
    	ctx.Values().Set("error", "creating user, read and parse form failed. "+err.Error())
    	ctx.StatusCode(xeon.StatusInternalServerError)
    	return
    }
    // renders "./views/users/create_verification.html"
    // with {{ . }} equals to the User object, i.e {{ .Username }} , {{ .Firstname}} etc...
    ctx.ViewData("", user)
    ctx.View("users/create_verification.html")
}
```

> Psst: Wanna go to [_examples](https://github.com/radicalmind/xeon/tree/master/_examples) to see more code-snippets?


## Support

- [Post](https://github.com/radicalmind/xeon/issues) a feature request or report a bug, will help to make the framework even better
- :star: and watch [the project](https://github.com/radicalmind/xeon/stargazers), will notify you about updates
- :earth_americas: publish [an article](https://medium.com/) or share a [tweet](https://twitter.com/) about Xeon

Thanks in advance!