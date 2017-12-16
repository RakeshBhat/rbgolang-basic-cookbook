package calculator

import "testing"

func TestMultiply(t *testing.T){
	result := multiply(5, 5)

	if result != 25 {
		t.Errorf("Multiplication failed, actual: %d, expected: %d.", result, 25)
	}
}

func TestSubtract(t *testing.T){
	result := subtract(37, 8)

	if result != 29 {
		t.Errorf("Subtraction failed, actual: %d, expected: %d.", result, 29)
	}
}

func TestAddition(t *testing.T)  {
	result := addition(28, 762)

	if result != 790 {
		t.Errorf("Addition failed, actual: %d, expected: %d.", result, 790)
	}
}

func TestDivide(t *testing.T)  {
	result := divide(144 , 2)

	if result != 72 {
		t.Errorf("Division failed, actual: %d, expected: %d.", result, 72)
	}
}
