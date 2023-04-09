package diagonaltraverse

import (
	"reflect"
	"testing"
)

func Test_findDiagonalOrder(t *testing.T) {
	type args struct {
		mat [][]int
	}
	tests := []struct {
		name       string
		args       args
		wantResult []int
	}{
		{"test1", args{[][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}}, []int{1, 2, 4, 7, 5, 3, 6, 8, 9}},
		{"test2", args{[][]int{{1, 2, 3, 4, 5}}}, []int{1, 2, 3, 4, 5}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotResult := DiagonalOrder(tt.args.mat); !reflect.DeepEqual(gotResult, tt.wantResult) {
				t.Errorf("findDiagonalOrder() = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}

// Benchmark_findDiagonalOrder-8   	21130629	        58.40 ns/op	      80 B/op	       1 allocs/op
func Benchmark_findDiagonalOrder(b *testing.B) {
	for i := 0; i < b.N; i++ {
		DiagonalOrder([][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}})
	}
}
