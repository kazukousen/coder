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
	for i := 0; i < n; i++ {
		as[i] = readN()
	}
	min := as[0]
	cnt := 1

	for _, a := range as {
		if min > a {
			cnt++
			min = a
		}
	}
	fmt.Println(cnt)
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
