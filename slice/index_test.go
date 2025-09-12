/**
 * Description：
 * FileName：index_test.go.go
 * Author：CJiaの用心
 * Create：2025/9/12 14:33:36
 * Remark：
 */

package slice

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIndex(t *testing.T) {
	testCases := []struct {
		name string
		src  []int
		dst  int
		want int
	}{
		{
			name: "first one",
			src:  []int{1, 1, 3, 5},
			dst:  1,
			want: 0,
		},
		{
			name: "the length of src is 0",
			src:  []int{},
			dst:  1,
			want: -1,
		},
		{
			name: "src nil",
			dst:  1,
			want: -1,
		},
		{
			name: "dst not exist",
			src:  []int{1, 4, 6},
			dst:  7,
			want: -1,
		},
		{
			name: "last one",
			src:  []int{1, 3, 4, 2, 0},
			dst:  0,
			want: 4,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			index := Index[int](tc.src, tc.dst)
			assert.Equal(t, tc.want, index)
		})
	}
}

func TestIndexFunc(t *testing.T) {
	testCases := []struct {
		name string
		src  []int
		dst  int
		want int
	}{
		{
			name: "first one",
			src:  []int{1, 1, 3, 5},
			dst:  1,
			want: 0,
		},
		{
			name: "the length of src is 0",
			src:  []int{},
			dst:  1,
			want: -1,
		},
		{
			name: "src nil",
			dst:  1,
			want: -1,
		},
		{
			name: "dst not exist",
			src:  []int{1, 4, 6},
			dst:  7,
			want: -1,
		},
		{
			name: "last one",
			src:  []int{1, 3, 4, 2, 0},
			dst:  0,
			want: 4,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.want, IndexFunc[int](tc.src, func(src int) bool {
				return src == tc.dst
			}))
		})
	}
}

func TestLastIndex(t *testing.T) {
	testCases := []struct {
		name string
		src  []int
		dst  int
		want int
	}{
		{
			name: "first one",
			src:  []int{1, 1, 3, 5},
			dst:  1,
			want: 1,
		},
		{
			name: "the length of src is 0",
			src:  []int{},
			dst:  1,
			want: -1,
		},
		{
			name: "src nil",
			dst:  1,
			want: -1,
		},
		{
			name: "dst not exist",
			src:  []int{1, 4, 6},
			dst:  7,
			want: -1,
		},
		{
			name: "last one",
			src:  []int{0, 1, 3, 4, 2, 0},
			dst:  0,
			want: 5,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			index := LastIndex[int](tc.src, tc.dst)
			assert.Equal(t, tc.want, index)
		})
	}
}

func TestLastIndexFunc(t *testing.T) {
	testCases := []struct {
		name string
		src  []int
		dst  int
		want int
	}{
		{
			name: "first one",
			src:  []int{1, 1, 3, 5},
			dst:  1,
			want: 1,
		},
		{
			name: "the length of src is 0",
			src:  []int{},
			dst:  1,
			want: -1,
		},
		{
			name: "src nil",
			dst:  1,
			want: -1,
		},
		{
			name: "dst not exist",
			src:  []int{1, 4, 6},
			dst:  7,
			want: -1,
		},
		{
			name: "last one",
			src:  []int{0, 1, 3, 4, 2, 0},
			dst:  0,
			want: 5,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.want, LastIndexFunc[int](tc.src, func(src int) bool {
				return src == tc.dst
			}))
		})
	}
}

func TestIndexAll(t *testing.T) {
	tests := []struct {
		name string
		src  []int
		dst  int
		want []int
	}{
		{
			name: "normal test",
			src:  []int{1, 1, 3, 5},
			dst:  1,
			want: []int{0, 1},
		},
		{
			name: "the length of src is 0",
			src:  []int{},
			dst:  1,
			want: []int{},
		},
		{
			name: "dst not exist",
			src:  []int{1, 4, 6},
			dst:  7,
			want: []int{},
		},
		{
			name: "normal test",
			src:  []int{0, 1, 3, 4, 2, 0},
			dst:  0,
			want: []int{0, 5},
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			res := IndexAll[int](tc.src, tc.dst)
			assert.ElementsMatch(t, tc.want, res)
		})
	}
}

func TestIndexAllFunc(t *testing.T) {
	tests := []struct {
		name string
		src  []int
		dst  int
		want []int
	}{
		{
			name: "normal test",
			src:  []int{1, 1, 3, 5},
			dst:  1,
			want: []int{0, 1},
		},
		{
			name: "the length of src is 0",
			src:  []int{},
			dst:  1,
			want: []int{},
		},
		{
			name: "dst not exist",
			src:  []int{1, 4, 6},
			dst:  7,
			want: []int{},
		},
		{
			name: "normal test",
			src:  []int{0, 1, 3, 4, 2, 0},
			dst:  0,
			want: []int{0, 5},
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			res := IndexAllFunc[int](tc.src, func(src int) bool {
				return src == tc.dst
			})
			assert.ElementsMatch(t, tc.want, res)
		})
	}
}
