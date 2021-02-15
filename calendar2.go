package main

import (
	"fmt"
	"log"

	"github.com/YevgenyBeast/calendar"
)

func main() {
	date := calendar.Date{}
	err := date.SetYear(2021)
	if err != nil {
		log.Fatal(err)
	}

	err = date.SetMonth(02)
	if err != nil {
		log.Fatal(err)
	}

	err = date.SetDay(15)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(date)
	fmt.Println(date.Year())
	fmt.Println(date.Month())
	fmt.Println(date.Day())
}
