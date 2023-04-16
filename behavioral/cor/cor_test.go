package cor

import "testing"

func TestCor(t *testing.T) {
	cashier := &Cashier{}
	medical := &Medical{}
	doctor := &Doctor{}
	reception := &Reception{}

	reception.setNext(doctor).setNext(medical).setNext(cashier)

	patient := &Patient{
		name: "sihaixianyu",
	}
	reception.execute(patient)
}
