package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	run()
}

func run() {
	n := readN()
	var sum float64
	for i := 0; i < n; i++ {
		xu := strings.Split(read(), " ")
		x, _ := strconv.ParseFloat(xu[0], 64)
		switch xu[1] {
		case "JPY":
			sum += x
		case "BTC":
			x *= 380000
			sum += x
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
