package main

import "fmt"

func main() {

	const a int = 5
	fmt.Println(a)

	//iota

	const (
		first = iota
		second
		third
	)
	fmt.Println(first, second, third)

}
