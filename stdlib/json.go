package stdlib

import (
	"bytes"
	gojson "encoding/json"

	"github.com/gslang/gslang"
	"github.com/gslang/gslang/stdlib/json"
)

var jsonModule = map[string]gslang.Object{
	"decode": &gslang.UserFunction{
		Name:  "decode",
		Value: jsonDecode,
	},
	"encode": &gslang.UserFunction{
		Name:  "encode",
		Value: jsonEncode,
	},
	"indent": &gslang.UserFunction{
		Name:  "encode",
		Value: jsonIndent,
	},
	"html_escape": &gslang.UserFunction{
		Name:  "html_escape",
		Value: jsonHTMLEscape,
	},
}

func jsonDecode(args ...gslang.Object) (ret gslang.Object, err error) {
	if len(args) != 1 {
		return nil, gslang.ErrWrongNumArguments
	}

	switch o := args[0].(type) {
	case *gslang.Bytes:
		v, err := json.Decode(o.Value)
		if err != nil {
			return &gslang.Error{
				Value: &gslang.String{Value: err.Error()},
			}, nil
		}
		return v, nil
	case *gslang.String:
		v, err := json.Decode([]byte(o.Value))
		if err != nil {
			return &gslang.Error{
				Value: &gslang.String{Value: err.Error()},
			}, nil
		}
		return v, nil
	default:
		return nil, gslang.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "bytes/string",
			Found:    args[0].TypeName(),
		}
	}
}

func jsonEncode(args ...gslang.Object) (ret gslang.Object, err error) {
	if len(args) != 1 {
		return nil, gslang.ErrWrongNumArguments
	}

	b, err := json.Encode(args[0])
	if err != nil {
		return &gslang.Error{Value: &gslang.String{Value: err.Error()}}, nil
	}

	return &gslang.Bytes{Value: b}, nil
}

func jsonIndent(args ...gslang.Object) (ret gslang.Object, err error) {
	if len(args) != 3 {
		return nil, gslang.ErrWrongNumArguments
	}

	prefix, ok := gslang.ToString(args[1])
	if !ok {
		return nil, gslang.ErrInvalidArgumentType{
			Name:     "prefix",
			Expected: "string(compatible)",
			Found:    args[1].TypeName(),
		}
	}

	indent, ok := gslang.ToString(args[2])
	if !ok {
		return nil, gslang.ErrInvalidArgumentType{
			Name:     "indent",
			Expected: "string(compatible)",
			Found:    args[2].TypeName(),
		}
	}

	switch o := args[0].(type) {
	case *gslang.Bytes:
		var dst bytes.Buffer
		err := gojson.Indent(&dst, o.Value, prefix, indent)
		if err != nil {
			return &gslang.Error{
				Value: &gslang.String{Value: err.Error()},
			}, nil
		}
		return &gslang.Bytes{Value: dst.Bytes()}, nil
	case *gslang.String:
		var dst bytes.Buffer
		err := gojson.Indent(&dst, []byte(o.Value), prefix, indent)
		if err != nil {
			return &gslang.Error{
				Value: &gslang.String{Value: err.Error()},
			}, nil
		}
		return &gslang.Bytes{Value: dst.Bytes()}, nil
	default:
		return nil, gslang.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "bytes/string",
			Found:    args[0].TypeName(),
		}
	}
}

func jsonHTMLEscape(args ...gslang.Object) (ret gslang.Object, err error) {
	if len(args) != 1 {
		return nil, gslang.ErrWrongNumArguments
	}

	switch o := args[0].(type) {
	case *gslang.Bytes:
		var dst bytes.Buffer
		gojson.HTMLEscape(&dst, o.Value)
		return &gslang.Bytes{Value: dst.Bytes()}, nil
	case *gslang.String:
		var dst bytes.Buffer
		gojson.HTMLEscape(&dst, []byte(o.Value))
		return &gslang.Bytes{Value: dst.Bytes()}, nil
	default:
		return nil, gslang.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "bytes/string",
			Found:    args[0].TypeName(),
		}
	}
}
