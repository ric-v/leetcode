package middleofthelinkedlist

import (
	"reflect"
	"testing"
)

func Test_middleNode(t *testing.T) {
	type args struct {
		head *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{"test1", args{&ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, nil}}}}}}, &ListNode{3, &ListNode{4, &ListNode{5, nil}}}},
		{"test2", args{&ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, &ListNode{6, nil}}}}}}}, &ListNode{4, &ListNode{5, &ListNode{6, nil}}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MiddleNode(tt.args.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("middleNode() = %v, want %v", got, tt.want)
			}
		})
	}
}

// Benchmark_middleNode-8   	169061060	         6.480 ns/op	       0 B/op	       0 allocs/op
func Benchmark_middleNode(b *testing.B) {
	for i := 0; i < b.N; i++ {
		MiddleNode(&ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, nil}}}}})
	}
}
