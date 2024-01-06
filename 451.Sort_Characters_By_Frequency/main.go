package main

import (
	"fmt"
	"sort"
)

var s = "Aabb" // Output: "bbAa"

// var s = "tree"
// Output: "eert"

func main() {
	res := frequencySort(s)
	fmt.Println(res)
}

func frequencySort(s string) string {
	m := make(map[rune]int)
	// Структура для хранения пары (ключ, значение)
	type KeyValue struct {
		Key   rune
		Value int
	}
	// записываю в словарь
	for _, char := range s {
		m[char]++
	}
	// Создаем слайс для хранения пар (ключ, значение)
	var keyValuePairs []KeyValue

	// Заполняем слайс парами (ключ, значение) из мапы
	for k, v := range m {
		keyValuePairs = append(keyValuePairs, KeyValue{k, v})
	}

	// Сортируем слайс по значениям (от большего к меньшему)
	sort.Slice(keyValuePairs, func(i, j int) bool {
		return keyValuePairs[i].Value > keyValuePairs[j].Value
	})

	// Формируем результат
	result := []rune{}
	for _, pair := range keyValuePairs {
		for i := 0; i < pair.Value; i++ {
			result = append(result, pair.Key)
		}
	}

	return string(result)

}
