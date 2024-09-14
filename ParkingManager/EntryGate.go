package ParkingManager

type EntryGate struct {
	manager *Manager
}

func NewEntryGate(pm *Manager) *EntryGate {
	return &EntryGate{
		manager: pm,
	}
}

func (eg *EntryGate) ParkVehicle(vehicle *Vehicle) (ticket *Ticket, errorMessage string) {
	if eg.manager.FreeSpace == 0 {
		return nil, "No Space left"
	}
	ticket, errorString := eg.GetTicket(vehicle)
	if errorString != "" {
		return nil, "Error in ticket Generation"
	}
	eg.manager.FreeSpace--
	eg.manager.Occupied++
	eg.manager.TicketsDB = append(eg.manager.TicketsDB, ticket)
	// return result
	return ticket, ""

}

func (eg *EntryGate) GetTicket(vehicle *Vehicle) (*Ticket, string) {
	if eg.manager.FreeSpace == 0 {
		return &Ticket{}, "No Empty spots!!"
	}

	ticket := eg.CreateTicket(vehicle)
	return ticket, ""

}

// Other function that we can have in Entry Gate

// FindSpot
// sdf
