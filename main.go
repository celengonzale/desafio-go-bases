package main

import (
	"fmt"
	"github.com/bootcamp-go/desafio-go-bases/internal/tickets"
)

func main() {
	ticketsByCountry, err := tickets.GetTotalTickets("Japan")
	fmt.Println(ticketsByCountry, err)
	person, err := tickets.GetCountByPeriod("0:22")
	fmt.Println(person, err)
	average, err := tickets.AverageDestination("Brazil")
	fmt.Println(average, err)

}
