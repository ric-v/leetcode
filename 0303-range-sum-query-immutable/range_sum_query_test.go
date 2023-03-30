package rangesumqueryimmutable

import (
	"reflect"
	"testing"
)

func TestConstructor(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want NumArray
	}{
		{"test1", args{[]int{-2, 0, 3, -5, 2, -1}}, NumArray{[]int{-2, -2, 1, -4, -2, -3}}},
		{"test2", args{[]int{1, 2, 3, 4, 5}}, NumArray{[]int{1, 3, 6, 10, 15}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Constructor(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Constructor() = %v, want %v", got, tt.want)
			}
		})
	}
}

// BenchmarkConstructor-8   	11353176	        93.80 ns/op	     112 B/op	       3 allocs/op
func BenchmarkConstructor(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Constructor([]int{-2, 0, 3, -5, 2, -1})
	}
}

func TestNumArray_SumRange(t *testing.T) {
	type fields struct {
		Arr []int
	}
	type args struct {
		left  int
		right int
	}
	tests := []struct {
		name   string
		fields fields
		args   []args
		want   []int
	}{
		{"test1", fields{[]int{-2, 0, 3, -5, 2, -1}}, []args{{0, 2}, {2, 5}, {0, 5}}, []int{1, -1, -3}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			this := Constructor(tt.fields.Arr)
			for i, arg := range tt.args {
				if got := this.SumRange(arg.left, arg.right); got != tt.want[i] {
					t.Errorf("SumRange() = %v, want %v", got, tt.want[i])
				}
			}
		})
	}
}
