package main

import (
	"fmt"

	"github.com/go-zen-chu/delve-debug-demo/count"
)

func main() {
	c := count.Parallel1000count()

	fmt.Println("1000 にならんわけないやろww")
	fmt.Printf("-> %d\n", c)
	// should not be happen
	if c != 1000 {
		fmt.Println("なんでや!!")
	}
}
