package main 

import "testing"

func singleNumber(nums []int) int {
    m := make(map[int]int,0)
    nLen := len(nums)
    for i:=0;i<nLen;i++ {
        if _,ok := m[nums[i]];ok {
            m[nums[i]] ++
        }else{
            m[nums[i]] = 1
        }
    }
    for k,v := range m {
        if v == 1 {
            return k
        }
    }
    return 0
}

func TestSingleNumber(t *testing.T) {
	a := []int{0,1,0,1,0,1,99}
	exp := 99

	ret := singleNumber(a)
	if exp != ret {
		t.Error(exp,"not equal",ret)
	}
}