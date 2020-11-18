package main

import (
	"unsafe"
	"testing"
	"functions"
)

var test_haystack []int = []int{15, 2, 1, 123, 272, 72, 727, 17, 8222, 1, 17, 25, 4}
var test_needle int = 72

type args struct {
	haystack []int
	needle int
}

type test struct {
	name string
	args args
	want int
}

func Test_LinearSearch(t *testing.T) {
	test_success := test{
		name: "success",
		args: args {
			haystack: test_haystack,
			needle: test_needle,
		},
		want: 5,
	}
	t.Run(test_success.name, func(t *testing.T) {
		res := functions.Linear(test_success.args.haystack, test_success.args.needle)
		if res != test_success.want {
			t.Errorf("LinearSearch() = %v, want %v", res, test_success.want)
		}
	})
}
func Test_BinarySearch(t *testing.T) {
	sorted_haystack := []int{1, 1, 2, 4, 15, 17, 17, 25, 72, 123, 272, 727, 8222}
	test_success := test{
		name: "success",
		args: args {
			haystack: sorted_haystack,
			needle: test_needle,
		},
		want: 8,
	}
	t.Run(test_success.name, func(t *testing.T) {
		left_index, right_index := 0, int(unsafe.Sizeof(test_success.args.haystack))
		res := functions.Binary(test_success.args.haystack, left_index, right_index, test_success.args.needle)
		if res != test_success.want {
			t.Errorf("BinarySearch() = %v, want %v", res, test_success.want)
		}
	})
}
func Test_ExponentialSearch(t *testing.T) {
	sorted_haystack := []int{1, 1, 2, 4, 15, 17, 17, 25, 72, 123, 272, 727, 8222}
	test_success := test{
		name: "success",
		args: args {
			haystack: sorted_haystack,
			needle: test_needle,
		},
		want: 8,
	}
	t.Run(test_success.name, func(t *testing.T) {
		res := functions.Exponential(test_success.args.haystack, test_success.args.needle)
		if res != test_success.want {
			t.Errorf("ExponentialSearch() = %v, want %v", res, test_success.want)
		}
	})
}
func Test_FibonacciSearch(t *testing.T) {
	test_success := test{
		name: "success",
		args: args {
			haystack: test_haystack,
			needle: test_needle,
		},
		want: 5,
	}
	t.Run(test_success.name, func(t *testing.T) {
		res := functions.Fibonacci(test_success.args.haystack, test_success.args.needle)
		if res != test_success.want {
			t.Errorf("FibonacciSearch() = %v, want %v", res, test_success.want)
		}
	})
}
func Test_InterpolationSearch(t *testing.T) {
	sorted_haystack := []int{1, 1, 2, 4, 15, 17, 17, 25, 72, 123, 272, 727, 8222}
	test_success := test{
		name: "success",
		args: args {
			haystack: sorted_haystack,
			needle: test_needle,
		},
		want: 8,
	}
	t.Run(test_success.name, func(t *testing.T) {
		res := functions.Interpolation(test_success.args.haystack, test_success.args.needle)
		if res != test_success.want {
			t.Errorf("InterpolationSearch() = %v, want %v", res, test_success.want)
		}
	})
}
func Test_JumpSearch(t *testing.T) {
	sorted_haystack := []int{1, 1, 2, 4, 15, 17, 17, 25, 72, 123, 272, 727, 8222}
	test_success := test{
		name: "success",
		args: args {
			haystack: sorted_haystack,
			needle: test_needle,
		},
		want: 8,
	}
	t.Run(test_success.name, func(t *testing.T) {
		res := functions.Jump(test_success.args.haystack, test_success.args.needle)
		if res != test_success.want {
			t.Errorf("JumpSearch() = %v, want %v", res, test_success.want)
		}
	})
}