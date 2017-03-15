package log

// package log prints logs inline
// Author: Achillesss
// date: 2017-03-10 23:12:00

import (
	"flag"
)

var (
	infoOn = flag.Bool("info", false, "whether print 'info'")
	warnOn = flag.Bool("warn", false, "whether print 'warning'")
	errOn  = flag.Bool("err", false, "whether print 'error'")

	pkgNameOn = flag.Bool("pn", false, "whether PRINT with package name")
)

type logAgent struct {
	// only '/', '.' supported, '/' prints with package name while '.' not
	symble string
	// prints with a skip. 0 means the function that first calls a print
	skip int
	// only '\t' and '\n' supported
	end string

	printType string

	err error
}

func newLogAgent() *logAgent {
	return &logAgent{symble: slash, end: inline, printType: logInformation}
}

func (a *logAgent) isNil() bool {
	return a == nil
}

func (a *logAgent) setErr(err error) *logAgent {
	a.err = err
	return a
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
	if e == "\n" {
		a.end = e
	}
	return a
}

func (a *logAgent) setPrintType(p string) *logAgent {
	if in(p, logWarning, logError) {
		a.printType = p
	}
	return a
}

var defaultSymble = "/"

// Parse parses flags
func Parse() {
	flag.Parse()
}

// FuncNameN returns name of function skipped by n
func FuncNameN(n int) string {
	return funcName(n+2, ".")
}

// PkgFuncNameN returns name of function with package name by n
func PkgFuncNameN(n int) string {
	return funcName(n+2, "/")
}

// FuncName returns name of the function which calls it
func FuncName() string {
	return funcName(2, ".")
}

// PkgFuncName returns name of the function which calls if with package name
func PkgFuncName() string {
	return funcName(2, "/")
}

// FmtErr formats an error with name of the function which calls it
func FmtErr(err *error) {
	formatErr(1, false, err)
}

// FmtErrP formats an error with name of the function which calls it, and with package name
func FmtErrP(err *error) {
	formatErr(1, true, err)
}

// func FError(err error) error {
// 	return formatErr(err, 0, slash)
// }

// // FErrorN formats an error with name of the function which calls it skipped by n
// func FErrorN(err error, n int) error {
// 	return formatErr(err, n, slash)
// }

// func l(ok bool, desc string, n int) {
// 	if !ok {
// 		print(n+3, "\t", desc)
// 	}
// }

// // L log a description when a function response is not ok
// func L(ok bool, desc string) {
// 	l(ok, desc, 1)
// }

// // Ln logs a description when a function response is not ok, skipped by n
// func Ln(ok bool, desc string, n int) {
// 	l(ok, desc, n)
// }

// Infof prints information inline
func Infof(format string, arg ...interface{}) {
	if *infoOn {
		print(1, "", "", format, arg...)
	}
}

// Infofln prints information and create new line
func Infofln(format string, arg ...interface{}) {
	if *infoOn {
		print(1, "", newline, format, arg...)
	}
}

// Warningf prints information inline
func Warningf(format string, arg ...interface{}) {
	if *warnOn {
		print(1, logWarning, "", format, arg...)
	}
}

// Warningfln prints information and create new line
func Warningfln(format string, arg ...interface{}) {
	if *warnOn {
		print(1, logWarning, newline, format, arg...)
	}
}

// Errorf prints information inline
func Errorf(format string, arg ...interface{}) {
	if *errOn {
		print(1, logError, "", format, arg...)
	}
}

// Errorfln prints information and create new line
func Errorfln(format string, arg ...interface{}) {
	if *errOn {
		print(1, logError, newline, format, arg...)
	}
}

func formatErr(skip int, pkgName bool, err *error) {
	newLogAgent().setSkip(skip + 1).setSymble(pkgName).formatErr(err)
}

func print(skip int, printType, end string, format string, arg ...interface{}) {
	newLogAgent().setSkip(skip+1).setPrintType(printType).setEnd(end).print(format, arg...)
}
