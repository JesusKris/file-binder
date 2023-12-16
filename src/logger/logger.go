package logger

import (
	"fmt"
	"os"
	"runtime"
	"strconv"
	"time"
)

var reset = "\033[0m"
var red = "\033[31m"
var yellow = "\033[33m"
var purple = "\033[35m"
var cyan = "\033[36m"

func callerLine() string {
	_, _, line, _ := runtime.Caller(2)
	return strconv.Itoa(line)
}
func callerLocation() string {
	_, location, _, _ := runtime.Caller(2)
	return location
}

//currentTime gets current time in "01-02-2006 15:04:05.00000" format
func currentTime() string {
	dt := time.Now()
	formatted := dt.Format("01-02-2006 15:04:05.00000")
	return formatted
}

//Log prints given message to stdout
func Log(msg string) {
	fmt.Println(cyan + "[LOG]" + purple + "[" + currentTime() + "][" + callerLocation() + ":" + callerLine() + "]: " + cyan + msg + reset)
}

//Warning prints given message to stdout as a warning
func Warning(msg string) {
	fmt.Println(yellow + "[WARNING]" + purple + "[" + currentTime() + "][" + callerLocation() + ":" + callerLine() + "]: " + yellow + msg + reset)
}

//Error prints given err to stdout as a error
func Error(err error) {
	fmt.Println(red + "[ERROR]" + purple + "[" + currentTime() + "][" + callerLocation() + ":" + callerLine() + "]: " + red + err.Error() + reset)
}

//Fatal prints given message and err to stdout and then exits safely
func Fatal(err error) {
	fmt.Println(red + "[FATAL ERROR]" + purple + "[" + currentTime() + "][" + callerLocation() + ":" + callerLine() + "]: " + red + err.Error() + reset)
	fmt.Println(red + "[EXITING] " + reset)
	os.Exit(1)
}
