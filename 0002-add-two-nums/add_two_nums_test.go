package addtwonums

import (
	"reflect"
	"testing"
)

func TestAddTwoNumbers(t *testing.T) {
	type args struct {
		l1 *ListNode
		l2 *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{"test1", args{&ListNode{2, &ListNode{4, &ListNode{3, nil}}}, &ListNode{5, &ListNode{6, &ListNode{4, nil}}}}, &ListNode{7, &ListNode{0, &ListNode{8, nil}}}},
		{"test2", args{&ListNode{0, nil}, &ListNode{0, nil}}, &ListNode{0, nil}},
		{"test3", args{&ListNode{9, &ListNode{9, &ListNode{9, &ListNode{9, &ListNode{9, &ListNode{9, &ListNode{9, nil}}}}}}}, &ListNode{9, &ListNode{9, &ListNode{9, &ListNode{9, nil}}}}}, &ListNode{8, &ListNode{9, &ListNode{9, &ListNode{9, &ListNode{0, &ListNode{0, &ListNode{0, &ListNode{1, nil}}}}}}}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := AddTwoNumbers(tt.args.l1, tt.args.l2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AddTwoNumbers() = %v, want %v", got, tt.want)
			}
		})
	}
}

// BenchmarkAddTwoNumbers-8   	11910528	        99.35 ns/op	      48 B/op	       3 allocs/op
func BenchmarkAddTwoNumbers(b *testing.B) {
	for i := 0; i < b.N; i++ {
		AddTwoNumbers(&ListNode{2, &ListNode{4, &ListNode{3, nil}}}, &ListNode{5, &ListNode{6, &ListNode{4, nil}}})
	}
}
