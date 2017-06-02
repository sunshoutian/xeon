// Copyright 2017 ΓΜ. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package rule

import (
	"github.com/radicalmind/xeon/context"
)

// Rule a superset of validators
type Rule interface {
	Claim(ctx context.Context) bool
	Valid(ctx context.Context) bool
}
