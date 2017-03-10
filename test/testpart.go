package logtest

import (
	log "bitbucket.org/magmeng/flashcore/log"
	"fmt"
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
}
