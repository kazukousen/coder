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
	as := make([]int, n)
	bs := make([]int, n)
	cs := make([]int, n)
	for i := 0; i < n; i++ {
		as[i], bs[i], cs[i] = readN(), readN(), readN()
	}

	dp := make([][]int, n)
	dp[0] = []int{as[0], bs[0], cs[0]}

	for i := 1; i < n; i++ {
		prevA, prevB, prevC := dp[i-1][0], dp[i-1][1], dp[i-1][2]
		dp[i] = []int{0, 0, 0}
		// if A was selected previous,
		chMax(&dp[i][1], prevA+bs[i])
		chMax(&dp[i][2], prevA+cs[i])

		// if B was selected previous,
		chMax(&dp[i][0], prevB+as[i])
		chMax(&dp[i][2], prevB+cs[i])

		// if C was selected previous
		chMax(&dp[i][0], prevC+as[i])
		chMax(&dp[i][1], prevC+bs[i])
	}

	fmt.Println(max(max(dp[n-1][0], dp[n-1][1]), dp[n-1][2]))
}

func chMax(a *int, b int) {
	if *a > b {
		return
	}
	*a = b
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
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
