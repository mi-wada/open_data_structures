package main

import "fmt"

func main() {
	s := make([]int, 0)
	s = append(s, 0)
	s = append(s, 1)

	fmt.Printf("s[1:1]: %v\n", s[1:1])
	fmt.Printf("s[0:1]: %v\n", s[0:1])
	fmt.Printf("s[0:2]: %v\n", s[0:2])
}
