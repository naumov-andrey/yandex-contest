package main

import (
	"fmt"
	"sort"
)

func buildValueToIndexMap(nums []int) map[int][]int {
	m := make(map[int][]int)
	for i, v := range nums {
		m[v] = append(m[v], i)
	}
	return m
}

func all(vs []bool) bool {
	for _, v := range vs {
		if !v {
			return false
		}
	}
	return true
}

func main() {
	var a [3]int
	_, _ = fmt.Scan(&a[0], &a[1], &a[2])

	canBecomeMedian := make([]bool, len(a))

	sortedA := make([]int, 3)
	copy(sortedA, a[:])
	sort.Ints(sortedA)

	for _, idx := range buildValueToIndexMap(a[:])[sortedA[1]] {
		canBecomeMedian[idx] = true
	}

	for i := 0; i < len(a) && !all(canBecomeMedian); i++ {
		for j := 0; j < len(a) && !all(canBecomeMedian); j++ {
			if j == i {
				continue
			}

			b := make([]int, 3)
			copy(b, a[:])
			b[i] = b[i] - b[j]

			sortedB := make([]int, 3)
			copy(sortedB, b[:])
			sort.Ints(sortedB)

			for _, idx := range buildValueToIndexMap(b)[sortedB[1]] {
				canBecomeMedian[idx] = true
			}
		}
	}

	for _, v := range canBecomeMedian {
		if v {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
	}
}
