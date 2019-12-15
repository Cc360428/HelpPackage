package UtilsHelp

// 取出最小的value
func Min(m map[string]int64) (v string) {
	for k1 := range m {
		for k2 := range m {
			if m[k1] < m[k2] {
				v = k1
			}
		}
	}
	return
}

// 去除最大的
func Max(m map[string]int64) (key string) {
	var maxNumber int64
	for k := range m {
		if m[k] > maxNumber {
			maxNumber = m[k]
			key = k
		}
	}
	return
}
