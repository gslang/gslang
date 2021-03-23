package stdlib

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"os/exec"

	"github.com/gslang/gslang"
)

var osModule = map[string]gslang.Object{
	"o_rdonly":            &gslang.Int{Value: int64(os.O_RDONLY)},
	"o_wronly":            &gslang.Int{Value: int64(os.O_WRONLY)},
	"o_rdwr":              &gslang.Int{Value: int64(os.O_RDWR)},
	"o_append":            &gslang.Int{Value: int64(os.O_APPEND)},
	"o_create":            &gslang.Int{Value: int64(os.O_CREATE)},
	"o_excl":              &gslang.Int{Value: int64(os.O_EXCL)},
	"o_sync":              &gslang.Int{Value: int64(os.O_SYNC)},
	"o_trunc":             &gslang.Int{Value: int64(os.O_TRUNC)},
	"mode_dir":            &gslang.Int{Value: int64(os.ModeDir)},
	"mode_append":         &gslang.Int{Value: int64(os.ModeAppend)},
	"mode_exclusive":      &gslang.Int{Value: int64(os.ModeExclusive)},
	"mode_temporary":      &gslang.Int{Value: int64(os.ModeTemporary)},
	"mode_symlink":        &gslang.Int{Value: int64(os.ModeSymlink)},
	"mode_device":         &gslang.Int{Value: int64(os.ModeDevice)},
	"mode_named_pipe":     &gslang.Int{Value: int64(os.ModeNamedPipe)},
	"mode_socket":         &gslang.Int{Value: int64(os.ModeSocket)},
	"mode_setuid":         &gslang.Int{Value: int64(os.ModeSetuid)},
	"mode_setgui":         &gslang.Int{Value: int64(os.ModeSetgid)},
	"mode_char_device":    &gslang.Int{Value: int64(os.ModeCharDevice)},
	"mode_sticky":         &gslang.Int{Value: int64(os.ModeSticky)},
	"mode_type":           &gslang.Int{Value: int64(os.ModeType)},
	"mode_perm":           &gslang.Int{Value: int64(os.ModePerm)},
	"path_separator":      &gslang.Char{Value: os.PathSeparator},
	"path_list_separator": &gslang.Char{Value: os.PathListSeparator},
	"dev_null":            &gslang.String{Value: os.DevNull},
	"seek_set":            &gslang.Int{Value: int64(io.SeekStart)},
	"seek_cur":            &gslang.Int{Value: int64(io.SeekCurrent)},
	"seek_end":            &gslang.Int{Value: int64(io.SeekEnd)},
	"args": &gslang.UserFunction{
		Name:  "args",
		Value: osArgs,
	}, // args() => array(string)
	"chdir": &gslang.UserFunction{
		Name:  "chdir",
		Value: FuncASRE(os.Chdir),
	}, // chdir(dir string) => error
	"chmod": osFuncASFmRE("chmod", os.Chmod), // chmod(name string, mode int) => error
	"chown": &gslang.UserFunction{
		Name:  "chown",
		Value: FuncASIIRE(os.Chown),
	}, // chown(name string, uid int, gid int) => error
	"clearenv": &gslang.UserFunction{
		Name:  "clearenv",
		Value: FuncAR(os.Clearenv),
	}, // clearenv()
	"environ": &gslang.UserFunction{
		Name:  "environ",
		Value: FuncARSs(os.Environ),
	}, // environ() => array(string)
	"exit": &gslang.UserFunction{
		Name:  "exit",
		Value: FuncAIR(os.Exit),
	}, // exit(code int)
	"expand_env": &gslang.UserFunction{
		Name:  "expand_env",
		Value: osExpandEnv,
	}, // expand_env(s string) => string
	"getegid": &gslang.UserFunction{
		Name:  "getegid",
		Value: FuncARI(os.Getegid),
	}, // getegid() => int
	"getenv": &gslang.UserFunction{
		Name:  "getenv",
		Value: FuncASRS(os.Getenv),
	}, // getenv(s string) => string
	"geteuid": &gslang.UserFunction{
		Name:  "geteuid",
		Value: FuncARI(os.Geteuid),
	}, // geteuid() => int
	"getgid": &gslang.UserFunction{
		Name:  "getgid",
		Value: FuncARI(os.Getgid),
	}, // getgid() => int
	"getgroups": &gslang.UserFunction{
		Name:  "getgroups",
		Value: FuncARIsE(os.Getgroups),
	}, // getgroups() => array(string)/error
	"getpagesize": &gslang.UserFunction{
		Name:  "getpagesize",
		Value: FuncARI(os.Getpagesize),
	}, // getpagesize() => int
	"getpid": &gslang.UserFunction{
		Name:  "getpid",
		Value: FuncARI(os.Getpid),
	}, // getpid() => int
	"getppid": &gslang.UserFunction{
		Name:  "getppid",
		Value: FuncARI(os.Getppid),
	}, // getppid() => int
	"getuid": &gslang.UserFunction{
		Name:  "getuid",
		Value: FuncARI(os.Getuid),
	}, // getuid() => int
	"getwd": &gslang.UserFunction{
		Name:  "getwd",
		Value: FuncARSE(os.Getwd),
	}, // getwd() => string/error
	"hostname": &gslang.UserFunction{
		Name:  "hostname",
		Value: FuncARSE(os.Hostname),
	}, // hostname() => string/error
	"lchown": &gslang.UserFunction{
		Name:  "lchown",
		Value: FuncASIIRE(os.Lchown),
	}, // lchown(name string, uid int, gid int) => error
	"link": &gslang.UserFunction{
		Name:  "link",
		Value: FuncASSRE(os.Link),
	}, // link(oldname string, newname string) => error
	"lookup_env": &gslang.UserFunction{
		Name:  "lookup_env",
		Value: osLookupEnv,
	}, // lookup_env(key string) => string/false
	"mkdir":     osFuncASFmRE("mkdir", os.Mkdir),        // mkdir(name string, perm int) => error
	"mkdir_all": osFuncASFmRE("mkdir_all", os.MkdirAll), // mkdir_all(name string, perm int) => error
	"readlink": &gslang.UserFunction{
		Name:  "readlink",
		Value: FuncASRSE(os.Readlink),
	}, // readlink(name string) => string/error
	"remove": &gslang.UserFunction{
		Name:  "remove",
		Value: FuncASRE(os.Remove),
	}, // remove(name string) => error
	"remove_all": &gslang.UserFunction{
		Name:  "remove_all",
		Value: FuncASRE(os.RemoveAll),
	}, // remove_all(name string) => error
	"rename": &gslang.UserFunction{
		Name:  "rename",
		Value: FuncASSRE(os.Rename),
	}, // rename(oldpath string, newpath string) => error
	"setenv": &gslang.UserFunction{
		Name:  "setenv",
		Value: FuncASSRE(os.Setenv),
	}, // setenv(key string, value string) => error
	"symlink": &gslang.UserFunction{
		Name:  "symlink",
		Value: FuncASSRE(os.Symlink),
	}, // symlink(oldname string newname string) => error
	"temp_dir": &gslang.UserFunction{
		Name:  "temp_dir",
		Value: FuncARS(os.TempDir),
	}, // temp_dir() => string
	"truncate": &gslang.UserFunction{
		Name:  "truncate",
		Value: FuncASI64RE(os.Truncate),
	}, // truncate(name string, size int) => error
	"unsetenv": &gslang.UserFunction{
		Name:  "unsetenv",
		Value: FuncASRE(os.Unsetenv),
	}, // unsetenv(key string) => error
	"create": &gslang.UserFunction{
		Name:  "create",
		Value: osCreate,
	}, // create(name string) => imap(file)/error
	"open": &gslang.UserFunction{
		Name:  "open",
		Value: osOpen,
	}, // open(name string) => imap(file)/error
	"open_file": &gslang.UserFunction{
		Name:  "open_file",
		Value: osOpenFile,
	}, // open_file(name string, flag int, perm int) => imap(file)/error
	"find_process": &gslang.UserFunction{
		Name:  "find_process",
		Value: osFindProcess,
	}, // find_process(pid int) => imap(process)/error
	"start_process": &gslang.UserFunction{
		Name:  "start_process",
		Value: osStartProcess,
	}, // start_process(name string, argv array(string), dir string, env array(string)) => imap(process)/error
	"exec_look_path": &gslang.UserFunction{
		Name:  "exec_look_path",
		Value: FuncASRSE(exec.LookPath),
	}, // exec_look_path(file) => string/error
	"exec": &gslang.UserFunction{
		Name:  "exec",
		Value: osExec,
	}, // exec(name, args...) => command
	"stat": &gslang.UserFunction{
		Name:  "stat",
		Value: osStat,
	}, // stat(name) => imap(fileinfo)/error
	"read_file": &gslang.UserFunction{
		Name:  "read_file",
		Value: osReadFile,
	}, // readfile(name) => array(byte)/error
}

func osReadFile(args ...gslang.Object) (ret gslang.Object, err error) {
	if len(args) != 1 {
		return nil, gslang.ErrWrongNumArguments
	}
	fname, ok := gslang.ToString(args[0])
	if !ok {
		return nil, gslang.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "string(compatible)",
			Found:    args[0].TypeName(),
		}
	}
	bytes, err := ioutil.ReadFile(fname)
	if err != nil {
		return wrapError(err), nil
	}
	if len(bytes) > gslang.MaxBytesLen {
		return nil, gslang.ErrBytesLimit
	}
	return &gslang.Bytes{Value: bytes}, nil
}

func osStat(args ...gslang.Object) (ret gslang.Object, err error) {
	if len(args) != 1 {
		return nil, gslang.ErrWrongNumArguments
	}
	fname, ok := gslang.ToString(args[0])
	if !ok {
		return nil, gslang.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "string(compatible)",
			Found:    args[0].TypeName(),
		}
	}
	stat, err := os.Stat(fname)
	if err != nil {
		return wrapError(err), nil
	}
	fstat := &gslang.Map{
		Value: map[string]gslang.Object{
			"name":  &gslang.String{Value: stat.Name()},
			"mtime": &gslang.Time{Value: stat.ModTime()},
			"size":  &gslang.Int{Value: stat.Size()},
			"mode":  &gslang.Int{Value: int64(stat.Mode())},
		},
	}
	if stat.IsDir() {
		fstat.Value["directory"] = gslang.TrueValue
	} else {
		fstat.Value["directory"] = gslang.FalseValue
	}
	return fstat, nil
}

func osCreate(args ...gslang.Object) (gslang.Object, error) {
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
	res, err := os.Create(s1)
	if err != nil {
		return wrapError(err), nil
	}
	return makeOSFile(res), nil
}

func osOpen(args ...gslang.Object) (gslang.Object, error) {
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
	res, err := os.Open(s1)
	if err != nil {
		return wrapError(err), nil
	}
	return makeOSFile(res), nil
}

func osOpenFile(args ...gslang.Object) (gslang.Object, error) {
	if len(args) != 3 {
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
	i2, ok := gslang.ToInt(args[1])
	if !ok {
		return nil, gslang.ErrInvalidArgumentType{
			Name:     "second",
			Expected: "int(compatible)",
			Found:    args[1].TypeName(),
		}
	}
	i3, ok := gslang.ToInt(args[2])
	if !ok {
		return nil, gslang.ErrInvalidArgumentType{
			Name:     "third",
			Expected: "int(compatible)",
			Found:    args[2].TypeName(),
		}
	}
	res, err := os.OpenFile(s1, i2, os.FileMode(i3))
	if err != nil {
		return wrapError(err), nil
	}
	return makeOSFile(res), nil
}

func osArgs(args ...gslang.Object) (gslang.Object, error) {
	if len(args) != 0 {
		return nil, gslang.ErrWrongNumArguments
	}
	arr := &gslang.Array{}
	for _, osArg := range os.Args {
		if len(osArg) > gslang.MaxStringLen {
			return nil, gslang.ErrStringLimit
		}
		arr.Value = append(arr.Value, &gslang.String{Value: osArg})
	}
	return arr, nil
}

func osFuncASFmRE(
	name string,
	fn func(string, os.FileMode) error,
) *gslang.UserFunction {
	return &gslang.UserFunction{
		Name: name,
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
			i2, ok := gslang.ToInt64(args[1])
			if !ok {
				return nil, gslang.ErrInvalidArgumentType{
					Name:     "second",
					Expected: "int(compatible)",
					Found:    args[1].TypeName(),
				}
			}
			return wrapError(fn(s1, os.FileMode(i2))), nil
		},
	}
}

func osLookupEnv(args ...gslang.Object) (gslang.Object, error) {
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
	res, ok := os.LookupEnv(s1)
	if !ok {
		return gslang.FalseValue, nil
	}
	if len(res) > gslang.MaxStringLen {
		return nil, gslang.ErrStringLimit
	}
	return &gslang.String{Value: res}, nil
}

func osExpandEnv(args ...gslang.Object) (gslang.Object, error) {
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
	var vlen int
	var failed bool
	s := os.Expand(s1, func(k string) string {
		if failed {
			return ""
		}
		v := os.Getenv(k)

		// this does not count the other texts that are not being replaced
		// but the code checks the final length at the end
		vlen += len(v)
		if vlen > gslang.MaxStringLen {
			failed = true
			return ""
		}
		return v
	})
	if failed || len(s) > gslang.MaxStringLen {
		return nil, gslang.ErrStringLimit
	}
	return &gslang.String{Value: s}, nil
}

func osExec(args ...gslang.Object) (gslang.Object, error) {
	if len(args) == 0 {
		return nil, gslang.ErrWrongNumArguments
	}
	name, ok := gslang.ToString(args[0])
	if !ok {
		return nil, gslang.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "string(compatible)",
			Found:    args[0].TypeName(),
		}
	}
	var execArgs []string
	for idx, arg := range args[1:] {
		execArg, ok := gslang.ToString(arg)
		if !ok {
			return nil, gslang.ErrInvalidArgumentType{
				Name:     fmt.Sprintf("args[%d]", idx),
				Expected: "string(compatible)",
				Found:    args[1+idx].TypeName(),
			}
		}
		execArgs = append(execArgs, execArg)
	}
	return makeOSExecCommand(exec.Command(name, execArgs...)), nil
}

func osFindProcess(args ...gslang.Object) (gslang.Object, error) {
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
	proc, err := os.FindProcess(i1)
	if err != nil {
		return wrapError(err), nil
	}
	return makeOSProcess(proc), nil
}

func osStartProcess(args ...gslang.Object) (gslang.Object, error) {
	if len(args) != 4 {
		return nil, gslang.ErrWrongNumArguments
	}
	name, ok := gslang.ToString(args[0])
	if !ok {
		return nil, gslang.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "string(compatible)",
			Found:    args[0].TypeName(),
		}
	}
	var argv []string
	var err error
	switch arg1 := args[1].(type) {
	case *gslang.Array:
		argv, err = stringArray(arg1.Value, "second")
		if err != nil {
			return nil, err
		}
	default:
		return nil, gslang.ErrInvalidArgumentType{
			Name:     "second",
			Expected: "array",
			Found:    arg1.TypeName(),
		}
	}

	dir, ok := gslang.ToString(args[2])
	if !ok {
		return nil, gslang.ErrInvalidArgumentType{
			Name:     "third",
			Expected: "string(compatible)",
			Found:    args[2].TypeName(),
		}
	}

	var env []string
	switch arg3 := args[3].(type) {
	case *gslang.Array:
		env, err = stringArray(arg3.Value, "fourth")
		if err != nil {
			return nil, err
		}
	default:
		return nil, gslang.ErrInvalidArgumentType{
			Name:     "fourth",
			Expected: "array",
			Found:    arg3.TypeName(),
		}
	}

	proc, err := os.StartProcess(name, argv, &os.ProcAttr{
		Dir: dir,
		Env: env,
	})
	if err != nil {
		return wrapError(err), nil
	}
	return makeOSProcess(proc), nil
}

func stringArray(arr []gslang.Object, argName string) ([]string, error) {
	var sarr []string
	for idx, elem := range arr {
		str, ok := elem.(*gslang.String)
		if !ok {
			return nil, gslang.ErrInvalidArgumentType{
				Name:     fmt.Sprintf("%s[%d]", argName, idx),
				Expected: "string",
				Found:    elem.TypeName(),
			}
		}
		sarr = append(sarr, str.Value)
	}
	return sarr, nil
}
