package log

import (
	"flag"
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	flag.Parse()
	test_funcName()
	test_funcNameN()
}

func test_funcName() {
	func() {
		fmt.Printf("function name: %v\n", FuncName())
	}()
}

func test_funcNameN() {

	func() {
		fmt.Printf("function name: %v\n", FuncNameN(0))
	}()
	func() {
		fmt.Printf("function name: %v\n", FuncNameN(1))
	}()
	func() {
		fmt.Printf("function name: %v\n", FuncNameNP(2))
	}()
	func() {
		fmt.Printf("function name: %v\n", FuncNameNP(3))
	}()
	warningfln()
	warningf()

	errorfln()
	errorf()

	infofln()
	infof()
	l()
	fmtErr()

}

func infof() {
	Infof("")
}
func infofln() {
	Infofln("")
}

func warningf() {
	Warningf("")
}
func warningfln() {
	Warningfln("")
}

func errorf() {
	Errorf("errorf\n")
}
func errorfln() {
	Errorfln("errorfln")
}

func l() {
	L(true, "desc")
	Lln(true, "%s Failed. Error: %v. Resp: %#v", FuncName(), "ERROR", "RESP")
}
func fmtErr() {
	err1 := fmt.Errorf("error1")
	err2 := fmt.Errorf("error2")
	Infofln("ERR: %v", err1)
	Infofln("ERR: %v", err2)
	e1 := FmtErr(&err1)
	e2 := FmtErrP(&err2)
	Infofln("ERR: %v\t%v", err1, e1)
	Infofln("ERR: %v\t%v", err2, e2)
}
