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
		dp[i] = make([]int, 3)

		prevA := dp[i-1][0]
		if bs[i] < cs[i] {
			dp[i][2] = max(prevA+cs[i], dp[i][2])
		} else {
			dp[i][1] = max(prevA+bs[i], dp[i][1])
		}
		prevB := dp[i-1][1]
		if as[i] < cs[i] {
			dp[i][2] = max(prevB+cs[i], dp[i][2])
		} else {
			dp[i][0] = max(prevB+as[i], dp[i][0])
		}
		prevC := dp[i-1][2]
		if as[i] < bs[i] {
			dp[i][1] = max(prevC+bs[i], dp[i][1])
		} else {
			dp[i][0] = max(prevC+as[i], dp[i][0])
		}
	}
	fmt.Println(max(max(dp[n-1][0], dp[n-1][1]), dp[n-1][2]))
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
