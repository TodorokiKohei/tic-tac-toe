package main

import "testing"

func TestPutToken01(t *testing.T) {
	b := &Board{
		tokens: []int{0, 0, 0, 0, 0, 0, 0, 0, 0},
	}
	b.put(1, 1, "o")
	if b.get(1, 1) != "o" {
		t.Errorf("Test fail expected: %s, result: %s\n", "o", b.get(1, 1))
	}
}

func TestPutToken02(t *testing.T) {
	b := &Board{
		tokens: []int{0, 0, 0, 0, 0, 0, 0, 0, 0},
	}
	b.put(1, 1, "x")
	if b.get(1, 1) != "x" {
		t.Errorf("Test fail expected: %s, result: %s\n", "x", b.get(1, 1))
	}
}

func TestWinLose01(t *testing.T) {
	b := &Board{
		tokens: []int{0, 0, 0, 0, 0, 0, 0, 0, 0},
	}
	// 勝負がついていない状態での判定
	if b.judge() {
		t.Errorf("Test fail expected: %t, result: %t\n", false, b.judge())
	}
}

func TestWinLose02(t *testing.T) {
	b := &Board{
		tokens: []int{0, 0, 0, 0, 0, 0, 0, 0, 0},
	}
	// 縦列で勝負がついた状態での判定
	b.put(0, 0, "o")
	b.put(0, 1, "o")
	b.put(0, 2, "o")
	if !b.judge() {
		t.Errorf("Test fail expected: %t, result: %t\n", true, b.judge())
	}
}

func TestWinLose03(t *testing.T) {
	b := &Board{
		tokens: []int{0, 0, 0, 0, 0, 0, 0, 0, 0},
	}
	// 横列で勝負がついた状態での判定
	b.put(0, 0, "o")
	b.put(1, 0, "o")
	b.put(2, 0, "o")
	if !b.judge() {
		t.Errorf("Test fail expected: %t, result: %t\n", true, b.judge())
	}
}
