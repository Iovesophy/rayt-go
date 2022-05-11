package main

import (
	"fmt"
	"math/rand"
	"rayt-go/pkg/utils"
)

func main() {
	fmt.Println(utils.Pow2(0.1))
	fmt.Println(utils.Pow3(0.1))
	fmt.Println(utils.Pow4(0.1))
	fmt.Println(utils.Pow5(0.1))
	for i := 0; i < 3; i++ {
		fmt.Println(rand.Float64())
	}
	fmt.Println(utils.Clamp(150, 100, 200)) // 150
}
