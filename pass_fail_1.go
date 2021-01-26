// pass_fail сообщает, сдал ли пользователь экзамнен с использованием пакетов
package main

import (
	"fmt"
	"keyboard"
	"log"
)

func main() {
	fmt.Print("Enter a grade: ")
	var status string
	grade, err := keyboard.GetFloat()
	if err != nil {
		log.Fatal(err)
	}
	if grade >= 60 {
		status = "Сдал"
	} else {
		status = "Не сдал"
	}
	fmt.Println("A grade of", grade, "is", status)
}
