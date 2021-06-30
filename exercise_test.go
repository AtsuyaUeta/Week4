package main

import "testing"

func TestPutToken01(t *testing.T) {
	b := &Board{
		tokens: []int{0, 0, 0, 0, 0, 0, 0, 0, 0},
	}
	b.put(1, 1, "o")
	if b.get(1, 1) != "o" {
		t.Errorf("....")
	}
}

func TestPutToken02(t *testing.T) {
	b := &Board{
		tokens: []int{0, 0, 0, 0, 0, 0, 0, 0, 0},
	}
	b.put(2, 2, "x")
	if b.get(2, 2) != "x" {
		t.Errorf("....")
	}
}

func TestPutToken03(t *testing.T) {
	b := &Board{
		tokens: []int{1, 10, 1, 10, 1, 10, 10, 1, 1},
	}
	v := judge(b)
	if v != 1 {
		t.Errorf("....")
	}
}

func TestPutToken04(t *testing.T) {
	b := &Board{
		tokens: []int{1, 10, 1, 1, 10, 0, 0, 10, 0},
	}
	v := judge(b)
	if v != 2 {
		t.Errorf("....")
	}
}

func TestPutToken05(t *testing.T) {
	b := &Board{
		tokens: []int{1, 10, 1, 1, 10, 10, 10, 1, 1},
	}
	v := judge(b)
	if v != 0 {
		t.Errorf("....")
	}
}