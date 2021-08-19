package debug

import (
	"fmt"
	"log"

	"github.com/bonagin/config"
)

type NewDebug string

var (
	Warn  = Yellow
	Fata  = Red
	Succ  = Green
	Paylt = Teal
	Paylr = Magenta
)

var (
	Black   = Color("\033[1;30m%s\033[0m")
	Red     = Color("\033[1;31m%s\033[0m")
	Green   = Color("\033[1;32m%s\033[0m")
	Yellow  = Color("\033[1;33m%s\033[0m")
	Purple  = Color("\033[1;34m%s\033[0m")
	Magenta = Color("\033[1;35m%s\033[0m")
	Teal    = Color("\033[1;36m%s\033[0m")
	White   = Color("\033[1;37m%s\033[0m")
)

func Color(colorString string) func(...interface{}) string {
	sprint := func(args ...interface{}) string {
		if config.Get("DEBUG_MODE") != "ON" {
			return fmt.Sprint(args...)
		}

		return fmt.Sprintf(colorString,
			fmt.Sprint(args...))
	}
	return sprint
}

func (title NewDebug) Success(args ...interface{}) {
	Log(Succ("[SUCCESS] : "), Succ(title+" - "), Succ(args))
}

func (title NewDebug) Info(args ...interface{}) {
	Log("[INFO]    : ", title+" - ", args)
}

func (title NewDebug) InfoT(args ...interface{}) {
	Log(Teal("[INFO]    : "), Teal(title+" - "), Teal(args))
}

func (title NewDebug) InfoP(args ...interface{}) {
	Log(Purple("[INFO]    : "), Purple(title+" - "), Purple(args))
}

func (title NewDebug) Dump(args ...interface{}) {
	Log("[INFO]    : ", title+" - ", args)
}

func (title NewDebug) Warning(args ...interface{}) {
	Log(Warn("[WARNING] : "), Warn(title+" - "), Warn(args))
}

func (title NewDebug) Alert(args ...interface{}) {
	Log(Fata("[ALERT]   : "), Fata(title+" - "), Fata(args))
}

func (title NewDebug) Fatal(args ...interface{}) {
	Log(Fata("[FATAL]   : "), Fata(title+" - "), Fata(args))
	log.Fatal(args)
}

func (title NewDebug) Printf(format string, v ...interface{}) {
	if config.Get("DEBUG_MODE") == "ON" {
		log.Printf(format, v)
		return
	}

	fmt.Printf(format, v)
}

func Log(args ...interface{}) {
	if config.Get("DEBUG_MODE") == "ON" {
		log.Println(args)
		return
	}

	fmt.Println(args)
}
