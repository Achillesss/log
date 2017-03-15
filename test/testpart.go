package logtest

import (
	"fmt"

	"time"

	log "github.com/achillesss/log"
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
		fmt.Printf("function name: %v\n", log.FuncNameN(2))
	}()
	func() {
		fmt.Printf("function name: %v\n", log.FuncNameN(3))
	}()
	go warningf()
	go warningfln()
	go errorf()
	go errorfln()

	infof()
	infofln()
	go fmtErr()

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

func fmtErr() {
	err := fmt.Errorf("error")
	log.Infofln("ERR: %v", err)
	log.FmtErr(&err)
	log.Infofln("ERR: %v", err)
}
