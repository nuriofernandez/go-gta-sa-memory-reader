package gtaSa

type Vehicle struct {
	gtaSa GtaSa
}

func (gta GtaSa) GetVehicle() *Vehicle {
	return &Vehicle{gtaSa: gta}
}

func (ped Vehicle) GetLocation() Location {
	gtaApi := ped.gtaSa.GtaApi

	return Location{
		X: gtaApi.GetPedVehiclePositionX(),
		Y: gtaApi.GetPedVehiclePositionY(),
		Z: gtaApi.GetPedVehiclePositionZ(),
	}
}
