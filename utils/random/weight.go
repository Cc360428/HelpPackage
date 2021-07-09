// 权重随机出结果
// 返回下标
package random

func WeightedRandomIndex(weights []float32) int {
	if len(weights) == 1 {
		return 0
	}
	var sum float32 = 0.0
	for _, w := range weights {
		sum += w
	}
	r := Float32() * sum
	var t float32 = 0.0
	for i, w := range weights {
		t += w
		if t > r {
			return i
		}
	}
	return len(weights) - 1
}

// And must be 1
func WeightF(weight []float64) int {

	if len(weight) == 0 {
		return 0
	}

	tmp := make([]float64, len(weight), len(weight))
	index := 0

	for _, v := range weight {
		if index != 0 {
			tmp[index] = tmp[index-1] + v
		} else {
			tmp[index] = v
		}
		index++
	}

	//随机一个数
	r := Float64()
	bucket := 0
	for r > tmp[bucket] {
		bucket++
	}

	return bucket
}

// 随机数权重
func Weight(weight []uint64) int64 {
	if len(weight) == 0 {
		return 0
	}

	tmp := make([]uint64, len(weight), len(weight))
	var index int64
	var total int64

	for _, v := range weight {
		if index != 0 {
			tmp[index] = tmp[index-1] + v
		} else {
			tmp[index] = v
		}
		total += int64(v)
		index++
	}

	r := RandSectionContain(0, total)
	var rt int64

	for uint64(r) > tmp[rt] {
		rt++
	}

	return rt
}
