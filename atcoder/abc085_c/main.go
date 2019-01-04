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
	N, M := readN(), readN()

	for man := 0; man < N; man++ {
		for gosen := 0; gosen < (N - man); gosen++ {
			sen := N - man - gosen
			total := man*10000 + gosen*5000 + sen*1000
			if man+gosen+sen == N && total == M {
				fmt.Printf("%d %d %d\n", man, gosen, sen)
				return
			}
		}
	}
	fmt.Println("-1 -1 -1")
}

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
