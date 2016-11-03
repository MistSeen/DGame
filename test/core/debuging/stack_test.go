package debuging

import (
	"fmt"
	"io/ioutil"
	"path"
	"path/filepath"
	"reflect"
	"runtime"
	. "server/core/debuging"
	"strings"
	"testing"
)

var importPath = "server/test/core/debuging"

type testType struct{}

func (tt testType) testMethod() (c StackFrame, pc uintptr, file string, line int, ok bool) {
	c = GetStackFrame(0)
	pc, file, line, ok = runtime.Caller(0)
	line--
	return
}

func TestStackInfo(t *testing.T) {
	t.Parallel()

	c := GetStackFrame(0)
	_, file, line, ok := runtime.Caller(0)
	line--
	if !ok {
		t.Fatal("runtime.Caller(0) failed")
	}

	if got, want := c.File(), file; got != want {
		t.Errorf("test file:got == %v, want == %v", got, want)
	}

	if got, want := c.Line(), line; got != want {
		t.Errorf("test line:got == %v, want == %v", got, want)
	}
}

type fholder struct {
	f func() Stack
}

func (fh *fholder) labyrinth() Stack {
	for {
		return fh.f()
	}
	panic("this line only needed for go 1.0")
}

func TestTrace(t *testing.T) {
	t.Parallel()

	fh := fholder{
		f: func() Stack {
			cs := GetStack()
			return cs
		},
	}

	cs := fh.labyrinth()

	lines := []int{61, 51, 66}

	for i, line := range lines {
		if got, want := cs[i].Line(), line; got != want {
			t.Errorf("got line[%d] == %v, want line[%d] == %v", i, got, i, want)
		}
	}
}

func TestCallFormat(t *testing.T) {
	t.Parallel()

	c := GetStackFrame(0)
	pc, file, line, ok := runtime.Caller(0)
	line--
	if !ok {
		t.Fatal("runtime.Caller(0) failed")
	}
	relFile := path.Join(importPath, filepath.Base(file))

	c2, pc2, file2, line2, ok2 := testType{}.testMethod()
	if !ok2 {
		t.Fatal("runtime.Caller(0) failed")
	}
	relFile2 := path.Join(importPath, filepath.Base(file2))
	runtime.FuncForPC(pc2).Name()
	runtime.FuncForPC(pc - 1).Name()
	data := []struct {
		c    StackFrame
		desc string
		fmt  string
		out  string
	}{
		{StackFrame{}, "error", "%s", "%!s(NO_FUNC)"},

		{c, "func", "%s", path.Base(file)},
		{c, "func", "%+s", relFile},
		{c, "func", "%#s", file},
		{c, "func", "%d", fmt.Sprint(line)},
		{c, "func", "%n", "TestCallFormat"},
		{c, "func", "%+n", "TestCallFormat" /*runtime.FuncForPC(pc - 1).Name()*/},
		{c, "func", "%v", fmt.Sprint(path.Base(file), ":", line)},
		{c, "func", "%+v", fmt.Sprint(relFile, ":", line)},
		{c, "func", "%#v", fmt.Sprint(file, ":", line)},

		{c2, "meth", "%s", path.Base(file2)},
		{c2, "meth", "%+s", relFile2},
		{c2, "meth", "%#s", file2},
		{c2, "meth", "%d", fmt.Sprint(line2)},
		{c2, "meth", "%n", "testType.testMethod"},
		{c2, "meth", "%+n", "testType.testMethod" /*runtime.FuncForPC(pc2).Name()*/},
		{c2, "meth", "%v", fmt.Sprint(path.Base(file2), ":", line2)},
		{c2, "meth", "%+v", fmt.Sprint(relFile2, ":", line2)},
		{c2, "meth", "%#v", fmt.Sprint(file2, ":", line2)},
	}

	for _, d := range data {
		got := fmt.Sprintf(d.fmt, d.c)
		if got != d.out {
			t.Errorf("fmt.Sprintf(%q, Call(%s)) =[%s], want:[%s]", d.fmt, d.desc, got, d.out)
		}
	}
}

func TestCallString(t *testing.T) {
	t.Parallel()

	c := GetStackFrame(0)
	_, file, line, ok := runtime.Caller(0)
	line--
	if !ok {
		t.Fatal("runtime.Caller(0) failed")
	}

	c2, _, file2, line2, ok2 := testType{}.testMethod()
	if !ok2 {
		t.Fatal("runtime.Caller(0) failed")
	}

	data := []struct {
		c    StackFrame
		desc string
		out  string
	}{
		{StackFrame{}, "error", "%!v(NO_FUNC)"},
		{c, "func", fmt.Sprint(path.Base(file), ":", line)},
		{c2, "meth", fmt.Sprint(path.Base(file2), ":", line2)},
	}

	for _, d := range data {
		got := d.c.String()
		if got != d.out {
			t.Errorf("got %s, want %s", got, d.out)
		}
	}
}

func TestCallMarshalText(t *testing.T) {
	t.Parallel()

	c := GetStackFrame(0)
	_, file, line, ok := runtime.Caller(0)
	line--
	if !ok {
		t.Fatal("runtime.Caller(0) failed")
	}

	c2, _, file2, line2, ok2 := testType{}.testMethod()
	if !ok2 {
		t.Fatal("runtime.Caller(0) failed")
	}

	data := []struct {
		c    StackFrame
		desc string
		out  []byte
		err  error
	}{
		{StackFrame{}, "error", nil, ErrNoFunc},
		{c, "func", []byte(fmt.Sprint(path.Base(file), ":", line)), nil},
		{c2, "meth", []byte(fmt.Sprint(path.Base(file2), ":", line2)), nil},
	}

	for _, d := range data {
		text, err := d.c.MarshalText()
		if got, want := err, d.err; got != want {
			t.Errorf("%s: got err %v, want err %v", d.desc, got, want)
		}
		if got, want := text, d.out; !reflect.DeepEqual(got, want) {
			t.Errorf("%s: got %s, want %s", d.desc, got, want)
		}
	}
}

func TestCallStackString(t *testing.T) {
	cs, line0 := getTrace(t)
	_, file, line1, ok := runtime.Caller(0)
	line1--
	if !ok {
		t.Fatal("runtime.Caller(0) failed")
	}
	file = path.Base(file)
	if got, want := cs.String(), fmt.Sprintf("[%s:%d %s:%d]", file, line0, file, line1); got != want {
		t.Errorf("\n got %v\nwant %v", got, want)
	}
}

func TestCallStackMarshalText(t *testing.T) {
	cs, line0 := getTrace(t)
	_, file, line1, ok := runtime.Caller(0)
	line1--
	if !ok {
		t.Fatal("runtime.Caller(0) failed")
	}
	file = path.Base(file)
	text, _ := cs.MarshalText()
	if got, want := text, []byte(fmt.Sprintf("[%s:%d %s:%d]", file, line0, file, line1)); !reflect.DeepEqual(got, want) {
		t.Errorf("\n got %v\nwant %v", got, want)
	}
}
func getTrace(t *testing.T) (Stack, int) {
	cs := GetStack().ExcludeRuntime()
	_, _, line, ok := runtime.Caller(0)
	line--
	if !ok {
		t.Fatal("runtime.Caller(0) failed")
	}
	return cs, line
}

func TestTrimAbove(t *testing.T) {
	trace := trimAbove()
	if got, want := len(trace), 2; got != want {
		t.Errorf("got len(trace) == %v, want %v, trace: %n", got, want, trace)
	}
	if got, want := fmt.Sprintf("%n", trace[1]), "TestTrimAbove"; got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func trimAbove() Stack {
	call := GetStackFrame(1)
	trace := GetStack()
	return trace.TrimAbove(call)
}

func TestTrimBelow(t *testing.T) {
	trace := trimBelow()
	if got, want := fmt.Sprintf("%n", trace[0]), "TestTrimBelow"; got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func trimBelow() Stack {
	call := GetStackFrame(1)
	trace := GetStack()
	return trace.TrimBelow(call)
}

func TestTrimRuntime(t *testing.T) {
	trace := GetStack().ExcludeRuntime()
	if got, want := len(trace), 1; got != want {
		t.Errorf("got len(trace) == %v, want %v, goroot: %q, trace: %#v", got, want, runtime.GOROOT(), trace)
	}
}

func BenchmarkCallVFmt(b *testing.B) {
	c := GetStackFrame(0)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		fmt.Fprint(ioutil.Discard, c)
	}
}

func BenchmarkCallPlusVFmt(b *testing.B) {
	c := GetStackFrame(0)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		fmt.Fprintf(ioutil.Discard, "%+v", c)
	}
}

func BenchmarkCallSharpVFmt(b *testing.B) {
	c := GetStackFrame(0)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		fmt.Fprintf(ioutil.Discard, "%#v", c)
	}
}

func BenchmarkCallSFmt(b *testing.B) {
	c := GetStackFrame(0)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		fmt.Fprintf(ioutil.Discard, "%s", c)
	}
}

func BenchmarkCallPlusSFmt(b *testing.B) {
	c := GetStackFrame(0)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		fmt.Fprintf(ioutil.Discard, "%+s", c)
	}
}

func BenchmarkCallSharpSFmt(b *testing.B) {
	c := GetStackFrame(0)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		fmt.Fprintf(ioutil.Discard, "%#s", c)
	}
}

func BenchmarkCallDFmt(b *testing.B) {
	c := GetStackFrame(0)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		fmt.Fprintf(ioutil.Discard, "%d", c)
	}
}

func BenchmarkCallNFmt(b *testing.B) {
	c := GetStackFrame(0)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		fmt.Fprintf(ioutil.Discard, "%n", c)
	}
}

func BenchmarkCallPlusNFmt(b *testing.B) {
	c := GetStackFrame(0)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		fmt.Fprintf(ioutil.Discard, "%+n", c)
	}
}

func BenchmarkCaller(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GetStackFrame(0)
	}
}

func BenchmarkTrace(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GetStack()
	}
}

func deepStack(depth int, b *testing.B) Stack {
	if depth > 0 {
		return deepStack(depth-1, b)
	}
	b.StartTimer()
	s := GetStack()
	return s
}

func BenchmarkTrace10(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		deepStack(10, b)
	}
}

func BenchmarkTrace50(b *testing.B) {
	b.StopTimer()
	for i := 0; i < b.N; i++ {
		deepStack(50, b)
	}
}

func BenchmarkTrace100(b *testing.B) {
	b.StopTimer()
	for i := 0; i < b.N; i++ {
		deepStack(100, b)
	}
}

////////////////
// Benchmark functions followed by formatting
////////////////

func BenchmarkCallerAndVFmt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fmt.Fprint(ioutil.Discard, GetStackFrame(0))
	}
}

func BenchmarkTraceAndVFmt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fmt.Fprint(ioutil.Discard, GetStack())
	}
}

func BenchmarkTrace10AndVFmt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		fmt.Fprint(ioutil.Discard, deepStack(10, b))
	}
}

////////////////
// Baseline against package runtime.
////////////////

func BenchmarkRuntimeCaller(b *testing.B) {
	for i := 0; i < b.N; i++ {
		runtime.Caller(0)
	}
}

func BenchmarkRuntimeCallerAndFmt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, file, line, _ := runtime.Caller(0)
		const sep = "/"
		if i := strings.LastIndex(file, sep); i != -1 {
			file = file[i+len(sep):]
		}
		fmt.Fprint(ioutil.Discard, file, ":", line)
	}
}

func BenchmarkFuncForPC(b *testing.B) {
	pc, _, _, _ := runtime.Caller(0)
	pc--
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		runtime.FuncForPC(pc)
	}
}

func BenchmarkFuncFileLine(b *testing.B) {
	pc, _, _, _ := runtime.Caller(0)
	pc--
	fn := runtime.FuncForPC(pc)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		fn.FileLine(pc)
	}
}
