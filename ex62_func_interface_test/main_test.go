package ex62_func_interface_test

import "testing"

func TestArea(t *testing.T) {
	squareArea := square{width: 10, height: 10}.area()
	if squareArea != 100 {
		t.Errorf("square area is wrong")
	}
}
