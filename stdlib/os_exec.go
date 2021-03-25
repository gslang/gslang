package stdlib

import (
	"os/exec"

	"github.com/gslang/gslang"
)

func makeOSExecCommand(cmd *exec.Cmd) *gslang.Map {
	return &gslang.Map{
		Value: map[string]gslang.Object{
			// combined_output() => bytes/error
			"combined_output": &gslang.UserFunction{
				Name:  "combined_output",
				Value: FuncARYE(cmd.CombinedOutput),
			},
			// output() => bytes/error
			"output": &gslang.UserFunction{
				Name:  "output",
				Value: FuncARYE(cmd.Output),
			}, //
			// run() => error
			"run": &gslang.UserFunction{
				Name:  "run",
				Value: FuncARE(cmd.Run),
			}, //
			// start() => error
			"start": &gslang.UserFunction{
				Name:  "start",
				Value: FuncARE(cmd.Start),
			}, //
			// wait() => error
			"wait": &gslang.UserFunction{
				Name:  "wait",
				Value: FuncARE(cmd.Wait),
			}, //
			// set_path(path string)
			"set_path": &gslang.UserFunction{
				Name: "set_path",
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
					cmd.Path = s1
					return gslang.NilValue, nil
				},
			},
			// set_dir(dir string)
			"set_dir": &gslang.UserFunction{
				Name: "set_dir",
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
					cmd.Dir = s1
					return gslang.NilValue, nil
				},
			},
			// set_env(env array(string))
			"set_env": &gslang.UserFunction{
				Name: "set_env",
				Value: func(args ...gslang.Object) (gslang.Object, error) {
					if len(args) != 1 {
						return nil, gslang.ErrWrongNumArguments
					}

					var env []string
					var err error
					switch arg0 := args[0].(type) {
					case *gslang.Array:
						env, err = stringArray(arg0.Value, "first")
						if err != nil {
							return nil, err
						}
					default:
						return nil, gslang.ErrInvalidArgumentType{
							Name:     "first",
							Expected: "array",
							Found:    arg0.TypeName(),
						}
					}
					cmd.Env = env
					return gslang.NilValue, nil
				},
			},
			// process() => imap(process)
			"process": &gslang.UserFunction{
				Name: "process",
				Value: func(args ...gslang.Object) (gslang.Object, error) {
					if len(args) != 0 {
						return nil, gslang.ErrWrongNumArguments
					}
					return makeOSProcess(cmd.Process), nil
				},
			},
		},
	}
}
