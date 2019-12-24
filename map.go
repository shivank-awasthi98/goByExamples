package main

import "fmt"

func main() {
	m := make(map[int]int)
	m[1] = 2
	m['2'] = 2
	fmt.Print(m)
}
