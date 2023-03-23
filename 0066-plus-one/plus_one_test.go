package plusone

import (
	"reflect"
	"testing"
)

func TestPlusOne(t *testing.T) {
	type args struct {
		digits []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"Test 1", args{[]int{1, 2, 3}}, []int{1, 2, 4}},
		{"Test 2", args{[]int{4, 3, 2, 1}}, []int{4, 3, 2, 2}},
		{"Test 3", args{[]int{0}}, []int{1}},
		{"Test 4", args{[]int{9}}, []int{1, 0}},
		{"Test 5", args{[]int{8, 9, 9, 9}}, []int{9, 0, 0, 0}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := PlusOne(tt.args.digits); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PlusOne() = %v, want %v", got, tt.want)
			}
		})
	}
}
