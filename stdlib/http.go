package stdlib

import (
	"time"
	"bytes"
	"net/http"
	"io/ioutil"
	"encoding/json"
	"github.com/gslang/gslang"
)

var httpModule = map[string]gslang.Object{
	"request": &gslang.UserFunction{
		Name:  "request",
		Value: httpRequest,
	},
}

func httpRequest(args ...gslang.Object) (gslang.Object, error) {
	if len(args) != 2 {
		return nil, gslang.ErrWrongNumArguments
	}
	s1, ok := gslang.ToString(args[0])
	if !ok {
		return nil, gslang.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "string(compatible)",
			Found:    args[0].TypeName(),
		}
	}
	s2, ok := gslang.ToString(args[1])
	if !ok {
		return nil, gslang.ErrInvalidArgumentType{
			Name:     "second",
			Expected: "string(compatible)",
			Found:    args[1].TypeName(),
		}
	}
	var req *http.Request
	var res *http.Response
	var err error
	req, err = http.NewRequest(s1, s2, nil)
	if err != nil {
		return nil, err
	}
	cli := &http.Client{
		Timeout: time.Duration(30) * time.Second,
	}
	return &gslang.Map{
		Value: map[string]gslang.Object{
			"set_timeout": &gslang.UserFunction{
				Name:  "set_timeout",
				Value: func(args ...gslang.Object) (gslang.Object, error) {
					if len(args) != 1 {
						return nil, gslang.ErrWrongNumArguments
					}
					i1, ok := gslang.ToInt(args[0])
					if !ok {
						return nil, gslang.ErrInvalidArgumentType{
							Name:     "first",
							Expected: "int(compatible)",
							Found:    args[0].TypeName(),
						}
					}
					
					cli.Timeout = time.Duration(int(i1)) * time.Second
					return nil, nil
				},
			},
			"set_header": &gslang.UserFunction{
				Name:  "set_header",
				Value: func(args ...gslang.Object) (gslang.Object, error) {
					if len(args) != 2 {
						return nil, gslang.ErrWrongNumArguments
					}
					s1, ok := gslang.ToString(args[0])
					if !ok {
						return nil, gslang.ErrInvalidArgumentType{
							Name:     "first",
							Expected: "string(compatible)",
							Found:    args[0].TypeName(),
						}
					}
					s2, ok := gslang.ToString(args[1])
					if !ok {
						return nil, gslang.ErrInvalidArgumentType{
							Name:     "second",
							Expected: "string(compatible)",
							Found:    args[1].TypeName(),
						}
					}
					req.Header.Set(s1,s2)
					return nil, nil
				},
			},
			"set_body": &gslang.UserFunction{
				Name:  "set_body",
				Value: func(args ...gslang.Object) (gslang.Object, error) {
					if len(args) != 1 {
						return nil, gslang.ErrWrongNumArguments
					}
					m1 := gslang.ToInterface(args[0])
					if m1 != nil {
						body := make(map[string]interface{})
						if m2, ok := m1.(map[string]interface{}); ok {
							for key, val := range m2 {
								body[key] = val
							}
						}
						byte, err := json.Marshal(body)
						if err != nil {
							return nil, err
						}
						req.Body = ioutil.NopCloser(bytes.NewReader(byte))
					}
					return nil, nil
				},
			},
			"get_response": &gslang.UserFunction{
				Name:  "get_response",
				Value: func(args ...gslang.Object) (gslang.Object, error) {
					res, err = cli.Do(req)
					if err != nil {
						return nil, err
					}
					return &gslang.Map{
						Value: map[string]gslang.Object{
							"get_header": &gslang.UserFunction{
								Name:  "get_header",
								Value: func(args ...gslang.Object) (gslang.Object, error) {
									if len(args) != 1 {
										return nil, gslang.ErrWrongNumArguments
									}
									s1, ok := gslang.ToString(args[0])
									if !ok {
										return nil, gslang.ErrInvalidArgumentType{
											Name:     "first",
											Expected: "string(compatible)",
											Found:    args[0].TypeName(),
										}
									}
									header := res.Header.Get(s1)
									return &gslang.String{Value: header}, nil
								},
							},
							"get_body": &gslang.UserFunction{
								Name:  "get_body",
								Value: func(args ...gslang.Object) (gslang.Object, error) {
									var body []byte
									body, err = ioutil.ReadAll(res.Body)
									defer res.Body.Close()
									if err != nil {
										return nil, err
									}
									return &gslang.String{Value: string(body)}, nil
								},
							},
						},
					}, nil 
				},
			},
		},
	}, nil
}