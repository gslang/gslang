package stdlib

import (
	"github.com/gslang/gslang"
)

func wrapError(err error) gslang.Object {
	if err == nil {
		return gslang.TrueValue
	}
	return &gslang.Error{Value: &gslang.String{Value: err.Error()}}
}
