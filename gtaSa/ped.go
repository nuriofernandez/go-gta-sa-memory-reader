package gtaSa

type Ped struct {
	gtaSa GtaSa
}

func (gta GtaSa) GetPed() *Ped {
	return &Ped{gtaSa: gta}
}

func (ped Ped) GetLocation() Location {
	gtaApi := ped.gtaSa.GtaApi

	return Location{
		X: gtaApi.GetPedPositionX(),
		Y: gtaApi.GetPedPositionY(),
		Z: gtaApi.GetPedPositionZ(),
	}
}

func (ped Ped) GetStatus() PedStatus {
	return PedStatus(ped.gtaSa.GtaApi.GetPedStatus())
}
