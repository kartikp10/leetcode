package leetcode

import "testing"

func Test_breakPalindrome(t *testing.T) {
	type args struct {
		palindrome string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
        {
        	name: "ex 1",
        	args: args{"aabbbbaa"},
        	want: "aaabbbaa",
        },
        {
        	name: "ex 2",
        	args: args{"a"},
        	want: "",
        },
        {
        	name: "ex 3",
        	args: args{"aaaa"},
        	want: "aaab",
        },
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := breakPalindrome(tt.args.palindrome); got != tt.want {
				t.Errorf("breakPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}
