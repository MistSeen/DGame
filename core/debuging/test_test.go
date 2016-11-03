package debuging

import "testing"

func TestFyindSigpanic(t *testing.T) {
	t.Parallel()
	sp := findSigpanic()
	if got, want := sp.Name(), "runtime.sigpanic"; got != want {
		t.Errorf("got: %v \nwant: %v", got, want)
	}
}

type MyData struct {
	IntField   int
	FloatField float64
	StrField   string
	MapField   map[int]string
	SliceField []int
	PointField *MyData
}

func Test_Dump(t *testing.T) {
	data := &MyData{
		1234,
		77.88,
		"xyz",
		map[int]string{
			1: "abc",
			2: "def",
			3: "ghi",
		},
		[]int{
			3,
			7,
			11,
			13,
			17,
		},
		nil,
	}
	data.PointField = data

	Trace(data)
}
