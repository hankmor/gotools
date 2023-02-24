package slice_test

import (
	"fmt"
	"github.com/huzhouv/gotools/collection/slice"
	"github.com/huzhouv/gotools/testool"
	"reflect"
	"testing"
)

var testSlice = []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

func TestNew(t *testing.T) {
	s := slice.New[int]()
	fmt.Println(len(s), cap(s))
	s = append(s, 1)
	s = append(s, 2)
	s = append(s, 3)
	s = append(s, 4)
	// s = append(s, 5)
	fmt.Println(len(s), cap(s))

	s = slice.NewSize[int](5)
	fmt.Println(len(s), cap(s))
	s = append(s, 1)
	fmt.Println(len(s), cap(s))
}

func TestRetain(t *testing.T) {
	logger := testool.Wrap(t)

	ret := slice.Wrap(testSlice).Retain(func(a int) bool {
		return a > 5
	})
	fmt.Println(testSlice)
	expect := []int{6, 7, 8, 9}
	logger.Require(reflect.DeepEqual(expect, ret), "expect %#v, actual %#v", expect, ret)

	// ret1 := slice.Wrap([]any{"a", 1, 3.14}).Retain(func(a any) bool {
	// 	switch a.(type) {
	// 	case int:
	// 		return a.(int) > 1
	// 	default:
	// 		return true
	// 	}
	// })
	// expect1 := []any{"a", 3.14}
	// logger.Require(reflect.DeepEqual(expect1, ret1), "expect %v, actual %v", expect1, ret1)
}

func TestJoin(t *testing.T) {
	logger := testool.Wrap(t)

	s := slice.Wrap(testSlice).Join(",")
	logger.Require(s == "1,2,3,4,5,6,7,8,9", "%v join result should be %s", testSlice, s)

	join := slice.Wrap([]string{"a", "b"}).Join(".")
	logger.Require(join == "a.b", "a join b should be %s", join)

	// join = slice.Wrap([]TestData{"a", 1, "3.14"}).Join(",")
	// logger.Require(join == "a,1,3.14", "%s join %d join %.2f should be %s", "a", 1, 3.14, join)
}

func TestUnion(t *testing.T) {
	var before = testSlice

	logger := testool.Wrap(t)
	sl := []int{1, 2, 3, 4, 5, 6, 10, 11}
	want := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}
	ret := slice.Wrap(testSlice).Union(sl)
	logger.Require(reflect.DeepEqual(testSlice, before), "raw slice should not be changed")
	logger.Require(reflect.DeepEqual(ret, want), "result is correct")
}

func TestIntersect(t *testing.T) {
	var before = testSlice

	logger := testool.Wrap(t)
	sl := []int{1, 2, 3, 10, 11}
	want := []int{1, 2, 3}
	ret := slice.Wrap(testSlice).Intersect(sl)
	logger.Require(reflect.DeepEqual(testSlice, before), "raw slice should not be changed")
	logger.Require(reflect.DeepEqual(ret, want), "result is correct")
}

func TestRemove(t *testing.T) {
	var before = testSlice

	logger := testool.Wrap(t)
	sl := []int{1, 2, 3, 10, 11}
	want := []int{4, 5, 6, 7, 8, 9}
	ret := slice.Wrap(testSlice).Remove(sl)
	logger.Require(reflect.DeepEqual(testSlice, before), "raw slice should not be changed")
	logger.Require(reflect.DeepEqual(ret, want), "result is correct")
}

func TestDiff(t *testing.T) {
	var before = testSlice

	logger := testool.Wrap(t)
	sl := []int{1, 2, 3, 10, 11}
	want := []int{4, 5, 6, 7, 8, 9, 10, 11}
	ret := slice.Wrap(testSlice).Diff(sl)
	logger.Require(reflect.DeepEqual(testSlice, before), "raw slice should not be changed")
	logger.Require(reflect.DeepEqual(ret, want), "result is correct")
}

func TestDelete(t *testing.T) {
	raw := testSlice

	logger := testool.Wrap(t)
	want := []int{4, 5, 6, 7, 8, 9}
	var ret = slice.Wrap(testSlice).Delete(1, 2, 3)
	logger.Require(reflect.DeepEqual(testSlice, raw), "raw slice should not be changed")
	logger.Require(reflect.DeepEqual(ret, want), "result correct")

	want = []int{2, 4, 6, 8}
	ret = slice.Wrap(testSlice).Delete(1, 3, 5, 7, 9)
	logger.Require(reflect.DeepEqual(testSlice, raw), "raw slice should not be changed")
	logger.Require(reflect.DeepEqual(ret, want), "result correct")
}

func TestSort(t *testing.T) {
	var ss = []int{5, 6, 7, 8, 9, 1, 2, 3, 4}
	s := slice.Wrap(ss)
	sr := s.Sort(func(i, j int) bool {
		if i < j {
			return true
		}
		return false
	})
	fmt.Printf("%v\n", ss)
	sr.Reverse()
	fmt.Printf("%v\n", ss)
}

type user struct {
	name  string
	age   int
	score float32
}

func TestSortObj(t *testing.T) {
	users := []user{
		{name: "huzhou", age: 18, score: 99.5},
		{name: "huzhou", age: 16, score: 100},
		{name: "zhangsan", age: 17, score: 99.5},
		{name: "abbc", age: 17, score: 99.5},
	}

	s := slice.Wrap(users)
	sr := s.Sort(func(a, b user) bool {
		if a.score != b.score {
			return a.score > b.score
		}
		if a.age != b.age {
			return a.age < b.age
		}
		if a.name != b.name {
			return a.name < b.name
		}
		return false
	})
	fmt.Printf("%v\n", s)
	sr.Reverse()
	fmt.Printf("%v\n", s)
}
