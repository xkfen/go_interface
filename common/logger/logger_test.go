package logger

import (
	"testing"
)

func TestLog(t *testing.T) {
	InitLogger(LvlDebug, nil)
	Debug("asf", "fwef", "123", "ok", "sss")
	Warn("123")
	Error("err")
}