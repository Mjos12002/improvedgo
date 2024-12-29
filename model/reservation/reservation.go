package reservation

import (
	model "motorparking/model/car"
	"motorparking/model/customer"
	"motorparking/model/parking"
	"time"
)

type Reservation struct {
	Date       time.Time
	Customer   customer.CustomerModel
	Car        model.CarModel
	ParkingLot int
	Parking    parking.Parking
}
