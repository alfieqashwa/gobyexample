package main

import (
	"fmt"
)

type EmployeeOperation interface {
	SalaryCalculator
	LeaveCalculator
}

type SalaryCalculator interface {
	DisplaySalary()
}

type LeaveCalculator interface {
	CalculateLeavesLeft() int
}

type Employee struct {
	fname, lname                           string
	basicPay, pf, totalLeaves, leavesTaken int
}

func (e Employee) DisplaySalary() {
	fmt.Printf("%s %s has salary $%d", e.fname, e.lname, (e.basicPay + e.pf))
}

func (e Employee) CalculateLeavesLeft() int {
	return e.totalLeaves - e.leavesTaken
}

func main() {
	e := Employee{
		fname:       "Alfie",
		lname:       "Qashwa",
		basicPay:    7000,
		pf:          3000,
		totalLeaves: 30,
		leavesTaken: 5,
	}
	var empOp EmployeeOperation = e
	empOp.DisplaySalary()
	fmt.Println("\nLeaves left =", empOp.CalculateLeavesLeft())
}
