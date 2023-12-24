package utils

import (
	"fmt"
	"math/rand"
)

func Say(a string, b string) {
	fmt.Println("===go mod test case===")
	fmt.Printf("you say:%s and %s", a, b)
}

func SayOne(a float64) {
	fmt.Println("===go mod test case===")
	fmt.Printf("you only say (double):%v\n", a)
	//panic("there is a bug:\n ")  //the bug has been fixed after v3.0.0
}

func SayRandomly() {
	content := rand.Intn(100)
	fmt.Println(content)
	//if content < 5 {
	//	panic("an error happens because of the unstable version")
	//} //the unstable bug has been fixed after v4.0.0-pre
}

func SayThree(a, b, c string) {
	fmt.Println("===go mod test case===")
	fmt.Printf("you say three words:%s,%s and %s\n", a, b, c)
}
