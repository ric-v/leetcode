package validparanthesis

import (
	"testing"
)

func TestIsValid(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"Test 1", args{"()"}, true},
		{"Test 2", args{"()[]{}"}, true},
		{"Test 3", args{"(]"}, false},
		{"Test 4", args{"([)]"}, false},
		{"Test 5", args{"{[]}"}, true},
		{"Test 6", args{"{"}, false},
		{"Test 7", args{"}"}, false},
		{"Test large", args{"(({}()[[]]))"}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsValid(tt.args.s); got != tt.want {
				t.Errorf("IsValid() = %v, want %v", got, tt.want)
			}
		})
	}
}

// BenchmarkIsValid-8   	14139670	        82.97 ns/op	      24 B/op	       2 allocs/op
func BenchmarkIsValid(b *testing.B) {
	for i := 0; i < b.N; i++ {
		IsValid("(({}()[[]]))")
	}
}

func Test_matchingPair(t *testing.T) {
	type args struct {
		char rune
	}
	tests := []struct {
		name string
		args args
		want rune
	}{
		{"Test 1", args{')'}, '('},
		{"Test 2", args{']'}, '['},
		{"Test 3", args{'}'}, '{'},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := matchingPair(tt.args.char); got != tt.want {
				t.Errorf("matchingPair() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isOpenBracket(t *testing.T) {
	type args struct {
		char rune
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"Test 1", args{'('}, true},
		{"Test 2", args{'['}, true},
		{"Test 3", args{'{'}, true},
		{"Test 4", args{')'}, false},
		{"Test 5", args{']'}, false},
		{"Test 6", args{'}'}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isOpenBracket(tt.args.char); got != tt.want {
				t.Errorf("isOpenBracket() = %v, want %v", got, tt.want)
			}
		})
	}
}
