package main

import (
	"bufio"
	"os"
	"strconv"
)

func main() {
	run()
}

func run() {
	setSpace()
	N, _ := readN(), readN()
	m := make(map[int]int, N)
	min := 0
	for i := 0; i < N; i++ {
		n := readN()
		m[n] = readN()
		if min > n {
			min = 0
		}
	}
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
