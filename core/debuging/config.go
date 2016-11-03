package debuging

import (
	"fmt"
	"io"
)

type DebugConfig struct {
	MaxDepth int

	DisableMethods bool

	DisablePointerMethods bool

	ContinueOnMethod bool

	SortKeys bool

	SpewKeys bool
}

var Config = DebugConfig{}

func DefautConfig() *DebugConfig {
	return &DebugConfig{}
}

func (config *DebugConfig) Errorf(format string, p ...interface{}) (err error) {
	return fmt.Errorf(format, config.formatParams(p)...)
}

func (config *DebugConfig) Fprint(w io.Writer, p ...interface{}) (n int, err error) {
	return fmt.Fprint(w, config.formatParams(p)...)
}

func (config *DebugConfig) Fprintf(w io.Writer, format string, p ...interface{}) (n int, err error) {
	return fmt.Fprintf(w, format, config.formatParams(p)...)
}
func (config *DebugConfig) Fprintln(w io.Writer, p ...interface{}) (n int, err error) {
	return fmt.Fprintln(w, config.formatParams(p)...)
}

func (config *DebugConfig) Printf(format string, p ...interface{}) (n int, err error) {
	return fmt.Printf(format, config.formatParams(p))
}

func (config *DebugConfig) Print(p ...interface{}) (n int, err error) {
	return fmt.Print(config.formatParams(p)...)
}
func (config *DebugConfig) Println(p ...interface{}) (n int, err error) {
	return fmt.Println(config.formatParams(p)...)
}

func (config *DebugConfig) Sprintln(p ...interface{}) string {
	return fmt.Sprintln(config.formatParams(p)...)
}
func (config *DebugConfig) Sprintf(format string, p ...interface{}) string {
	return fmt.Sprintf(format, config.formatParams(p)...)
}
func (config *DebugConfig) Sprint(p ...interface{}) string {
	return fmt.Sprint(config.formatParams(p)...)
}

func (config *DebugConfig) formatParams(params []interface{}) (formatters []interface{}) {
	formatters = make([]interface{}, len(params))
	for index, p := range params {
		formatters[index] = NewFmtPrintWriter(config, p)
	}
	return formatters
}
