package other

import (
	"testing"
)

func TestEmail(t *testing.T) {
	email := "li_chao_cheng@163.com"
	gotB := Email(email)
	t.Log(gotB)
}
