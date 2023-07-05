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

func TestWinLose04(t *testing.T) {
	b := &Board{
		tokens: []int{0, 0, 0, 0, 0, 0, 0, 0, 0},
	}
	// 斜め列で勝負がついた状態での判定
	b.put(0, 0, "o")
	b.put(1, 1, "o")
	b.put(2, 2, "o")
	if !b.judge() {
		t.Errorf("Test fail expected: %t, result: %t\n", true, b.judge())
	}
}

func TestWinLose05(t *testing.T) {
	b := &Board{
		tokens: []int{0, 0, 0, 0, 0, 0, 0, 0, 0},
	}
	// 斜め列で勝負がついた状態での判定
	b.put(0, 2, "x")
	b.put(1, 1, "x")
	b.put(2, 0, "x")
	if !b.judge() {
		t.Errorf("Test fail expected: %t, result: %t\n", true, b.judge())
	}
}

func TestInvalidOperation01(t *testing.T) {
	b := &Board{
		tokens: []int{0, 0, 0, 0, 0, 0, 0, 0, 0},
	}
	// 配列外操作
	err := b.put(4, 4, "x")
	if err == nil {
		t.Errorf("Invalid Operation")
	}

	// 重複操作
	err = b.put(1, 1, "x")
	if err != nil {
		if err == nil {
			t.Errorf("Test fail %T\n", err)
		}
	}
	err = b.put(1, 1, "o")
	if err == nil {
		t.Errorf("Invalid Operation")
	}
}
