package parking

import (
	"motorparking/model/parking"
	park "motorparking/repository/parking"
)

func CreateParking(parking parking.Parking) {
	park.CreateParking(parking)
}
