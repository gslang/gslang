package stdlib

import (
	"io/ioutil"
	"crypto/md5"
	"crypto/sha1"
	"encoding/hex"
	"github.com/gslang/gslang"
)
var cryptoModule = map[string]gslang.Object{
	"md5": &gslang.UserFunction{
		Name:  "md5",
		Value: cryptoMd5,
	},
	"md5_file": &gslang.UserFunction{
		Name:  "md5_file",
		Value: cryptoMd5File,
	},
	"sha1": &gslang.UserFunction{
		Name:  "sha1",
		Value: cryptoSha1,
	},
	"sha1_file": &gslang.UserFunction{
		Name:  "sha1_file",
		Value: cryptoSha1File,
	},
}

func cryptoMd5(args ...gslang.Object) (
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
	h := md5.New()
	h.Write([]byte(s1))
	ret = &gslang.String{Value: hex.EncodeToString(h.Sum(nil))}
	return
}

func cryptoMd5File(args ...gslang.Object) (
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
	var str []byte
	str, err = ioutil.ReadFile(s1)
	if err != nil {
		return
	}
	h := md5.New()
	h.Write([]byte(str))
	ret = &gslang.String{Value: hex.EncodeToString(h.Sum(nil))}
	return
}


func cryptoSha1(args ...gslang.Object) (
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
	h := sha1.New()
	h.Write([]byte(s1))
	ret = &gslang.String{Value: hex.EncodeToString(h.Sum(nil))}
	return
}

func cryptoSha1File(args ...gslang.Object) (
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
	var str []byte
	str, err = ioutil.ReadFile(s1)
	if err != nil {
		return
	}
	h := sha1.New()
	h.Write([]byte(str))
	ret = &gslang.String{Value: hex.EncodeToString(h.Sum(nil))}
	return
}