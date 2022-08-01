package logger

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"runtime"
	"time"
)

type Level int8
type Fields map[string]any

const (
	LevelDebug Level = iota
	LevelInfo
	LevelWarn
	LevelError
	LevelFatal
	LevelPanic
)

func (l Level) String() string {
	switch l {
	case LevelInfo:
		return "info"
	case LevelWarn:
		return "warn"
	case LevelError:
		return "error"
	case LevelDebug:
		return "debug"
	case LevelFatal:
		return "fatal"

	default:
		return ""
	}
}

type Logger struct {
	newLogger *log.Logger
	ctx       context.Context
	level     Level
	fields    Fields
	callers   []string
}

func NewLogger(w io.Writer, prefix string, flag int) *Logger {
	l := log.New(w, prefix, flag)
	return &Logger{newLogger: l}
}

func (l *Logger) clone() *Logger {
	cl := *l
	return &cl
}
func (l *Logger) WithLevel(level Level) *Logger {
	clone := l.clone()
	clone.level = level
	return clone
}

func (l *Logger) WithFields(f Fields) *Logger {
	clone := l.clone()
	if clone.fields == nil {
		clone.fields = make(Fields)
	}
	for k, v := range f {
		clone.fields[k] = v
	}
	return clone
}

func (l *Logger) WithContext(ctx context.Context) *Logger {
	clone := l.clone()
	clone.ctx = ctx
	return clone
}

func (l *Logger) WithCaller(skip int) *Logger {
	clone := l.clone()
	pc, file, line, ok := runtime.Caller(skip)
	if ok {
		f := runtime.FuncForPC(pc)
		clone.callers = []string{fmt.Sprintf("%s; %d %s", file, line, f.Name())}
	}
	return clone
}

func (l *Logger) WithCallersFrames() *Logger {
	maxCallerDepth := 25
	minCallerDepth := 1
	var callers []string

	pcs := make([]uintptr, maxCallerDepth)
	depth := runtime.Callers(minCallerDepth, pcs)
	frames := runtime.CallersFrames(pcs[:depth])
	for frame, more := frames.Next(); more; frame, more = frames.Next() {
		callers = append(callers, fmt.Sprintf("%s: %d :%s", frame.File, frame.Line, frame.Function))
		if !more {
			break
		}
	}
	clone := l.clone()
	clone.callers = callers
	return clone
}

func (l *Logger) JsonFormat(msg string) map[string]any {
	data := make(Fields, len(l.fields)+4)
	data["level"] = l.level.String()
	data["time"] = time.Now().Local().UnixNano()
	data["msg"] = msg
	data["callers"] = l.callers
	if len(l.fields) > 0 {
		for k, v := range l.fields {
			if _, ok := data[k]; ok {
				data[k] = v
			}
		}
	}
	return data
}

func (l *Logger) Output(msg string) {
	body, _ := json.Marshal(l.JsonFormat(msg))
	content := string(body)
	switch l.level {
	case LevelInfo:
		l.newLogger.Print(content)
	case LevelError:
		l.newLogger.Print(content)
	case LevelWarn:
		l.newLogger.Print(content)
	case LevelFatal:
		l.newLogger.Fatal(content)
	case LevelPanic:
		l.newLogger.Panic(content)
	}
}

func (l *Logger) Debug(v ...any) {
	l.WithLevel(LevelDebug).Output(fmt.Sprint(v...))
}

func (l *Logger) DebugF(format string, v ...any) {
	l.WithLevel(LevelDebug).Output(fmt.Sprintf(format, v...))
}

func (l *Logger) Info(v ...any) {
	l.WithLevel(LevelInfo).Output(fmt.Sprint(v...))
}

func (l *Logger) InfoF(format string, v ...any) {
	l.WithLevel(LevelInfo).Output(fmt.Sprintf(format, v...))
}

func (l *Logger) Error(v ...any) {
	l.WithLevel(LevelError).Output(fmt.Sprint(v...))
}

func (l *Logger) ErrorF(format string, v ...any) {
	l.WithLevel(LevelError).Output(fmt.Sprintf(format, v...))
}

func (l *Logger) Fatal(v ...any) {
	l.WithLevel(LevelFatal).Output(fmt.Sprint(v...))
}

func (l *Logger) Fatalf(format string, v ...any) {
	l.WithLevel(LevelFatal).Output(fmt.Sprintf(format, v...))
}
func (l *Logger) Panic(v ...any) {
	l.WithLevel(LevelPanic).Output(fmt.Sprint(v...))
}

func (l *Logger) PanicF(format string, v ...any) {
	l.WithLevel(LevelPanic).Output(fmt.Sprintf(format, v...))
}
