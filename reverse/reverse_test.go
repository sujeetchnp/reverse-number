package reverse_test

import (
	"fmt"
	rn "reverse-number/reverse"
	"testing"
)

func TestReverseNumber(t *testing.T) {
	var number int64 = 1234
	var expected int64 = 4321
	got := rn.ReverseNumber(number)
	if got != expected {
		t.Errorf("ReverseNumber(%d) = %d; want %d", number, got, expected)
		fmt.Println(number, got, expected)
	}
	number = 24356879
	expected = 97865342
	got = rn.ReverseNumber(number)
	if got != expected {
		t.Errorf("ReverseNumber(%d) = %d; want %d", number, got, expected)
	}
}
