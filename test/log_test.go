package logtest

import (
	"testing"

	"github.com/achillesss/log"
)

func TestLog(t *testing.T) {
	log.Parse()
	funcName()
	funcNameN()
}
