package pivotindex

import "testing"

func Test_findMiddleIndex(t *testing.T) {
	type args struct {
		num []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"test1", args{[]int{1, 7, 3, 6, 5, 6}}, 3},
		{"test2", args{[]int{1, 2, 3}}, -1},
		{"test3", args{[]int{2, 1, -1}}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := PivotIndex(tt.args.num); got != tt.want {
				t.Errorf("findMiddleIndex() = %v, want %v", got, tt.want)
			}
		})
	}
}

// Benchmark_findMiddleIndex-8   	140607044	         7.544 ns/op	       0 B/op	       0 allocs/op
func Benchmark_findMiddleIndex(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PivotIndex([]int{1, 7, 3, 6, 5, 6})
	}
}
