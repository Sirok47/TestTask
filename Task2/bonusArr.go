package main

import "fmt"

func CreateBonusArr(arr []int)  {
	var bonusArr [MaxArrSize]int
	copy(bonusArr[:],arr)
	for i:=0;i<MaxArrSize-len(arr);i++{
		bonusArr[i+len(arr)]=105-i
	}
	fmt.Println("Bonus array: ",bonusArr)
}
