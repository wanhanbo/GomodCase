package main

import "github.com/wanhanbo/GomodCase/v4/utils"

func main() {
	utils.SayOne(float64(2))
	for i := 0; i < 10; i++ {
		utils.SayRandomly()
	}

}
