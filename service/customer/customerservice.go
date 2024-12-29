package customer

import (
	"motorparking/model/customer"
	cust "motorparking/repository/customer"
)

func CreateCustomer(customer customer.CustomerModel) {
	cust.CreateCustomer(customer)
}
