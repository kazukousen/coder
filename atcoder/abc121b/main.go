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
	N, M, C := readN(), readN(), readN()
	bs := make([]int, M)
	for i := 0; i < M; i++ {
		bs[i] = readN()
	}
	var sum int
	for i := 0; i < N; i++ {
		m := 0
		for ii := 0; ii < M; ii++ {
			m += readN() * bs[ii]
		}
		if m+C > 0 {
			sum++
		}
	}
	fmt.Println(sum)
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
