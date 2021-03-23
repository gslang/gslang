package stdlib

import (
	"os"
	"syscall"

	"github.com/gslang/gslang"
)

func makeOSProcessState(state *os.ProcessState) *gslang.Map {
	return &gslang.Map{
		Value: map[string]gslang.Object{
			"exited": &gslang.UserFunction{
				Name:  "exited",
				Value: FuncARB(state.Exited),
			},
			"pid": &gslang.UserFunction{
				Name:  "pid",
				Value: FuncARI(state.Pid),
			},
			"string": &gslang.UserFunction{
				Name:  "string",
				Value: FuncARS(state.String),
			},
			"success": &gslang.UserFunction{
				Name:  "success",
				Value: FuncARB(state.Success),
			},
		},
	}
}

func makeOSProcess(proc *os.Process) *gslang.Map {
	return &gslang.Map{
		Value: map[string]gslang.Object{
			"kill": &gslang.UserFunction{
				Name:  "kill",
				Value: FuncARE(proc.Kill),
			},
			"release": &gslang.UserFunction{
				Name:  "release",
				Value: FuncARE(proc.Release),
			},
			"signal": &gslang.UserFunction{
				Name: "signal",
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
					return wrapError(proc.Signal(syscall.Signal(i1))), nil
				},
			},
			"wait": &gslang.UserFunction{
				Name: "wait",
				Value: func(args ...gslang.Object) (gslang.Object, error) {
					if len(args) != 0 {
						return nil, gslang.ErrWrongNumArguments
					}
					state, err := proc.Wait()
					if err != nil {
						return wrapError(err), nil
					}
					return makeOSProcessState(state), nil
				},
			},
		},
	}
}
