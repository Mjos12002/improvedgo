package reservation

import (
	"fmt"
	carmodel "motorparking/model/car"
	customermodel "motorparking/model/customer"
	parkingmodel "motorparking/model/parking"
	reservationmodel "motorparking/model/reservation"
	reservationservice "motorparking/service/reservation"
	"time"
)

func CreateReservation() {
	car := carmodel.CarModel{
		Type:        "SUV",
		NumberPlate: "RAC399D",
	}
	parking := parkingmodel.Parking{
		ID:   1233,
		Name: "Rubangura Parking",
	}
	customer := customermodel.CustomerModel{
		Name:        "Joseph",
		PhoneNumber: "25088833",
		Email:       "email@gmail.com",
	}
	reservation := reservationmodel.Reservation{
		Date:       time.Now(),
		Customer:   customer,
		Car:        car,
		ParkingLot: 2,
		Parking:    parking,
	}
	fmt.Println(reservation.CalculateAge(reservation.Date))
	reservationservice.CreateReservation(reservation)
}
