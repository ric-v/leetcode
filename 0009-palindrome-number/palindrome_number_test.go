package palindromenumber

import (
	"testing"
)

func TestPalindrome(t *testing.T) {
	type args struct {
		x int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"test1", args{121}, true},
		{"test2", args{-121}, false},
		{"test3", args{10}, false},
		{"test_large", args{12345678987654321}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsPalindrome(tt.args.x); got != tt.want {
				t.Errorf("Palindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}

// BenchmarkPalindrome-8   	57600091	        21.51 ns/op	       0 B/op	       0 allocs/op
func BenchmarkPalindrome(b *testing.B) {
	for i := 0; i < b.N; i++ {
		IsPalindrome(12345678987654321)
	}
}

func TestIsPalindrome(t *testing.T) {
	type args struct {
		x int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"test1", args{121}, true},
		{"test2", args{-121}, false},
		{"test3", args{10}, false},
		{"test_large", args{12345678987654321}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsPalindrome_Unoptimized(tt.args.x); got != tt.want {
				t.Errorf("IsPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}

// BenchmarkIsPalindrome-8   	 8347639	       149.9 ns/op	      80 B/op	       4 allocs/op
func BenchmarkIsPalindrome(b *testing.B) {
	for i := 0; i < b.N; i++ {
		IsPalindrome_Unoptimized(12345678987654321)
	}
}
