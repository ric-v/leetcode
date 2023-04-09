package richestcustomerwealth

import "testing"

func Test_maximumWealth(t *testing.T) {
	type args struct {
		accounts [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"test1", args{[][]int{{1, 2, 3}, {3, 2, 1}}}, 6},
		{"test2", args{[][]int{{1, 5}, {7, 3}, {3, 5}}}, 10},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MaximumWealth(tt.args.accounts); got != tt.want {
				t.Errorf("maximumWealth() = %v, want %v", got, tt.want)
			}
		})
	}
}

// Benchmark_maximumWealth-8   	142014598	         7.939 ns/op	       0 B/op	       0 allocs/op
func Benchmark_maximumWealth(b *testing.B) {
	for i := 0; i < b.N; i++ {
		MaximumWealth([][]int{{1, 2, 3}, {3, 2, 1}})
	}
}
