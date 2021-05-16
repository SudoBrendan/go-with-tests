package main

import "testing"

func TestSum(t *testing.T) {

	nums := []int{1, 2, 3, 4, 5}
	got := Sum(nums)
	want := 15
	if want != got {
		t.Errorf("wanted '%d' but got '%d', given %v", want, got, nums)
	}
}
