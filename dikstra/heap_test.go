package dikstra

import "testing"

func Test_visiualHeap(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			name: "h1",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			visiualHeap()
		})
	}
}
