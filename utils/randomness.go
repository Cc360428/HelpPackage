/**
 * @Author cc
 * @Date 2021/4/1 10:27
 * @Description $ 条件随机
 **/
package utils

import (
	"github.com/Cc360428/HelpPackage/utils/logs"
	"math/rand"
	"time"
)

// 总和==1000 的数组
// 返回概率数组下标（下标从0开始的）
func RandomSection(probability []int64) (section int) {
	if len(probability) == 0 {
		logs.Error("错误 下标")
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
	randomNum := random(total)
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

// 最大返回随机数
func random(value int64) int64 {
	rand.Seed(time.Now().UnixNano())
	if value == 0 {
		return 0
	}
	return rand.Int63n(value)
}
