package main

import (
	"flag"
	"fmt"

	"github.com/zchee/dispatch"
)

var max = flag.Int("max", 0, "enter integer or bust!")

func main() {
	flag.Parse()

	queue := dispatch
	dispatch.Apply(*max, queue, fizzbuzz(i))
}

func fizzbuzz(i int) {
	fizz := "fizz"
	buzz := "buzz"

	if i%3 == 0 && i%5 == 0 {
		fmt.Println(i, fizz+buzz)
	} else if i%3 == 0 {
		fmt.Println(i, fizz)
	} else if i%5 == 0 {
		fmt.Println(i, buzz)
	} else {
		fmt.Println(i)
	}
}
