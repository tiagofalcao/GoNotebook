package log

import (
	"flag"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

const (
	DefaultLevel = "Default"
	DebugLevel   = "Debug"
	InfoLevel    = "Info"
	WarningLevel = "Warning"
	ErrorLevel   = "Error"
	NoneLevel    = "None"
)

type LogLevel string

var (
	Debug   *log.Logger = log.New(ioutil.Discard, DebugLevel+" :", log.Lshortfile)
	Info    *log.Logger = log.New(ioutil.Discard, InfoLevel+" :", log.Lshortfile)
	Warning *log.Logger = log.New(ioutil.Discard, WarningLevel+" :", log.Lshortfile)
	Error   *log.Logger = log.New(ioutil.Discard, ErrorLevel+" :", log.Lshortfile)
)

var optLog string = WarningLevel

func init() {
	flag.StringVar(&optLog, "log", WarningLevel, "Log Level")
}

var initLock bool = false

func Init(l LogLevel) {
	if initLock {
		return
	}
	initLock = true

	level := string(l)
	if strings.EqualFold(NoneLevel, level) {
		return
	}

	if strings.EqualFold(DefaultLevel, level) {
		level = optLog
	}

	Error = log.New(os.Stderr, ErrorLevel+" :", log.Lshortfile)

	if strings.EqualFold(ErrorLevel, level) {
		return
	}

	Warning = log.New(os.Stdout, WarningLevel+" :", log.Lshortfile)

	if strings.EqualFold(WarningLevel, level) {
		return
	}

	Info = log.New(os.Stdout, InfoLevel+" :", log.Lshortfile)

	if strings.EqualFold(InfoLevel, level) {
		return
	}

	Debug = log.New(os.Stdout, DebugLevel+" :", log.Lshortfile)

}
