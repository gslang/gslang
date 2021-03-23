package stdlib

import (
	"fmt"

	"github.com/gslang/gslang"
)

var fmtModule = map[string]gslang.Object{
	"print":   &gslang.UserFunction{Name: "print", Value: fmtPrint},
	"printf":  &gslang.UserFunction{Name: "printf", Value: fmtPrintf},
	"println": &gslang.UserFunction{Name: "println", Value: fmtPrintln},
	"sprintf": &gslang.UserFunction{Name: "sprintf", Value: fmtSprintf},
}

func fmtPrint(args ...gslang.Object) (ret gslang.Object, err error) {
	printArgs, err := getPrintArgs(args...)
	if err != nil {
		return nil, err
	}
	_, _ = fmt.Print(printArgs...)
	return nil, nil
}

func fmtPrintf(args ...gslang.Object) (ret gslang.Object, err error) {
	numArgs := len(args)
	if numArgs == 0 {
		return nil, gslang.ErrWrongNumArguments
	}

	format, ok := args[0].(*gslang.String)
	if !ok {
		return nil, gslang.ErrInvalidArgumentType{
			Name:     "format",
			Expected: "string",
			Found:    args[0].TypeName(),
		}
	}
	if numArgs == 1 {
		fmt.Print(format)
		return nil, nil
	}

	s, err := gslang.Format(format.Value, args[1:]...)
	if err != nil {
		return nil, err
	}
	fmt.Print(s)
	return nil, nil
}

func fmtPrintln(args ...gslang.Object) (ret gslang.Object, err error) {
	printArgs, err := getPrintArgs(args...)
	if err != nil {
		return nil, err
	}
	printArgs = append(printArgs, "\n")
	_, _ = fmt.Print(printArgs...)
	return nil, nil
}

func fmtSprintf(args ...gslang.Object) (ret gslang.Object, err error) {
	numArgs := len(args)
	if numArgs == 0 {
		return nil, gslang.ErrWrongNumArguments
	}

	format, ok := args[0].(*gslang.String)
	if !ok {
		return nil, gslang.ErrInvalidArgumentType{
			Name:     "format",
			Expected: "string",
			Found:    args[0].TypeName(),
		}
	}
	s, err := gslang.Format(format.Value, args[1:]...)
	if err != nil {
		return nil, err
	}
	return &gslang.String{Value: s}, nil
}

func getPrintArgs(args ...gslang.Object) ([]interface{}, error) {
	var printArgs []interface{}
	l := 0
	for _, arg := range args {
		s, _ := gslang.ToString(arg)
		slen := len(s)
		// make sure length does not exceed the limit
		if l+slen > gslang.MaxStringLen {
			return nil, gslang.ErrStringLimit
		}
		l += slen
		printArgs = append(printArgs, s)
	}
	return printArgs, nil
}
