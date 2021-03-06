// Copyright 2015 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

/*
Package log implements a logger package for net-upnp-go.
*/
package log

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

type LoggerOutpter func(file string, level int, msg string) (int, error)

type Logger struct {
	File     string
	Level    int
	outputer LoggerOutpter
}

const (
	format           = "%s %s %s"
	lf               = "\n"
	FilePerm         = 0644
	LoggerLevelTrace = (1 << 5)
	LoggerLevelInfo  = (1 << 4)
	LoggerLevelWarn  = (1 << 3)
	LoggerLevelError = (1 << 2)
	LoggerLevelFatal = (1 << 1)
	LoggerLevelNone  = 0

	loggerLevelUnknownString = "UNKNOWN"
	loggerStdout             = "stdout"
)

var sharedLogger *Logger

func SetSharedLogger(logger *Logger) {
	sharedLogger = logger
}

func GetSharedLogger() *Logger {
	return sharedLogger
}

var sharedLogStrings = map[int]string{
	LoggerLevelTrace: "TRACE",
	LoggerLevelInfo:  "INFO",
	LoggerLevelWarn:  "WARN",
	LoggerLevelError: "ERROR",
	LoggerLevelFatal: "FATAL",
}

func GetLevelString(level int) string {
	logString, hasString := sharedLogStrings[level]
	if !hasString {
		return loggerLevelUnknownString
	}
	return logString
}

func (logger *Logger) GetLevel() int {
	return logger.Level
}

func (logger *Logger) GetLevelString() string {
	logString, hasString := sharedLogStrings[logger.Level]
	if !hasString {
		return loggerLevelUnknownString
	}
	return logString
}

func NewStdoutLogger(level int) *Logger {
	logger := &Logger{
		File:     loggerStdout,
		Level:    level,
		outputer: outputStrount}
	return logger
}

func outputStrount(file string, level int, msg string) (int, error) {
	fmt.Println(msg)
	return len(msg), nil
}

func NewFileLogger(file string, level int) *Logger {
	logger := &Logger{
		File:     file,
		Level:    level,
		outputer: outputToFile}
	return logger
}

func outputToFile(file string, level int, msg string) (int, error) {
	msgBytes := []byte(msg + lf)
	fd, err := os.OpenFile(file, (os.O_WRONLY | os.O_CREATE | os.O_APPEND), FilePerm)
	if err != nil {
		return 0, err
	}

	writer := bufio.NewWriter(fd)
	writer.Write(msgBytes)
	writer.Flush()

	fd.Close()

	return len(msgBytes), nil
}

func outputString(outputLevel int, msg string) int {
	if sharedLogger == nil {
		return 0
	}

	logLevel := sharedLogger.GetLevel()
	if (logLevel < outputLevel) || (logLevel <= LoggerLevelFatal) || (LoggerLevelTrace < logLevel) {
		return 0
	}

	t := time.Now()
	logDate := fmt.Sprintf("%d-%02d-%02d %02d:%02d:%02d",
		t.Year(), t.Month(), t.Day(),
		t.Hour(), t.Minute(), t.Second())

	headerString := fmt.Sprintf("[%s]", GetLevelString(outputLevel))
	logMsg := fmt.Sprintf(format, logDate, headerString, msg)
	logMsgLen := len(logMsg)

	if 0 < logMsgLen {
		logMsgLen, _ = sharedLogger.outputer(sharedLogger.File, logLevel, logMsg)
	}

	return logMsgLen
}

func output(outputLevel int, v interface{}) int {
	switch msg := v.(type) {
	case string:
		return outputString(outputLevel, msg)
	case error:
		return outputString(outputLevel, msg.Error())
	}
	return 0
}

func Trace(msg interface{}) int {
	return output(LoggerLevelTrace, msg)
}

func Info(msg interface{}) int {
	return output(LoggerLevelInfo, msg)
}

func Warn(msg interface{}) int {
	return output(LoggerLevelWarn, msg)
}

func Error(msg interface{}) int {
	return output(LoggerLevelError, msg)
}

func Fatal(msg interface{}) int {
	return output(LoggerLevelFatal, msg)
}
