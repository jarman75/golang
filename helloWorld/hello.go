package main

import (
	"fmt"

	"rsc.io/quote"
)

func main() {
	Hello()
	Go()
}
func Hello() {
	fmt.Println(quote.Hello())
}
func Go() {
	fmt.Println(quote.Go())
}
