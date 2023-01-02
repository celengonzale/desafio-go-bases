package tickets

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetTotalTickets(t *testing.T) {
	//Arrange
	destination := "Japan"
	//Act
	ticketsByDestination, _ := GetTotalTickets(destination)
	//Assert
	expectedValue := 14
	assert.Equal(t, expectedValue, ticketsByDestination)
}
