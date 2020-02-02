package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
	run()
}

func run() {
	setSpace()
	n, k := readN(), readN()
	as := make([]int, n)
	for i := 0; i < n; i++ {
		as[i] = readN()
	}
	if n <= k {
		fmt.Println("0")
		return
	}
	sort.Ints(as)
	sum := 0
	for _, a := range as[:(n - k)] {
		sum += a
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
