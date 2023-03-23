package longestcommonprefix

import "testing"

func TestLongestCommonPrefix(t *testing.T) {
	type args struct {
		strs []string
	}
	tests := []struct {
		name       string
		args       args
		wantPrefix string
	}{
		{"Test 1", args{[]string{"flower", "flow", "flight"}}, "fl"},
		{"Test 2", args{[]string{"dog", "racecar", "car"}}, ""},
		{"Test 3", args{[]string{"dog", "dog", "dog"}}, "dog"},
		{"Test 4", args{[]string{"dog", "dog", "dog", "dog"}}, "dog"},
		{"Test 5", args{[]string{"a"}}, "a"},
		{"Test 6", args{[]string{"flower", "flower", "flower", "flower"}}, "flower"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotPrefix := LongestCommonPrefix(tt.args.strs); gotPrefix != tt.wantPrefix {
				t.Errorf("LongestCommonPrefix() = %v, want %v", gotPrefix, tt.wantPrefix)
			}
		})
	}
}

// BenchmarkLongestCommonPrefix-8   	30076696	        33.33 ns/op	       0 B/op	       0 allocs/op
func BenchmarkLongestCommonPrefix(b *testing.B) {
	for i := 0; i < b.N; i++ {
		LongestCommonPrefix([]string{"flower", "flow", "flight"})
	}
}
