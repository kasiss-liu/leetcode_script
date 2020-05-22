package main

import "testing"

func containsDuplicate(nums []int) bool {
    max := 0
    for _,v := range nums {
        if max < v {
            max = v
        }
    }
    contains := make(map[int]bool,max+1)
    for i:=0;i < len(nums);i ++ {
       k := nums[i]
       if contains[k] {
           return true
       }else{
           contains[k] = true
       }
    }
    return false
}


func TestContains(t *testing.T) {
	a := []int{1,2,3,4,1}
	exp := true

	ret := containsDuplicate(a)
	if exp != ret {
		t.Error(exp,"not equal",ret)
	}
}