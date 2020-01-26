package main

import "testing"

func TestSeparator(t *testing.T) {
	inputString := "a23b43d3njs3"
	rightResult := []string{"a", "23", "b", "43", "d", "3", "n", "j", "s", "3"}
	s := Separator(inputString)
	if len(rightResult) == len(s) {
		for i:=0; i<len(s); i++{
			if rightResult[i] != s[i]{
				t.Errorf("Ожидается: %q, получено: %q", rightResult, s)
			}
		}
	}
}

func TestMultiplySymbol(t *testing.T) {
	inputSlice := []string{"a", "1", "b", "2", "c", "4", "f", "1"}
	rightResult := "abbccccf"
	if m, _ := MultiplySymbol(inputSlice); m != rightResult {
		t.Errorf("Ожидается: %q, получено: %q", rightResult, m)
	}
}



