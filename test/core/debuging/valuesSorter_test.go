package debuging

import (
	"fmt"
	"reflect"
	. "server/core/debuging"
	"testing"
)

type embed struct {
	a string
}

type sortTestCase struct {
	input    []reflect.Value
	expected []reflect.Value
}
type sortableStruct struct {
	x int
}
type unsortableStruct struct {
	y int
}

func (ss sortableStruct) String() string {
	return fmt.Sprintf("ss.x = %d", ss.x)
}

func TestSortValue(t *testing.T) {
	v := reflect.ValueOf

	a := v("a")
	b := v("b")
	c := v("c")

	embedA := v(embed{"a"})
	embedB := v(embed{"b"})
	embedC := v(embed{"c"})

	tests := []sortTestCase{
		//No vlaue
		{
			[]reflect.Value{},
			[]reflect.Value{},
		},
		{
			[]reflect.Value{v(false), v(true), v(false)},
			[]reflect.Value{v(false), v(false), v(true)},
		},
		//ints
		{
			[]reflect.Value{v(2), v(1), v(3)},
			[]reflect.Value{v(1), v(2), v(3)},
		},
		//uInts
		{
			[]reflect.Value{v(uint8(22)), v(uint8(21)), v(uint8(23))},
			[]reflect.Value{v(uint8(21)), v(uint8(22)), v(uint8(23))},
		},
		//strings
		{
			[]reflect.Value{b, a, c},
			[]reflect.Value{a, b, c},
		},
		//float
		{
			[]reflect.Value{v(2.0), v(1.0), v(3.0)},
			[]reflect.Value{v(1.0), v(2.0), v(3.0)},
		},
		//array
		{
			[]reflect.Value{v([3]int{3, 2, 1}), v([3]int{1, 3, 2}), v([3]int{1, 2, 3})},
			[]reflect.Value{v([3]int{1, 2, 3}), v([3]int{1, 3, 2}), v([3]int{3, 2, 1})},
		},
		//sortableStruct
		{
			[]reflect.Value{v(sortableStruct{2}), v(sortableStruct{1}), v(sortableStruct{3})},
			[]reflect.Value{v(sortableStruct{2}), v(sortableStruct{1}), v(sortableStruct{3})},
		},
		//unSortableStruct
		{
			[]reflect.Value{v(unsortableStruct{2}), v(unsortableStruct{1}), v(unsortableStruct{3})},
			[]reflect.Value{v(unsortableStruct{2}), v(unsortableStruct{1}), v(unsortableStruct{3})},
		},
		//invalid
		{
			[]reflect.Value{embedA, embedB, embedC},
			[]reflect.Value{embedA, embedB, embedC},
		},
	}
	config := DebugConfig{DisableMethods: true, SpewKeys: false}
	sortValueTesting(tests, &config, t)
}

func TestArraySortValue(t *testing.T) {
	v := reflect.ValueOf

	tests := []sortTestCase{
		//array
		{
			[]reflect.Value{v([3]int{3, 2, 1}), v([3]int{1, 3, 2}), v([3]int{1, 2, 3})},
			[]reflect.Value{v([3]int{1, 2, 3}), v([3]int{1, 3, 2}), v([3]int{3, 2, 1})},
		},
	}
	config := DebugConfig{DisableMethods: true, SpewKeys: false}
	sortValueTesting(tests, &config, t)
}

func TestSortValueWithMethods(t *testing.T) {
	v := reflect.ValueOf

	a := v("a")
	b := v("b")
	c := v("c")

	embedA := v(embed{"a"})
	embedB := v(embed{"b"})
	embedC := v(embed{"c"})

	tests := []sortTestCase{
		//No vlaue
		{
			[]reflect.Value{},
			[]reflect.Value{},
		},
		{
			[]reflect.Value{v(false), v(true), v(false)},
			[]reflect.Value{v(false), v(false), v(true)},
		},
		//ints
		{
			[]reflect.Value{v(2), v(1), v(3)},
			[]reflect.Value{v(1), v(2), v(3)},
		},
		//uInts
		{
			[]reflect.Value{v(uint8(22)), v(uint8(21)), v(uint8(23))},
			[]reflect.Value{v(uint8(21)), v(uint8(22)), v(uint8(23))},
		},
		//strings
		{
			[]reflect.Value{b, a, c},
			[]reflect.Value{a, b, c},
		},
		//float
		{
			[]reflect.Value{v(2.0), v(1.0), v(3.0)},
			[]reflect.Value{v(1.0), v(2.0), v(3.0)},
		},
		//array
		{
			[]reflect.Value{v([3]int{3, 2, 1}), v([3]int{1, 3, 2}), v([3]int{1, 2, 3})},
			[]reflect.Value{v([3]int{1, 2, 3}), v([3]int{1, 3, 2}), v([3]int{3, 2, 1})},
		},
		//sortableStruct
		{
			[]reflect.Value{v(sortableStruct{2}), v(sortableStruct{1}), v(sortableStruct{3})},
			[]reflect.Value{v(sortableStruct{1}), v(sortableStruct{2}), v(sortableStruct{3})},
		},
		//unSortableStruct
		{
			[]reflect.Value{v(unsortableStruct{2}), v(unsortableStruct{1}), v(unsortableStruct{3})},
			[]reflect.Value{v(unsortableStruct{1}), v(unsortableStruct{2}), v(unsortableStruct{3})},
		},
		//invalid
		{
			[]reflect.Value{embedA, embedB, embedC},
			[]reflect.Value{embedA, embedB, embedC},
		},
	}
	config := DebugConfig{DisableMethods: true, SpewKeys: true}
	sortValueTesting(tests, &config, t)
}

func TestSortValueWithSpew(t *testing.T) {
	v := reflect.ValueOf

	a := v("a")
	b := v("b")
	c := v("c")

	embedA := v(embed{"a"})
	embedB := v(embed{"b"})
	embedC := v(embed{"c"})

	tests := []sortTestCase{
		//No vlaue
		{
			[]reflect.Value{},
			[]reflect.Value{},
		},
		{
			[]reflect.Value{v(false), v(true), v(false)},
			[]reflect.Value{v(false), v(false), v(true)},
		},
		//ints
		{
			[]reflect.Value{v(2), v(1), v(3)},
			[]reflect.Value{v(1), v(2), v(3)},
		},
		//uInts
		{
			[]reflect.Value{v(uint8(22)), v(uint8(21)), v(uint8(23))},
			[]reflect.Value{v(uint8(21)), v(uint8(22)), v(uint8(23))},
		},
		//strings
		{
			[]reflect.Value{b, a, c},
			[]reflect.Value{a, b, c},
		},
		//float
		{
			[]reflect.Value{v(2.0), v(1.0), v(3.0)},
			[]reflect.Value{v(1.0), v(2.0), v(3.0)},
		},
		//array
		{
			[]reflect.Value{v([3]int{3, 2, 1}), v([3]int{1, 3, 2}), v([3]int{1, 2, 3})},
			[]reflect.Value{v([3]int{1, 2, 3}), v([3]int{1, 3, 2}), v([3]int{3, 2, 1})},
		},
		//sortableStruct
		{
			[]reflect.Value{v(sortableStruct{2}), v(sortableStruct{1}), v(sortableStruct{3})},
			[]reflect.Value{v(sortableStruct{1}), v(sortableStruct{2}), v(sortableStruct{3})},
		},
		//unSortableStruct
		{
			[]reflect.Value{v(unsortableStruct{2}), v(unsortableStruct{1}), v(unsortableStruct{3})},
			[]reflect.Value{v(unsortableStruct{2}), v(unsortableStruct{1}), v(unsortableStruct{3})},
		},
		//invalid
		{
			[]reflect.Value{embedA, embedB, embedC},
			[]reflect.Value{embedA, embedB, embedC},
		},
	}
	config := DebugConfig{DisableMethods: false, SpewKeys: false}
	sortValueTesting(tests, &config, t)
}

func sortValueTesting(tests []sortTestCase, config *DebugConfig, t *testing.T) {
	getInterface := func(values []reflect.Value) []interface{} {
		ifaces := []interface{}{}
		for _, v := range values {
			ifaces = append(ifaces, v.Interface())
		}
		return ifaces
	}
	for _, i := range tests {
		SortValue(i.input, config)
		// reflect.DeepEqual cannot really make sense of reflect.Value,
		// probably because of all the pointer tricks. For instance,
		// v(2.0) != v(2.0) on a 32-bits system. Turn them into interface{}
		// instead.
		input := getInterface(i.input)
		expected := getInterface(i.expected)

		if !reflect.DeepEqual(input, expected) {
			t.Errorf("SortValue mismatch:\n %v != %v", input, expected)
		}
	}
}
