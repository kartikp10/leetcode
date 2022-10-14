package leetcode

import (
    "reflect"
	"testing"
)

func Test_deleteMiddle(t *testing.T) {
    type args struct {
        head *ListNode
    }
    tests := []struct {
        name string
        args args
        want *ListNode
    }{
        // TODO: Add test cases.
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            if got := deleteMiddle(tt.args.head); !reflect.DeepEqual(got, tt.want) {
                t.Errorf("deleteMiddle() = %v, want %v", got, tt.want)
            }
        })
    }
}

func Test_getNodeNum(t *testing.T) {
    type args struct {
        n    int
        head *ListNode
    }
    tests := []struct {
        name string
        args args
        want *ListNode
    }{
        // TODO: Add test cases.
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            if got := getNodeNum(tt.args.n, tt.args.head); !reflect.DeepEqual(got, tt.want) {
                t.Errorf("getNodeNum() = %v, want %v", got, tt.want)
            }
        })
    }
}

func Test_getLen(t *testing.T) {
	type args struct {
		l *ListNode
	}
	tests := []struct {
		name string
		args args
		want int
	}{
        {
        	name: "ex 1",
        	args: args{createList([]int{1,2,3,4,5})},
        	want: 5,
        },
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getLen(tt.args.l); got != tt.want {
				t.Errorf("getLen() = %v, want %v", got, tt.want)
			}
		})
	}
}

// func Test_deleteNode(t *testing.T) {
//     type args struct {
//         node *ListNode
//     }
//     tests := []struct {
//         name string
//         args args
//     }{
//         // TODO: Add test cases.
//     }
//     for _, tt := range tests {
//         t.Run(tt.name, func(t *testing.T) {
//             deleteNode(tt.args.node)
//         })
//     }
// }
