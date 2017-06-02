```go
// Copyright 2017 ΓΜ. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Alias-related code. Keep for now.
// On August: 
// 1. This should be transferred at the top of the current context.go file,
//    which contains the default, unexported, xeon/context.Context implementation.
// 2. Remove the ignore build flag
// 3. Remove this context_go.1.8.1.typealias.md file
// 4. Update docs in order to use only the main import path instead of two.
// +build ignore

package xeon

import (
	"github.com/radicalmind/xeon/context"
)

type Context = context.Context

/* When new go version will be released: 
import "github.com/radicalmind/xeon"
func main(){
	app := xeon.New()
	app.Handle("GET", "/", func(ctx xeon.Context){
		ctx.HTML("<h2> Hello World! </h2>")
	})
	app.Run(xeon.Addr(":8080"))
}
*/

/* Instead of, current need of import the context package,
   which contains the interface: 
import (
	"github.com/radicalmind/xeon"
	"github.com/radicalmind/xeon/context"
)
func main(){
	app := xeon.New()
	app.Handle("GET", "/", func(ctx context.Context){
		ctx.HTML("<h2> Hello World! </h2>")
	})
	app.Run(xeon.Addr(":8080"))
}
*/
```