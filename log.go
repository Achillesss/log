package log

// package log prints logs inline
// Author: Achillesss
// date: 2017-03-10 23:12:00

import (
	"fmt"
	"runtime"
	"strings"
)

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

func funcName(n int, smb string) string {
	pc, _, _, _ := runtime.Caller(n)
	name := runtime.FuncForPC(pc).Name()
	strs := strings.Split(name, smb)
	return strs[len(strs)-1]
}

func print(skip int, end string, format string, arg ...interface{}) {
	if _, file, line, ok := runtime.Caller(skip); ok {
		fileName := strings.Split(file, "/")
		fmt.Printf("[%v:%v", fileName[len(fileName)-1], line)
		if format != "" {
			fmt.Printf(" ")
			fmt.Printf(format, arg...)
		}
		fmt.Printf("]%s", end)
	}
}

func fError(err error, n int) error {
	if err != nil {
		return fmt.Errorf("%s failed. Error: %v", funcName(3+n, "/"), err)
	}
	return nil
}

// FError formats an error with name of the function which calls it
func FError(err error) error {
	return fError(err, 0)
}

// FErrorN formats an error with name of the function which calls it skipped by n
func FErrorN(err error, n int) error {
	return fError(err, n)
}

// Log prints a description when a function response is not ok
func Log(ok bool, desc string, n int) {
	Logn(ok, desc, 1)
}

// Logn prints a description when a function response is not ok skipped by n
func Logn(ok bool, desc string, n int) {
	if !ok {
		print(n+2, "\t", desc)
	}
}

// Printf prints information inline
func Printf(format string, arg ...interface{}) {
	print(2, "\t", format, arg...)
}

// Printfln prints information and create new line
func Printfln(format string, arg ...interface{}) {
	print(2, "\n", format, arg...)
}
