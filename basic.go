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
func newLogAgent() *logAgent {
	return &logAgent{symble: slash, end: inline, printType: logInformation}
}

func (a *logAgent) isNil() bool {
	return a == nil
}

func (a *logAgent) setSymble(on bool) *logAgent {
	a.symble = slash
	if !on {
		a.symble = point
	}
	return a
}

func (a *logAgent) setSkip(s int) *logAgent {
	a.skip = s
	return a
}

func (a *logAgent) setEnd(e string) *logAgent {
	if e == newline {
		a.end = e
	}
	return a
}

func (a *logAgent) setPrintType(p string) *logAgent {
	if in(p, logWarning, logError, logLog) {
		a.printType = p
	}
	return a
}

func (a *logAgent) funcName() string {
	pc, _, _, _ := runtime.Caller(a.skip + 1)
	name := runtime.FuncForPC(pc).Name()
	strs := strings.Split(name, a.symble)
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
		*err = fmt.Errorf("%s fail. Desc: %v", a.funcName(), *err)
	}
}

func (a *logAgent) log(ok bool, desc string) {
	if !ok {
		a.setSkip(a.skip + 1).print(desc)
	}
}
