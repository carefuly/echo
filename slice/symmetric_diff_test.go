/**
 * Description：
 * FileName：symmetric_diff_test.go.go
 * Author：CJiaの用心
 * Create：2025/9/16 11:56:59
 * Remark：
 */

package slice

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSymmetricDiffSet(t *testing.T) {
	testCases := []struct {
		name string
		src  []int
		dst  []int
		want []int
	}{
		{
			name: "no inter",
			src:  []int{1, 2, 3},
			dst:  []int{4, 5, 6},
			want: []int{1, 2, 3, 4, 5, 6},
		},
		{
			name: "part inter",
			src:  []int{1, 2, 3},
			dst:  []int{3, 4, 5},
			want: []int{1, 2, 4, 5},
		},
		{
			name: "src contain dst",
			src:  []int{1, 2, 3},
			dst:  []int{2, 3},
			want: []int{1},
		},
		{
			name: "dst contain src",
			src:  []int{4},
			dst:  []int{4, 5, 6},
			want: []int{5, 6},
		},
		{
			name: "equal",
			src:  []int{1, 2, 3},
			dst:  []int{1, 2, 3},
			want: []int{},
		},
		{
			name: "dst empty",
			src:  []int{1, 2, 3},
			dst:  []int{},
			want: []int{1, 2, 3},
		},
		{
			name: "src empty",
			src:  []int{},
			dst:  []int{4, 5, 6},
			want: []int{4, 5, 6},
		},
		{
			name: "all empty",
			src:  []int{},
			dst:  []int{},
			want: []int{},
		},
		{
			name: "src nil",
			src:  nil,
			dst:  []int{4, 5, 6},
			want: []int{4, 5, 6},
		},
		{
			name: "dst nil",
			src:  []int{4, 5, 6},
			dst:  nil,
			want: []int{4, 5, 6},
		},
		{
			name: "both nil",
			src:  nil,
			dst:  nil,
			want: []int{},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			res := SymmetricDiffSet[int](tc.src, tc.dst)
			assert.ElementsMatch(t, tc.want, res)
		})
	}
}

func TestSymmetricDiffSetFunc(t *testing.T) {
	testCases := []struct {
		name string
		src  []int
		dst  []int
		want []int
	}{
		{
			name: "no inter",
			src:  []int{1, 2, 3},
			dst:  []int{4, 5, 6},
			want: []int{1, 2, 3, 4, 5, 6},
		},
		{
			name: "part inter",
			src:  []int{1, 2, 3},
			dst:  []int{3, 4, 5},
			want: []int{1, 2, 4, 5},
		},
		{
			name: "src contain dst",
			src:  []int{1, 2, 3},
			dst:  []int{2, 3},
			want: []int{1},
		},
		{
			name: "dst contain src",
			src:  []int{4},
			dst:  []int{4, 5, 6},
			want: []int{5, 6},
		},
		{
			name: "equal",
			src:  []int{1, 2, 3},
			dst:  []int{1, 2, 3},
			want: []int{},
		},
		{
			name: "dst empty",
			src:  []int{1, 2, 3},
			dst:  []int{},
			want: []int{1, 2, 3},
		},
		{
			name: "src empty",
			src:  []int{},
			dst:  []int{4, 5, 6},
			want: []int{4, 5, 6},
		},
		{
			name: "all empty",
			src:  []int{},
			dst:  []int{},
			want: []int{},
		},
		{
			name: "src nil",
			src:  nil,
			dst:  []int{4, 5, 6},
			want: []int{4, 5, 6},
		},
		{
			name: "dst nil",
			src:  []int{4, 5, 6},
			dst:  nil,
			want: []int{4, 5, 6},
		},
		{
			name: "both nil",
			src:  nil,
			dst:  nil,
			want: []int{},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			res := SymmetricDiffSetFunc[int](tc.src, tc.dst, func(src, dst int) bool {
				return src == dst
			})
			assert.ElementsMatch(t, tc.want, res)
		})
	}
}
