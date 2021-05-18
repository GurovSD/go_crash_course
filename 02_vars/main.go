package main

import "fmt"

func main() {
	var name = "Brad"
	var age int32 = 37
	const isCool = true

	fmt.Println(name, age, isCool)
	fmt.Printf("%T\n", age)
}
