package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	var name string = "World"
	if len(os.Args) == 2 {
		name = os.Args[1]
	}
	fmt.Println("Hello, " + name)
	time.Sleep(200000 * time.Second)
	fmt.Println("Finished")
}
