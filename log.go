package log

import (
	"fmt"
	"log"
	"os"
	"runtime"
	"strconv"
	"strings"
)

type Logger struct {
}

type logInfo struct {
}
type logDebug struct {
}
type logError struct {
}

func (writer logInfo) Write(bytes []byte) (int, error) {
	if os.Getenv("LOGGING") != "ERROR_ONLY" {
		return fmt.Print("[INFO] " + trace() + " " + string(bytes))
	}
	return 0, nil
}
func (writer logDebug) Write(bytes []byte) (int, error) {
	if os.Getenv("LOGGING") == "ALL" {
		return fmt.Print("[DEBUG] " + trace() + " " + string(bytes))
	}
	return 0, nil

}
func (writer logError) Write(bytes []byte) (int, error) {

	return fmt.Print("[ERROR] " + trace() + " " + string(bytes))
}

//Info wraps a custom log
func (l *Logger) Info(v ...interface{}) {
	_, _, line, _ := runtime.Caller(1)
	log.SetOutput(new(logInfo))
	aString := ""
	for _, aV := range v {
		aString += " " + fmt.Sprintf("%v", aV)
	}
	log.Println("[l:"+strconv.Itoa(line)+"]", aString)
}

//Debug wraps a custom log
func (l *Logger) Debug(v ...interface{}) {
	_, _, line, _ := runtime.Caller(1)
	log.SetOutput(new(logDebug))
	aString := ""
	for _, aV := range v {
		aString += " " + fmt.Sprintf("%v", aV)
	}
	log.Println("[l:"+strconv.Itoa(line)+"]", aString)
}

//Error wraps a custom log
func (l *Logger) Error(v ...interface{}) {
	_, _, line, _ := runtime.Caller(1)
	log.SetOutput(new(logError))
	aString := ""
	for _, aV := range v {
		aString += " " + fmt.Sprintf("%v", aV)
	}
	log.Println("[l:"+strconv.Itoa(line)+"]", aString)
}

func trace() string {
	pc := make([]uintptr, 10) // at least 1 entry needed
	runtime.Callers(7, pc)

	f := runtime.FuncForPC(pc[0])

	file, _ := f.FileLine(pc[0])

	ss := strings.Split(file, "/")
	fileShort := ss[len(ss)-1]
	//fmt.Printf("%s:%d %s\n", file, line, f.Name())
	return fileShort + " > " + strings.Split(f.Name(), ".")[1]
}
