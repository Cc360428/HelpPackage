package random

import "testing"

func TestWeightedRandomIndex(t *testing.T) {
	t.Log(WeightedRandomIndex([]float32{0.1, 0.2, 0.3, 0.4}))
	t.Log(WeightF([]float64{0.1, 0.2, 0.3, 0.4}))
	t.Log(Weight([]uint32{1122, 121, 212}))
}
