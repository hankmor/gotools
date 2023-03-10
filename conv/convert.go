package conv

import (
	"fmt"
	"github.com/hankmor/gotools/errs"
	"math/big"
	"strconv"
	"strings"
)

// StrToInt covert string to int
func StrToInt(str string) int {
	n, err := strconv.Atoi(str)
	errs.Throw(err)
	return n
}

// StrToInt64 covert string to int64
func StrToInt64(str string) int64 {
	n, err := strconv.ParseInt(str, 10, 64)
	errs.Throw(err)
	return n
}

// IntToStr covert int to string
func IntToStr(src int) string {
	return strconv.Itoa(src)
}

// Int64ToStr covert int64 to string
func Int64ToStr(src int64) string {
	return strconv.FormatInt(src, 10)
}

// JoinBigInt join a slice of big.Int to a string with ","
func JoinBigInt(ints []*big.Int) string {
	var temp = make([]string, len(ints))
	for k, v := range ints {
		temp[k] = fmt.Sprintf("%d", v.Int64())
	}
	var result = strings.Join(temp, ",")
	return result
}

// StrToFloat64 convert string to float64
func StrToFloat64(amount string) float64 {
	float, err := strconv.ParseFloat(amount, 64)
	errs.Throw(err)
	return float
}

// Int64ToHex convert int64 to hex string
func Int64ToHex(src int64) string {
	return strconv.FormatInt(src, 16)
}

// HexToInt64 convert hex string to int64
func HexToInt64(src string) uint64 {
	id, err := strconv.ParseUint(src, 16, 64)
	errs.Throw(err)
	return id
}

// SplitStrToInt split a string to int slice with given separator.
func SplitStrToInt(s string, sep string) []int64 {
	ss := strings.Split(s, sep)
	var is []int64
	for _, i := range ss {
		it, err := strconv.ParseInt(i, 10, 64)
		errs.Throw(err)
		is = append(is, it)
	}
	return is
}

// IntsToStr convert a int slice to a string slice.
func IntsToStr(is []int64) []string {
	if is == nil || len(is) == 0 {
		return nil
	}
	var ss = make([]string, len(is))
	for _, i := range is {
		ss = append(ss, strconv.FormatInt(i, 10))
	}
	return ss
}

// StrsToInt convert a str slice to a int slice.
func StrsToInt(ss []string) []int64 {
	if ss == nil || len(ss) == 0 {
		return nil
	}
	var is = make([]int64, len(ss))
	for _, s := range ss {
		i, err := strconv.ParseInt(s, 10, 64)
		if err != nil {
			return nil
		}
		is = append(is, i)
	}
	return is
}
