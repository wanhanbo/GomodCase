package utils

import "fmt"

func Say(a string, b string) {
	fmt.Println("===go mod test case===")
	fmt.Printf("you say:%s and %s", a, b)
}

func SayOne(a float64) {
	fmt.Println("===go mod test case===")
	fmt.Printf("you only say (double):%v\n", a)
	//panic("there is a bug:\n ")  //the bug has been fixed
}
