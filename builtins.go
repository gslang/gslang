package gslang

import (
    "sort"
	"time"
	"math/rand"
)

var builtinFuncs = []*BuiltinFunction{
	{
		Name:  "len",
		Value: builtinLen,
	},
	{
		Name:  "type",
		Value: builtinType,
	},
	{
		Name:  "range",
		Value: builtinRange,
	},
	{
		Name:  "format",
		Value: builtinFormat,
	},
	{
		Name:  "copy",
		Value: builtinCopy,
	},
	{
		Name:  "map_keys",
		Value: builtinMapKeys,
	},
	{
		Name:  "map_values",
		Value: builtinMapValues,
	},
	{
		Name:  "array_sort",
		Value: builtinArraySort,
	},
	{
		Name:  "array_rand",
		Value: builtinArrayRand,
	},
	{
		Name:  "array_push",
		Value: builtinArrayPush,
	},
	{
		Name:  "array_pop",
		Value: builtinArrayPop,
	},
	{
		Name:  "array_unshift",
		Value: builtinArrayUnShift,
	},
	{
		Name:  "array_shift",
		Value: builtinArrayShift,
	},
	{
		Name:  "array_reverse",
		Value: builtinArrayReverse,
	},
	{
		Name:  "array_unique",
		Value: builtinArrayUnique,
	},
	{
		Name:  "array_column",
		Value: builtinArrayColumn,
	},
	{
		Name:  "array_splice",
		Value: builtinArraySplice,
	},
	{
		Name:  "append",
		Value: builtinAppend,
	},
	{
		Name:  "delete",
		Value: builtinDelete,
	},
	{
		Name:  "exists",
		Value: builtinExists,
	},
	{
		Name:  "string",
		Value: builtinString,
	},
	{
		Name:  "int",
		Value: builtinInt,
	},
	{
		Name:  "bool",
		Value: builtinBool,
	},
	{
		Name:  "float",
		Value: builtinFloat,
	},
	{
		Name:  "char",
		Value: builtinChar,
	},
	{
		Name:  "bytes",
		Value: builtinBytes,
	},
	{
		Name:  "is_int",
		Value: builtinIsInt,
	},
	{
		Name:  "is_float",
		Value: builtinIsFloat,
	},
	{
		Name:  "is_string",
		Value: builtinIsString,
	},
	{
		Name:  "is_bool",
		Value: builtinIsBool,
	},
	{
		Name:  "is_char",
		Value: builtinIsChar,
	},
	{
		Name:  "is_bytes",
		Value: builtinIsBytes,
	},
	{
		Name:  "is_array",
		Value: builtinIsArray,
	},
	{
		Name:  "is_map",
		Value: builtinIsMap,
	},
	{
		Name:  "is_function",
		Value: builtinIsFunction,
	},
	{
		Name:  "is_callable",
		Value: builtinIsCallable,
	},
	{
		Name:  "is_iterable",
		Value: builtinIsIterable,
	},
	{
		Name:  "is_error",
		Value: builtinIsError,
	},
	{
		Name:  "is_nil",
		Value: builtinIsNil,
	},
}

// GetAllBuiltinFunctions returns all builtin function objects.
func GetAllBuiltinFunctions() []*BuiltinFunction {
	return append([]*BuiltinFunction{}, builtinFuncs...)
}

// len(obj object) => int
func builtinLen(args ...Object) (Object, error) {
	if len(args) != 1 {
		return nil, ErrWrongNumArguments
	}
	switch arg := args[0].(type) {
	case *Array:
		return &Int{Value: int64(len(arg.Value))}, nil
	case *String:
		return &Int{Value: int64(len(arg.Value))}, nil
	case *Bytes:
		return &Int{Value: int64(len(arg.Value))}, nil
	case *Map:
		return &Int{Value: int64(len(arg.Value))}, nil
	default:
		return nil, ErrInvalidArgumentType{
			Name:     "first",
			Expected: "array/string/bytes/map",
			Found:    arg.TypeName(),
		}
	}
}

func builtinType(args ...Object) (Object, error) {
	if len(args) != 1 {
		return nil, ErrWrongNumArguments
	}
	return &String{Value: args[0].TypeName()}, nil
}

//range(start, stop[, step])
func builtinRange(args ...Object) (Object, error) {
	numArgs := len(args)
	if numArgs < 2 || numArgs > 3 {
		return nil, ErrWrongNumArguments
	}
	var start, stop, step *Int

	for i, arg := range args {
		v, ok := args[i].(*Int)
		if !ok {
			var name string
			switch i {
			case 0:
				name = "start"
			case 1:
				name = "stop"
			case 2:
				name = "step"
			}

			return nil, ErrInvalidArgumentType{
				Name:     name,
				Expected: "int",
				Found:    arg.TypeName(),
			}
		}
		if i == 2 && v.Value <= 0 {
			return nil, ErrInvalidRangeStep
		}
		switch i {
		case 0:
			start = v
		case 1:
			stop = v
		case 2:
			step = v
		}
	}

	if step == nil {
		step = &Int{Value: int64(1)}
	}

	return buildRange(start.Value, stop.Value, step.Value), nil
}

func buildRange(start, stop, step int64) *Array {
	array := &Array{}
	if start <= stop {
		for i := start; i < stop; i += step {
			array.Value = append(array.Value, &Int{
				Value: i,
			})
		}
	} else {
		for i := start; i > stop; i -= step {
			array.Value = append(array.Value, &Int{
				Value: i,
			})
		}
	}
	return array
}

func builtinFormat(args ...Object) (Object, error) {
	numArgs := len(args)
	if numArgs == 0 {
		return nil, ErrWrongNumArguments
	}
	format, ok := args[0].(*String)
	if !ok {
		return nil, ErrInvalidArgumentType{
			Name:     "format",
			Expected: "string",
			Found:    args[0].TypeName(),
		}
	}
	s, err := Format(format.Value, args[1:]...)
	if err != nil {
		return nil, err
	}
	return &String{Value: s}, nil
}

func builtinCopy(args ...Object) (Object, error) {
	if len(args) != 1 {
		return nil, ErrWrongNumArguments
	}
	return args[0].Copy(), nil
}

func builtinMapKeys(args ...Object) (Object, error) {
	if len(args) != 1 {
		return nil, ErrWrongNumArguments
	}
	m, ok := args[0].(*Map)
	if !ok {
		return nil, ErrInvalidArgumentType{
			Name:     "map",
			Expected: "Map",
			Found:    m.TypeName(),
		}
	}
	idx, keys := 0, make([]Object, len(m.Value))
	for key := range m.Value {
		keys[idx] = &String{Value: key}
		idx++
	}
	return &Array{Value:keys}, nil
}

func builtinMapValues(args ...Object) (Object, error) {
	if len(args) != 1 {
		return nil, ErrWrongNumArguments
	}
	m, ok := args[0].(*Map)
	if !ok {
		return nil, ErrInvalidArgumentType{
			Name:     "map",
			Expected: "Map",
			Found:    m.TypeName(),
		}
	}
	idx, values := 0, make([]Object, len(m.Value))
	for _, val := range m.Value {
		values[idx] = val
		idx++
	}
	return &Array{Value:values}, nil
}

func builtinArraySort(args ...Object) (Object, error) {
	argsLen := len(args)
	if argsLen < 1 || argsLen > 2 {
		return nil, ErrWrongNumArguments
	}
	a, ok := args[0].(*Array)
	if !ok {
		return nil, ErrInvalidArgumentType{
			Name:     "array",
			Expected: "Array",
			Found:    a.TypeName(),
		}
	}
	var c = false
	if argsLen == 2 {
		b, ok := args[1].(*Bool)
		if !ok {
			return nil, ErrInvalidArgumentType{
				Name:     "type",
				Expected: "Bool",
				Found:    a.TypeName(),
			}
		}
		c = b.Value
	}
	res := make([]Object, 0, len(a.Value))
	switch a.Value[0].(type) {
	case *Int:
		for i := 0; i < len(a.Value); i++ {
			res = append(res,a.Value[i])
		}
		sort.Slice(res, func(i int, j int) bool {
			if c {
				return res[i].(*Int).Value > res[j].(*Int).Value
			}
			return res[i].(*Int).Value < res[j].(*Int).Value
		})
	case *Float:
		for i := 0; i < len(a.Value); i++ {
			res = append(res,a.Value[i])
		}
		sort.Slice(res, func(i int, j int) bool {
			if c {
				return res[i].(*Float).Value > res[j].(*Float).Value
			}
			return res[i].(*Float).Value < res[j].(*Float).Value
		})
	case *String:
		for i := 0; i < len(a.Value); i++ {
			res = append(res,a.Value[i])
		}
		sort.Slice(res, func(i int, j int) bool {
			if c {
				return res[i].(*String).Value > res[j].(*String).Value
			}
			return res[i].(*String).Value < res[j].(*String).Value
		})
	}

	return &Array{Value:res}, nil
}

func builtinArrayRand(args ...Object) (Object, error) {
	if len(args) != 1 {
		return nil, ErrWrongNumArguments
	}
	a, ok := args[0].(*Array)
	if !ok {
		return nil, ErrInvalidArgumentType{
			Name:     "array",
			Expected: "Array",
			Found:    a.TypeName(),
		}
	}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	n := make([]Object, len(a.Value))
	for i, v := range r.Perm(len(a.Value)) {
		n[i] = a.Value[v]
	}
	return &Array{Value:n}, nil
}

func builtinArrayPush(args ...Object) (Object, error) {
	if len(args) < 1 {
		return nil, ErrWrongNumArguments
	}
	a, ok := args[0].(*Array)
	if !ok {
		return nil, ErrInvalidArgumentType{
			Name:     "array",
			Expected: "Array",
			Found:    a.TypeName(),
		}
	}
	return &Array{Value: append(a.Value, args[1:]...)}, nil
}

func builtinArrayPop(args ...Object) (Object, error) {
	if len(args) != 1 {
		return nil, ErrWrongNumArguments
	}
	a, ok := args[0].(*Array)
	if !ok {
		return nil, ErrInvalidArgumentType{
			Name:     "array",
			Expected: "Array",
			Found:    a.TypeName(),
		}
	}
	l := len(a.Value) - 1
	e := a.Value[l]
	a.Value = a.Value[:l]
	return e, nil
}

func builtinArrayUnShift(args ...Object) (Object, error) {
	if len(args) < 1 {
		return nil, ErrWrongNumArguments
	}
	a, ok := args[0].(*Array)
	if !ok {
		return nil, ErrInvalidArgumentType{
			Name:     "array",
			Expected: "Array",
			Found:    a.TypeName(),
		}
	}
	return &Array{Value: append(args[1:],a.Value...)}, nil
}

func builtinArrayShift(args ...Object) (Object, error) {
	if len(args) != 1 {
		return nil, ErrWrongNumArguments
	}
	a, ok := args[0].(*Array)
	if !ok {
		return nil, ErrInvalidArgumentType{
			Name:     "array",
			Expected: "Array",
			Found:    a.TypeName(),
		}
	}
	f := a.Value[0]
	a.Value = a.Value[1:]
	return f, nil
}

func builtinArrayReverse(args ...Object) (Object, error) {
	if len(args) != 1 {
		return nil, ErrWrongNumArguments
	}
	a, ok := args[0].(*Array)
	if !ok {
		return nil, ErrInvalidArgumentType{
			Name:     "array",
			Expected: "Array",
			Found:    a.TypeName(),
		}
	}
	for i, j := 0, len(a.Value)-1; i < j; i, j = i+1, j-1 {
		a.Value[i], a.Value[j] = a.Value[j], a.Value[i]
	}
	return &Array{Value:a.Value}, nil
}

func builtinArrayUnique(args ...Object) (Object, error) {
	if len(args) != 1 {
		return nil, ErrWrongNumArguments
	}
	a, ok := args[0].(*Array)
	if !ok {
		return nil, ErrInvalidArgumentType{
			Name:     "array",
			Expected: "Array",
			Found:    a.TypeName(),
		}
	}
	set := make(map[string]int)
	res := make([]Object,0,0)
	for i := 0; i < len(a.Value); i++ {
		if _, ok := set[a.Value[i].String()]; !ok {
			set[a.Value[i].String()] = i
			res = append(res, a.Value[i])
		}
	}
	return &Array{Value:res}, nil
}

func builtinArrayColumn(args ...Object) (Object, error) {
	if len(args) != 2 {
		return nil, ErrWrongNumArguments
	}
	a, ok := args[0].(*Array)
	if !ok {
		return nil, ErrInvalidArgumentType{
			Name:     "array",
			Expected: "Array",
			Found:    a.TypeName(),
		}
	}
	s, ok := args[1].(*String)
	if !ok {
		return nil, ErrInvalidArgumentType{
			Name:     "key",
			Expected: "String",
			Found:    s.TypeName(),
		}
	}
	c := make([]Object, 0, len(a.Value))
	for _, val := range a.Value {
		m, ok := val.(*Map)
		if !ok {
			return nil, ErrInvalidArgumentType{
				Name:     "map...",
				Expected: "Map",
				Found:    m.TypeName(),
			}
		}
		if v, ok := m.Value[s.Value]; ok {
			c = append(c, v)
		}
	}
	return &Array{Value:c}, nil
}

// builtinArraySplice deletes and changes given Array, returns deleted items.
// usage:
// deleted_items := splice(array[,start[,delete_count[,item1[,item2[,...]]]])
func builtinArraySplice(args ...Object) (Object, error) {
	argsLen := len(args)
	if argsLen == 0 {
		return nil, ErrWrongNumArguments
	}

	array, ok := args[0].(*Array)
	if !ok {
		return nil, ErrInvalidArgumentType{
			Name:     "first",
			Expected: "array",
			Found:    args[0].TypeName(),
		}
	}
	arrayLen := len(array.Value)

	var startIdx int
	if argsLen > 1 {
		arg1, ok := args[1].(*Int)
		if !ok {
			return nil, ErrInvalidArgumentType{
				Name:     "second",
				Expected: "int",
				Found:    args[1].TypeName(),
			}
		}
		startIdx = int(arg1.Value)
		if startIdx < 0 || startIdx > arrayLen {
			return nil, ErrIndexOutOfBounds
		}
	}

	delCount := len(array.Value)
	if argsLen > 2 {
		arg2, ok := args[2].(*Int)
		if !ok {
			return nil, ErrInvalidArgumentType{
				Name:     "third",
				Expected: "int",
				Found:    args[2].TypeName(),
			}
		}
		delCount = int(arg2.Value)
		if delCount < 0 {
			return nil, ErrIndexOutOfBounds
		}
	}
	// if count of to be deleted items is bigger than expected, truncate it
	if startIdx+delCount > arrayLen {
		delCount = arrayLen - startIdx
	}
	// delete items
	endIdx := startIdx + delCount
	deleted := append([]Object{}, array.Value[startIdx:endIdx]...)

	head := array.Value[:startIdx]
	var items []Object
	if argsLen > 3 {
		items = make([]Object, 0, argsLen-3)
		for i := 3; i < argsLen; i++ {
			items = append(items, args[i])
		}
	}
	items = append(items, array.Value[endIdx:]...)
	array.Value = append(head, items...)

	// return deleted items
	return &Array{Value: deleted}, nil
}

// append(arr, items...)
func builtinAppend(args ...Object) (Object, error) {
	if len(args) < 2 {
		return nil, ErrWrongNumArguments
	}
	switch arg := args[0].(type) {
	case *Map:
		for _, m := range args[1:] {
			m1, ok := m.(*Map)
			if !ok {
				return nil, ErrInvalidArgumentType{
					Name:     "map...",
					Expected: "Map",
					Found:    m1.TypeName(),
				}
			}
			for k, v := range m1.Value {
				arg.Value[k] = v
			}
		}
		return &Map{Value:arg.Value}, nil
	case *Array:
		return &Array{Value: append(arg.Value, args[1:]...)}, nil
	default:
		return nil, ErrInvalidArgumentType{
			Name:     "first",
			Expected: "array",
			Found:    arg.TypeName(),
		}
	}
}

// builtinDelete deletes Map keys
// usage: delete(map, "key")
func builtinDelete(args ...Object) (Object, error) {
	argsLen := len(args)
	if argsLen != 2 {
		return nil, ErrWrongNumArguments
	}
	switch arg := args[0].(type) {
	case *Map:
		if key, ok := args[1].(*String); ok {
			delete(arg.Value, key.Value)
			return NilValue, nil
		}
		return nil, ErrInvalidArgumentType{
			Name:     "second",
			Expected: "string",
			Found:    args[1].TypeName(),
		}
	case *Array:
		if key, ok := args[1].(*Int); ok {
			arg.Value = append(arg.Value[:key.Value], arg.Value[key.Value+1:]...)
			return NilValue, nil
		}
		return nil, ErrInvalidArgumentType{
			Name:     "second",
			Expected: "int",
			Found:    args[1].TypeName(),
		}
	default:
		return nil, ErrInvalidArgumentType{
			Name:     "first",
			Expected: "map",
			Found:    arg.TypeName(),
		}
	}
}

func builtinExists(args ...Object) (Object, error) {
	argsLen := len(args)
	if argsLen != 2 {
		return nil, ErrWrongNumArguments
	}
	switch arg := args[0].(type) {
	case *Map:
		for _, i := range arg.Value {
			if i.String() == args[1].String() {
				return &Bool{Value:true}, nil
			}
		}
		return &Bool{Value:false}, nil
	case *Array:
		for _, i := range arg.Value {
			if i.String() == args[1].String() {
				return &Bool{Value:true}, nil
			}
		}
		return &Bool{Value:false}, nil
	default:
		return nil, ErrInvalidArgumentType{
			Name:     "first",
			Expected: "map|array",
			Found:    arg.TypeName(),
		}
	}
}

func builtinString(args ...Object) (Object, error) {
	argsLen := len(args)
	if !(argsLen == 1 || argsLen == 2) {
		return nil, ErrWrongNumArguments
	}
	if _, ok := args[0].(*String); ok {
		return args[0], nil
	}
	v, ok := ToString(args[0])
	if ok {
		if len(v) > MaxStringLen {
			return nil, ErrStringLimit
		}
		return &String{Value: v}, nil
	}
	if argsLen == 2 {
		return args[1], nil
	}
	return NilValue, nil
}

func builtinInt(args ...Object) (Object, error) {
	argsLen := len(args)
	if !(argsLen == 1 || argsLen == 2) {
		return nil, ErrWrongNumArguments
	}
	if _, ok := args[0].(*Int); ok {
		return args[0], nil
	}
	v, ok := ToInt64(args[0])
	if ok {
		return &Int{Value: v}, nil
	}
	if argsLen == 2 {
		return args[1], nil
	}
	return NilValue, nil
}

func builtinBool(args ...Object) (Object, error) {
	if len(args) != 1 {
		return nil, ErrWrongNumArguments
	}
	if _, ok := args[0].(*Bool); ok {
		return args[0], nil
	}
	v, ok := ToBool(args[0])
	if ok {
		if v {
			return TrueValue, nil
		}
		return FalseValue, nil
	}
	return NilValue, nil
}

func builtinFloat(args ...Object) (Object, error) {
	argsLen := len(args)
	if !(argsLen == 1 || argsLen == 2) {
		return nil, ErrWrongNumArguments
	}
	if _, ok := args[0].(*Float); ok {
		return args[0], nil
	}
	v, ok := ToFloat64(args[0])
	if ok {
		return &Float{Value: v}, nil
	}
	if argsLen == 2 {
		return args[1], nil
	}
	return NilValue, nil
}

func builtinChar(args ...Object) (Object, error) {
	argsLen := len(args)
	if !(argsLen == 1 || argsLen == 2) {
		return nil, ErrWrongNumArguments
	}
	if _, ok := args[0].(*Char); ok {
		return args[0], nil
	}
	v, ok := ToRune(args[0])
	if ok {
		return &Char{Value: v}, nil
	}
	if argsLen == 2 {
		return args[1], nil
	}
	return NilValue, nil
}

func builtinBytes(args ...Object) (Object, error) {
	argsLen := len(args)
	if !(argsLen == 1 || argsLen == 2) {
		return nil, ErrWrongNumArguments
	}

	// bytes(N) => create a new bytes with given size N
	if n, ok := args[0].(*Int); ok {
		if n.Value > int64(MaxBytesLen) {
			return nil, ErrBytesLimit
		}
		return &Bytes{Value: make([]byte, int(n.Value))}, nil
	}
	v, ok := ToByteSlice(args[0])
	if ok {
		if len(v) > MaxBytesLen {
			return nil, ErrBytesLimit
		}
		return &Bytes{Value: v}, nil
	}
	if argsLen == 2 {
		return args[1], nil
	}
	return NilValue, nil
}

func builtinIsString(args ...Object) (Object, error) {
	if len(args) != 1 {
		return nil, ErrWrongNumArguments
	}
	if _, ok := args[0].(*String); ok {
		return TrueValue, nil
	}
	return FalseValue, nil
}

func builtinIsInt(args ...Object) (Object, error) {
	if len(args) != 1 {
		return nil, ErrWrongNumArguments
	}
	if _, ok := args[0].(*Int); ok {
		return TrueValue, nil
	}
	return FalseValue, nil
}

func builtinIsFloat(args ...Object) (Object, error) {
	if len(args) != 1 {
		return nil, ErrWrongNumArguments
	}
	if _, ok := args[0].(*Float); ok {
		return TrueValue, nil
	}
	return FalseValue, nil
}

func builtinIsBool(args ...Object) (Object, error) {
	if len(args) != 1 {
		return nil, ErrWrongNumArguments
	}
	if _, ok := args[0].(*Bool); ok {
		return TrueValue, nil
	}
	return FalseValue, nil
}

func builtinIsChar(args ...Object) (Object, error) {
	if len(args) != 1 {
		return nil, ErrWrongNumArguments
	}
	if _, ok := args[0].(*Char); ok {
		return TrueValue, nil
	}
	return FalseValue, nil
}

func builtinIsBytes(args ...Object) (Object, error) {
	if len(args) != 1 {
		return nil, ErrWrongNumArguments
	}
	if _, ok := args[0].(*Bytes); ok {
		return TrueValue, nil
	}
	return FalseValue, nil
}

func builtinIsArray(args ...Object) (Object, error) {
	if len(args) != 1 {
		return nil, ErrWrongNumArguments
	}
	if _, ok := args[0].(*Array); ok {
		return TrueValue, nil
	}
	return FalseValue, nil
}

func builtinIsMap(args ...Object) (Object, error) {
	if len(args) != 1 {
		return nil, ErrWrongNumArguments
	}
	if _, ok := args[0].(*Map); ok {
		return TrueValue, nil
	}
	return FalseValue, nil
}

func builtinIsFunction(args ...Object) (Object, error) {
	if len(args) != 1 {
		return nil, ErrWrongNumArguments
	}
	switch args[0].(type) {
	case *CompiledFunction:
		return TrueValue, nil
	}
	return FalseValue, nil
}

func builtinIsCallable(args ...Object) (Object, error) {
	if len(args) != 1 {
		return nil, ErrWrongNumArguments
	}
	if args[0].CanCall() {
		return TrueValue, nil
	}
	return FalseValue, nil
}

func builtinIsIterable(args ...Object) (Object, error) {
	if len(args) != 1 {
		return nil, ErrWrongNumArguments
	}
	if args[0].CanIterate() {
		return TrueValue, nil
	}
	return FalseValue, nil
}


func builtinIsError(args ...Object) (Object, error) {
	if len(args) != 1 {
		return nil, ErrWrongNumArguments
	}
	if _, ok := args[0].(*Error); ok {
		return TrueValue, nil
	}
	return FalseValue, nil
}

func builtinIsNil(args ...Object) (Object, error) {
	if len(args) != 1 {
		return nil, ErrWrongNumArguments
	}
	if args[0] == NilValue {
		return TrueValue, nil
	}
	return FalseValue, nil
}