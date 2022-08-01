/**
 * @Author: Cc
 * @Description: 描述
 * @File: randomness_test.go
 * @Version: 1.0.0
 * @Date: 2022/7/27 16:51
 * @Software : GoLand
 */

package randc

import "testing"

func TestRandomSection(t *testing.T) {
	p := make([]int64, 10)
	p[0] = 100
	p[1] = 50
	p[2] = 10
	p[3] = 40
	p[4] = 300
	p[5] = 30
	p[6] = 25
	p[7] = 5
	p[8] = 50
	p[9] = 400
	t.Log(RandomSection(p))
}
