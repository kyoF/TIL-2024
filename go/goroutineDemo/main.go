package main

import "fmt"

func main() {
	for i := 0; i < 3; i++ {
		go func(i int) {
			fmt.Println(i)
		}(i)
	}
}
