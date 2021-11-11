package recursion

import "testing"

func Test_wrapperSubset(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "first",
			args: args{
				input: "ABC",
			},
			want: 8,	
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := wrapperSubset(tt.args.input); got != tt.want {
				t.Errorf("wrapperSubset() = %v, want %v", got, tt.want)
			}
		})
	}
}
