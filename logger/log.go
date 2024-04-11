package logger

import (
	"fmt"
	"log"
	"os"

	"github.com/jedib0t/go-pretty/v6/text"
)

var Level LogLevel = Error

var (
	tranLogger = log.New(os.Stdout, text.BgWhite.Sprint("[tran]")+" ", log.Ldate|log.Ltime)
	debuLogger = log.New(os.Stdout, text.BgBlue.Sprint("[debu]")+" ", log.Ldate|log.Ltime)
	infoLogger = log.New(os.Stdout, text.BgGreen.Sprint("[info]")+" ", log.Ldate|log.Ltime)
	warnLogger = log.New(os.Stdout, text.BgYellow.Sprint("[warn]")+" ", log.Ldate|log.Ltime)
	erroLogger = log.New(os.Stdout, text.BgRed.Sprint("[erro]")+" ", log.Ldate|log.Ltime)
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
		tranLogger.Println(fmt.Sprint(a...))
	}
}
func Debu(a ...any) {
	if Level >= Debug {
		debuLogger.Println(fmt.Sprint(a...))
	}
}
func Info(a ...any) {
	if Level >= Infomation {
		infoLogger.Println(fmt.Sprint(a...))
	}
}
func Warn(a ...any) {
	if Level >= Warning {
		warnLogger.Println(fmt.Sprint(a...))
	}
}
func Erro(a ...any) {
	if Level >= Error {
		erroLogger.Println(fmt.Sprint(a...))
	}
}
