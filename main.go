package main

import "fmt"

func main() {
	events := []Event{}
	events = append(events, getTiketti(searhStrings)...)
	events = append(events, getLippu(searhStrings)...)
	printEvents(events)
}

func printEvents(events []Event) {
	fmt.Println()
	for _, event := range events {
		fmt.Println(event.name)
		fmt.Println("Price: " + event.price)
		fmt.Println("Date: " + event.date)
		fmt.Println("Location: " + event.location)
		fmt.Println()
	}
}
