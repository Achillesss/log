package logtest

import (
	"fmt"

	"time"

	"github.com/achillesss/log"
)

func funcName() {
	func() {
		fmt.Printf("function name: %v\n", log.FuncName())
	}()
}

func funcNameN() {

	func() {
		fmt.Printf("function name: %v\n", log.FuncNameN(0))
	}()
	func() {
		fmt.Printf("function name: %v\n", log.FuncNameN(1))
	}()
	func() {
		fmt.Printf("function name: %v\n", log.FuncNameNP(2))
	}()
	func() {
		fmt.Printf("function name: %v\n", log.FuncNameNP(3))
	}()
	go warningf()
	go warningfln()
	go errorf()
	go errorfln()

	infof()
	infofln()
	l()
	// go fmtErr()

	time.Sleep(time.Millisecond * 50)
}

func infof() {
	log.Infof("")
}
func infofln() {
	log.Infofln("")
}

func warningf() {
	log.Warningf("")
}
func warningfln() {
	log.Warningfln("")
}

func errorf() {
	log.Errorf("")
}
func errorfln() {
	log.Errorfln("")
}
func l() {
	log.L(true, "desc")
	log.Ln(true, "desc", 1)
}
func fmtErr() {
	err1 := fmt.Errorf("error1")
	err2 := fmt.Errorf("error2")
	log.Infofln("ERR: %v", err1)
	log.Infofln("ERR: %v", err2)
	log.FmtErr(&err1)
	log.FmtErrP(&err2)
	log.Infofln("ERR: %v", err1)
	log.Infofln("ERR: %v", err2)
}
