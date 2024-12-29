package parking

import (
	parkingmodel "motorparking/model/parking"
	parkingservice "motorparking/service/parking"
)

func CreateParking() {
	parking := parkingmodel.Parking{
		ID:   1234,
		Name: "Rubangura Parking",
	}
	parkingservice.CreateParking(parking)
}
