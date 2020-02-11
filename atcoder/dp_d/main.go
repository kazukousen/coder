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
	N, W := int(readN()), readN()

	dp := make([][]int, N+1)
	// dp[i][sumW]
	// i 番目まで品物を選んで、合計の重さが sumW のときの価値の最大値

	dp[0] = make([]int, W+1)
	for i := 0; i < N; i++ {
		w, v := readN(), readN()
		dp[i+1] = make([]int, W+1)
		// i + 1 番目の更新は、
		for sumW := 0; sumW <= W; sumW++ {
			// i 番目までの価値をそのまま更新する
			dp[i+1][sumW] = dp[i][sumW]
			if sumW-w >= 0 {
				// もし、選べるとしたら, 比較して、価値があがるのであれば更新する
				chmax(&dp[i+1][sumW], dp[i][sumW-w]+v)
			}
		}
	}

	fmt.Println(dp[N][W])
}

func chmax(a *int, b int) {
	if *a > b {
		return
	}
	*a = b
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
