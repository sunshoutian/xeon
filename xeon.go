// Copyright 2017 ΓΜ. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package xeon

import (
	// std packages
	"io"
	"log"
	"net"
	"net/http"
	"sync"
	"time"

	// context for the handlers
	"github.com/radicalmind/xeon/context"
	// core packages, needed to build the application
	"github.com/radicalmind/xeon/core/errors"
	"github.com/radicalmind/xeon/core/host"
	"github.com/radicalmind/xeon/core/logger"
	"github.com/radicalmind/xeon/core/nettools"
	"github.com/radicalmind/xeon/core/router"
	// sessions and view
	"github.com/radicalmind/xeon/sessions"
	"github.com/radicalmind/xeon/view"
	// middleware used in Default method
	requestLogger "github.com/radicalmind/xeon/middleware/logger"
	"github.com/radicalmind/xeon/middleware/recover"
)

const banner = `____  ___                   
\   \/  /____  ____   ____  
 \     // __ \/  _ \ /    \ 
 /     \  ___(  <_> )   |  \
/___/\  \___  >____/|___|  /
      \_/   \/           \/`

const (

	// Version is the current version number of the Xeon Web framework.
	Version = "1.8.3"
)

// MethodNone is a Virtual method
// to store the "offline" routes.
//
// Conversion for router.MethodNone.
const MethodNone = router.MethodNone

// Application is responsible to manage the state of the application.
// It contains and handles all the necessary parts to create a fast web server.
type Application struct {
	Scheduler host.Scheduler
	// routing embedded | exposing APIBuilder's and Router's public API.
	*router.APIBuilder
	*router.Router
	ContextPool *context.Pool

	// config contains the configuration fields
	// all fields defaults to something that is working, developers don't have to set it.
	config *Configuration

	// logger logs to the defined logger.
	// Use AttachLogger to change the default which prints messages to the os.Stdout.
	// It's just an io.Writer, period.
	logger io.Writer

	// view engine
	view view.View

	// sessions and flash messages
	sessions sessions.Sessions
	// used for build
	once sync.Once
}

// New creates and returns a fresh empty Xeon *Application instance.
func New() *Application {
	config := DefaultConfiguration()

	app := &Application{
		config:     &config,
		logger:     logger.NewDevLogger(),
		APIBuilder: router.NewAPIBuilder(),
		Router:     router.NewRouter(),
	}

	app.ContextPool = context.New(func() context.Context {
		return context.NewContext(app)
	})

	return app
}

func Default() *Application {
	app := New()

	app.AttachView(view.HTML("./templates", ".html"))
	app.AttachSessionManager(sessions.New(sessions.Config{
		Cookie:  "irissessionid",
		Expires: 7 * (24 * time.Hour), // 1 week
	}))

	app.Use(recover.New())
	app.Use(requestLogger.New())

	return app
}

// Configure can called when modifications to the framework instance needed.
// It accepts the framework instance
// and returns an error which if it's not nil it's printed to the logger.
// See configuration.go for more.
//
// Returns itself in order to be used like app:= New().Configure(...)
func (app *Application) Configure(configurators ...Configurator) *Application {
	for _, cfg := range configurators {
		cfg(app)
	}

	return app
}

func (app *Application) Build() (err error) {
	app.once.Do(func() {
		// view engine
		// here is where we declare the closed-relative framework functions.
		// Each engine has their defaults, i.e yield,render,render_r,partial, params...
		rv := router.NewRoutePathReverser(app.APIBuilder)
		app.view.AddFunc("urlpath", rv.Path)
		// app.view.AddFunc("url", rv.URL)
		err = app.view.Load()
		if err != nil {
			return // if view engine loading failed then don't continue
		}

		var routerHandler router.RequestHandler
		// router
		// create the request handler, the default routing handler
		routerHandler = router.NewDefaultHandler()

		err = app.Router.BuildRouter(app.ContextPool, routerHandler, app.APIBuilder)
		// re-build of the router from outside can be done with;
		// app.RefreshRouter()
	})

	return
}

// NewHost accepts a standar *http.Server object,
// completes the necessary missing parts of that "srv"
// and returns a new, ready-to-use, host (supervisor).
//
// other names that could describe this function are:
// NewServer
// ProduceServer
// ConstructServer
// Bind
// BindHost
func (app *Application) NewHost(srv *http.Server) *host.Supervisor {
	// set the server's handler to the framework's router
	if srv.Handler == nil {
		srv.Handler = app.Router
	}

	// check if different ErrorLog provided, if not bind it with the framework's logger
	if srv.ErrorLog == nil {
		srv.ErrorLog = log.New(app.logger, "[HTTP Server] ", 0)
	}

	if srv.Addr == "" {
		srv.Addr = ":8080"
	}

	// create the new host supervisor
	// bind the constructed server and return it
	su := host.New(srv)

	if app.config.vhost == "" { // vhost now is useful for router subdomain on wildcard subdomains,
		// in order to correct decide what to do on:
		// mydomain.com -> invalid
		// localhost -> invalid
		// sub.mydomain.com -> valid
		// sub.localhost -> valid
		// we need the host (without port if 80 or 443) in order to validate these, so:
		app.config.vhost = nettools.ResolveVHost(srv.Addr)
	}
	// the below schedules some tasks that will run among the server

	// I was thinking to have them on Default or here and if user not wanted these, could use a custom core/host
	// but that's too much for someone to just disable the banner for example,
	// so I will bind them to a configuration field, although is not direct to the *Application,
	// host is de-coupled from *Application as the other features too, it took me 2 months for this design.

	// copy the registered schedule tasks from the scheduler, if any will be copied to this host supervisor's scheduler.
	app.Scheduler.CopyTo(&su.Scheduler)

	if !app.config.DisableBanner {
		// show the banner and the available keys to exit from app.
		su.Schedule(host.WriteBannerTask(app.logger, banner+Version))
	}

	// give 5 seconds to the server to wait for the (idle) connections.
	shutdownTimeout := 5 * time.Second
	if app.config.EnableTray {
		// start the tray icon to the taskbar (cross-platform) when server started.
		su.Schedule(host.ShowTrayTask(Version, shutdownTimeout))
	}

	if !app.config.DisableInterruptHandler {
		// when control/cmd+C pressed.
		su.Schedule(host.ShutdownOnInterruptTask(shutdownTimeout))
	}

	return su
}

// Runner is just an interface which accepts the framework instance
// and returns an error.
//
// It can be used to register a custom runner with `Run` in order
// to set the framework's server listen action.
//
// Currently Runner is being used to declare the built'n server listeners.
//
// See `Run` for more.
type Runner func(*Application) error

func Listener(l net.Listener) Runner {
	return func(app *Application) error {
		app.config.vhost = nettools.ResolveVHost(l.Addr().String())
		return app.NewHost(new(http.Server)).
			Serve(l)
	}
}

func Server(srv *http.Server) Runner {
	return func(app *Application) error {
		return app.NewHost(srv).
			ListenAndServe()
	}
}

func Addr(addr string) Runner {
	return func(app *Application) error {
		return app.NewHost(&http.Server{Addr: addr}).
			ListenAndServe()
	}
}

func TLS(addr string, certFile, keyFile string) Runner {
	return func(app *Application) error {
		return app.NewHost(&http.Server{Addr: addr}).
			ListenAndServeTLS(certFile, keyFile)
	}
}

func AutoTLS(addr string) Runner {
	return func(app *Application) error {
		return app.NewHost(&http.Server{Addr: addr}).
			ListenAndServeAutoTLS()
	}
}

// or that to accept just a func like host.ListenAndServe instead of func(*Application) error { return host.ListenAndServe() }:
func Raw(f func() error) Runner {
	return func(*Application) error {
		return f()
	}
}

// first a server runner because it should be only one
// secondly the optional configuration.
// If more than one server needed to run on the same xeon instance
// then users should use the `go NewHost(*http.Server).Serve/ListenAndServe/ListenAndServeTLS/ListenAndServeAutoTLS`.
// or run with custom function and configs, i.e an already created host
// h := NewHost(*http.Server)
// Run(Raw(h.ListenAndServe), WithoutBanner, WithTray, WithCharset("UTF-8"))
func (app *Application) Run(serve Runner, withOrWithout ...Configurator) error {
	// first Build because it doesn't need anything from configuration,
	//  this give the user the chance to modify the router inside a configurator as well.
	if err := app.Build(); err != nil {
		return err
	}

	app.Configure(withOrWithout...)

	// this will block until an error(unless supervisor's DeferFlow called from a Task).
	return serve(app)
}

// AttachLogger attachs a new logger to the framework.
func (app *Application) AttachLogger(logWriter io.Writer) {
	if logWriter == nil {
		// convert that to an empty writerFunc
		logWriter = logger.NoOpLogger
	}
	app.logger = logWriter
}

// Log sends a message to the defined io.Writer logger, it's
// just a help function for internal use but it can be used to a cusotom middleware too.
//
// See AttachLogger too.
func (app *Application) Log(format string, a ...interface{}) {
	logger.Log(app.logger, format, a...)
}

// AttachView should be used to register view engines mapping to a root directory
// and the template file(s) extension.
// Returns an error on failure, otherwise nil.
func (app *Application) AttachView(viewEngine view.Engine) error {
	return app.view.Register(viewEngine)
}

var errViewEngineIsMissing = errors.New(
	`
Please attach one of the available view engines.
By-default Xeon supports five template engines:
 - standard html  | view.HTML(...)
 - django         | view.Django(...)
 - handlebars     | view.Handlebars(...)
 - pug(jade)      | view.Pug(...)
 - amber          | view.Amber(...)

Edit your main .go source file to attach one of these and restart your app.
	i.e: lines (<---) were missing.
	-------------------------------------------------------------------
	import (
		"github.com/radicalmind/xeon"
		"github.com/radicalmind/xeon/view" // <--- this line
	)

	func main(){
		app := xeon.New()
		[...]
		app.AttachView(view.HTML("./templates", ".html")) // <--- and this line, were missing.
		[...]

		app.Start()
	}
	-------------------------------------------------------------------
 `)

// View executes and writes the result of a template file to the writer.
//
// First parameter is the writer to write the parsed template.
// Second parameter is the relative, to templates directory, template filename, including extension.
// Third parameter is the layout, can be empty string.
// Forth parameter is the bindable data to the template, can be nil.
//
// Use context.View to render templates to the client instead.
// Returns an error on failure, otherwise nil.
func (app *Application) View(writer io.Writer, filename string, layout string, bindingData interface{}) error {
	if app.view.Len() == 0 {
		return errViewEngineIsMissing
	}
	return app.view.ExecuteWriter(writer, filename, layout, bindingData)
}

// AttachSessionManager registers a session manager to the framework which is used for flash messages too.
//
// See context.Session too.
func (app *Application) AttachSessionManager(manager sessions.Sessions) {
	app.sessions = manager
}

var errSessionsManagerIsMissing = errors.New(
	`
SessionsManager called without specified sessions manager!
Please follow the below to fix that.

Edit your main .go source file to adapt one of these and restart your app.
	i.e: lines (<---) were missing.
	-------------------------------------------------------------------
	import (
		"github.com/radicalmind/xeon"
		"github.com/radicalmind/xeon/sessions" // <--- this line
	)

	func main(){
		app := xeon.New()
		// [...]
		mySessions := sessions.New(sessions.config{
			// Cookie string, the session's client cookie name, for example: "mysessionid"
			//
			// Defaults to "gosessionid"
			Cookie: "mysessionid",
			// it's time.Duration, from the time cookie is created, how long it can be alive?
			// 0 means no expire.
			Expires: 0,
			// if you want to invalid cookies on different subdomains
			// of the same host, then enable it
			DisableSubdomainPersistence: false,
		})

		app.AttachSessionManager(mySessions) // <--- and this line were, missing.
		// [...]

		app.Start()
	}
	-------------------------------------------------------------------
 `)

// SessionManager returns the session manager which contain a Start and Destroy methods
// used inside the context.Session().
//
// It's ready to use after the RegisterSessions.
func (app *Application) SessionManager() (sessions.Sessions, error) {
	if app.sessions == nil {
		return nil, errSessionsManagerIsMissing
	}
	return app.sessions, nil
}

// ConfigurationReadOnly returns a structure which doesn't allow writing.
func (app *Application) ConfigurationReadOnly() context.ConfigurationReadOnly {
	return app.config
}
