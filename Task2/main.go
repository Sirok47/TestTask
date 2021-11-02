package main

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	MaxArrSize = 104
	MaxValue = 105
)

func main(){
	rand.Seed(time.Now().UnixNano())
	arrLen:=1+rand.Intn(MaxArrSize)
	var arr [MaxArrSize]int
	for i:=0;i<arrLen;i++{
		arr[i]=rand.Intn(MaxValue)+1
	}
	fmt.Println("Before modification: ",arr[:arrLen])
	res:=TransformArray(arr[:arrLen])
	fmt.Println("After modification: ",res)
	if arrLen%2==0 {
		CreateBonusArr(res)
	}
}