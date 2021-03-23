package stdlib

import (
	"math/rand"

	"github.com/gslang/gslang"
)

var randModule = map[string]gslang.Object{
	"int": &gslang.UserFunction{
		Name:  "int",
		Value: FuncARI64(rand.Int63),
	},
	"float": &gslang.UserFunction{
		Name:  "float",
		Value: FuncARF(rand.Float64),
	},
	"intn": &gslang.UserFunction{
		Name:  "intn",
		Value: FuncAI64RI64(rand.Int63n),
	},
	"exp_float": &gslang.UserFunction{
		Name:  "exp_float",
		Value: FuncARF(rand.ExpFloat64),
	},
	"norm_float": &gslang.UserFunction{
		Name:  "norm_float",
		Value: FuncARF(rand.NormFloat64),
	},
	"perm": &gslang.UserFunction{
		Name:  "perm",
		Value: FuncAIRIs(rand.Perm),
	},
	"seed": &gslang.UserFunction{
		Name:  "seed",
		Value: FuncAI64R(rand.Seed),
	},
	"read": &gslang.UserFunction{
		Name: "read",
		Value: func(args ...gslang.Object) (ret gslang.Object, err error) {
			if len(args) != 1 {
				return nil, gslang.ErrWrongNumArguments
			}
			y1, ok := args[0].(*gslang.Bytes)
			if !ok {
				return nil, gslang.ErrInvalidArgumentType{
					Name:     "first",
					Expected: "bytes",
					Found:    args[0].TypeName(),
				}
			}
			res, err := rand.Read(y1.Value)
			if err != nil {
				ret = wrapError(err)
				return
			}
			return &gslang.Int{Value: int64(res)}, nil
		},
	},
	"rand": &gslang.UserFunction{
		Name: "rand",
		Value: func(args ...gslang.Object) (gslang.Object, error) {
			if len(args) != 1 {
				return nil, gslang.ErrWrongNumArguments
			}
			i1, ok := gslang.ToInt64(args[0])
			if !ok {
				return nil, gslang.ErrInvalidArgumentType{
					Name:     "first",
					Expected: "int(compatible)",
					Found:    args[0].TypeName(),
				}
			}
			src := rand.NewSource(i1)
			return randRand(rand.New(src)), nil
		},
	},
}

func randRand(r *rand.Rand) *gslang.Map {
	return &gslang.Map{
		Value: map[string]gslang.Object{
			"int": &gslang.UserFunction{
				Name:  "int",
				Value: FuncARI64(r.Int63),
			},
			"float": &gslang.UserFunction{
				Name:  "float",
				Value: FuncARF(r.Float64),
			},
			"intn": &gslang.UserFunction{
				Name:  "intn",
				Value: FuncAI64RI64(r.Int63n),
			},
			"exp_float": &gslang.UserFunction{
				Name:  "exp_float",
				Value: FuncARF(r.ExpFloat64),
			},
			"norm_float": &gslang.UserFunction{
				Name:  "norm_float",
				Value: FuncARF(r.NormFloat64),
			},
			"perm": &gslang.UserFunction{
				Name:  "perm",
				Value: FuncAIRIs(r.Perm),
			},
			"seed": &gslang.UserFunction{
				Name:  "seed",
				Value: FuncAI64R(r.Seed),
			},
			"read": &gslang.UserFunction{
				Name: "read",
				Value: func(args ...gslang.Object) (
					ret gslang.Object,
					err error,
				) {
					if len(args) != 1 {
						return nil, gslang.ErrWrongNumArguments
					}
					y1, ok := args[0].(*gslang.Bytes)
					if !ok {
						return nil, gslang.ErrInvalidArgumentType{
							Name:     "first",
							Expected: "bytes",
							Found:    args[0].TypeName(),
						}
					}
					res, err := r.Read(y1.Value)
					if err != nil {
						ret = wrapError(err)
						return
					}
					return &gslang.Int{Value: int64(res)}, nil
				},
			},
		},
	}
}
