package main

import (
	// developers can use any library to add a custom cookie encoder/decoder.
	// At this example we use the gorilla's securecookie library:
	"github.com/gorilla/securecookie"

	"github.com/radicalmind/xeon"
	"github.com/radicalmind/xeon/context"
	"github.com/radicalmind/xeon/sessions"
)

func main() {
	app := xeon.New()
	// enable all (error) logs
	// select the httprouter as the servemux

	cookieName := "mycustomsessionid"
	// AES only supports key sizes of 16, 24 or 32 bytes.
	// You either need to provide exactly that amount or you derive the key from what you type in.
	hashKey := []byte("the-big-and-secret-fash-key-here")
	blockKey := []byte("lot-secret-of-characters-big-too")
	secureCookie := securecookie.New(hashKey, blockKey)

	mySessions := sessions.New(sessions.Config{
		Cookie: cookieName,
		Encode: secureCookie.Encode,
		Decode: secureCookie.Decode,
	})

	// OPTIONALLY:
	// import "github.com/radicalmind/xeon/sessions/sessiondb/redis"
	// or import "github.com/radicalmind/go-sessions/sessiondb/$any_available_community_database"
	// mySessions.UseDatabase(redis.New(...))

	app.AttachSessionManager(mySessions) // Attach the session manager we just created.

	app.Get("/", func(ctx context.Context) {
		ctx.Writef("You should navigate to the /set, /get, /delete, /clear,/destroy instead")
	})
	app.Get("/set", func(ctx context.Context) {

		//set session values
		ctx.Session().Set("name", "xeon")

		//test if setted here
		ctx.Writef("All ok session setted to: %s", ctx.Session().GetString("name"))
	})

	app.Get("/get", func(ctx context.Context) {
		// get a specific key, as string, if no found returns just an empty string
		name := ctx.Session().GetString("name")

		ctx.Writef("The name on the /set was: %s", name)
	})

	app.Get("/delete", func(ctx context.Context) {
		// delete a specific key
		ctx.Session().Delete("name")
	})

	app.Get("/clear", func(ctx context.Context) {
		// removes all entries
		ctx.Session().Clear()
	})

	app.Get("/destroy", func(ctx context.Context) {
		//destroy, removes the entire session data and cookie
		ctx.SessionDestroy()
	}) // Note about destroy:
	//
	// You can destroy a session outside of a handler too, using the:
	// mySessions.DestroyByID
	// mySessions.DestroyAll

	app.Run(xeon.Addr(":8080"))
}
