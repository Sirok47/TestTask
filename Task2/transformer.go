package main

func TransformArray(arr []int) []int {
	for i:=range arr{
		if i==len(arr)-1 {arr[i]=-1;break}
		max:=arr[i+1]
		for _,v:=range arr[i+1:]{
			if v>max {max=v}
		}
		arr[i]=max
	}
	return arr
}
