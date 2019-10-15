package mylib

import (
	"testing"
)

func Test_BasicChecks(t *testing.T) {
	t.Parallel()
	t.Run("Go can add", func(t *testing.T ) {
		if 1 + 1 != 2 {
			t.Fail()
		}
	})
	t.Run("Go can concatenate strings", func(t *testing.T) {
		if "Hello, " + "Go" != "Hello, Go" {
			t.Fail()
		}
	})
}

func TestAdder(t *testing.T) {
	if adder(2, 5) != 7 {
		t.Fail()
	}
}

func BenchmarkAdder(b *testing.B) {
	for i:=0; i < 10; i ++{
		adder(5, 7)
	}
}
