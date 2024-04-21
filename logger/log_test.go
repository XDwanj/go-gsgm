package logger

import "testing"

func TestPrintln(t *testing.T) {
	Level = Trance
	Tran("hello ", "world")
	Debu("hello ", "world")
	Info("hello ", "world")
	Warn("hello ", "world")
	Erro("hello ", "world")
}
