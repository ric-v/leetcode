package runningsumof1darray

import (
	"reflect"
	"testing"
)

func Test_runningSum(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"test1", args{[]int{1, 2, 3, 4}}, []int{1, 3, 6, 10}},
		{"test2", args{[]int{1, 1, 1, 1, 1}}, []int{1, 2, 3, 4, 5}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := PrefixSum(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("runningSum() = %v, want %v", got, tt.want)
			}
		})
	}
}

// Benchmark_runningSum-8   	342454414	         3.403 ns/op	       0 B/op	       0 allocs/op
func Benchmark_runningSum(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PrefixSum([]int{1, 2, 3, 4})
	}
}
