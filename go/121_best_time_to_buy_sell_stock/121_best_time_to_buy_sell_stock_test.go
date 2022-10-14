package leetcode

import "testing"

func Test_maxProfit(t *testing.T) {
	type args struct {
		prices []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
        {
        	name: "ex 1",
        	args: args{[]int{1,2,3,4,5}},
        	want: 4,
        },
        {
        	name: "ex 2",
        	args: args{[]int{7,6,4,3,1}},
        	want: 0,
        },
        {
        	name: "ex 3",
        	args: args{[]int{1,2,4}},
        	want: 3,
        },
        {
        	name: "ex 4",
        	args: args{[]int{4,1,3,7,8}},
        	want: 7,
        },
        {
        	name: "ex 5",
        	args: args{[]int{3,2,6,5,0,3}},
        	want: 4,
        },
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxProfit(tt.args.prices); got != tt.want {
				t.Errorf("maxProfit() = %v, want %v", got, tt.want)
			}
		})
	}
}
