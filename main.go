package main

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println(`expecting exactly one argument`)
		os.Exit(1)
	}
	i, err := strconv.ParseInt(os.Args[1], 10, 64)
	if err != nil || i < 0 {
		fmt.Println(`expecting a positive integer number of seconds to wait`)
		os.Exit(1)
	}
	<-time.After(time.Second * time.Duration(i))
}
