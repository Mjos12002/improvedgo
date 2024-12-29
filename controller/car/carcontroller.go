package car

import (
	carmodel "motorparking/model/car"
	carservice "motorparking/service/car"
)

func CreateCar() {
	car := carmodel.CarModel{
		Type:        "SUV",
		NumberPlate: "RAC499A",
	}
	carservice.CreateCar(car)
}
