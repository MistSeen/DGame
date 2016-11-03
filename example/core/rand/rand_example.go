package rand

import "fmt"

func ExampleRandGroup() {
	i := RandGroup(0, 0, 50, 50)
	switch i {
	case 2, 3:
		fmt.Println("ok")
	}

	// Output:
	// ok
}

func ExampleRandInterval() {
	v := RandInterval(-1, 1)
	switch v {
	case -1, 0, 1:
		fmt.Println("ok")
	}

	// Output:
	// ok
}

func ExampleRandIntervalN() {
	r := RandIntervalN(-1, 0, 2)
	if r[0] == -1 && r[1] == 0 ||
		r[0] == 0 && r[1] == -1 {
		fmt.Println("ok")
	}

	// Output:
	// ok
}
