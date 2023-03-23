package romantointeger

import "testing"

func TestRomanToInt(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name  string
		args  args
		wantN int
	}{
		{"test1", args{"III"}, 3},
		{"test2", args{"IV"}, 4},
		{"test3", args{"IX"}, 9},
		{"test4", args{"LVIII"}, 58},
		{"test5", args{"MCMXCIV"}, 1994},
		{"test6", args{"MMMCMXCIX"}, 3999},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotN := RomanToInt(tt.args.s); gotN != tt.wantN {
				t.Errorf("RomanToInt() = %v, want %v", gotN, tt.wantN)
			}
		})
	}
}

// BenchmarkRomanToInt-8   	100000000	        11.04 ns/op	       0 B/op	       0 allocs/op
func BenchmarkRomanToInt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		RomanToInt("MMMCMXCIX")
	}
}
