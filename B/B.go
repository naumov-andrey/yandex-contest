package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func printGlass(glass [][]rune) {
	for _, layer := range glass {
		fmt.Print(string(layer))
	}
}

func main() {
	in := bufio.NewReader(os.Stdin)

	var n, m int
	_, _ = fmt.Fscanln(in, &n, &m)

	glass := make([][]rune, n)

	for i := 0; i < n; i++ {
		layer, _ := in.ReadString('\n')
		glass[i] = []rune(layer)
	}

	var k int
	_, _ = fmt.Fscanln(in, &k)

	free := n - 1
	for i := 0; i < k; i++ {
		row, _ := in.ReadString('\n')
		params := strings.Split(row, " ")

		count, _ := strconv.Atoi(params[1])
		symbol := []rune(params[2])[0]

		for j := 0; j < count; j++ {
			idx := free - 1 - j
			for i := range glass[idx] {
				if glass[idx][i] == ' ' {
					glass[idx][i] = symbol
				}
			}
		}
		
		free -= count
	}

	printGlass(glass)
}
