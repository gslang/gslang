package stdlib

import (
	"github.com/gslang/gslang"
)

// BuiltinModules are builtin type standard library modules.
var BuiltinModules = map[string]map[string]gslang.Object{
	"math":   mathModule,
	"os":     osModule,
	"text":   textModule,
	"time":   timeModule,
	"http":   httpModule,
	"rand":   randModule,
	"fmt":    fmtModule,
	"json":   jsonModule,
	"base64": base64Module,
	"crypto": cryptoModule,
	"hex":    hexModule,
}
