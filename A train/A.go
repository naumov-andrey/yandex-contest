package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	rdr := bufio.NewReader(os.Stdin)

	var n int
	_, _ = fmt.Fscanln(rdr, &n)

	var first int
	_, _ = fmt.Fscan(rdr, &first)
	last := first
	for i := 1; i < n; i++ {
		var cur int
		_, _ = fmt.Fscan(rdr, &cur)

		if cur < last {
			fmt.Print(-1)
			return
		}
		
		last = cur
	}

	fmt.Print(last - first)
}
