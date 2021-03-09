package main

import (
	"flag"
	"fmt"
	// "math/rand"
)

func main() {
	// txTime := time.Now()
	// const layout = "2019/01/01"
	// const layout = "2006/01/02 15:04:05"
	// fmt.Println(txTime.Format(layout))
	flag.Parse()
	myargs := flag.Args()
	fmt.Println(myargs)
}
