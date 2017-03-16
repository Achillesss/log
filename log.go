package log

// Author: Achillesss
// date: 2017-03-10 23:12:00

import (
	"flag"
)

var (
	infoOn = flag.Bool("info", false, "whether print 'info', default off")
	warnOn = flag.Bool("warn", false, "whether print 'warning', default off")
	errOn  = flag.Bool("err", true, "whether print 'error', default on")
)

// Parse parses flags
func Parse() {
	flag.Parse()
}

func formatErr(skip int, pkgName bool, err *error) {
	newLogAgent().setSkip(skip + 1).setSymble(pkgName).formatErr(err)
}

func print(skip int, printType, end string, format string, arg ...interface{}) {
	newLogAgent().setSkip(skip+1).setPrintType(printType).setEnd(end).print(format, arg...)
}

func funcName(skip int, on bool) string {
	return newLogAgent().setSkip(skip + 1).setSymble(on).funcName()
}

// FuncName returns name of the function which calls it
func FuncName() string {
	return funcName(1, false)
}

// FuncNameP returns name of the function which calls if with package name
func FuncNameP() string {
	return funcName(1, true)
}

// FuncNameN returns name of function skipped by n
func FuncNameN(skip int) string {
	return funcName(skip+1, false)
}

// FuncNameNP returns name of function with package name by n
func FuncNameNP(skip int) string {
	return funcName(skip+1, true)
}

// FmtErr formats an error with name of the function which calls it
func FmtErr(err *error) {
	formatErr(1, false, err)
}

// FmtErrP differs from FmtErr in that it formats an error WITH PACKAGE NAME IN ADDITION
func FmtErrP(err *error) {
	formatErr(1, true, err)
}

// FmtErrN formats an error with name of the function which calls it skipped by skip
func FmtErrN(skip int, err *error) {
	formatErr(skip+1, false, err)
}

// FmtErrNP differs from FmtErrN in that it formats an error WITH PACKAGE NAME IN ADDITION
func FmtErrNP(skip int, err *error) {
	formatErr(skip+1, true, err)
}

// L logs a description when a function response is not ok
func L(ok bool, desc string) {
	if !ok {
		print(1, logLog, "", desc)
	}
}

// Ln differs from L in that it create a newline after loging
func Ln(ok bool, desc string, skip int) {
	if !ok {
		print(1, logLog, newline, desc)
	}
}

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
