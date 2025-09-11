/**
 * Description：
 * FileName：diff_test.go.go
 * Author：CJiaの用心
 * Create：2025/9/11 23:40:50
 * Remark：
 */

package slice

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDiffSet(t *testing.T) {
	testCases := []struct {
		name string
		src  []int
		dst  []int
		want []int
	}{
		{
			name: "diff 1",
			src:  []int{1, 3, 5, 7},
			dst:  []int{1, 3, 5},
			want: []int{7},
		},
		{
			name: "src less than dst",
			src:  []int{1, 3, 5},
			dst:  []int{1, 3, 5, 7},
			want: []int{},
		},
		{
			name: "diff deduplicate",
			src:  []int{1, 3, 5, 7, 7},
			dst:  []int{1, 3, 5},
			want: []int{7},
		},
		{
			name: "dst duplicate ele",
			src:  []int{1, 1, 3, 5, 7},
			dst:  []int{1, 3, 5, 5},
			want: []int{7},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			res := DiffSet[int](tc.src, tc.dst)
			assert.ElementsMatch(t, tc.want, res)
		})
	}
}

func TestDiffSetFunc(t *testing.T) {
	testCases := []struct {
		name string
		src  []int
		dst  []int
		want []int
	}{
		{
			name: "diff 1",
			want: []int{7},
			src:  []int{1, 3, 5, 7},
			dst:  []int{1, 3, 5},
		},
		{
			name: "src less than dst",
			src:  []int{1, 3, 5},
			dst:  []int{1, 3, 5, 7},
			want: []int{},
		},
		{
			name: "diff deduplicate",
			src:  []int{1, 3, 5, 7, 7},
			dst:  []int{1, 3, 5},
			want: []int{7},
		},
		{
			name: "dst duplicate ele",
			src:  []int{1, 1, 3, 5, 7},
			dst:  []int{1, 3, 5, 5},
			want: []int{7},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			res := DiffSetFunc[int](tc.src, tc.dst, func(src, dst int) bool {
				return src == dst
			})
			assert.ElementsMatch(t, tc.want, res)
		})
	}
}
