package debuging

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"server/core/debuging"
	"testing"
)

func TestParamFormat(t *testing.T) {
	// p := &ParamFormat{}

	// if p.isNextIngore(ingore_next_type) {
	// 	t.Error(" need ingore_next_type ")
	// }
	// if p.isNextIngore(ingore_next_lencap) {
	// 	t.Error(" do need ingore_next_lencap ")
	// }
	// if !p.isIngoreNone() {
	// 	t.Logf(" no  isIngoreNone()")
	// }

	// p.setNextIngoreState(ingore_next_type, true)

	// if !p.isNextIngore(ingore_next_type) {
	// 	t.Error(" need ingore_next_type ")
	// }
	// if p.isNextIngore(ingore_next_lencap) {
	// 	t.Error(" do need ingore_next_lencap ")
	// }
	// if p.isIngoreNone() {
	// 	t.Error(" no  isIngoreNone()")
	// }

	// p.setIngoreNone()

	// if p.isNextIngore(ingore_next_type) {
	// 	t.Error(" need ingore_next_type ")
	// }
	// if p.isNextIngore(ingore_next_lencap) {
	// 	t.Error(" do need ingore_next_lencap ")
	// }
	// if !p.isIngoreNone() {
	// 	t.Error(" no  isIngoreNone()")
	// }

	// p.setNextIngoreState(ingore_next_lencap, true)

	// if p.isNextIngore(ingore_next_type) {
	// 	t.Error(" need ingore_next_type ")
	// }
	// if !p.isNextIngore(ingore_next_lencap) {
	// 	t.Error(" do need ingore_next_lencap ")
	// }
	// if p.isIngoreNone() {
	// 	t.Error(" no  isIngoreNone()")
	// }

	// p.setIngoreNone()

	// if p.isNextIngore(ingore_next_type) {
	// 	t.Error(" need ingore_next_type ")
	// }
	// if p.isNextIngore(ingore_next_lencap) {
	// 	t.Error(" do need ingore_next_lencap ")
	// }
	// if !p.isIngoreNone() {
	// 	t.Error(" no  isIngoreNone()")
	// }
	// p.setNextIngoreState(ingore_next_all, true)
	// if !p.isNextIngore(ingore_next_type) {
	// 	t.Error(" need ingore_next_type ")
	// }
	// if !p.isNextIngore(ingore_next_lencap) {
	// 	t.Error(" do need ingore_next_lencap ")
	// }
	// if p.isIngoreNone() {
	// 	t.Error(" no  isIngoreNone()")
	// }

	// p.setNextIngoreState(ingore_next_type, true)
	// if !p.isNextIngore(ingore_next_type) {
	// 	t.Error(" need ingore_next_type ")
	// }
	// if !p.isNextIngore(ingore_next_lencap) {
	// 	t.Error(" do need ingore_next_lencap ")
	// }
	// if p.isIngoreNone() {
	// 	t.Error(" no  isIngoreNone()")
	// }

	// p.setNextIngoreState(ingore_next_lencap, true)
	// if !p.isNextIngore(ingore_next_type) {
	// 	t.Error(" need ingore_next_type ")
	// }
	// if !p.isNextIngore(ingore_next_lencap) {
	// 	t.Error(" do need ingore_next_lencap ")
	// }
	// if p.isIngoreNone() {
	// 	t.Error(" no  isIngoreNone()")
	// }

	// p.setNextIngoreState(ingore_next_type, false)
	// if p.isNextIngore(ingore_next_type) {
	// 	t.Error(" need ingore_next_type ")
	// }
	// if !p.isNextIngore(ingore_next_lencap) {
	// 	t.Error(" do need ingore_next_lencap ")
	// }
	// if p.isIngoreNone() {
	// 	t.Error(" no  isIngoreNone()")
	// }

	// p.setNextIngoreState(ingore_next_lencap, false)
	// if p.isNextIngore(ingore_next_type) {
	// 	t.Error(" need ingore_next_type ")
	// }
	// if p.isNextIngore(ingore_next_lencap) {
	// 	t.Error(" do need ingore_next_lencap ")
	// }
	// if !p.isIngoreNone() {
	// 	t.Error(" no  isIngoreNone()")
	// }

	// p.setNextIngoreState(ingore_next_all, false)
	// if p.isNextIngore(ingore_next_type) {
	// 	t.Error(" need ingore_next_type ")
	// }
	// if p.isNextIngore(ingore_next_lencap) {
	// 	t.Error(" do need ingore_next_lencap ")
	// }
	// if !p.isIngoreNone() {
	// 	t.Error(" no  isIngoreNone()")
	// }
	// p.setNextIngoreState(ingore_next_all, true)
	// p.setNextIngoreState(ingore_next_all, false)
	// if p.isNextIngore(ingore_next_type) {
	// 	t.Error(" need ingore_next_type ")
	// }
	// if p.isNextIngore(ingore_next_lencap) {
	// 	t.Error(" do need ingore_next_lencap ")
	// }
	// if !p.isIngoreNone() {
	// 	t.Error(" no  isIngoreNone()")
	// }

}

// debugingFunc is used to identify which public function of the debuging package or
// DebugConfig a test applies to.
type debugingFunc int

const (
	fCSFdump debugingFunc = iota
	fCSFprint
	fCSFprintf
	fCSFprintln
	fCSPrint
	fCSPrintln
	fCSSdump
	fCSSprint
	fCSSprintf
	fCSSprintln
	fCSErrorf
	fCSNewFormatter
	fErrorf
	fFprint
	fFprintln
	fPrint
	fPrintln
	fSdump
	fSprint
	fSprintf
	fSprintln
)

// Map of debugingFunc values to names for pretty printing.
var debugingFuncStrings = map[debugingFunc]string{
	fCSFdump:        "DebugConfig.Fdump",
	fCSFprint:       "DebugConfig.Fprint",
	fCSFprintf:      "DebugConfig.Fprintf",
	fCSFprintln:     "DebugConfig.Fprintln",
	fCSSdump:        "DebugConfig.Sdump",
	fCSPrint:        "DebugConfig.Print",
	fCSPrintln:      "DebugConfig.Println",
	fCSSprint:       "DebugConfig.Sprint",
	fCSSprintf:      "DebugConfig.Sprintf",
	fCSSprintln:     "DebugConfig.Sprintln",
	fCSErrorf:       "DebugConfig.Errorf",
	fCSNewFormatter: "DebugConfig.NewFormatter",
	fErrorf:         "FmtErrorf",
	fFprint:         "FmtFprint",
	fFprintln:       "FmtFprintln",
	fPrint:          "FmtPrint",
	fPrintln:        "FmtPrintln",
	fSdump:          "FmtSdump",
	fSprint:         "FmtSprint",
	fSprintf:        "FmtSprintf",
	fSprintln:       "FmtSprintln",
}

func (f debugingFunc) String() string {
	if s, ok := debugingFuncStrings[f]; ok {
		return s
	}
	return fmt.Sprintf("Unknown debugingFunc (%d)", int(f))
}

// debugingTest is used to describe a test to be performed against the public
// functions of the debuging package or DebugConfig.
type debugingTest struct {
	cs     *debuging.DebugConfig
	f      debugingFunc
	format string
	in     interface{}
	want   string
}

// debugingTests houses the tests to be performed against the public functions of
// the debuging package and DebugConfig.
//
// These tests are only intended to ensure the public functions are exercised
// and are intentionally not exhaustive of types.  The exhaustive type
// tests are handled in the dump and format tests.
var debugingTests []debugingTest

// redirStdout is a helper function to return the standard output from f as a
// byte slice.
func redirStdout(f func()) ([]byte, error) {
	tempFile, err := ioutil.TempFile("", "ss-test")
	if err != nil {
		return nil, err
	}
	fileName := tempFile.Name()
	defer os.Remove(fileName) // Ignore error

	origStdout := os.Stdout
	os.Stdout = tempFile
	f()
	os.Stdout = origStdout
	tempFile.Close()

	return ioutil.ReadFile(fileName)
}

type stringer string

func (s stringer) String() string {
	return string(s)
}

type pstringer string

func (p *pstringer) String() string {
	return string(*p)
}

type RefRef0 struct {
	ref1 *RefRef1
}
type RefRef1 struct {
	ref0 *RefRef0
}

type IndirCir1 struct {
	ref2 *IndirCir2
}
type IndirCir2 struct {
	ref3 *IndirCir3
}
type IndirCir3 struct {
	ref1 *IndirCir1
}

type customError int

func (e customError) Error() string {
	return fmt.Sprintf("error: %d", int(e))
}
func initSpewTests() {
	// Config states with various settings.
	scsDefault := debuging.DefautConfig()
	scsNoMethods := &debuging.DebugConfig{DisableMethods: true}
	scsNoPmethods := &debuging.DebugConfig{DisablePointerMethods: true}
	scsMaxDepth := &debuging.DebugConfig{MaxDepth: 1}
	scsContinue := &debuging.DebugConfig{ContinueOnMethod: true}

	// Variables for tests on types which implement Stringer interface with and
	// without a pointer receiver.
	ts := stringer("test")
	tps := pstringer("test")

	// depthTester is used to test max depth handling for structs, array, slices
	// and maps.
	type depthTester struct {
		ic    IndirCir1
		arr   [1]string
		slice []string
		m     map[string]int
	}
	dt := depthTester{IndirCir1{nil}, [1]string{"arr"}, []string{"slice"}, map[string]int{"one": 1}}

	// Variable for tests on types which implement error interface.
	te := customError(10)

	debugingTests = []debugingTest{
		{scsDefault, fCSFdump, "", int8(127), "(int8)127"},
		{scsDefault, fCSFprint, "", int16(32767), "32767"},
		{scsDefault, fCSFprintf, "%v", int32(2147483647), "2147483647"},
		{scsDefault, fCSFprintln, "", int(2147483647), "2147483647\n"},
		{scsDefault, fCSPrint, "", int64(9223372036854775807), "9223372036854775807"},
		{scsDefault, fCSPrintln, "", uint8(255), "255\n"},
		{scsDefault, fCSSdump, "", uint8(64), "(uint8)64"},
		{scsDefault, fCSSprint, "", complex(1, 2), "(1+2i)"},
		{scsDefault, fCSSprintf, "%v", complex(float32(3), 4), "(3+4i)"},
		{scsDefault, fCSSprintln, "", complex(float64(5), 6), "(5+6i)\n"},
		{scsDefault, fCSErrorf, "%#v", uint16(65535), "(uint16)65535"},
		{scsDefault, fCSNewFormatter, "%v", uint32(4294967295), "4294967295"},
		{scsDefault, fErrorf, "%v", uint64(18446744073709551615), "18446744073709551615"},
		{scsDefault, fFprint, "", float32(3.14), "3.14"},
		{scsDefault, fFprintln, "", float64(6.28), "6.28\n"},
		{scsDefault, fPrint, "", true, "true"},
		{scsDefault, fPrintln, "", false, "false\n"},
		{scsDefault, fSdump, "", complex(-10, -20), "(complex128)(-10-20i)"},
		{scsDefault, fSprint, "", complex(-1, -2), "(-1-2i)"},
		{scsDefault, fSprintf, "%v", complex(float32(-3), -4), "(-3-4i)"},
		{scsDefault, fSprintln, "", complex(float64(-5), -6), "(-5-6i)\n"},
		{scsNoMethods, fCSFprint, "", ts, "\"test\""},
		{scsNoMethods, fCSFprint, "", &ts, "<*>\"test\""},
		{scsNoMethods, fCSFprint, "", tps, "\"test\""},
		{scsNoMethods, fCSFprint, "", &tps, "<*>\"test\""},
		{scsNoPmethods, fCSFprint, "", ts, "test"},
		{scsNoPmethods, fCSFprint, "", &ts, "<*>test"},
		{scsNoPmethods, fCSFprint, "", tps, "\"test\""},
		{scsNoPmethods, fCSFprint, "", &tps, "<*>test"},
		{scsMaxDepth, fCSFprint, "", dt,
			"{\n" +
				" {......},\n" +
				" ([1]string)(len=1 cap=1) [......],\n" +
				" ([]string)(len=1 cap=1) [......],\n" +
				" map[......],\n" +
				"}"},
		{scsMaxDepth, fCSFdump, "", dt, "(debuging.depthTester){\n" +
			" ic:(debuging.IndirCir1){......},\n" +
			" arr:([1]string)(len=1 cap=1) [......],\n" +
			" slice:([]string)(len=1 cap=1) [......],\n" +
			" m:(map[string]int)(len=1) {......},\n}"},
		{scsContinue, fCSFprint, "", ts, "(test) \"test\""},
		{scsContinue, fCSFdump, "", ts,
			"(debuging.stringer)(len=4) (test) \"test\""},
		{scsContinue, fCSFprint, "", te, "(error: 10) 10"},
		{scsContinue, fCSFdump, "", te, "(debuging.customError)(error: 10) 10"},
	}
}

// TestSpew executes all of the tests described by debugingTests.
func TestSpew(t *testing.T) {
	initSpewTests()

	t.Logf("Running %d tests", len(debugingTests))
	for i, test := range debugingTests {
		buf := new(bytes.Buffer)
		switch test.f {
		case fCSFdump:
			debuging.DumpWithConfigOfWriter(test.cs, buf, test.in)

		case fCSFprint:
			test.cs.Fprint(buf, test.in)

		case fCSFprintf:
			test.cs.Fprintf(buf, test.format, test.in)

		case fCSFprintln:
			test.cs.Fprintln(buf, test.in)

		case fCSPrint:
			b, err := redirStdout(func() { test.cs.Print(test.in) })
			if err != nil {
				t.Errorf("%v #%d %v", test.f, i, err)
				continue
			}
			buf.Write(b)

		case fCSPrintln:
			b, err := redirStdout(func() { test.cs.Println(test.in) })
			if err != nil {
				t.Errorf("%v #%d %v", test.f, i, err)
				continue
			}
			buf.Write(b)

		case fCSSdump:
			str := debuging.DumpValue(test.in)
			buf.WriteString(str)

		case fCSSprint:
			str := test.cs.Sprint(test.in)
			buf.WriteString(str)

		case fCSSprintf:
			str := test.cs.Sprintf(test.format, test.in)
			buf.WriteString(str)

		case fCSSprintln:
			str := test.cs.Sprintln(test.in)
			buf.WriteString(str)

		case fCSErrorf:
			err := test.cs.Errorf(test.format, test.in)
			buf.WriteString(err.Error())

		case fCSNewFormatter:
			fmt.Fprintf(buf, test.format, debuging.NewFmtPrintWriter(test.cs, test.in))

		case fErrorf:
			err := debuging.FmtErrorf(test.format, test.in)
			buf.WriteString(err.Error())

		case fFprint:
			debuging.FmtFprint(buf, test.in)

		case fFprintln:
			debuging.FmtFprintln(buf, test.in)

		case fPrint:
			b, err := redirStdout(func() { debuging.FmtPrint(test.in) })
			if err != nil {
				t.Errorf("%v #%d %v", test.f, i, err)
				continue
			}
			buf.Write(b)

		case fPrintln:
			b, err := redirStdout(func() { debuging.FmtPrintln(test.in) })
			if err != nil {
				t.Errorf("%v #%d %v", test.f, i, err)
				continue
			}
			buf.Write(b)

		case fSdump:
			str := debuging.DumpValue(test.in)
			buf.WriteString(str)

		case fSprint:
			str := debuging.FmtSprint(test.in)
			buf.WriteString(str)

		case fSprintf:
			str := debuging.FmtSprintf(test.format, test.in)
			buf.WriteString(str)

		case fSprintln:
			str := debuging.FmtSprintln(test.in)
			buf.WriteString(str)

		default:
			t.Errorf("%v #%d unrecognized function", test.f, i)
			continue
		}
		s := buf.String()
		if test.want != s {
			t.Errorf("DebugConfig #%d %v\n got:[%s] \n want:[%s]", i, test.f, s, test.want)
			continue
		}
	}
}
