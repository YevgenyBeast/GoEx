package main

import "fmt"

func main() {
	for i := 0; i < 100; i++ {
		if (i+1)%3 == 0 && (i+1)%5 == 0 {
			fmt.Println("FizzBuzz")
		} else if (i+1)%3 == 0 {
			fmt.Println("Fizz")
		} else if (i+1)%5 == 0 {
			fmt.Println("Buzz")
		} else {
			fmt.Println(i + 1)
		}
	}
}
