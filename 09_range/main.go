package main

import "fmt"

func main() {

	ids := []int{33, 63, 45, 78, 33, 13, 4}

	// Loop trought ids
	for i, id := range ids {
		fmt.Printf("%d - ID: %d\n", i, id)
	}

	// Not using index
	for _, id := range ids {
		fmt.Printf("ID: %d\n", id)
	}

	// Declare map and add kv
	emails := map[string]string{"Bob": "bob@gmail.com", "Sharon": "sharon@gmail.com"}

	for k := range emails {
		fmt.Println("Name: " + k)
	}

	for _, v := range emails {
		fmt.Println("Name: " + v)
	}

}
