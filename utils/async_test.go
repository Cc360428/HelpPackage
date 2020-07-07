package utils

import (
	"fmt"
	"github.com/Cc360428/HelpPackage/utils/logs"
	"testing"
)

func TestNew(t *testing.T) {
	a := New()
	defer a.Clean()
	a.Add("test", test, "Cc", 21)
	if chans, ok := a.Run(); ok {
		res := <-chans
		for _, resultItems := range res {
			if resultItems[1] != nil {
				logs.Error(resultItems[1])
			} else {
				logs.Info(resultItems[0])
			}
		}
	}
}

func test(t string, age int) (string, error) {
	return fmt.Sprintf("%s-%d", t, age), nil
}
