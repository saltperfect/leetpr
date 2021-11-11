package recursion

import (
	"fmt"
	"testing"
)

func Test_wrapper(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		args args
		want int
	}{
		{
			args: args{
				input: "ABC",
			},
			want: 6,
		},
		{
			args: args{
				input: "ABCD",
			},
			want: 24,
		},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			if got := wrapper(tt.args.input); got != tt.want {
				t.Errorf("wrapper() = %v, want %v", got, tt.want)
			}
		})
	}
}
