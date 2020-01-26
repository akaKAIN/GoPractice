package main

import (
	"fmt"
	"log"
	"strconv"
)

/*
Домашнее задание
Распаковка строки
Создать Go функцию, осуществляющую примитивную распаковку строки, содержащую повторяющиеся символы / руны, например:
* "a4bc2d5e" => "aaaabccddddde"
* "abcd" => "abcd"
* "45" => "" (некорректная строка)
Дополнительное задание: поддержка escape - последовательности
* `qwe\4\5` => `qwe45` (*)
* `qwe\45` => `qwe44444` (*)
* `qwe\\5` => `qwe\\\\\` (*)
*/

func main() {
	var input string

	fmt.Println("Inter row")
	if _, err := fmt.Scan(&input); err != nil {
		log.Fatal(err)
	}

	sepRow := Separator(input)
	result, err := MultiplySymbol(sepRow)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(result)

}

func Separator(row string) []string {
	var tmp string
	var result []string
	for i := range row {
		switch {
		case row[i] >= 48 && row[i] <= 57:
			tmp += string(row[i])
		case row[i] > 57 && tmp != "":
			result = append(result, tmp, string(row[i]))
			tmp = ""
		default:
			result = append(result, string(row[i]))
		}
	}

	if tmp != ""{
		result = append(result, tmp)
	}
	return result
}

func MultiplySymbol(strSlice []string) (string, error) {
	var result string
	for i := range strSlice{
		num, err := strconv.Atoi(strSlice[i])
		if err != nil {
			result += strSlice[i]
			continue
		}
		for j:=0; j<num-1; j++{
			if i == 0{
				err = fmt.Errorf("Некорректная строка\n")
				return strSlice[i], err
			}
			result += strSlice[i-1]
		}
	}
	return result, nil
}
