package main

import (
	"fmt"
)

type Person struct {
	Education
	Name string
	Age int
}

type Education struct {
	EducationalInstitution string
	Course int
	AverageMark int
}

func (e *Education) incrementCourse() {
	if e.Course>=5 {
		fmt.Println("Out of bound")
		return
	}
	e.Course++
}

func reversedName(p *Person) {
	var result string
	for _,v := range p.Name {
		result = string(v) + result
	}
	fmt.Println("Name reversed: ",result)
}
