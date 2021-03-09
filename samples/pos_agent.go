package main

import (
	"flag"
	"fmt"
)

func main(){

	var i = flag.Int("int", 0, "int flag")
	var s = flag.String("str", "default", "string flag")
	var b = flag.Bool("bool", false, "bool flag")
	flag.Parse()

	args := flag.Args()
	fmt.Println(*i, *s, *b)
	fmt.Println(args)
}