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
func FuncNameN(n uint) string {
	return funcName(int(n)+2, ".")
}

// PkgFuncNameN returns name of function with package name by n
func PkgFuncNameN(n uint) string {
	return funcName(int(n)+2, "/")
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

func print(skip int, format string, arg ...interface{}) {
	if _, file, line, ok := runtime.Caller(skip); ok {
		fileName := strings.Split(file, "/")
		fmt.Printf("[%v:%v", fileName[len(fileName)-1], line)
		if format != "" {
			fmt.Printf(" ")
			fmt.Printf(format, arg...)
		}
		fmt.Printf("]\t")
	}
}

// FError formats an error with name of the function which calls it
func FError(err error) error {
	if err != nil {
		return fmt.Errorf("%s failed. Error: %v", funcName(2, "/"), err)
	}
	return nil
}

// Log prints a description when a function response is not ok
func Log(ok bool, desc string) {
	if !ok {
		print(2, desc)
	}
}

// Printf prints information inline
func Printf(format string, arg ...interface{}) {
	print(2, format, arg...)
}
