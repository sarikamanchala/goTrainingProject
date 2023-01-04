package main

import (
	"testing"
)

type mockSQUARE struct{ length int }

func (square mockSQUARE) AREA() float32 {
	return float32(square.length * square.length)
}

func TestArea1(t *testing.T) {
	s := mockSQUARE{2}
	result := s.AREA()
	expected := 4
	if result != float32(expected) {
		t.Errorf("Not matched")
	} else {
		t.Logf("matched")
	}
}
func TestArea(t *testing.T) {
	m := SQUARE{2}
	c := CIRCLE{2}
	result := m.AREA()
	expected := 4
	if result != float32(expected) {
		t.Errorf("Not matched")
	} else {
		t.Logf("matched")
	}
	result2 := c.AREA()
	expected2 := 4
	if result2 != float32(expected2) {
		t.Errorf("Not matched")
	} else {
		t.Logf("matched")
	}
}

func TestHello(t *testing.T) {
	expected := "Hello"
	result := Hello()
	if result != expected {
		t.Errorf("not matched")
	} else {
		t.Logf("matched")
	}
}
