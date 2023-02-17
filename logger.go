package stdlog

import (
	"log"
	"os"
)

func newInfo() *log.Logger {
	return log.New(os.Stdout, "[info] ", log.LstdFlags)
}

func newErr() *log.Logger {
	return log.New(os.Stderr, "[error] ", log.LstdFlags)
}

func newWarn() *log.Logger {
	return log.New(os.Stderr, "[warn] ", log.LstdFlags)
}

func Infoln(v ...any) {
	newInfo().Println(v...)
}

func Infof(format string, v ...any) {
	newInfo().Printf(format, v...)
}

func Warnln(v ...any) {
	newWarn().Println(v...)
}

func Warnf(format string, v ...any) {
	newWarn().Printf(format, v...)
}

func Errln(v ...any) {
	newErr().Println(v...)
}

func Errf(format string, v ...any) {
	newErr().Printf(format, v...)
}
