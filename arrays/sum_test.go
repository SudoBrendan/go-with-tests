package main

import (
	"reflect"
	"testing"
)

func assertSliceDeepEqual(t *testing.T, want []int, got []int) {
	t.Helper()
	if !reflect.DeepEqual(want, got) {
		t.Errorf("wanted %v but got %v", want, got)
	}
}

func benchmarkAllSliceFunction(b *testing.B, f func(...[]int) []int, slices ...[]int) {
	b.Helper()
	for i := 0; i < b.N; i++ {
		f(slices...)
	}
}

func TestSum(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5}
	got := Sum(nums)
	want := 15

	if want != got {
		t.Errorf("wanted '%d' but got '%d', given %v", want, got, nums)
	}
}

func BenchmarkSum(b *testing.B) {
	nums := []int{1, 2, 3, 4, 5}
	for i := 0; i < b.N; i++ {
		Sum(nums)
	}
}

func TestSumAll(t *testing.T) {
	got := SumAll([]int{1, 2}, []int{0, 9})
	want := []int{3, 9}
	assertSliceDeepEqual(t, want, got)
}

func BenchmarkSumAll(b *testing.B) {
	s1 := []int{1, 2}
	s2 := []int{3, 4}
	s3 := []int{5, 6}
	benchmarkAllSliceFunction(b, SumAll, s1, s2, s3)
}

func TestSumAllTails(t *testing.T) {

	t.Run("happy", func(t *testing.T) {
		got := SumAllTails([]int{1, 2, 3}, []int{0, 9})
		want := []int{5, 9}
		assertSliceDeepEqual(t, want, got)
	})

	t.Run("single element", func(t *testing.T) {
		got := SumAllTails([]int{1}, []int{0})
		want := []int{0, 0}
		assertSliceDeepEqual(t, want, got)
	})

	t.Run("empty slices", func(t *testing.T) {
		got := SumAllTails([]int{})
		want := []int{0}
		assertSliceDeepEqual(t, want, got)
	})
}

func BenchmarkSumAllTails(b *testing.B) {
	s1 := []int{1, 2}
	s2 := []int{3, 4}
	s3 := []int{5, 6}
	benchmarkAllSliceFunction(b, SumAllTails, s1, s2, s3)
}
