package ParkingManager

import (
	"math/rand"
	"time"
)

func (pm *Manager) RemoveVehicle(ticket *Ticket) {
    vehicleNumber := ticket.vehicle.VehicleNumber
    var newList []*Vehicle

    if pm.Occupied > len(pm.VehicleParked) {
        pm.Occupied = len(pm.VehicleParked)
    }

    for i := 0; i < pm.Occupied; i++ {
        if pm.VehicleParked[i].VehicleNumber != vehicleNumber {
            newList = append(newList, pm.VehicleParked[i])
        }
    }

    pm.VehicleParked = newList
    pm.Occupied = len(pm.VehicleParked)
}
func (eg *EntryGate) CreateTicket(vehicle *Vehicle) (*Ticket){
	return &Ticket{
		vehicle: vehicle,
		id: rand.Intn(10000),
		entryTime: time.Now(),
	}
}

