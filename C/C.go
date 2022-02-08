package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func packAddresses(n []int) []interface{} {
	p := make([]interface{}, len(n))
	for i := range n {
		p[i] = &n[i]
	}
	return p
}

func main() {
	in := bufio.NewReader(os.Stdin)

	var n, m int
	_, _ = fmt.Fscan(in, &n, &m)

	plates := make([]int, n)
	_, _ = fmt.Fscan(in, packAddresses(plates)...)

	guards := make([]int, m)
	_, _ = fmt.Fscan(in, packAddresses(guards)...)
	sort.Sort(sort.Reverse(sort.IntSlice(guards)))

	spaces := make([]int, 0, n)
	spaces = append(spaces, plates[n-1])
	max := plates[n-1]
	for i := n - 1; i >= 0; i-- {
		if plates[i] <= max {
			continue
		}
		spaces = append(spaces, plates[i]-max)
		max = plates[i]
	}
	sort.Sort(sort.Reverse(sort.IntSlice(spaces)))

	count := 0
	for i := 0; i < m && count < len(spaces); i++ {
		if guards[i] <= spaces[count] {
			count++
		}
	}

	fmt.Print(count)
}
