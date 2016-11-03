package debuging

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"runtime"
	"strconv"
	"strings"
)

const (
	dunnostr string = "???"
	sepStr   string = "/"
)

var (
	runtimePath string
	isWindowsOS bool = runtime.GOOS == "windows"
	ErrNoFunc        = errors.New("no call stack information")
	dunno            = []byte("???")
	dot              = []byte(".")
	slash            = []byte("/")

	sigpanic *runtime.Func
)

func init() {
	var pcs [1]uintptr

	runtime.Callers(0, pcs[:])
	fn := runtime.FuncForPC(pcs[0])

	file, _ := fn.FileLine(pcs[0])
	idx := pkgIndex(file, fn.Name())

	runtimePath = file[:idx]

	if isWindowsOS {
		runtimePath = strings.ToLower(runtimePath)
	}
}

type StackFrame struct {
	fn *runtime.Func
	pc uintptr
}
type Stack []StackFrame

func GetStack() Stack {
	var pcs [512]uintptr

	n := runtime.Callers(2, pcs[:])

	frames := make([]StackFrame, n)

	for index, pc := range pcs[:n] {
		fix := pc
		if index > 0 && !isSigpanic(frames[index-1].fn) {
			fix--
		}
		frames[index] = StackFrame{
			fn: runtime.FuncForPC(fix),
			pc: fix,
		}
	}

	return frames
}

func GetStackFrame(skip int) (item StackFrame) {
	var pcs [2]uintptr
	if runtime.Callers(skip+1, pcs[:]) < 2 {
		return item
	}

	item.pc = pcs[1]

	if !isSigpanic(runtime.FuncForPC(pcs[0])) {
		item.pc--
	}
	item.fn = runtime.FuncForPC(item.pc)
	return item
}

//////////////////////////////////////////////////////////////////////////////////////////////
func (i StackFrame) Format(fs fmt.State, verb rune) {
	if i.fn == nil {
		fmt.Fprintf(fs, "%%!%c(NO_FUNC)", verb)
		return
	}

	switch verb {
	case 's', 'v':
		file, line := i.FileLine()
		switch {
		case fs.Flag('#'): /*Do Nothing*/
		case fs.Flag('+'):
			file = file[pkgIndex(file, i.fn.Name()):]
		default:
			if i := strings.LastIndexByte(file, slash[0]); i != -1 {
				file = file[i+1:]
			}
		}
		io.WriteString(fs, file)
		switch verb {
		case 'v':
			buf := [7]byte{':'}
			fs.Write(strconv.AppendInt(buf[:1], int64(line), 10))
		}

		switch {
		case fs.Flag('#'):
			name := i.Name()
			if i := strings.LastIndexByte(name, slash[0]); i != -1 {
				name = name[i+1:]
			}
			if i := strings.IndexByte(name, dot[0]); i != -1 {
				name = name[i+1:]
			}
			fs.Write(pointerChainBytes)
			fs.Write(openBracketBytes)
			io.WriteString(fs, name)
			fs.Write(closeBracketBytes)
		default: /*DO nothing*/
		}

	case 'd':
		buf := [6]byte{}
		fs.Write(strconv.AppendInt(buf[:0], int64(i.Line()), 10))
	case 'n':
		name := i.Name()
		switch {
		case fs.Flag('#'): /*Do nonthing*/
		default:
			if i := strings.LastIndexByte(name, slash[0]); i != -1 {
				name = name[i+1:]
			}
			if i := strings.IndexByte(name, dot[0]); i != -1 {
				name = name[i+1:]
			}
		}
		io.WriteString(fs, name)
	}
}

// MarshalText implements encoding.TextMarshaler. It formats the Call the same
// as fmt.Sprintf("%v",i).
func (i StackFrame) MarshalText() ([]byte, error) {
	if i.fn == nil {
		return nil, ErrNoFunc
	}
	buf := bytes.Buffer{}
	fmt.Fprint(&buf, i)
	return buf.Bytes(), nil
}

func (i StackFrame) String() string {
	return fmt.Sprint(i)
}

func (i StackFrame) Name() string {
	if i.fn == nil {
		return dunnostr
	}
	return i.fn.Name()
}

func (i StackFrame) File() string {
	if i.fn == nil {
		return dunnostr
	}
	file, _ := i.fn.FileLine(i.pc)
	return file
}
func (i StackFrame) Line() int {
	if i.fn == nil {
		return 0
	}
	_, line := i.fn.FileLine(i.pc)
	return line
}

func (i StackFrame) FileLine() (string, int) {
	if i.fn == nil {
		return dunnostr, 0
	}
	return i.fn.FileLine(i.pc)
}

func (i StackFrame) PC() uintptr {
	if i.fn == nil {
		return 0
	}
	return i.pc
}
func (i StackFrame) isGoRoot() bool {
	file := i.File()
	if len(file) == 0 || file[0] == '?' {
		return true
	}
	if isWindowsOS {
		file = strings.ToLower(file)
	}
	return strings.HasPrefix(file, runtimePath) ||
		strings.HasPrefix(file, "/_testmain.go")
}

//////////////////////////////////////////////////////////////////////////////////////////////
func (s Stack) Format(fs fmt.State, verb rune) {
	fs.Write(openBracketBytes)
	numEntries := len(s)
	if numEntries > 0 {
		fs.Write(newlineBytes)
	}
	for i, item := range s {
		item.Format(fs, verb)
		if i < numEntries {
			fs.Write(commaNewlineBytes)
		} else {
			fs.Write(newlineBytes)
		}
	}
	fs.Write(closeBracketBytes)
}

// MarshalText implements encoding.TextMarshaler. It formats the CallStack the
// same as fmt.Sprintf("%v", s).
func (s Stack) MarshalText() ([]byte, error) {
	buf := bytes.Buffer{}
	buf.Write(openBracketBytes)
	for i, item := range s {
		if item.fn == nil {
			return nil, ErrNoFunc
		}
		if i > 0 {
			buf.Write(spaceBytes)
		}
		fmt.Fprint(&buf, item)
	}
	buf.Write(closeBracketBytes)
	return buf.Bytes(), nil
}
func (s Stack) String() string {
	return fmt.Sprint(s)
}
func (s Stack) ExcludeRuntime() Stack {
	for len(s) > 0 && s[len(s)-1].isGoRoot() {
		s = s[:len(s)-1]
	}
	return s
}

// TrimBelow returns a slice of the CallStack with all entries below c
// removed.
func (s Stack) TrimBelow(item StackFrame) Stack {
	for len(s) > 0 && s[0].pc != item.pc {
		s = s[1:]
	}
	return s
}

// TrimAbove returns a slice of the CallStack with all entries above c
// removed.
func (s Stack) TrimAbove(item StackFrame) Stack {
	for len(s) > 0 && s[len(s)-1].pc != item.pc {
		s = s[:len(s)-1]
	}
	return s
}

//-------------------------------------------------------------------------//

// pkgIndex returns the index that results in file[index:] being the path of
// file relative to the compile time GOPATH, and file[:index] being the
// $GOPATH/src/ portion of file. funcName must be the name of a function in
// file as returned by runtime.Func.Name.
func pkgIndex(file, funcName string) int {
	// As of Go 1.6.2 there is no direct way to know the compile time GOPATH
	// at runtime, but we can infer the number of path segments in the GOPATH.
	// We note that runtime.Func.Name() returns the function name qualified by
	// the import path, which does not include the GOPATH. Thus we can trim
	// segments from the beginning of the file path until the number of path
	// separators remaining is one more than the number of path separators in
	// the function name. For example, given:
	//
	//    GOPATH     /home/user
	//    file       /home/user/src/pkg/sub/file.go
	//    fn.Name()  pkg/sub.Type.Method
	//
	// We want to produce:
	//
	//    file[:idx] == /home/user/src/
	//    file[idx:] == pkg/sub/file.go
	//
	// From this we can easily see that fn.Name() has one less path separator
	// than our desired result for file[idx:]. We count separators from the
	// end of the file path until it finds two more than in the function name
	// and then move one character forward to preserve the initial path
	// segment without a leading separator.
	i := len(file)
	for n := strings.Count(funcName, sepStr) + 2; n > 0; n-- {
		i = strings.LastIndexByte(file[:i], slash[0])
		if i == -1 {
			i = -len(sepStr)
			break
		}
	}
	// get back to 0 or trim the leading separator
	return i + len(sepStr)
}

func isSigpanic(fn *runtime.Func) bool {
	if sigpanic == nil {
		//sigpanic = findSigpanic()
	}
	return fn == sigpanic
}

func findSigpanic() (fn *runtime.Func) {
	func() runtime.Func {
		defer func() {
			if err := recover(); err != nil {
				var pcs [512]uintptr
				n := runtime.Callers(2, pcs[:])
				for _, pc := range pcs[:n] {
					f := runtime.FuncForPC(pc)
					if f.Name() == "runtime.sigpanic" {
						fn = f
						break
					}
				}
			}
		}()

		return *fn
	}()
	return fn
}

//-------------------------------------------------------------------------//
