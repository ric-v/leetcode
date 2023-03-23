package fizzbuzz

import (
	"reflect"
	"testing"
)

func TestFizzBuzz(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{"Test 1", args{3}, []string{"1", "2", "Fizz"}},
		{"Test 2", args{5}, []string{"1", "2", "Fizz", "4", "Buzz"}},
		{"Test 3", args{15}, []string{"1", "2", "Fizz", "4", "Buzz", "Fizz", "7", "8", "Fizz", "Buzz", "11", "Fizz", "13", "14", "FizzBuzz"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FizzBuzz(tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FizzBuzz() = %v, want %v", got, tt.want)
			}
		})
	}
}

// BenchmarkFizzBuzz-8   	 2403470	       496.2 ns/op	     504 B/op	       6 allocs/op
func BenchmarkFizzBuzz(b *testing.B) {
	for i := 0; i < b.N; i++ {
		FizzBuzz(15)
	}
}
