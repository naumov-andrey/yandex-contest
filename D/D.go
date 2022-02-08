package main

import (
	"bufio"
	"fmt"
	"os"
)

type CellType int

const (
	Default CellType = iota
	InnerCorner
	OuterCorner
	Border
)

var dirVectorsToAngle = map[[4]int]int{
	[4]int{1, 0, 1, 0}:  0,
	[4]int{1, 0, 0, 1}:  90,
	[4]int{1, 0, -1, 0}: 180,
	[4]int{1, 0, 0, -1}: 270,

	[4]int{0, 1, 1, 0}:  270,
	[4]int{0, 1, 0, 1}:  0,
	[4]int{0, 1, -1, 0}: 90,
	[4]int{0, 1, 0, -1}: 180,

	[4]int{-1, 0, 1, 0}:  180,
	[4]int{-1, 0, 0, 1}:  270,
	[4]int{-1, 0, -1, 0}: 0,
	[4]int{-1, 0, 0, -1}: 90,

	[4]int{0, -1, 1, 0}:  90,
	[4]int{0, -1, 0, 1}:  180,
	[4]int{0, -1, -1, 0}: 270,
	[4]int{0, -1, 0, -1}: 0,
}

func markCells(field [][]CellType, prevX, x, prevY, y int) int {
	count := 0
	if prevX-x != 0 {
		for i := min(prevX, x) + 1; i < max(prevX, x); i++ {
			field[i-1][y-1] = Border
			count++
		}
	} else if prevY-y != 0 {
		for i := min(prevY, y) + 1; i < max(prevY, y); i++ {
			field[x-1][i-1] = Border
			count++
		}
	}
	return count
}

func main() {
	in := bufio.NewReader(os.Stdin)

	var n, m int
	_, _ = fmt.Fscan(in, &n, &m)

	field := make([][]CellType, n)
	for i := range field {
		field[i] = make([]CellType, m)
	}

	var k int
	_, _ = fmt.Fscan(in, &k)

	var firstX, firstY int
	_, _ = fmt.Fscan(in, &firstX, &firstY)
	prevX, prevY := firstX, firstY

	firstDirVectorX, firstDirVectorY := 0, 0
	prevDirVectorX, prevDirVectorY := 0, 0

	border, inner, outer := 0, 0, 0

	for i := 1; i < k; i++ {
		var x, y int
		_, _ = fmt.Fscan(in, &x, &y)

		dirVectorX, dirVectorY := norm(x-prevX), norm(y-prevY)

		if firstDirVectorX == firstDirVectorY {
			firstDirVectorX, firstDirVectorY = dirVectorX, dirVectorY
		}

		if prevDirVectorX != prevDirVectorY {
			angle := dirVectorsToAngle[[4]int{prevDirVectorX, prevDirVectorY, dirVectorX, dirVectorY}]
			if angle == 90 {
				field[prevX-1][prevY-1] = OuterCorner
				outer++
			} else if angle == 270 {
				field[prevX-1][prevY-1] = InnerCorner
				inner++
			}
		}

		border += markCells(field, prevX, x, prevY, y)
		prevX, prevY = x, y
		prevDirVectorX, prevDirVectorY = dirVectorX, dirVectorY
	}
	dirVectorX, dirVectorY := norm(firstX-prevX), norm(firstY-prevY)
	angle := dirVectorsToAngle[[4]int{dirVectorX, dirVectorY, firstDirVectorX, firstDirVectorY}]
	if angle == 90 {
		field[prevX-1][prevY-1] = OuterCorner
		outer++
	} else if angle == 270 {
		field[prevX-1][prevY-1] = InnerCorner
		inner++
	}
	border += markCells(field, prevX, firstX, prevY, firstY)

	//for i := range field {
	//	fmt.Println(field[i])
	//}

	full := 0

	fmt.Print(full, border, outer, inner)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func norm(x int) int {
	if x == 0 {
		return 0
	}
	return x / abs(x)
}
