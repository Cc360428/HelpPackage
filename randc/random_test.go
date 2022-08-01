package randc

import "testing"

func TestRandMax(t *testing.T) {
	t.Log(RandInt64(88))
	t.Log(RandInt32(88))
	t.Log(RandInt(88))
}

func TestRandInt(t *testing.T) {
	t.Log(RandSection(4, 5))
	t.Log(RandSectionContain(4, 5))
}
