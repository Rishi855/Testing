package main

import (
	"fmt"
	"sort"
)

func main() {
	data := map[int]int{
		10: 100,
		2:  30,
		7:  70,
		1:  10,
	}

	keys := make([]int, 0, len(data))

	for k := range data {
		keys = append(keys, k)
	}

	sort.Ints(keys)

	for _, k := range keys {
		fmt.Printf("key: %d value: %d\n",k,data[k])
	}
}