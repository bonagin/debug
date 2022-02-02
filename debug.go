package debug

import (
	"fmt"
	"log"
	"os"
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

func debugMode() bool {
	args := os.Args

	for _, arg := range args {
		if arg == "--debug" {
			return true
		}
	}

	return false
}

func Color(colorString string) func(...interface{}) string {
	sprint := func(args ...interface{}) string {
		if !debugMode() {
			return fmt.Sprint(args...)
		}

		return fmt.Sprintf(colorString,
			fmt.Sprint(args...))
	}
	return sprint
}

func (title NewDebug) Success(args ...interface{}) {
	if !debugMode() {
		return
	}
	Log(Succ("[SUCCESS] : "), Succ(title+" - "), Succ(args))
}

func (title NewDebug) Info(args ...interface{}) {
	if !debugMode() {
		return
	}
	Log("[INFO]    : ", title+" - ", args)
}

func (title NewDebug) InfoT(args ...interface{}) {
	if !debugMode() {
		return
	}
	Log(Teal("[INFO]    : "), Teal(title+" - "), Teal(args))
}

func (title NewDebug) InfoP(args ...interface{}) {
	if !debugMode() {
		return
	}
	Log(Purple("[INFO]    : "), Purple(title+" - "), Purple(args))
}

func (title NewDebug) Dump(args ...interface{}) {
	if !debugMode() {
		return
	}
	Log(Purple("[INFO]    : "), Purple(title+" - "), Purple(args))
}

func (title NewDebug) Warning(args ...interface{}) {
	if !debugMode() {
		return
	}
	Log(Warn("[WARNING] : "), Warn(title+" - "), Warn(args))
}

func (title NewDebug) Alert(args ...interface{}) {
	Log(Fata("[ALERT]   : "), Fata(title+" - "), Fata(args))
}

func (title NewDebug) Fatal(args ...interface{}) {
	Log(Fata("[FATAL]   : "), Fata(title+" - "), Fata(args))
	log.Fatal(args...)
}

func (title NewDebug) Printf(format string, v ...interface{}) {
	if !debugMode() {
		return
	}
	fmt.Printf(format, v...)
}

func (title NewDebug) Println(format string, v ...interface{}) {
	if !debugMode() {
		return
	}

	fmt.Printf(format, v...)
}

func Log(args ...interface{}) {
	if !debugMode() {
		log.Println(args...)
		return
	}

	fmt.Println(args...)
}
