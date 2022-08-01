package randc

import (
	"math"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

///////////////////////////////////////////// int

// RandInt [0,n)
func RandInt(n int) int {
	if n == 0 {
		return 0
	}
	return rand.Intn(n)
}

// RandSectionInt [min,max)
func RandSectionInt(min, max int) int {
	if min > max {
		max = min
	}
	if max-min == 0 {
		return max
	}
	return min + RandInt(max-min)
}

// RandSectionContainInt [min,max]
func RandSectionContainInt(min, max int) int {
	return RandSectionInt(min, max+1)
}

///////////////////////////////////////////// int32

// RandInt32 [0,n)
func RandInt32(n int32) int32 {
	if n == 0 {
		return 0
	}
	return rand.Int31n(n)
}

///////////////////////////////////////////// int64

// RandInt64 [0,n)
func RandInt64(n int64) int64 {
	if n == 0 {
		return 0
	}
	return rand.Int63n(n)
}

// RandSection [min,max)
func RandSection(min, max int64) int64 {
	if min > max {
		max = min
	}
	if max-min == 0 {
		return max
	}
	return min + RandInt64(max-min)
}

// RandSectionContain [min,max]
func RandSectionContain(min, max int64) int64 {
	return RandSection(min, max+1)
}

// GaussInt64 ...
// @Description: 正态(高斯)分布随机数生产器
// @param min 最小值
// @param max 最大值
// @param miu 期望值（均值）
// @param sigma 方差
// @return int64
func GaussInt64(min int64, max int64, miu int64, sigma int64) int64 {
	if min >= max {
		min, max = max, min
	}

	if miu < min {
		miu = min
	}
	if miu > max {
		miu = max
	}

	var x int64
	var y, dScope float64
	for i := 0; i < 10000; i++ {
		x = RandSectionContain(min, max)
		y = GaussFloat64(x, miu, sigma) * 100000
		dScope = float64(RandSectionContain(0, int64(GaussFloat64(miu, miu, sigma)*100000)))
		//注意下传的是两个miu
		if dScope <= y {
			break
		}
	}
	return x
}

///////////////////////////////////////////// float32

func Float32() float32 {
	return rand.Float32()
}

///////////////////////////////////////////// float64

func Float64() float64 {
	return rand.Float64()
}

// GaussFloat64 ...
// @Description:
// @param x
// @param miu
// @param sigma
// @return float64
func GaussFloat64(x int64, miu int64, sigma int64) float64 {
	randomNormal := 1 / (math.Sqrt(2*math.Pi) * float64(sigma)) * math.Pow(math.E, -math.Pow(float64(x-miu), 2)/(2*math.Pow(float64(sigma), 2)))
	return randomNormal
}
