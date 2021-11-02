package main

import (
	"fmt"
	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
	"os"
	"strconv"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}

func main(){
	envsize,_:= os.LookupEnv("ARRAY_SIZE")
	arrSize,err:=strconv.Atoi(envsize)
	if err != nil {
		log.Fatal("Non-int array size")
	}
	var arr = make([]*Person, arrSize)
	for i:=range arr{
		arr[i]=&Person{}
		fmt.Println("Enter name, age, educational institution, course and average mark:")
		_,err:=fmt.Scan(&arr[i].Name,&arr[i].Age,&arr[i].EducationalInstitution,&arr[i].Course,&arr[i].AverageMark)
		if err != nil {
			log.Fatal("Wrong argument type")
		}
	}
	for i,v:=range arr{
		v.incrementCourse()
		fmt.Printf("%v: Name: %v\n   Age: %v\n   Educating in: %v\n   Course: %v\n   Average mark: %v\n",i+1,v.Name,v.Age,v.EducationalInstitution,v.Course,v.AverageMark)
		reversedName(v)
	}
}