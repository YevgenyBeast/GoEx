package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func funcFizzBuzz(n int) {
	for i := 0; i < n; i++ {
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

func main() {
	funcFizzBuzz(100)
	fmt.Println("\nPlease enter \"q\" to exit")
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		if scanner.Text() == "q" {
			break
		}
	}
	if scanner.Err() != nil {
		log.Fatal()
	}
}
