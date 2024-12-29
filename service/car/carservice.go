package car

import (
	model "motorparking/model/car"
	carrepo "motorparking/repository/car"
)

func CreateCar(car model.CarModel) {
	carrepo.CreateCar(car)
}
