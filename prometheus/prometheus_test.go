/**
 * @Author: Cc
 * @Description: 描述
 * @File: prometheus_test.go
 * @Version: 1.0.0
 * @Date: 2022/8/2 10:24
 * @Software : GoLand
 */

package prometheus

import "testing"

func TestInitStart(t *testing.T) {
	InitStart("0.0.0.0:2112")
}
