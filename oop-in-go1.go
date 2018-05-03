package main

import (
	"fmt"
)

type author struct {
	fname string
	lname string
	bio   string
}

func (a author) fullname() string {
	return fmt.Sprintf("%s %s", a.fname, a.lname)
}

type post struct {
	title   string
	content string
	author
}

func (p post) details() {
	fmt.Println("Title: ", p.title)
	fmt.Println("Content: ", p.content)
	fmt.Println("Author: ", p.fullname())
	fmt.Println("Bio: ", p.bio)
}

func main() {
	author1 := author{
		"Qashwa",
		"Zhafira",
		"Golang Enthusiast",
	}
	post1 := post{
		"Inheritance Go",
		"Go supports composition instead of inheritance",
		author1,
	}
	post1.details()
}
