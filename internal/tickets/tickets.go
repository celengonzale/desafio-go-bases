package tickets

import (
	"encoding/csv"
	"errors"
	"log"
	"os"
	"strconv"
	"strings"
)

type Ticket struct {
	ID          string
	FullName    string
	Email       string
	Destination string
	Hour        string
	Price       string
}

var ALL_TICKETS []Ticket

func createTicketsList(data [][]string) []Ticket {
	var allTickets []Ticket
	for _, ticketSlc := range data {
		var itemTicket Ticket
		for i2, ticketFile := range ticketSlc {
			switch i2 {
			case 0:
				itemTicket.ID = ticketFile
			case 1:
				itemTicket.FullName = ticketFile
			case 2:
				itemTicket.Email = ticketFile
			case 3:
				itemTicket.Destination = ticketFile
			case 4:
				itemTicket.Hour = ticketFile
			case 5:
				itemTicket.Price = ticketFile
			}
		}
		allTickets = append(allTickets, itemTicket)

	}
	ALL_TICKETS = allTickets
	return ALL_TICKETS

}
func destinationByTickets(destination string, allTickets []Ticket) (int, error) {
	var count int
	for _, item := range allTickets {
		if destination == item.Destination {
			count++
		}
	}
	return count, nil
}

func GetTotalTickets(destination string) (int, error) {
	file, err := os.Open("tickets.csv")

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	csvReader := csv.NewReader(file)
	data, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}
	allTickets := createTicketsList(data)
	countTickets, err := destinationByTickets(destination, allTickets)
	if err != nil {
		log.Fatal(err)
	}
	return countTickets, nil
}

const (
	earlyMorningMin = 0.00
	earlyMorningMax = 6.00
	morningMin      = 7.00
	morningMax      = 12.00
	afternoonMin    = 13.00
	afternoonMax    = 19.00
	nightMin        = 20.00
	nightMax        = 23.00
)

func convertToFloat(str string) float64 {
	stringWithComma := strings.Replace(str, ":", ".", 1)
	resultFloat, err := strconv.ParseFloat(stringWithComma, 8)
	if err != nil {
		log.Fatal(err)
	}
	return resultFloat
}
func GetCountByPeriod(time string) (int, error) {
	parseTime := convertToFloat(time)
	switch {
	case parseTime >= earlyMorningMin && parseTime <= earlyMorningMax:
		return countByPeriod(earlyMorningMin, earlyMorningMax)
	case parseTime >= morningMin && parseTime <= morningMax:
		return countByPeriod(morningMin, morningMax)
	case parseTime >= afternoonMin && parseTime <= afternoonMax:
		return countByPeriod(afternoonMin, afternoonMax)
	case parseTime >= nightMin && parseTime <= nightMax:
		return countByPeriod(nightMin, nightMax)
	default:
		return 0, errors.New("error: Nonexistent time range")
	}
}

func countByPeriod(minimum, maximum float64) (int, error) {
	var count int
	for _, file := range ALL_TICKETS {
		hourTime := convertToFloat(file.Hour)
		if hourTime >= minimum && hourTime <= maximum {
			count++
		}

	}
	return count, nil
}

func AverageDestination(destination string) (float64, error) {
	totalTickets := len(ALL_TICKETS)
	var average float64
	personsByDestination, err := GetTotalTickets(destination)
	if err != nil {
		log.Fatal(err)
	}
	average = float64(personsByDestination) / float64(totalTickets)
	return average, nil

}
