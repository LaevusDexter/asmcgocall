package asmcgocall

import (
	"fmt"
	"strings"
)

var Debug bool = false

func bytes(i interface{}) string {
	var res int

	switch i.(type) {
	case int:
		res = i.(int)
	case uintptr:
		res = int(i.(uintptr))
	}

	if res > 1 {
		return "bytes"
	} else {
		return "byte"
	}
}

type LogFormat byte

const (
	LogFormatDefault = iota
	LogFormatC
	LogFormatGo
)

var logFormat LogFormat = LogFormatC

func SetLogFormat(logf LogFormat) {
	logFormat = logf
}

func logf(str string, args ...interface{}) {
	switch logFormat {
	case LogFormatC:
		if strings.Index(str, "__pad") != -1 {
			return
		}

		for i, a := range args {
			switch a.(type) {
			case string:
				args[i] = strings.ReplaceAll(args[i].(string), "_Ctype_", "")
			}
		}
	case LogFormatGo:
		if strings.Index(str, "__pad") != -1 {
			return
		}

		for i, a := range args {
			switch a.(type) {
			case string:
				args[i] = strings.ReplaceAll(args[i].(string), "_Ctype_", "C.")
			}
		}
	}

	fmt.Printf(str, args...)
}