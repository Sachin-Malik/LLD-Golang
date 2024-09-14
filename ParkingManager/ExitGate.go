package ParkingManager

import (
	"fmt"
	"time"
)


type ExitGate struct {
   manager *Manager;
}

func (eg *ExitGate) ExitVehicle(ticket *Ticket) string  {
	// find the vehicle, 
	eg.manager.RemoveVehicle(ticket)
    invoice:=eg.GenerateInvoice(ticket);
	// removed it from VehicleDB
	eg.manager.FreeSpace++;
	eg.manager.Occupied--;
	eg.manager.History = append(eg.manager.History, ticket);
	return invoice;
}

func (eg *ExitGate) GenerateInvoice(ticket *Ticket) string {
	enteredAt := ticket.entryTime;
	timeIn := time.Since(enteredAt)/(1000*60);
	invoice := timeIn*17;
    return fmt.Sprintf("You were here for %d and you have to pay %d",timeIn,invoice)

}

func NewExitGate(pm *Manager) *ExitGate{
  return &ExitGate{
	 manager:pm,
  }
}


func  GetTicketForVehicle( vehicle *Vehicle){
	// get vehicleNumber from vehicle
	// get ticket for this vehicle
   
}

