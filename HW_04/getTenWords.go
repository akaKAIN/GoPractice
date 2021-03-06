package main

import (
	"fmt"
	"strings"
)

/*
Написать функцию, которая получает на вход текст и возвращает
10 самых часто встречающихся слов без учета словоформ
*/

const textRow = "Когда вся эта ученая толпа успевала приходить несколько ранее или когда знали, что профессора будут позже обыкновенного, тогда, со всеобщего согласия, замышляли бой, и в этом бою должны были участвовать все, даже цензора, обязанные смотреть за порядком и нравственностию всего учащегося сословия. Два богослова обыкновенно решали, как происходить битве: каждый ли класс должен стоять за себя особенно или все должны разделиться на две половины: на бурсу и семинарию. Во всяком случае, грамматики начинали прежде всех, и как только вмешивались риторы, они уже бежали прочь и становились на возвышениях наблюдать битву."

func main() {
	result := ReverseSlice(GetTopKeys(GetMap(textRow)))
	fmt.Println(result)
}

func GetMap(text string) map[string]int {

	var (
		tmp    string
		result = make(map[string]int)
	)

	words := strings.Fields(text)

	for _, word := range words {
		switch {
		case strings.Contains(word, ","):
			tmp = strings.Replace(word, ",", "", -1)
		case strings.Contains(word, "."):
			tmp = strings.Replace(word, ".", "", -1)
		case strings.Contains(word, ":"):
			tmp = strings.Replace(word, ":", "", -1)
		default:
			tmp = word
		}
		tmp = strings.ToLower(tmp)
		result[tmp]++
	}
	return result
}

func GetTopKeys(data map[string]int) []string {
	var row = make([][]string, len(data))
	var resultRow []string
	var max int

	for word, count := range data {
		row[count-1] = append(row[count-1], word)
	}
	for i := range row {
		resultRow = append(resultRow, row[i]...)
	}
	max = len(resultRow)
	if max <= 5 {
		return resultRow
	} else {
		return resultRow[max-5:]
	}

}

func ReverseSlice(s []string) []string {
	for i := 0; i < len(s)/2; i++ {
		s[i], s[len(s)-i-1] = s[len(s)-i-1], s[i]
	}
	return s
}

func ReverseSliceTwo(s []string) []string{
	j := len(s)-1
	for i:=0; i<j; i++{
		s[i], s[j] = s[j], s[i]
		j--
	}
	return s
}
