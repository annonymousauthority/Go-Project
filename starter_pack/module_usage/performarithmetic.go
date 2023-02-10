package main

import (
	"fmt"

	"math"

	"tryporpra.com/starter_pack/arithmetic"
)

func main() {
	fmt.Println("Performing computation")
	addition, multiplication, division, subtraction := arithmetic.PerformArithmetic(math.MaxInt32, 4)
	fmt.Println("The added sum is ", addition)
	fmt.Println("The subtracted differece is ", subtraction)
	fmt.Println("The divided total is ", division)
	fmt.Println("The multiplied total is ", multiplication)
}
