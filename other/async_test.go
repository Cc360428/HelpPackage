package other

import (
	"fmt"
	"log"
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
				log.Println(resultItems[1])
			} else {
				log.Println(resultItems[0])
			}
		}
	}
}

func test(t string, age int) (string, error) {
	return fmt.Sprintf("%s-%d", t, age), nil
}
