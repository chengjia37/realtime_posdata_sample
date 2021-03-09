package main

import (
	"flag"
	"fmt"

	"github.com/hpcloud/tail"
)

func main() {
	flag.Parse()
	args := flag.Args()
	if len(args) != 1 {
		fmt.Println("Usage: read_csv [target file]")
		return
	}
	targetFile := args[0]
	t, err := tail.TailFile(targetFile, tail.Config{Follow: true, Poll: true})
	if err != nil {
		fmt.Println(err)
		return
	}

	for line := range t.Lines {
		fmt.Println(line.Text)
	}
}
