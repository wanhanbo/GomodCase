package utils

import "fmt"

func Say(a string, b string) {
	fmt.Println("===go mod test case===")
	fmt.Printf("you say:%s and %s", a, b)
}

func SayOne(a string) {
	fmt.Println("===go mod test case===")
	fmt.Printf("you only say:%s ", a)
}
