package main

import (
	"reflect"
	"sort"
	"strings"
)

/*
создаю массив мап, в ключах которых будут лежать
руны, а в значениях количество их в слове.
*/
var arrayOfSets []map[rune]int

/*функция, которая преобразует слово в структуру Set*/
func wordToSet(word string) map[rune]int {
	m := make(map[rune]int)
	for _, r := range word {
		_, exists := m[r]
		if exists {
			m[r] = m[r] + 1
		} else {
			m[r] = 1
		}
	}
	return m
}

/*
функция, проверяющая, есть ли представление символов слова
в массиве
*/
func toCheckPresence(word_splitted map[rune]int) int {
	for id, set := range arrayOfSets {
		eq := reflect.DeepEqual(set, word_splitted)
		if eq {
			//если будет найдено такое же множество, вернёт id
			return id
		} else {
			continue
		}
	}
	return -1
}

func FindAnnagrams(array *[]string) map[string][]string {
	//сначала создадим саму мапу
	result := make(map[string][]string)

	//теперь создадим мапу
	angrms := make(map[int][]string)
	for _, word := range *array {
		word = strings.ToLower(word)
		//разбиваем слово на массив рун
		word_splitted := wordToSet(word)
		//проверяем наличие такого же набора рун в предыдущих словах
		existed := toCheckPresence(word_splitted)
		//если такого набора не было, добавляем полученный набор
		if existed == -1 {
			arrayOfSets = append(arrayOfSets, word_splitted)
			angrms[len(angrms)] = []string{word}
		} else {
			//иначе обновляем массив аннаграмм(если такого слова ещё не было)
			for i, word_ := range angrms[existed] {
				if word == word_ {
					break
				}
				if i == (len(angrms[existed]) - 1) {
					angrms[existed] = append(angrms[existed], word)
				}
			}
		}
	}
	for _, words := range angrms {
		if len(words) > 1 {
			key := words[0]
			value := words[1:]
			sort.Strings(value)
			result[key] = value
		}
	}
	return result
}
