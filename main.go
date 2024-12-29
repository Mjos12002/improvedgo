package main

import (
	car "motorparking/controller/car"
	customer "motorparking/controller/customer"
	parking "motorparking/controller/parking"
	reservation "motorparking/controller/reservation"
)

func main() {
	car.CreateCar()
	customer.CreateCustomer()
	parking.CreateParking()
	reservation.CreateReservation()

}
