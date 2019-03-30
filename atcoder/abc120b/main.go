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
	a, b, k := readN(), readN(), readN()
	n := max(a, b)
	ks := make([]int, 0, n+1)
	for i := 1; i <= n; i++ {
		if a%i == 0 && b%i == 0 {
			ks = append(ks, i)
		}
	}
	fmt.Println(ks[len(ks)-k])
}

func max(a, b int) int {
	var n int
	if a < b {
		n = a
	} else {
		n = b
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
