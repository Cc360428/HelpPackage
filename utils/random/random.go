package random

import (
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func Float64() float64 {
	return rand.Float64()
}
func Float32() float32 {
	return rand.Float32()
}

// [0,n)
func RandInt(n int) int {
	if n == 0 {
		return 0
	}
	return rand.Intn(n)
}

// [0,n)
func RandInt32(n int32) int32 {
	if n == 0 {
		return 0
	}
	return rand.Int31n(n)
}

// [0,n)
func RandInt64(n int64) int64 {
	if n == 0 {
		return 0
	}
	return rand.Int63n(n)
}

// [min,max)
func RandSection(min, max int64) int64 {
	if min > max {
		max = min
	}
	if max-min == 0 {
		return max
	}
	return min + RandInt64(max-min)
}

// [min,max]
func RandSectionContain(min, max int64) int64 {
	return RandSection(min, max+1)
}

// [min,max)
func RandSectionInt(min, max int) int {
	if min > max {
		max = min
	}
	if max-min == 0 {
		return max
	}
	return min + RandInt(max-min)
}

// [min,max]
func RandSectionContainInt(min, max int) int {
	return RandSectionInt(min, max+1)
}
