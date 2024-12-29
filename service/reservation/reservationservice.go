package reservation

import (
	resmodell "motorparking/model/reservation"
	resrepo "motorparking/repository/reservation"
)

func CreateReservation(reservation resmodell.Reservation) {
	resrepo.CreateReservation(reservation)
}
