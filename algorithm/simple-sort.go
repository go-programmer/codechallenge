package algorithm

import (
	"fmt"
	"sort"
)

func sortByKeys() {

	unSortedMap := map[string]int{}
	unSortedMap["India"] = 20
	unSortedMap["Canada"] = 70
	unSortedMap["Germany"] = 15
	keys := make([]string, 0, len(unSortedMap))

	for k := range unSortedMap {
		keys = append(keys, k)
	}

	sort.Strings(keys)

	for _, k := range keys {
		fmt.Println(k, unSortedMap[k])
	}
}

func sortByValues() {
	unSortedMap := map[string]int{"India": 20, "Canada": 70, "Germany": 15}

	// Int slice to store values of map.
	values := make([]int, 0, len(unSortedMap))

	for _, v := range unSortedMap {
		values = append(values, v)
	}

	// Sort slice values.
	sort.Ints(values)

	// Print values of sorted Slice.
	for _, v := range values {
		fmt.Println(v)
	}
}

// SimpleSortClient sorting example client
func SimpleSortClient() {
	fmt.Println("Sort by Key")
	sortByKeys()
	fmt.Println("Sort by Value")
	sortByValues()
}
