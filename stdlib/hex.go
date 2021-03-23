package stdlib

import (
	"encoding/hex"

	"github.com/gslang/gslang"
)

var hexModule = map[string]gslang.Object{
	"encode": &gslang.UserFunction{Value: FuncAYRS(hex.EncodeToString)},
	"decode": &gslang.UserFunction{Value: FuncASRYE(hex.DecodeString)},
}
