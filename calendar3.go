package main

import (
	"fmt"
	"log"

	"github.com/YevgenyBeast/calendar"
)

func main() {
	event := calendar.Event{}
	err := event.SetTitle("Today")
	if err != nil {
		log.Fatal(err)
	}

	err = event.SetYear(2021)
	if err != nil {
		log.Fatal(err)
	}

	err = event.SetMonth(02)
	if err != nil {
		log.Fatal(err)
	}

	err = event.SetDay(20)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(event)
	fmt.Println(event.Title())
	fmt.Println(event.Date.Year())
	fmt.Println(event.Month())
	fmt.Println(event.Day())
}
