package main

import (
	"fmt"
	"reflect"
	"testing"
)

const TestText = "ШЕСТЬ Один, пять пять шесть ШесТЬ, два, ТрИ: дВа. Четыре, шесть Пять. ПЯТЬ: Шесть. Три Четыре ШЕстЬ. четыре пять трИ. ЧЕТЫРЕ"

type ResultMap struct {
	len int
	m   map[string]int
}

type ResultSlice struct {
	len int
	s   []string
}

func TestGetMap(t *testing.T) {
	var right = ResultMap{
		6,
		map[string]int{
			"один":   1,
			"два":    2,
			"три":    3,
			"четыре": 4,
			"пять":   5,
			"шесть":  6,
		}}
	result := GetMap(TestText)
	if right.len != len(result) {
		t.Errorf("Ожидается: %v, получено: %v", right.len, len(result))
	}
	for key, val := range result {
		if right.m[key] != val {
			t.Errorf("Ожидается: %v, получено: %v", right.m[key], val)
		}
	}
}

func TestGetTopKeys(t *testing.T) {
	var rigth = ResultSlice{5, []string{"шесть", "пять", "четыре", "три", "два"}}
	k := GetTopKeys(GetMap(TestText))
	if rigth.len != len(k){
		t.Errorf("Длина сайса неверна. Получено: %v  Ожидается: %v\n", len(k), rigth.len)
	}
	if reflect.DeepEqual(rigth.s, k){
		t.Errorf("Ошибка содержимого слайса.\nПолучено: %v\nОжидается: %v\n", k, rigth.s)
	}
	fmt.Printf("Ожидается:%v\nПолучено:%v\n", rigth.s, k)
}

func TestReverseSlice(t *testing.T) {

	var right = []string{"5","4","3","2","1"}
	s := ReverseSlice([]string{"1","2","3","4","5"})
	if len(right) != len(s) {
		t.Errorf("Длина сайса неверна. Получено: %v  Ожидается: %v\n", len(s), len(right))
	}
	for i:=0; i<len(right); i++{
		if right[i] != s[i]{
			t.Errorf("Ошибка содержимогo. Получено: %v  Ожидается: %v\n", s[i], right[i])
		}
	}
}