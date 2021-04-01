package main

import "fmt"

func main() {
	myarr := [...]int{1, 2, 3, 4}
	for _, x := range myarr {
		fmt.Printf("%d\n", x)
	}
	fmt.Println("Hello GitHub")
}
