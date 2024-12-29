package customer

import (
	customermodel "motorparking/model/customer"
	customerservice "motorparking/service/customer"
)

func CreateCustomer() {
	customer := customermodel.CustomerModel{
		Name:        "Joseph MANZI",
		PhoneNumber: "250788812144",
		Email:       "jmanzi@gmail.com",
	}
	customerservice.CreateCustomer(customer)
}
