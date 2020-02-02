package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	run()
}

func run() {
	setSpace()
	n := readN()
	hs := make([]int, n)
	for i := 0; i < n; i++ {
		hs[i] = readN()
	}
	dp := make([]int, n)
	dp[0] = 0
	dp[1] = abs(hs[0] - hs[1])

	for i := 2; i < n; i++ {
		var cost int
		costSingle := abs(hs[i-1]-hs[i]) + dp[i-1]
		costDouble := abs(hs[i-2]-hs[i]) + dp[i-2]
		if costSingle < costDouble {
			cost = costSingle
		} else {
			cost = costDouble
		}

		dp[i] = cost
	}
	fmt.Println(dp[n-1])
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

// ------以下、ユーティリティ------

var sc = bufio.NewScanner(os.Stdin)

func setSpace() {
	sc.Split(bufio.ScanWords)
}

func read() string {
	sc.Scan()
	return sc.Text()
}

func readN() int {
	n, err := strconv.Atoi(read())
	if err != nil {
		panic(err)
	}
	return n
}
