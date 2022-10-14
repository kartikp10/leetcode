package leetcode

import "testing"

func Test_isPalindrome(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
        {
        	name: "example 1",
        	args: args{"This si ht"},
        	want: true,
        },
        {
        	name: "example 2",
            args: args{"not a palindrome"},
        	want: false,
        },
        {
        	name: "example 3",
            args: args{"0110"},
        	want: true,
        },
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPalindrome(tt.args.s); got != tt.want {
				t.Errorf("isPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}
