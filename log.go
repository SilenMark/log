package log

import (
	"fmt"
	"log"
	"os"
)

const (
	GRAY, _           = ANSICode(iota + 30), ANSICode(iota + 30) //gray
	RED, erro_code                                               //red
	GREEN, succ_code                                             //green
	YELLOW, warn_code                                            //yellow
	BLUE, info_code                                              //blue
	MAGENTA, _                                                   // 粉色
	CYAN, _                                                      //青色
	WHITE, _
	info_text = "[INFO]"
	erro_text = "[ERRO]"
	succ_text = "[SUCC]"
	warn_text = "[WARN]"
)

var (
	info = flag{Code: info_code, Text: info_text}
	erro = flag{Code: erro_code, Text: erro_text}
	succ = flag{Code: succ_code, Text: succ_text}
	warn = flag{Code: warn_code, Text: warn_text}
)

var (
	l   = log.New(os.Stderr, "", log.LstdFlags)
	std = &Logger{true, l}
)

type Any = interface{}

type Logger struct {
	showColor bool
	*log.Logger
}

func (l *Logger) HasColor() bool {
	return l.showColor
}

func (l *Logger) Output(f flag, s string) error {
	if l.HasColor() {
		return l.Logger.Output(2, f.Code.Render(f.Text)+s)
	}
	return l.Logger.Output(2, f.Text+s)
}
func (l *Logger) Info(v ...Any) {
	l.Output(info, fmt.Sprint(v...))
}
func (l *Logger) Infoln(v ...Any) {
	l.Output(info, fmt.Sprintln(v...))
}
func (l *Logger) Infof(format string, v ...Any) {
	l.Output(info, fmt.Sprintf(format, v...))
}
func (l *Logger) Warn(v ...Any) {
	l.Output(warn, fmt.Sprint(v...))
}
func (l *Logger) Warnln(v ...Any) {
	l.Output(warn, fmt.Sprintln(v...))
}
func (l *Logger) Warnf(format string, v ...Any) {
	l.Output(warn, fmt.Sprintf(format, v...))
}
func (l *Logger) Erro(v ...Any) {
	l.Output(erro, fmt.Sprint(v...))
}
func (l *Logger) Erroln(v ...Any) {
	l.Output(erro, fmt.Sprintln(v...))
}
func (l *Logger) Errof(format string, v ...Any) {
	l.Output(erro, fmt.Sprintf(format, v...))
}
func (l *Logger) Succ(v ...Any) {
	l.Output(succ, fmt.Sprint(v...))
}
func (l *Logger) Succln(v ...Any) {
	l.Output(succ, fmt.Sprintln(v...))
}
func (l *Logger) Succf(format string, v ...Any) {
	l.Output(succ, fmt.Sprintf(format, v...))
}
func Info(v ...Any) {
	std.Info(v...)
}
func Infof(format string, v ...Any) {
	std.Infof(format, v...)
}
func Infoln(v ...Any) {
	std.Infoln(v...)
}
func Erro(v ...Any) {
	std.Erro(v...)
}
func Errof(format string, v ...Any) {
	std.Errof(format, v...)
}
func Erroln(v ...Any) {
	std.Erroln(v...)
}
func Warn(v ...Any) {
	std.Warn(v...)
}
func Warnf(format string, v ...Any) {
	std.Warnf(format, v...)
}
func Warnln(v ...Any) {
	std.Warnln(v...)
}
func Succ(v ...Any) {
	std.Succ(v...)
}
func Succf(format string, v ...Any) {
	std.Succf(format, v...)
}
func Succln(v ...Any) {
	std.Succln(v...)
}
