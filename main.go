package main

import (
	"fmt"
	"github.com/xXNurioXx/go-gta-sa-memory-reader/gtaSa"
)

func main() {
	gtaSa := new(gtaSa.GtaSa)
	gtaSa.Hook()

	money := gtaSa.GetMoney()
	pedLocation := gtaSa.GetPed().GetLocation()
	vehicleLocation := gtaSa.GetVehicle().GetLocation()
	pedStatus := gtaSa.GetPed().GetStatus()

	fmt.Println("Ped status ->", pedStatus)
	fmt.Println("User money is ->", money)
	fmt.Printf("Ped location is '%f', '%f', '%f'\n", pedLocation.X, pedLocation.Y, pedLocation.Z)
	fmt.Printf("Vehicle location is '%f', '%f', '%f'\n", vehicleLocation.X, vehicleLocation.Y, vehicleLocation.Z)
}
