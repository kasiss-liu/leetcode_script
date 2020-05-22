package main

import "testing"

func maxProfit(prices []int) int {
    profit := 0
    ln := len(prices)
    for i:= 1;i<ln;i++ {
        pro := prices[i] - prices[i-1]
        if  pro > 0 {
            profit += pro
        }
    }
    return profit
}


func TestMaxProfit(t *testing.T) {
	prices := []int{1,2,3,4,5}
	exp := 4

	ret := maxProfit(prices)
	if exp != ret {
		t.Error(exp,"not equal",ret)
	}
}