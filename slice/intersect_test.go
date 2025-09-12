/**
 * Description：
 * FileName：intersect_test.go.go
 * Author：CJiaの用心
 * Create：2025/9/12 16:21:08
 * Remark：
 */

package slice

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIntersectSet(t *testing.T) {
	testCases := []struct {
		name string
		src  []int
		dst  []int
		want []int
	}{
		{
			name: "normal test",
			src:  []int{1, 3, 5, 7},
			dst:  []int{1, 3, 5},
			want: []int{1, 3, 5},
		},
		{
			name: "length of src is 0",
			src:  []int{},
			dst:  []int{1, 3, 5, 7},
			want: []int{},
		},
		{
			name: "src nil",
			dst:  []int{1, 3, 5, 7},
			want: []int{},
		},
		{
			name: "exist the same ele in src",
			src:  []int{1, 3, 5, 5},
			dst:  []int{1, 3, 5},
			want: []int{1, 3, 5},
		},
		{
			name: "dst empty",
			src:  []int{1, 3, 5, 5},
			dst:  []int{},
			want: []int{},
		},
		{
			name: "dst nil",
			src:  []int{1, 3, 5, 5},
			dst:  []int{},
			want: []int{},
		},
		{
			name: "exist the same ele in src and dst",
			src:  []int{1, 1, 3, 5, 7},
			dst:  []int{1, 3, 5, 5},
			want: []int{1, 3, 5},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			res := IntersectSet[int](tc.src, tc.dst)
			assert.ElementsMatch(t, tc.want, res)
		})
	}
}

func TestIntersectSetFunc(t *testing.T) {
	testCases := []struct {
		name string
		src  []int
		dst  []int
		want []int
	}{
		{
			name: "normal test",
			src:  []int{1, 3, 5, 7},
			dst:  []int{1, 3, 5},
			want: []int{1, 3, 5},
		},
		{
			name: "length of src is 0",
			src:  []int{},
			dst:  []int{1, 3, 5, 7},
			want: []int{},
		},
		{
			name: "src nil",
			dst:  []int{1, 3, 5, 7},
			want: []int{},
		},
		{
			name: "exist the same ele in src",
			src:  []int{1, 3, 5, 5},
			dst:  []int{1, 3, 5},
			want: []int{1, 3, 5},
		},
		{
			name: "dst empty",
			src:  []int{1, 3, 5, 5},
			dst:  []int{},
			want: []int{},
		},
		{
			name: "dst nil",
			src:  []int{1, 3, 5, 5},
			dst:  []int{},
			want: []int{},
		},
		{
			name: "exist the same ele in src and dst",
			src:  []int{1, 1, 3, 5, 7},
			dst:  []int{1, 3, 5, 5},
			want: []int{1, 3, 5},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			res := IntersectSetFunc[int](tc.src, tc.dst, func(src, dst int) bool {
				return src == dst
			})
			assert.ElementsMatch(t, tc.want, res)
		})
	}
}
