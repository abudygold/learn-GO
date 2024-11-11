package main

import (
	"fmt"
	"math/rand"
)

func init() {
	fmt.Println("This is from the init function")
}

func main() {
	for i := 0; i < 100; i++ {
		if i == 10 {
			break
		}

		x := rand.Intn(10)
		y := rand.Intn(10)

		fmt.Printf("iteration %v \t x and y are %v and %v\t", i, x, y)

		/* if x < 4 && y < 4 {
			fmt.Println("both are less than four")
		} else if x > 6 && y > 6 {
			fmt.Println("both are greater than six")
		} else if x >= 4 && y <= 6 {
			fmt.Println("x is from 4 to 6 inclusive of both numbers")
		} else if y != 5 {
			fmt.Println("y is not 5")
		} else {
			fmt.Println("none of the previous were met")
		} */

		switch {
		case x < 4 && y < 4:
			fmt.Println("both are less than four")
		case x > 6 && y > 6:
			fmt.Println("both are greater than six")
		case x >= 4 && y <= 6:
			fmt.Println("x is from 4 to 6 inclusive of both numbers")
		case y != 5:
			fmt.Println("y is not 5")
		default:
			fmt.Println("none of the previous were met")
		}

		/* x1 := rand.Intn(250)

		fmt.Printf("value of x is %v\t", x1)

		if x1 <= 100 {
			fmt.Println("less than 100")
		} else if x1 >= 101 && x1 <= 200 {
			fmt.Println("between 101 and 200")
		} else if x1 >= 201 && x1 <= 250 {
			fmt.Println("between 201 and 300")
		} else {
			fmt.Println("this was more than 250")
		}

		switch {
		case x1 <= 100:
			fmt.Println("less than 100")
		case x1 >= 101 && x1 <= 200:
			fmt.Println("between 101 and 200")
		case x1 > 201 && x1 <= 250:
			fmt.Println("between 201 and 300")
		default:
			fmt.Println("this was more than 250")
		}

		fmt.Println(rand.Intn(3))
		fmt.Println(rand.Intn(3))
		fmt.Println(rand.Intn(3)) */
	}
}
