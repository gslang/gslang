package stdlib

import (
	"math"

	"github.com/gslang/gslang"
)

var mathModule = map[string]gslang.Object{
	"e":       &gslang.Float{Value: math.E},
	"pi":      &gslang.Float{Value: math.Pi},
	"phi":     &gslang.Float{Value: math.Phi},
	"sqrt2":   &gslang.Float{Value: math.Sqrt2},
	"sqrtE":   &gslang.Float{Value: math.SqrtE},
	"sqrtPi":  &gslang.Float{Value: math.SqrtPi},
	"sqrtPhi": &gslang.Float{Value: math.SqrtPhi},
	"ln2":     &gslang.Float{Value: math.Ln2},
	"log2E":   &gslang.Float{Value: math.Log2E},
	"ln10":    &gslang.Float{Value: math.Ln10},
	"log10E":  &gslang.Float{Value: math.Log10E},
	"abs": &gslang.UserFunction{
		Name:  "abs",
		Value: FuncAFRF(math.Abs),
	},
	"acos": &gslang.UserFunction{
		Name:  "acos",
		Value: FuncAFRF(math.Acos),
	},
	"acosh": &gslang.UserFunction{
		Name:  "acosh",
		Value: FuncAFRF(math.Acosh),
	},
	"asin": &gslang.UserFunction{
		Name:  "asin",
		Value: FuncAFRF(math.Asin),
	},
	"asinh": &gslang.UserFunction{
		Name:  "asinh",
		Value: FuncAFRF(math.Asinh),
	},
	"atan": &gslang.UserFunction{
		Name:  "atan",
		Value: FuncAFRF(math.Atan),
	},
	"atan2": &gslang.UserFunction{
		Name:  "atan2",
		Value: FuncAFFRF(math.Atan2),
	},
	"atanh": &gslang.UserFunction{
		Name:  "atanh",
		Value: FuncAFRF(math.Atanh),
	},
	"cbrt": &gslang.UserFunction{
		Name:  "cbrt",
		Value: FuncAFRF(math.Cbrt),
	},
	"ceil": &gslang.UserFunction{
		Name:  "ceil",
		Value: FuncAFRF(math.Ceil),
	},
	"copysign": &gslang.UserFunction{
		Name:  "copysign",
		Value: FuncAFFRF(math.Copysign),
	},
	"cos": &gslang.UserFunction{
		Name:  "cos",
		Value: FuncAFRF(math.Cos),
	},
	"cosh": &gslang.UserFunction{
		Name:  "cosh",
		Value: FuncAFRF(math.Cosh),
	},
	"dim": &gslang.UserFunction{
		Name:  "dim",
		Value: FuncAFFRF(math.Dim),
	},
	"erf": &gslang.UserFunction{
		Name:  "erf",
		Value: FuncAFRF(math.Erf),
	},
	"erfc": &gslang.UserFunction{
		Name:  "erfc",
		Value: FuncAFRF(math.Erfc),
	},
	"exp": &gslang.UserFunction{
		Name:  "exp",
		Value: FuncAFRF(math.Exp),
	},
	"exp2": &gslang.UserFunction{
		Name:  "exp2",
		Value: FuncAFRF(math.Exp2),
	},
	"expm1": &gslang.UserFunction{
		Name:  "expm1",
		Value: FuncAFRF(math.Expm1),
	},
	"floor": &gslang.UserFunction{
		Name:  "floor",
		Value: FuncAFRF(math.Floor),
	},
	"gamma": &gslang.UserFunction{
		Name:  "gamma",
		Value: FuncAFRF(math.Gamma),
	},
	"hypot": &gslang.UserFunction{
		Name:  "hypot",
		Value: FuncAFFRF(math.Hypot),
	},
	"ilogb": &gslang.UserFunction{
		Name:  "ilogb",
		Value: FuncAFRI(math.Ilogb),
	},
	"inf": &gslang.UserFunction{
		Name:  "inf",
		Value: FuncAIRF(math.Inf),
	},
	"is_inf": &gslang.UserFunction{
		Name:  "is_inf",
		Value: FuncAFIRB(math.IsInf),
	},
	"is_nan": &gslang.UserFunction{
		Name:  "is_nan",
		Value: FuncAFRB(math.IsNaN),
	},
	"j0": &gslang.UserFunction{
		Name:  "j0",
		Value: FuncAFRF(math.J0),
	},
	"j1": &gslang.UserFunction{
		Name:  "j1",
		Value: FuncAFRF(math.J1),
	},
	"jn": &gslang.UserFunction{
		Name:  "jn",
		Value: FuncAIFRF(math.Jn),
	},
	"ldexp": &gslang.UserFunction{
		Name:  "ldexp",
		Value: FuncAFIRF(math.Ldexp),
	},
	"log": &gslang.UserFunction{
		Name:  "log",
		Value: FuncAFRF(math.Log),
	},
	"log10": &gslang.UserFunction{
		Name:  "log10",
		Value: FuncAFRF(math.Log10),
	},
	"log1p": &gslang.UserFunction{
		Name:  "log1p",
		Value: FuncAFRF(math.Log1p),
	},
	"log2": &gslang.UserFunction{
		Name:  "log2",
		Value: FuncAFRF(math.Log2),
	},
	"logb": &gslang.UserFunction{
		Name:  "logb",
		Value: FuncAFRF(math.Logb),
	},
	"max": &gslang.UserFunction{
		Name:  "max",
		Value: FuncAFFRF(math.Max),
	},
	"min": &gslang.UserFunction{
		Name:  "min",
		Value: FuncAFFRF(math.Min),
	},
	"mod": &gslang.UserFunction{
		Name:  "mod",
		Value: FuncAFFRF(math.Mod),
	},
	"nan": &gslang.UserFunction{
		Name:  "nan",
		Value: FuncARF(math.NaN),
	},
	"nextafter": &gslang.UserFunction{
		Name:  "nextafter",
		Value: FuncAFFRF(math.Nextafter),
	},
	"pow": &gslang.UserFunction{
		Name:  "pow",
		Value: FuncAFFRF(math.Pow),
	},
	"pow10": &gslang.UserFunction{
		Name:  "pow10",
		Value: FuncAIRF(math.Pow10),
	},
	"remainder": &gslang.UserFunction{
		Name:  "remainder",
		Value: FuncAFFRF(math.Remainder),
	},
	"signbit": &gslang.UserFunction{
		Name:  "signbit",
		Value: FuncAFRB(math.Signbit),
	},
	"sin": &gslang.UserFunction{
		Name:  "sin",
		Value: FuncAFRF(math.Sin),
	},
	"sinh": &gslang.UserFunction{
		Name:  "sinh",
		Value: FuncAFRF(math.Sinh),
	},
	"sqrt": &gslang.UserFunction{
		Name:  "sqrt",
		Value: FuncAFRF(math.Sqrt),
	},
	"tan": &gslang.UserFunction{
		Name:  "tan",
		Value: FuncAFRF(math.Tan),
	},
	"tanh": &gslang.UserFunction{
		Name:  "tanh",
		Value: FuncAFRF(math.Tanh),
	},
	"trunc": &gslang.UserFunction{
		Name:  "trunc",
		Value: FuncAFRF(math.Trunc),
	},
	"y0": &gslang.UserFunction{
		Name:  "y0",
		Value: FuncAFRF(math.Y0),
	},
	"y1": &gslang.UserFunction{
		Name:  "y1",
		Value: FuncAFRF(math.Y1),
	},
	"yn": &gslang.UserFunction{
		Name:  "yn",
		Value: FuncAIFRF(math.Yn),
	},
}
