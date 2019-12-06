package basic

import "fmt"

func FizzBuzzWithSwitch(n int) {
	for index := 1; index < n; index++ {
		switch {
		case index%15 == 0:
			fmt.Println("FizzBuzz")
		case index%3 == 0:
			fmt.Println("Fizz")
		case index%5 == 0:
			fmt.Println("Buzz")
		default:
			fmt.Println(index)
		}
	}
}

func FizzBuzzWithIf(n int) {
	// Write your code here
	for i := 1; i <= n; i++ {
		if i%15 == 0 {
			fmt.Println("FizzBuzz")
		} else if i%3 == 0 {
			fmt.Println("Fizz")
		} else if i%5 == 0 {
			fmt.Println("Buzz")
		} else {
			fmt.Println(i)
		}
	}
}
