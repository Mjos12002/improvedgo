package car

import (
	"fmt"
	carmodel "motorparking/model/car"
	carservice "motorparking/service/car"
)

func CreateCar() {
	car := carmodel.CarModel{
		Type:        "SUV",
		NumberPlate: "RAC499A",
	}
	if car.CheckType(car.Type) {
		fmt.Println("Type is existing")
	}
	carservice.CreateCar(car)
}
