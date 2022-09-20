package testlog

import (
	"strconv"
	"testing"
)

const (
	checkMark = "\u2713" // √
	ballotX   = "\u2717" // ×
)

type Logger struct {
	caseNum int
	prefix  string
	*testing.T
}

func New(t *testing.T) *Logger {
	return &Logger{T: t}
}

func (l *Logger) Title(format string, args ...any) {
	l.Logf("Test Case => "+format, args...)
}

func (l *Logger) Case(format string, args ...any) {
	l.caseNum++
	l.prefix = "Case " + strconv.Itoa(l.caseNum) + " -> "
	l.Logf(l.prefix+format, args...)
}

func (l *Logger) Pass(format string, args ...any) {
	l.Logf("\t%s "+format, prependTag(checkMark, args)...)
}

func (l *Logger) Fail(format string, args ...any) {
	l.Errorf("\t%s "+format, prependTag(ballotX, args)...)
}

func (l *Logger) Quit(format string, args ...any) {
	prependTag(ballotX, args)
	l.Fatalf("\t%s "+format, prependTag(ballotX, args)...)
}

func prependTag(tag any, args []any) []any { // TODO 这里的args如果为可变参数，则传递nil是args长度为1？
	if args == nil {
		args = make([]any, 1)
		args[0] = tag
	} else {
		args = append([]any{tag}, args...)
	}
	return args
}

func (l *Logger) Require(cond bool, desc string, args ...any) {
	if cond {
		l.Pass(desc, args...)
	} else {
		l.Fail(desc, args...)
	}
}
