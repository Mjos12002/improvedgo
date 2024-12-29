package customer

import (
	"fmt"
	customermodel "motorparking/model/customer"
)

func CreateCustomer(customer customermodel.CustomerModel) {
	fmt.Println(customer)
}
