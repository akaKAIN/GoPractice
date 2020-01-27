package main

import (
	"fmt"
	"testing"
)

const TestText = "Один, два, ТрИ: дВа. Четыре, Три Четыре четыре трИ. ЧЕТЫРЕ"

type Result struct {
	len int
	m   map[string]int
}

func TestGetMap(t *testing.T) {
	var right = Result{
		4,
		map[string]int{
			"один":   1,
			"два":    2,
			"три":    3,
			"четыре": 4,
		}}
	result := GetMap(TestText)
	fmt.Println(result)
	fmt.Println(right.m)
	if right.len != len(result) {
		t.Errorf("Ожидается: %v, получено: %v", right.len, len(result))
	}
	for key, val := range result {
		if right.m[key] != val {
			t.Errorf("Ожидается: %v, получено: %v", right.m[key], val)
		}
	}
}
