package slice

import (
	"github.com/huzhouv/gotools/str"
	"reflect"
	"sort"
	"strings"
)

// S is a generic slice type.
type S[T comparable] []T

func New[T comparable]() S[T] {
	return make(S[T], 0)
}

func NewSize[T comparable](size int) S[T] {
	return make(S[T], size)
}

func Wrap[T comparable](s []T) S[T] {
	if reflect.TypeOf(s).Kind() != reflect.Slice {
		panic("require slice")
	}
	return s
}

// Retain is a method that retain the elements that matched the condition the function param defined,
// and not matched elements will be removed.
// This method will return a new slice and the original slice will not be changed
func (s S[T]) Retain(cond func(a T) bool) []T {
	var ret []T
	for _, a := range s {
		if cond(a) { // 符合条件
			ret = append(ret, a)
		}
	}
	return ret
}

// Filter is a method that opposite to the Retain method, it will filter the elements that not matched the condition
// the function param defined, and not matched elements will be retained.
// This method will return a new slice and the original slice will not be changed
func (s S[T]) Filter(cond func(a T) bool) []T {
	var ret []T
	for _, a := range s {
		if !cond(a) { // 不符合条件
			ret = append(ret, a)
		}
	}
	return ret
}

// Join is a method that join all the elements by the string the sep param defined.
func (s S[T]) Join(sep string) string {
	var ret []string
	for _, a := range s {
		ret = append(ret, str.String(a))
	}
	return strings.Join(ret, sep)
}

// Union returns all elements of the two slices, i.e. the result is and union set.
func (s S[T]) Union(dest []T) []T {
	var ret []T = s
	var d = Wrap(dest)
	for _, a := range s {
		d = d.Delete(a)
	}
	ret = append(ret, d...)
	return ret
}

// Intersect returns those elements that both the source and the dest slice have. i.e. the result is and intersection set.
func (s S[T]) Intersect(dest []T) []T {
	var ret []T
	for _, v := range s {
		find := false
		for _, el := range dest {
			if v == el {
				find = true
				break
			}
		}
		if find {
			ret = append(ret, v)
		}
	}
	return ret
}

// Diff returns the different elements between source and dest slice.
func (s S[T]) Diff(dest []T) []T {
	var ret = s.Union(dest)
	var it = s.Intersect(dest)
	return Wrap(ret).Remove(it)
}

// Remove method will remove elements which the dest slice have from the source slice.
func (s S[T]) Remove(dest []T) []T {
	return s.Delete(dest...)
}

// Delete will delete the give elements.
func (s S[T]) Delete(elem ...T) []T {
	var ret []T
	for _, v := range s {
		find := false
		for _, el := range elem {
			if v == el {
				find = true
				break
			}
		}
		if !find {
			ret = append(ret, v)
		}
	}
	return ret
}

// sortable slice

// SortableSlice is a struct to define a sortable slice, it implements sort.Interface.
type SortableSlice[T comparable] struct {
	slice S[T]
	less  func(x, y T) bool // 比较的方法，参数为T的指针，直接更改原始slice的顺序
}

func (s SortableSlice[T]) Len() int {
	return len(s.slice)
}

func (s SortableSlice[T]) Less(i, j int) bool {
	return s.less(s.slice[i], s.slice[j])
}

func (s SortableSlice[T]) Swap(i, j int) {
	s.slice[i], s.slice[j] = s.slice[j], s.slice[i]
}

// Sort is a method to sort the original slice by the given argument less, which is a method to define how to compare
// the elements in it, this less method receive two T type so Sort method will NOT change the original slice.
// Sort will return a SortableSlice instance, so you could invoke its other api such as SortableSlice.Reverse etc.
func (s S[T]) Sort(less func(x, y T) bool) SortableSlice[T] {
	v := &SortableSlice[T]{s, less}
	sort.Sort(v)
	return *v
}

// Reverse is a method to reverse the sequence that the Slice.Sort method sorted, so you should has invoked the S.Sort
// first before to call Reverse.
func (s SortableSlice[T]) Reverse() S[T] {
	sort.Sort(sort.Reverse(s))
	return s.slice
}
