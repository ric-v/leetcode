package rootequalssumofchildren

import "testing"

func Test_checkTree(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"test1", args{&TreeNode{Val: 10, Left: &TreeNode{Val: 5}, Right: &TreeNode{Val: 5}}}, true},
		{"test2", args{&TreeNode{Val: 10, Left: &TreeNode{Val: 5}, Right: &TreeNode{Val: 6}}}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RootEqualSumOfChildren(tt.args.root); got != tt.want {
				t.Errorf("checkTree() = %v, want %v", got, tt.want)
			}
		})
	}
}

// Benchmark_checkTree-8   	309439689	         3.615 ns/op	       0 B/op	       0 allocs/op
func Benchmark_checkTree(b *testing.B) {
	for i := 0; i < b.N; i++ {
		RootEqualSumOfChildren(&TreeNode{Val: 10, Left: &TreeNode{Val: 5}, Right: &TreeNode{Val: 5}})
	}
}
