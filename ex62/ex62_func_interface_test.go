package ex62

import "testing"

func TestArea(t *testing.T) {
	squareArea := square{width: 10, height: 10}.area()
	if squareArea != 100 {
		t.Errorf("square area is wrong")
	}
}
