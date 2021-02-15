package main

import (
	"fmt"

	"github.com/YevgenyBeast/magazine"
)

func main() {
	subscriber := magazine.Subscriber{Name: "Aman Singh"}
	subscriber.Address.Street = "123 Oak St"
	subscriber.Address.City = "Omaha"
	subscriber.Address.State = "NE"
	subscriber.Address.PostalCode = "68111"
	fmt.Println("Subscriber Name:", subscriber.Name)
	fmt.Println("Street:", subscriber.Street)
	fmt.Println("City:", subscriber.City)
	fmt.Println("State:", subscriber.State)
	fmt.Println("Postal Code:", subscriber.PostalCode)

	employee := magazine.Employee{Name: "Joy Carr"}
	address := magazine.Address{Street: "456 Elm St", City: "Portland", State: "OR", PostalCode: "97222"}
	employee.HomeAddress = address
	fmt.Println("Employee Name:", employee.Name)
	fmt.Println("Street:", employee.HomeAddress.Street)
	fmt.Println("City:", employee.HomeAddress.City)
	fmt.Println("State:", employee.HomeAddress.State)
	fmt.Println("Postal Code:", employee.HomeAddress.PostalCode)
}
