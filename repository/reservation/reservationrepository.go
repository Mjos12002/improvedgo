package reservation

import (
	"fmt"
	reservationmodel "motorparking/model/reservation"
)

func CreateReservation(reservation reservationmodel.Reservation) {
	fmt.Println(reservation)
}
