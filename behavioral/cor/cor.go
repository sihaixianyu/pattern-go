package cor

import "fmt"

type Department interface {
	execute(*Patient)
	setNext(next Department) Department
}

type Patient struct {
	name              string
	registrationDone  bool
	doctorCheckUpDone bool
	medicineDone      bool
	paymentDone       bool
}

type Reception struct {
	next Department
}

func (r *Reception) execute(p *Patient) {
	if p.registrationDone {
		fmt.Println("Patient registration has done.")
		r.next.execute(p)
		return
	}

	fmt.Println("Reception is registering patient...")
	p.registrationDone = true
	r.next.execute(p)
}

func (r *Reception) setNext(next Department) Department {
	r.next = next
	return next
}

type Doctor struct {
	next Department
}

func (d *Doctor) execute(p *Patient) {
	if p.doctorCheckUpDone {
		fmt.Println("Doctor checkup has done.")
		d.next.execute(p)
		return
	}

	fmt.Println("Doctor is checking the patient...")
	p.doctorCheckUpDone = true
	d.next.execute(p)
}

func (d *Doctor) setNext(next Department) Department {
	d.next = next
	return next
}

type Medical struct {
	next Department
}

func (m *Medical) execute(p *Patient) {
	if p.medicineDone {
		fmt.Println("Medical has given medicine to patient.")
		m.next.execute(p)
		return
	}
	fmt.Println("Medical is giving medicine to patient...")
	p.medicineDone = true
	m.next.execute(p)
}

func (m *Medical) setNext(next Department) Department {
	m.next = next
	return next
}

type Cashier struct {
	next Department
}

func (c *Cashier) execute(p *Patient) {
	if p.paymentDone {
		fmt.Println("Cashier has finished the payment.")
	}
	fmt.Println("Cashier is taking cash from patient...")
}

func (c *Cashier) setNext(next Department) Department {
	c.next = next
	return next
}
