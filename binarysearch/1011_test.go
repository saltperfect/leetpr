package main

import (
	"fmt"
	"math/rand"
	"testing"
)

func Test_shipWithinDays(t *testing.T) {
	type args struct {
		weights []int
		days    int
	}
	tests := []struct {
		args args
		want int
	}{
		{
			args: args{
				weights: []int{1,2,3,4,5,6,7,8,9,10},
				days: 5,
			},
			want: 15,			
		},
		{
			args: args{
				weights: []int{3,2,2,4,1,4},
				days: 3,
			},
			want: 6,
		},
		{
			args: args{
				weights: []int{1,2,3,1,1},
				days: 4,
			},
			want: 3,
		},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			if got := shipWithinDays(tt.args.weights, tt.args.days); got != tt.want {
				t.Errorf("shipWithinDays() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Benchmark_shipWithindays(b *testing.B) {
	for _, size := range []int {
		100, 200, 500,
	} {
		b.Run(fmt.Sprint(size), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				b.StopTimer()
				list := rand.Perm(size)
				b.StartTimer()
				shipWithinDays(list, rand.Intn(size))
			}
		})
	}
}
