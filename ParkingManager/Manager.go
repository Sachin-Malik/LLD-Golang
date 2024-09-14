package ParkingManager

import (
	"fmt"
	"time"
)

type Ticket struct {
	id int
	entryTime time.Time;
	vehicle *Vehicle
}

type Vehicle struct {
	VehicleNumber string;
	VehicleType string; // 2wheeler 4wheeler
}

type Manager struct {
   FreeSpace int
   Occupied int
   History  []*Ticket
   VehicleParked []*Vehicle
   TicketsDB []*Ticket
}

func NewManager(free int) *Manager {
	return &Manager{
		FreeSpace: free,
		Occupied: 0,
		TicketsDB: []*Ticket{},
		History: []*Ticket{},
		VehicleParked: []*Vehicle{},
	}
}

func (pm *Manager) GetInfo() string {
	formattedString := fmt.Sprintf("This is the current Info , %v",pm)
	return formattedString;
}





