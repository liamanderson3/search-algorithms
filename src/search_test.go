package main

import (
	"functions"
	"testing"
)

var testhaystack []int = []int{15, 2, 1, 123, 272, 72, 727, 17, 8222, 1, 17, 25, 4}
var sortedHaystack []int = []int{1, 1, 2, 4, 15, 17, 17, 25, 72, 123, 272, 727, 8222}
var testNeedle int = 72

type args struct {
	haystack []int
	needle   int
}

type test struct {
	name string
	args args
	want int
}

func Test_LinearSearch(t *testing.T) {
	testSuccess := test{
		name: "success",
		args: args{
			haystack: testhaystack,
			needle:   testNeedle,
		},
		want: 5,
	}
	t.Run(testSuccess.name, func(t *testing.T) {
		res := functions.Linear(testSuccess.args.haystack, testSuccess.args.needle)
		if res != testSuccess.want {
			t.Errorf("LinearSearch() = %v, want %v", res, testSuccess.want)
		}
	})
}
func Test_BinarySearch(t *testing.T) {
	testSuccess := test{
		name: "success",
		args: args{
			haystack: sortedHaystack,
			needle:   testNeedle,
		},
		want: 8,
	}
	t.Run(testSuccess.name, func(t *testing.T) {
		leftIndex, rightIndex := 0, len(testSuccess.args.haystack)
		res := functions.Binary(testSuccess.args.haystack, leftIndex, rightIndex, testSuccess.args.needle)
		if res != testSuccess.want {
			t.Errorf("BinarySearch() = %v, want %v", res, testSuccess.want)
		}
	})
}
func Test_ExponentialSearch(t *testing.T) {
	testSuccess := test{
		name: "success",
		args: args{
			haystack: sortedHaystack,
			needle:   testNeedle,
		},
		want: 8,
	}
	t.Run(testSuccess.name, func(t *testing.T) {
		res := functions.Exponential(testSuccess.args.haystack, testSuccess.args.needle)
		if res != testSuccess.want {
			t.Errorf("ExponentialSearch() = %v, want %v", res, testSuccess.want)
		}
	})
}
func Test_FibonacciSearch(t *testing.T) {
	testSuccess := test{
		name: "success",
		args: args{
			haystack: sortedHaystack,
			needle:   testNeedle,
		},
		want: 8,
	}
	t.Run(testSuccess.name, func(t *testing.T) {
		res := functions.Fibonacci(testSuccess.args.haystack, testSuccess.args.needle)
		if res != testSuccess.want {
			t.Errorf("FibonacciSearch() = %v, want %v", res, testSuccess.want)
		}
	})
}
func Test_InterpolationSearch(t *testing.T) {
	testSuccess := test{
		name: "success",
		args: args{
			haystack: sortedHaystack,
			needle:   testNeedle,
		},
		want: 8,
	}
	t.Run(testSuccess.name, func(t *testing.T) {
		res := functions.Interpolation(testSuccess.args.haystack, testSuccess.args.needle)
		if res != testSuccess.want {
			t.Errorf("InterpolationSearch() = %v, want %v", res, testSuccess.want)
		}
	})
}
func Test_JumpSearch(t *testing.T) {
	testSuccess := test{
		name: "success",
		args: args{
			haystack: sortedHaystack,
			needle:   testNeedle,
		},
		want: 8,
	}
	t.Run(testSuccess.name, func(t *testing.T) {
		res := functions.Jump(testSuccess.args.haystack, testSuccess.args.needle)
		if res != testSuccess.want {
			t.Errorf("JumpSearch() = %v, want %v", res, testSuccess.want)
		}
	})
}
