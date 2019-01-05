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
	N := readN()
	dim := N + 1
	as := make([]int, dim)
	as[0] = readN()
	for i := 1; i < dim; i++ {
		as[i] = (as[i-1] + readN()) % (1e9 + 7)
	}

	sumC := 0
	for ib := 0; ib < dim; ib++ {
		b := readN()
		c := (as[N-ib] * b) % (1e9 + 7)
		sumC = (sumC + c) % (1e9 + 7)
	}
	fmt.Println(sumC)
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
