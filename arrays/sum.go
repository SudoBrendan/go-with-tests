package main

func Sum(nums []int) (sum int) {
	for _, val := range nums {
		sum += val
	}
	return
}

func SumAll(slices ...[]int) []int {
	numResults := len(slices)
	sums := make([]int, numResults)
	for i, slice := range slices {
		sums[i] = Sum(slice)
	}
	return sums
}

func SumAllTails(slices ...[]int) []int {
	numResults := len(slices)
	sums := make([]int, numResults)
	for i, slice := range slices {
		if len(slice) <= 1 {
			sums[i] = 0
		} else {
			sums[i] = Sum(slice[1:])
		}
	}
	return sums
}
