package helloworld

import "testing"

func TestPuts(t *testing.T) {
	for i := 0; i < 3; i++ {
		t.Log(Puts(i))
	}
}
