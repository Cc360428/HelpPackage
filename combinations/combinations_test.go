/**
 * @Author cc
 * @Date 2022/3/26 09:34
 * @Description $
 **/
package combinations

import (
	"testing"
)

func TestAll(t *testing.T) {
	for i, value := range All([]string{"1", "2", "3", "4"}) {
		t.Log(i, value)
	}
}

func TestCombinations(t *testing.T) {
	for i, value := range Combinations([]string{"1", "2", "3", "4"}, 3) {
		t.Log(i, value)
	}
}
