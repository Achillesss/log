package log

import (
	"fmt"
	"runtime"
	"strings"
)

func in(value interface{}, src ...interface{}) (ok bool) {
	for i := range src {
		if ok = value == src[i]; ok {
			break
		}
	}
	return
}

func checkSplit(split string) string {
	if !in(split, slash, point) {
		split = slash
	}
	return split
}

func funcName(skip int, symb string) string {
	pc, _, _, _ := runtime.Caller(skip)
	name := runtime.FuncForPC(pc).Name()
	strs := strings.Split(name, symb)
	return strs[len(strs)-1]
}

func (a *logAgent) print(format string, arg ...interface{}) {
	if _, file, line, ok := runtime.Caller(a.skip + 1); ok {
		fileName := strings.Split(file, slash)
		arg = append([]interface{}{a.printType, fileName[len(fileName)-1], line}, arg...)
		arg = append(arg, a.end)
		fmt.Printf("[%s_%v_%v] "+format+"%v", arg...)
	}
}

func (a *logAgent) formatErr(err *error) {
	if err != nil && *err != nil {
		*err = fmt.Errorf("%s fail. Desc: %v", funcName(a.skip+1, a.symble), *err)
	}
}
