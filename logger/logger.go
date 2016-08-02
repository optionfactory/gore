package logger

import (
	"fmt"
	"io"
	"log"
	"strings"
)

type Logger struct {
	ndc       []string
	mdc       map[string]string
	ndcprefix string
	mdcprefix string
	real      *log.Logger
}

func New(out io.Writer, prefix string, flag int) *Logger {
	return &Logger{
		make([]string, 0),
		make(map[string]string),
		"",
		"",
		log.New(out, prefix, flag),
	}
}

func (self *Logger) Mdc(values map[string]string) *Logger {
	newmdc := make(map[string]string)
	for k, v := range self.mdc {
		newmdc[k] = v
	}
	for k, v := range values {
		newmdc[k] = v
	}

	newmdcprefix := ""
	first := true
	for k, v := range newmdc {
		if !first {
			newmdcprefix += ","
		}
		newmdcprefix += k + "=" + v
		first = false
	}

	if newmdcprefix != "" {
		newmdcprefix = "[" + newmdcprefix + "]"
	}

	return &Logger{
		self.ndc,
		newmdc,
		self.ndcprefix,
		newmdcprefix,
		self.real,
	}
}

func (self *Logger) Ndc(values ...string) *Logger {
	newndc := make([]string, 0)
	newndc = append(newndc, self.ndc...)
	newndc = append(newndc, values...)
	newndcprefix := strings.Join(newndc, "][")
	if newndcprefix != "" {
		newndcprefix = "[" + newndcprefix + "]"
	}
	return &Logger{
		newndc,
		self.mdc,
		newndcprefix,
		self.mdcprefix,
		self.real,
	}
}

func (self *Logger) Trace(format string, args ...interface{}) {
	prefix := fmt.Sprintf("[TRACE]%s%s ", self.ndcprefix, self.mdcprefix)
	suffix := fmt.Sprintf(format, args...)
	self.real.Print(prefix, suffix, "\n")
}

func (self *Logger) Info(format string, args ...interface{}) {
	prefix := fmt.Sprintf("[INFO ]%s%s ", self.ndcprefix, self.mdcprefix)
	suffix := fmt.Sprintf(format, args...)
	self.real.Print(prefix, suffix, "\n")
}

func (self *Logger) Warning(format string, args ...interface{}) {
	prefix := fmt.Sprintf("[WARN ]%s%s ", self.ndcprefix, self.mdcprefix)
	suffix := fmt.Sprintf(format, args...)
	self.real.Print(prefix, suffix, "\n")
}

func (self *Logger) Error(format string, args ...interface{}) {
	prefix := fmt.Sprintf("[ERROR]%s%s ", self.ndcprefix, self.mdcprefix)
	suffix := fmt.Sprintf(format, args...)
	self.real.Print(prefix, suffix, "\n")
}
