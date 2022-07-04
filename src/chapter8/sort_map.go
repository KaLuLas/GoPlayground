package chapter8

import (
	"fmt"
	"sort"
)

func SortMap() {
	rawMap := map[string]int{"alpha": 34, "bravo": 56, "charlie": 23,
		"delta": 87, "echo": 56, "foxtrot": 12,
		"golf": 34, "hotel": 16, "indio": 87,
		"juliet": 65, "kili": 43, "lima": 98}

	fmt.Println("unsorted:")
	for k, v := range rawMap {
		fmt.Printf("origninal key %v with value %v \n", k, v)
	}

	keys := make([]string, len(rawMap))
	i := 0
	for k := range rawMap {
		keys[i] = k
		i++
	}

	sort.Strings(keys)
	fmt.Println("sorted")
	for _, key := range keys {
		fmt.Printf("sorted key %v, value %v \n", key, rawMap[key])
	}
}
