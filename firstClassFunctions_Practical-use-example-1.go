// Practical use of first class functions

// Till now we have defined what first class functions are and we have seen a few contrived examples to learn how they work.
// Now lets write a concrete program which shows the practical use of first class functions.

// We will create a program which filters a slice of students based on some criteria.
package main

import "fmt"

type student struct {
	firstname, lastname, grade, country string
}

func filter(s []student, f func(student) bool) []student {
	var r []student
	for _, v := range s {
		if f(v) == true {
			r = append(r, v)
		}
	}
	return r
}

func main() {
	s1 := student{
		firstname: "Qashwa",
		lastname:  "Zhafira",
		grade:     "A",
		country:   "Kazakhtan",
	}
	s2 := student{
		firstname: "Cello",
		lastname:  "el-Zhafran",
		grade:     "B",
		country:   "Iran",
	}
	s := []student{s1, s2}
	f := filter(s, func(s student) bool {
		if s.grade == "B" {
			return true
		}
		return false
	})
	fmt.Println(f)

	c := filter(s, func(s student) bool {
		if s.country == "Kazakhtan" {
			return true
		}
		return false
	})
	fmt.Printf("\n%v\n", c)
}
