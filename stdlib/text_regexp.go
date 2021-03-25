package stdlib

import (
	"regexp"

	"github.com/gslang/gslang"
)

func makeTextRegexp(re *regexp.Regexp) *gslang.Map {
	return &gslang.Map{
		Value: map[string]gslang.Object{
			// match(text) => bool
			"match": &gslang.UserFunction{
				Value: func(args ...gslang.Object) (
					ret gslang.Object,
					err error,
				) {
					if len(args) != 1 {
						err = gslang.ErrWrongNumArguments
						return
					}

					s1, ok := gslang.ToString(args[0])
					if !ok {
						err = gslang.ErrInvalidArgumentType{
							Name:     "first",
							Expected: "string(compatible)",
							Found:    args[0].TypeName(),
						}
						return
					}

					if re.MatchString(s1) {
						ret = gslang.TrueValue
					} else {
						ret = gslang.FalseValue
					}

					return
				},
			},

			// find(text) 			=> array(array({text:,begin:,end:}))/nil
			// find(text, maxCount) => array(array({text:,begin:,end:}))/nil
			"find": &gslang.UserFunction{
				Value: func(args ...gslang.Object) (
					ret gslang.Object,
					err error,
				) {
					numArgs := len(args)
					if numArgs != 1 && numArgs != 2 {
						err = gslang.ErrWrongNumArguments
						return
					}

					s1, ok := gslang.ToString(args[0])
					if !ok {
						err = gslang.ErrInvalidArgumentType{
							Name:     "first",
							Expected: "string(compatible)",
							Found:    args[0].TypeName(),
						}
						return
					}

					if numArgs == 1 {
						m := re.FindStringSubmatchIndex(s1)
						if m == nil {
							ret = gslang.NilValue
							return
						}

						arr := &gslang.Array{}
						for i := 0; i < len(m); i += 2 {
							arr.Value = append(arr.Value,
								&gslang.Map{
									Value: map[string]gslang.Object{
										"text": &gslang.String{
											Value: s1[m[i]:m[i+1]],
										},
										"begin": &gslang.Int{
											Value: int64(m[i]),
										},
										"end": &gslang.Int{
											Value: int64(m[i+1]),
										},
									}})
						}

						ret = &gslang.Array{Value: []gslang.Object{arr}}

						return
					}

					i2, ok := gslang.ToInt(args[1])
					if !ok {
						err = gslang.ErrInvalidArgumentType{
							Name:     "second",
							Expected: "int(compatible)",
							Found:    args[1].TypeName(),
						}
						return
					}
					m := re.FindAllStringSubmatchIndex(s1, i2)
					if m == nil {
						ret = gslang.NilValue
						return
					}

					arr := &gslang.Array{}
					for _, m := range m {
						subMatch := &gslang.Array{}
						for i := 0; i < len(m); i += 2 {
							subMatch.Value = append(subMatch.Value,
								&gslang.Map{
									Value: map[string]gslang.Object{
										"text": &gslang.String{
											Value: s1[m[i]:m[i+1]],
										},
										"begin": &gslang.Int{
											Value: int64(m[i]),
										},
										"end": &gslang.Int{
											Value: int64(m[i+1]),
										},
									}})
						}

						arr.Value = append(arr.Value, subMatch)
					}

					ret = arr

					return
				},
			},

			// replace(src, repl) => string
			"replace": &gslang.UserFunction{
				Value: func(args ...gslang.Object) (
					ret gslang.Object,
					err error,
				) {
					if len(args) != 2 {
						err = gslang.ErrWrongNumArguments
						return
					}

					s1, ok := gslang.ToString(args[0])
					if !ok {
						err = gslang.ErrInvalidArgumentType{
							Name:     "first",
							Expected: "string(compatible)",
							Found:    args[0].TypeName(),
						}
						return
					}

					s2, ok := gslang.ToString(args[1])
					if !ok {
						err = gslang.ErrInvalidArgumentType{
							Name:     "second",
							Expected: "string(compatible)",
							Found:    args[1].TypeName(),
						}
						return
					}

					s, ok := doTextRegexpReplace(re, s1, s2)
					if !ok {
						return nil, gslang.ErrStringLimit
					}

					ret = &gslang.String{Value: s}

					return
				},
			},

			// split(text) 			 => array(string)
			// split(text, maxCount) => array(string)
			"split": &gslang.UserFunction{
				Value: func(args ...gslang.Object) (
					ret gslang.Object,
					err error,
				) {
					numArgs := len(args)
					if numArgs != 1 && numArgs != 2 {
						err = gslang.ErrWrongNumArguments
						return
					}

					s1, ok := gslang.ToString(args[0])
					if !ok {
						err = gslang.ErrInvalidArgumentType{
							Name:     "first",
							Expected: "string(compatible)",
							Found:    args[0].TypeName(),
						}
						return
					}

					var i2 = -1
					if numArgs > 1 {
						i2, ok = gslang.ToInt(args[1])
						if !ok {
							err = gslang.ErrInvalidArgumentType{
								Name:     "second",
								Expected: "int(compatible)",
								Found:    args[1].TypeName(),
							}
							return
						}
					}

					arr := &gslang.Array{}
					for _, s := range re.Split(s1, i2) {
						arr.Value = append(arr.Value,
							&gslang.String{Value: s})
					}

					ret = arr

					return
				},
			},
		},
	}
}

// Size-limit checking implementation of regexp.ReplaceAllString.
func doTextRegexpReplace(re *regexp.Regexp, src, repl string) (string, bool) {
	idx := 0
	out := ""
	for _, m := range re.FindAllStringSubmatchIndex(src, -1) {
		var exp []byte
		exp = re.ExpandString(exp, repl, src, m)
		if len(out)+m[0]-idx+len(exp) > gslang.MaxStringLen {
			return "", false
		}
		out += src[idx:m[0]] + string(exp)
		idx = m[1]
	}
	if idx < len(src) {
		if len(out)+len(src)-idx > gslang.MaxStringLen {
			return "", false
		}
		out += src[idx:]
	}
	return out, true
}
