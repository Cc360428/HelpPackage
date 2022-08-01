/**
 * @Author: Cc
 * @Description: 随机函数
 * @File: randomness
 * @Version: 1.0.0
 * @Date: 2022/7/27 16:50
 * @Software : GoLand
 */

package randc

import "log"

// RandomSection 总和==1000 的数组
// 返回概率数组下标（下标从0开始的）
func RandomSection(probability []int64) (section int) {
	if len(probability) == 0 {
		log.Fatal("下标错误")
		return 0
	}
	var (
		total          int64 // 总和
		minProbability = make([]int64, 0)
	)
	for _, item := range probability {
		total += item
		minProbability = append(minProbability, total)
	}
	randomNum := RandInt64(total)
	for index := 0; index < len(minProbability); index++ {
		if index == 0 {
			if randomNum <= minProbability[index] {
				return index
			}
		} else {
			if (randomNum > minProbability[index-1]) && (randomNum <= minProbability[index]) {
				return index
			}
		}
	}
	return 0
}
