package loopback

import (
	"fmt"
	"testing"
)

func TestLoopback6(t *testing.T) {
	a, err := Loopback6()
	fmt.Println("err", err)
	fmt.Println("addr", a)
	t.Fail()
}
