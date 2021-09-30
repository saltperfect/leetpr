package greedy

import (
	"fmt"
	"testing"
)

func Test_balancedStringSplit(t *testing.T) {
	tests := []struct {
		s    string
		want int
	}{
		{
			s:    "RLRRLLRLRL",
			want: 4,
		},
		{
			s:    "RLLLLRRRLR",
			want: 3,
		},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			if got := balancedStringSplit(tt.s); got != tt.want {
				t.Errorf("balancedStringSplit() = %v, want %v", got, tt.want)
			}
		})
	}
}
