package debuging

import (
	"bytes"
	"reflect"
	"sort"
)

type valuesSorter struct {
	values []reflect.Value
	keys   []string
	config *DebugConfig
}

/*
Len() int :Len is the number of elements in the collection.
Less(i, j int) bool:
	  Less reports whether the element with
	  index i should sort before the element with index j.
 Swap:swaps the elements with indexes i and j.
	  Swap(i, j int)
*/
func (sorter *valuesSorter) Len() int {
	return len(sorter.values)
}
func (sorter *valuesSorter) Less(i, j int) bool {
	if sorter.keys != nil {
		return sorter.keys[i] < sorter.keys[j]
	}
	return sorter.valueLess(sorter.values[i], sorter.values[j])
}
func (sorter *valuesSorter) Swap(i, j int) {
	if sorter.keys != nil {
		sorter.keys[i], sorter.keys[j] = sorter.keys[j], sorter.keys[i]
	}
	sorter.values[i], sorter.values[j] = sorter.values[j], sorter.values[i]
}

//=======================================================================//

//SortValue is SortValue
func SortValue(values []reflect.Value, config *DebugConfig) {
	if len(values) == 0 {
		return
	}
	sort.Sort(valueSortCall(values, config))
}

//=======================================================================//

func valueSortCall(values []reflect.Value, config *DebugConfig) sort.Interface {

	sorter := &valuesSorter{values: values, config: config}
	if sorter.isNeedComplexKey() {
		sorter.complexSortKey()
	}
	return sorter
}

func (sorter *valuesSorter) isNeedComplexKey() bool {
	kind := sorter.values[0].Kind()
	switch kind {
	case reflect.Bool:
		return false
	case reflect.Int, reflect.Uint,
		reflect.Int8, reflect.Uint8,
		reflect.Int16, reflect.Uint16,
		reflect.Int32, reflect.Uint32,
		reflect.Int64, reflect.Uint64:
		return false
	case reflect.Float32, reflect.Float64:
		return false
	case reflect.String, reflect.Array, reflect.Uintptr:
		return false
	}
	return true
}
func (sorter *valuesSorter) complexSortKey() {
	if !sorter.config.DisableMethods {
		sorter.keys = make([]string, len(sorter.values))
		for i := range sorter.values {
			buf := bytes.Buffer{}
			if !handleMethods(sorter.config, &buf, sorter.values[i]) {
				sorter.keys = nil
				break
			}
			sorter.keys[i] = buf.String()
		}
	}
	if sorter.keys == nil && sorter.config.SpewKeys {
		sorter.keys = make([]string, len(sorter.values))
		for index := range sorter.values {
			sorter.keys[index] = sorter.config.Sprintf("%#v", sorter.values[index].Interface())
		}
	}
}

func (sorter *valuesSorter) valueLess(a reflect.Value, b reflect.Value) bool {
	switch a.Kind() {
	case reflect.Bool:
		return !a.Bool() && b.Bool()
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return a.Int() < b.Int()
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return a.Uint() < b.Uint()
	case reflect.Float32, reflect.Float64:
		return a.Float() < b.Float()
	case reflect.String:
		return a.String() < b.String()
	case reflect.Array:
		len := a.Len()
		for i := 0; i < len; i++ {
			av := a.Index(i)
			bv := b.Index(i)
			if av.Interface() == bv.Interface() {
				continue
			}
			return sorter.valueLess(av, bv)
		}
	}
	return a.String() < b.String()
}
