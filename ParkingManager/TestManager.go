package ParkingManager

import (
	"fmt"
	"time"
)
func TestManager(){
	pm:= NewManager(69);
	entryGate := NewEntryGate(pm);
	exitGage := NewExitGate(pm);
	ticket1,error := entryGate.ParkVehicle(&Vehicle{VehicleNumber: "ABC123",VehicleType: "4wheeler"})

	if(error!=""){
		return;
	}
	fmt.Println(pm.GetInfo())
    time.Sleep(time.Second*1)
	invoice:= exitGage.ExitVehicle(ticket1);

	fmt.Println(invoice)
}