package combinations

import "math/bits"

func All(set []string) (subsets [][]string) {
	length := uint(len(set))

	for subsetBits := 1; subsetBits < (1 << length); subsetBits++ {
		var subset []string

		for object := uint(0); object < length; object++ {

			if (subsetBits>>object)&1 == 1 {

				subset = append(subset, set[object])
			}
		}

		subsets = append(subsets, subset)
	}
	return subsets
}

func Combinations(set []string, n int) (subsets [][]string) {
	length := uint(len(set))

	if n > len(set) {
		n = len(set)
	}

	for subsetBits := 1; subsetBits < (1 << length); subsetBits++ {
		if n > 0 && bits.OnesCount(uint(subsetBits)) != n {
			continue
		}

		var subset []string

		for object := uint(0); object < length; object++ {

			if (subsetBits>>object)&1 == 1 {

				subset = append(subset, set[object])
			}
		}

		subsets = append(subsets, subset)
	}
	return subsets
}
