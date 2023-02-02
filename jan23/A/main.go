package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	inputBuf := bufio.NewReader(os.Stdin)
	outputBuf := bufio.NewWriter(os.Stdout)

	fmt.Fprintln(outputBuf, Work(inputBuf))

	outputBuf.Flush()
}

func Work(in *bufio.Reader) (outputStr string) {
	var testCount int

	fmt.Fscan(in, &testCount)

	for i := 0; i < testCount; i++ {
		var n, m int
		fmt.Fscan(in, &n, &m)

		outputStr += strconv.Itoa(n+m) + "\n"

	}

	return outputStr
}
