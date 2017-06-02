# Examples

This folder provides easy to understand code snippets on how to get started with web development with the Go programming language using the  [Xeon](https://github.com/radicalmind/xeon) web framework.

It doesn't contains "best ways" neither explains all its features. It's just a simple, practical cookbook for young Gophers!

## 40+2 examples

* [Level: Beginner](beginner)
    * [Overview](beginner/overview/main.go)
    * [Hello World](beginner/hello-world/main.go)
    * [Configuration](beginner/configuration)
        * [Basic way](beginner/configuration/basic/main.go)
        * [Functional way](beginner/configuration/functional/main.go)
        * [Import from YAML file](beginner/configuration/from-yaml-file/main.go)
        * [Import from TOML file](beginner/configuration/from-toml-file/main.go)
    * [Routing](beginner/routing)
        * [Overview](beginner/routing/main.go)
        * [Basic](beginner/routing/basic/main.go)
        * [Dynamic Path](beginner/routing/dynamic-path/main.go)
        * [Reverse routing](beginner/routing/reverse/main.go)
    * [Internal Application File Logger](beginner/file-logger/main.go)
    * [Custom HTTP Errors](beginner/http-errors/main.go)
    * [Write JSON](beginner/write-json/main.go)
    * [Read JSON](beginner/read-json/main.go)
    * [Read Form](beginner/read-form/main.go)
    * [Favicon](beginner/favicon/main.go)
    * [File Server](beginner/file-server/main.go)
    * [Send Files](beginner/send-files/main.go)
    * [Stream Writer](beginner/stream-writer/main.go)
    * [Send An E-mail](beginner/e-mail/main.go)
    * [Upload/Read Files](beginner/upload-files/main.go)
    * [Recovery](beginner/recover/main.go)
    * [Profiling (pprof)](beginner/pprof/main.go)
    * [Request Logger](beginner/request-logger/main.go)
    * [Basic Authentication](beginner/basicauth/main.go)
    * [Listen UNIX Socket](beginner/listen-unix/main.go)
    * [Listen TLS](beginner/listen-tls/main.go)
    * [Listen Letsencrypt (Automatic Certifications)](beginner/listen-letsencrypt/main.go)
* [Level: Intermediate](intermediate)
    * [Transactions](intermediate/transactions/main.go)
    * [HTTP Testing](intermediate/httptest/main_test.go)
    * [Watch & Compile Typescript source files](intermediate/typescript/main.go)
    * [Cloud Editor](intermediate/cloud-editor/main.go)
    * [Serve Embedded Files](intermediate/serve-embedded-files/main.go)
    * [HTTP Access Control](intermediate/cors/main.go)
    * [Cache Markdown](intermediate/cache-markdown/main.go)
    * [Localization and Internationalization](intermediate/i18n/main.go)
    * [Graceful Shutdown](intermediate/graceful-shutdown)
        * [Basic and simple](intermediate/graceful-shutdown/basic/main.go)
        * [Custom Host](intermediate/graceful-shutdown/custom-host/main.go)
    * [Custom HTTP Server](intermediate/custom-httpserver)
        * [Xeon way](intermediate/custom-httpserver/xeon-way/main.go)
        * [Standar way](intermediate/custom-httpserver/std-way/main.go)
        * [More than one server](intermediate/custom-httpserver/multi/main.go)
    * [Custom TCP Listener](intermediate/custom-listener/main.go)
    * [Custom Context](intermediate/custom-context)
        * [Method Overriding](intermediate/custom-context/method-overriding/main.go)
    * [Route State](intermediate/route-state/main.go)
    * [View Engine](intermediate/view)
        * [Overview](intermediate/view/overview/main.go)
        * [Hi](intermediate/view/template_html_0/main.go)
        * [Showcase one simple Layout](intermediate/view/template_html_1/main.go)
        * [Layouts `yield` and `render` tmpl funcs](intermediate/view/template_html_2/main.go)
        * [Showcase of the `urlpath` tmpl func](intermediate/view/template_html_3/main.go)
        * [Showcase of the `url` tmpl func](intermediate/view/template_html_4/main.go)
        * [Inject Data Between Handlers](intermediate/view/context-view-data/main.go)
        * [Embedding Templates Into App Executable File](intermediate/view/embedding-templates-into-app)
    * [Sessions](intermediate/sessions)
        * [Overview](intermediate/sessions/overview/main.go)
        * [Encoding & Decoding the Session ID: Secure Cookie](intermediate/sessions/securecookie/main.go)
        * [Standalone](intermediate/sessions/standalone/main.go)
        * [With A Back-End Database](intermediate/sessions/database/main.go)
        * [Password Hashing](intermediate/sessions/password-hashing/main.go)
    * [Flash Messages](intermediate/flash-messages/main.go)
    * [Websockets](intermediate/websockets)
        * [Ridiculous Simple](intermediate/websockets/ridiculous-simple/main.go)
        * [Overview](intermediate/websockets/overview/main.go)
        * [Connection List](intermediate/websockets/connectionlist/main.go)
        * [Native Messages](intermediate/websockets/naive-messages/main.go)
        * [Secure](intermediate/websockets/secure/main.go)
        * [Custom Go Client](intermediate/websockets/custom-go-client/main.go)
    * [Subdomains](intermediate/subdomains)
        * [Single](intermediate/subdomains/single/main.go)
        * [Multi](intermediate/subdomains/multi/main.go)
        * [Wildcard](intermediate/subdomains/wildcard/main.go)
* [Level: Advanced](advanced)
    * [Online Visitors](advanced/online-visitors/main.go)
    * [URL Shortener using BoltDB](advanced/url-shortener/main.go)

> Do not forget to [star or watch the project](https://github.com/radicalmind/xeon/stargazers) in order to stay updated with the latest tech trends, it takes some seconds for the sake of go!

> Developers should read the official [documentation](https://godoc.org/github.com/radicalmind/xeon), in depth, for better understanding.