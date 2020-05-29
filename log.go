package echo

import (
	"github.com/labstack/gommon/log"
	"io"
)

type (
	Logger interface {
		Output() io.Writer
		SetOutput(w io.Writer)

		Prefix() string
		SetPrefix(p string)

		Level() log.Lvl
		SetLevel(log.Lvl)

		Print(i ...interface{})
		Printf(format string, args ...interface{})
		Printj(j log.JSON)

		Debug(i ...interface{})
		Debugf(format string, args ...interface{})
		Debugj(j log.JSON)

		Info(i ...interface{})
		Infof(format string, args ...interface{})
		Infoj(j log.JSON)

		Warn(i ...interface{})
		Warnf(format string, args ...interface{})
		Warnj(j log.JSON)

		Error(i ...interface{})
		Errorf(format string, args ...interface{})
		Errorj(l log.JSON)

		Fatal(i ...interface{})
		Fatalf(format string, args ...interface{})
		Fatalj(l log.JSON)

		Panic(i ...interface{})
		Panicf(format string, args ...interface{})
		Panicj(l log.JSON)
	}
)
