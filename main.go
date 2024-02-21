package main

import (
	myLogger "example/hello/loggerTest"
	"fmt"

	"rsc.io/quote"
)

func main() {
	fmt.Println("Hello, World!")
	fmt.Println(quote.Go())
	myLogger.BasicLogger()

}
