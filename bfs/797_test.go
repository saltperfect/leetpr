package bfs

import (
	"fmt"
	"reflect"
	"testing"
)

func Test_allPathsSourceTarget(t *testing.T) {
	tests := []struct {
		graph [][]int
		want  [][]int
	}{
		{
			graph: [][]int{
				{1, 2},
				{3},
				{3},
				{},
			},
			want: [][]int{
				{0, 1, 3},
				{0, 2, 3},
			},
		},
		{
			graph: [][]int{
				{4, 3, 1},
				{3, 2, 4},
				{3},
				{4},
				{},
			},
			want: [][]int{
				{0, 4},
				{0, 3, 4},
				{0, 1, 3, 4},
				{0, 1, 2, 3, 4},
				{0, 1, 4},
			},
		},
		{
			graph: [][]int{
				{3, 1},
				{4, 6, 7, 2, 5},
				{4, 6, 3},
				{6, 4},
				{7, 6, 5},
				{6},
				{7},
				{},
			},
			want: [][]int{
				{0, 3, 6, 7},
				{0, 3, 4, 7},
				{0, 3, 4, 6, 7},
				{0, 3, 4, 5, 6, 7},
				{0, 1, 4, 7},
				{0, 1, 4, 6, 7},
				{0, 1, 4, 5, 6, 7},
				{0, 1, 6, 7},
				{0, 1, 7},
				{0, 1, 2, 4, 7},
				{0, 1, 2, 4, 6, 7},
				{0, 1, 2, 4, 5, 6, 7},
				{0, 1, 2, 6, 7},
				{0, 1, 2, 3, 6, 7},
				{0, 1, 2, 3, 4, 7},
				{0, 1, 2, 3, 4, 6, 7},
				{0, 1, 2, 3, 4, 5, 6, 7},
				{0, 1, 5, 6, 7},
			},
		},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			got := allPathsSourceTarget(tt.graph)
			count := 0
			for _, v := range got {
				for _, l := range tt.want {
					if reflect.DeepEqual(v, l) {
						count++
					}

				}
			}
			if len(got) != count {
				t.Errorf("allPathsSourceTarget() = %v, want %v", got, tt.want)
			}
		})
	}
}
