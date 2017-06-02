package main

import (
	"github.com/radicalmind/xeon"
	"github.com/radicalmind/xeon/context"
	"github.com/radicalmind/xeon/middleware/i18n"
)

func main() {
	app := xeon.New()

	app.Use(i18n.New(i18n.Config{
		Default:      "en-US",
		URLParameter: "lang",
		Languages: map[string]string{
			"en-US": "./locales/locale_en-US.ini",
			"el-GR": "./locales/locale_el-GR.ini",
			"zh-CN": "./locales/locale_zh-CN.ini"}}))

	app.Get("/", func(ctx context.Context) {

		// it tries to find the language by:
		// ctx.Values().GetString("language")
		// if that was empty then
		// it tries to find from the URLParameter setted on the configuration
		// if not found then
		// it tries to find the language by the "language" cookie
		// if didn't found then it it set to the Default setted on the configuration

		// hi is the key, 'radicalmind' is the %s on the .ini file
		// the second parameter is optional

		// hi := ctx.Translate("hi", "radicalmind")
		// or:
		hi := i18n.Translate(ctx, "hi", "radicalmind")

		language := ctx.Values().GetString(ctx.Application().ConfigurationReadOnly().GetTranslateLanguageContextKey()) // return is form of 'en-US'

		// The first succeed language found saved at the cookie with name ("language"),
		//  you can change that by changing the value of the:  xeon.TranslateLanguageContextKey
		ctx.Writef("From the language %s translated output: %s", language, hi)
	})

	// go to http://localhost:8080/?lang=el-GR
	// or http://localhost:8080
	// or http://localhost:8080/?lang=zh-CN
	app.Run(xeon.Addr(":8080"))

}
