/**
 * Description：
 * FileName：union_test.go.go
 * Author：CJiaの用心
 * Create：2025/9/16 12:01:25
 * Remark：
 */

package slice

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUnionSet(t *testing.T) {
	testCases := []struct {
		name string
		src  []int
		dst  []int
		want []int
	}{
		{
			name: "not empty",
			src:  []int{1, 2, 3},
			dst:  []int{4, 5, 6, 1},
			want: []int{1, 2, 3, 4, 5, 6},
		},
		{
			name: "src is empty",
			src:  []int{},
			dst:  []int{1, 3},
			want: []int{1, 3},
		},
		{

			name: "dst is empty",
			src:  []int{1, 3},
			dst:  []int{},
			want: []int{1, 3},
		},
		{
			name: "src and dst are empty",
			src:  []int{},
			dst:  []int{},
			want: []int{},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			res := UnionSet[int](tc.src, tc.dst)
			assert.ElementsMatch(t, tc.want, res)
		})
	}
}

func TestUnionSetFunc(t *testing.T) {
	testCases := []struct {
		name string
		src  []int
		dst  []int
		want []int
	}{
		{
			src:  []int{1, 2, 3},
			dst:  []int{4, 5, 6, 1},
			want: []int{1, 2, 3, 4, 5, 6},
			name: "not empty",
		},
		{
			src:  []int{},
			dst:  []int{1, 3},
			want: []int{1, 3},
			name: "src is empty",
		},
		{

			src:  []int{1, 3},
			dst:  []int{},
			want: []int{1, 3},
			name: "dst is empty",
		},
		{
			src:  []int{},
			dst:  []int{},
			want: []int{},
			name: "src and dst are empty",
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			res := UnionSetFunc[int](tc.src, tc.dst, func(src, dst int) bool {
				return src == dst
			})
			assert.ElementsMatch(t, tc.want, res)
		})
	}
}
