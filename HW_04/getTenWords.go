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
	result := GetMap(textRow)
	fmt.Println(result)
}

func GetMap(text string) map[string]int {
	//var row []Word
	var tmp string

	words := strings.Fields(text)
	fmt.Printf("Длина строки: %d\n", len(words))
	var result = make(map[string]int, len(words))
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
	fmt.Println("Длина словаря:", len(result))

	return result
}
