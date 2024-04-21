package logger

import (
	"fmt"
	"log"
	"os"

	"github.com/jedib0t/go-pretty/v6/text"
)

var Level LogLevel = Error

var (
	tranLogger = log.New(os.Stdout, text.BgWhite.Sprint("[tran]")+" ", log.Ltime|log.Lshortfile)
	debuLogger = log.New(os.Stdout, text.BgBlue.Sprint("[debu]")+" ", log.Ltime|log.Lshortfile)
	infoLogger = log.New(os.Stdout, text.BgGreen.Sprint("[info]")+" ", log.Ltime|log.Lshortfile)
	warnLogger = log.New(os.Stdout, text.BgYellow.Sprint("[warn]")+" ", log.Ltime|log.Lshortfile)
	erroLogger = log.New(os.Stdout, text.BgRed.Sprint("[erro]")+" ", log.Ltime|log.Lshortfile)
)

type LogLevel int

const (
	NoLog LogLevel = iota + 1
	Error
	Warning
	Infomation
	Debug
	Trance
)

func Tran(a ...any) {
	if Level >= Trance {
		tranLogger.Output(2, fmt.Sprint(a...))
	}
}
func Debu(a ...any) {
	if Level >= Debug {
		debuLogger.Output(2, fmt.Sprint(a...))
	}
}
func Info(a ...any) {
	if Level >= Infomation {
		infoLogger.Output(2, fmt.Sprint(a...))
	}
}
func Warn(a ...any) {
	if Level >= Warning {
		warnLogger.Output(2, fmt.Sprint(a...))
	}
}
func Erro(a ...any) {
	if Level >= Error {
		erroLogger.Output(2, fmt.Sprint(a...))
	}
}
