package main

import (
	"testing"
)

func TestCalculate(t *testing.T) {
	tests := []struct {
		a         int
		op        string
		b         int
		expected  int
		expectErr bool
	}{
		{a: 2, op: "+", b: 3, expected: 5},
		{a: 5, op: "-", b: 2, expected: 3},
		{a: 4, op: "*", b: 2, expected: 8},
		{a: 10, op: "/", b: 0, expectErr: true},
	}
	for _, test := range tests {
		result, err := calculate(test.a, test.op, test.b)
		if test.expectErr {
			if err == nil {
				t.Errorf("expected error, got nil\n\a=%d op=%s b=%d",
					test.a, test.op, test.b)
			}
		} else if err != nil || result != test.expected {
			t.Errorf("expected %v, got %v, err %v", test.expected, result, err)
		}
	}

}
