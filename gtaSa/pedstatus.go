package gtaSa

type PedStatus byte

const (
	AirOrWater       PedStatus = 0
	InVehicle                  = 1
	EnteringInterior           = 2
	InFloor                    = 3
)

func (s PedStatus) String() string {
	switch s {
	case AirOrWater:
		return "AirOrWater"
	case InVehicle:
		return "InVehicle"
	case EnteringInterior:
		return "EnteringInterior"
	case InFloor:
		return "InFloor"
	}
	return "unknown"
}
