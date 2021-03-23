package stdlib

import (
	"os"

	"github.com/gslang/gslang"
)

func makeOSFile(file *os.File) *gslang.Map {
	return &gslang.Map{
		Value: map[string]gslang.Object{
			// chdir() => true/error
			"chdir": &gslang.UserFunction{
				Name:  "chdir",
				Value: FuncARE(file.Chdir),
			}, //
			// chown(uid int, gid int) => true/error
			"chown": &gslang.UserFunction{
				Name:  "chown",
				Value: FuncAIIRE(file.Chown),
			}, //
			// close() => error
			"close": &gslang.UserFunction{
				Name:  "close",
				Value: FuncARE(file.Close),
			}, //
			// name() => string
			"name": &gslang.UserFunction{
				Name:  "name",
				Value: FuncARS(file.Name),
			}, //
			// readdirnames(n int) => array(string)/error
			"readdirnames": &gslang.UserFunction{
				Name:  "readdirnames",
				Value: FuncAIRSsE(file.Readdirnames),
			}, //
			// sync() => error
			"sync": &gslang.UserFunction{
				Name:  "sync",
				Value: FuncARE(file.Sync),
			}, //
			// write(bytes) => int/error
			"write": &gslang.UserFunction{
				Name:  "write",
				Value: FuncAYRIE(file.Write),
			}, //
			// write(string) => int/error
			"write_string": &gslang.UserFunction{
				Name:  "write_string",
				Value: FuncASRIE(file.WriteString),
			}, //
			// read(bytes) => int/error
			"read": &gslang.UserFunction{
				Name:  "read",
				Value: FuncAYRIE(file.Read),
			}, //
			// chmod(mode int) => error
			"chmod": &gslang.UserFunction{
				Name: "chmod",
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
					return wrapError(file.Chmod(os.FileMode(i1))), nil
				},
			},
			// seek(offset int, whence int) => int/error
			"seek": &gslang.UserFunction{
				Name: "seek",
				Value: func(args ...gslang.Object) (gslang.Object, error) {
					if len(args) != 2 {
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
					i2, ok := gslang.ToInt(args[1])
					if !ok {
						return nil, gslang.ErrInvalidArgumentType{
							Name:     "second",
							Expected: "int(compatible)",
							Found:    args[1].TypeName(),
						}
					}
					res, err := file.Seek(i1, i2)
					if err != nil {
						return wrapError(err), nil
					}
					return &gslang.Int{Value: res}, nil
				},
			},
			// stat() => imap(fileinfo)/error
			"stat": &gslang.UserFunction{
				Name: "stat",
				Value: func(args ...gslang.Object) (gslang.Object, error) {
					if len(args) != 0 {
						return nil, gslang.ErrWrongNumArguments
					}
					return osStat(&gslang.String{Value: file.Name()})
				},
			},
		},
	}
}
