package logging

/**
 * ***********************************************
 * Author: RichardKnop
 * Description: A simple leveled logging library with coloured output.
 * Address: go get github.com/RichardKnop/logging
 * ************************************************
 */
import (
	"io"
	"log"
	"os"
)

// Level type
type level int

const (
	// DEBUG level
	DEBUG level = iota
	// INFO level
	INFO
	// WARNING level
	WARNING
	// ERROR level
	ERROR
	// FATAL level
	FATAL

	flag = log.Ldate | log.Ltime
)

// Log level prefix map
var prefix = map[level]string{
	DEBUG:   "DEBUG: ",
	INFO:    "INFO: ",
	WARNING: "WARNING: ",
	ERROR:   "ERROR: ",
	FATAL:   "FATAL: ",
}

// Logger ...
type Logger map[level]LoggerInterface

// New returns instance of Logger
func New(out, errOut io.Writer, f Formatter) Logger {
	// Fall back to stdout if out not set
	if out == nil {
		out = os.Stdout
	}

	// Fall back to stderr if errOut not set
	if errOut == nil {
		errOut = os.Stderr
	}

	// Fall back to DefaultFormatter if f not set
	if f == nil {
		f = new(DefaultFormatter)
	}

	l := make(map[level]LoggerInterface, 5)
	l[DEBUG] = &Wrapper{lvl: DEBUG, formatter: f, logger: log.New(out, GetPrefix(DEBUG)+prefix[DEBUG], flag)}
	l[INFO] = &Wrapper{lvl: INFO, formatter: f, logger: log.New(out, GetPrefix(INFO)+prefix[INFO], flag)}
	l[WARNING] = &Wrapper{lvl: INFO, formatter: f, logger: log.New(out, GetPrefix(WARNING)+prefix[WARNING], flag)}
	l[ERROR] = &Wrapper{lvl: INFO, formatter: f, logger: log.New(errOut, GetPrefix(ERROR)+prefix[ERROR], flag)}
	l[FATAL] = &Wrapper{lvl: INFO, formatter: f, logger: log.New(errOut, GetPrefix(FATAL)+prefix[FATAL], flag)}

	return Logger(l)
}

// Wrapper ...
type Wrapper struct {
	lvl       level
	formatter Formatter
	logger    LoggerInterface
}

// Print ...
func (w *Wrapper) Print(v ...interface{}) {
	v = Format(w.lvl, v...)
	v = append(v, GetSuffix(w.lvl))
	Print(v...)
}

// Printf ...
func (w *Wrapper) Printf(format string, v ...interface{}) {
	suffix := GetSuffix(w.lvl)
	v = Format(w.lvl, v...)
	Printf("%s"+format+suffix, v...)
}

// Println ...
func (w *Wrapper) Println(v ...interface{}) {
	v = Format(w.lvl, v...)
	v = append(v, GetSuffix(w.lvl))
	Println(v...)
}

// Fatal ...
func (w *Wrapper) Fatal(v ...interface{}) {
	v = Format(w.lvl, v...)
	v = append(v, GetSuffix(w.lvl))
	Fatal(v...)
}

// Fatalf ...
func (w *Wrapper) Fatalf(format string, v ...interface{}) {
	suffix := GetSuffix(w.lvl)
	v = Format(w.lvl, v...)
	Fatalf("%s"+format+suffix, v...)
}

// Fatalln ...
func (w *Wrapper) Fatalln(v ...interface{}) {
	v = Format(w.lvl, v...)
	v = append(v, GetSuffix(w.lvl))
	Fatalln(v...)
}

// Panic ...
func (w *Wrapper) Panic(v ...interface{}) {
	v = Format(w.lvl, v...)
	v = append(v, GetSuffix(w.lvl))
	Fatal(v...)
}

// Panicf ...
func (w *Wrapper) Panicf(format string, v ...interface{}) {
	suffix := GetSuffix(w.lvl)
	v = Format(w.lvl, v...)
	Panicf("%s"+format+suffix, v...)
}

// Panicln ...
func (w *Wrapper) Panicln(v ...interface{}) {
	v = Format(w.lvl, v...)
	v = append(v, GetSuffix(w.lvl))
	Panicln(v...)
}
