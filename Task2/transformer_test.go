package main

import (
	"reflect"
	"testing"
)

func TestTransformArray(t *testing.T) {
	res := TransformArray([]int{17,18,5,4,6,1})
	if !reflect.DeepEqual(res,[]int{18,6,6,6,1,-1}) {
	t.Errorf("TransformArray unexpected result: %v, expected %v", res, []int{18,6,6,6,1,-1})
	}
	res = TransformArray([]int{400})
	if !reflect.DeepEqual(res,[]int{-1}) {
		t.Errorf("TransformArray unexpected result: %v, expected %v", res, []int{-1})
	}
}
