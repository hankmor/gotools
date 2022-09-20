package random

import (
	"math/rand"
	"time"
)

func Int(max int) int {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return r.Intn(max)
}

// Between 生成 min(包含) 到 max(不包括) 之间的随机数，min 必须 小于等于 max，否则抛出异常
func Between(min, max int) int {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	switch {
	case min == max:
		return min
	case min > max:
		panic("min must be less than or equal to max")
	}
	return min + r.Intn(max-min)
}
