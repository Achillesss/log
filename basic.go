package log

import (
	"fmt"
	"runtime"
	"strings"
	"time"
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
		timeTag := ""
		if *timeOn {
			timeTag = " " + time.Now().UTC().Format(time.RFC3339)[:19]
		}
		arg = append([]interface{}{a.printType, fileName[len(fileName)-1], line, timeTag}, arg...)
		arg = append(arg, a.end)
		format = "[%s_%v_%v%s] " + format + "%v"
		fmt.Printf(format, arg...)
	}
}

func (a *logAgent) formatErr(err *error) error {
	if err != nil && *err != nil {
		*err = fmt.Errorf("%s fail. Desc: %v", a.funcName(), *err)
		return *err
	}
	return nil
}
