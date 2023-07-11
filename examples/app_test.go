package gorgonzola

import (
	"testing"
)

func TestIntMinBasic(t *testing.T) {
	ans := seven()
	if ans != 7 {
		t.Errorf("seven() = %d; want 7", ans)
	}
}
