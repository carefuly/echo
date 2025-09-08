/**
 * Description：
 * FileName：contains_test.go.go
 * Author：CJiaの用心
 * Create：2025/9/9 00:14:22
 * Remark：
 */

package slice

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestContains(t *testing.T) {
	testCases := []struct {
		name string
		src  []int
		dst  int
		want bool
	}{
		{
			name: "dst exist",
			src:  []int{1, 2, 3, 4},
			dst:  4,
			want: true,
		},
		{
			name: "dst not exist",
			src:  []int{1, 2, 3, 4},
			dst:  6,
			want: false,
		},
		{
			name: "length of src is 0",
			src:  []int{},
			dst:  4,
			want: false,
		},
		{
			want: false,
			dst:  4,
			name: "src nil",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.want, Contains[int](tc.src, tc.dst))
		})
	}
}

func TestContainsFunc(t *testing.T) {
	testCases := []struct {
		name string
		src  []int
		dst  int
		want bool
	}{
		{
			name: "dst exist",
			src:  []int{1, 2, 3, 4},
			dst:  4,
			want: true,
		},
		{
			name: "dst not exist",
			src:  []int{1, 2, 3, 4},
			dst:  6,
			want: false,
		},
		{
			name: "length of src is 0",
			src:  []int{},
			dst:  4,
			want: false,
		},
		{
			want: false,
			dst:  4,
			name: "src nil",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.want, ContainsFunc[int](tc.src, func(src int) bool {
				return src == tc.dst
			}))
		})
	}
}

func TestContainsAny(t *testing.T) {
	testCases := []struct {
		name string
		src  []int
		dst  []int
		want bool
	}{
		{
			name: "exist two ele",
			src:  []int{1, 4, 6, 2, 6},
			dst:  []int{1, 6},
			want: true,
		},
		{
			want: false,
			src:  []int{1, 4, 6, 2, 6},
			dst:  []int{7, 0},
			name: "not exist the same",
		},
		{
			want: true,
			src:  []int{1, 1, 8},
			dst:  []int{1, 1},
			name: "exist two same ele",
		},
		{
			want: false,
			src:  []int{},
			dst:  []int{1},
			name: "length of src is 0",
		},
		{
			want: false,
			dst:  []int{1},
			name: "src nil",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.want, ContainsAny[int](tc.src, tc.dst))
		})
	}
}

func TestContainsAnyFunc(t *testing.T) {
	testCases := []struct {
		name string
		src  []int
		dst  []int
		want bool
	}{
		{
			name: "exist two ele",
			src:  []int{1, 4, 6, 2, 6},
			dst:  []int{1, 6},
			want: true,
		},
		{
			want: false,
			src:  []int{1, 4, 6, 2, 6},
			dst:  []int{7, 0},
			name: "not exist the same",
		},
		{
			want: true,
			src:  []int{1, 1, 8},
			dst:  []int{1, 1},
			name: "exist two same ele",
		},
		{
			want: false,
			src:  []int{},
			dst:  []int{1},
			name: "length of src is 0",
		},
		{
			want: false,
			dst:  []int{1},
			name: "src nil",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.want, ContainsAnyFunc[int](tc.src, tc.dst, func(src, dst int) bool {
				return src == dst
			}))
		})
	}
}

func TestContainsAll(t *testing.T) {
	tests := []struct {
		want bool
		src  []int
		dst  []int
		name string
	}{
		{
			want: true,
			src:  []int{1, 4, 6, 2, 6},
			dst:  []int{1, 4, 6, 2},
			name: "src exist one not in dst",
		},
		{
			want: false,
			src:  []int{1, 4, 6, 2, 6},
			dst:  []int{1, 4, 6, 2, 6, 7},
			name: "src not include the whole ele",
		},
		{
			want: false,
			src:  []int{},
			dst:  []int{1},
			name: "length of src is 0",
		},
		{
			want: true,
			src:  nil,
			dst:  []int{},
			name: "src nil dst empty",
		},
		{
			want: true,
			src:  nil,
			name: "src and dst nil",
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.want, ContainsAll[int](test.src, test.dst))
		})
	}
}

func TestContainsAllFunc(t *testing.T) {
	tests := []struct {
		want bool
		src  []int
		dst  []int
		name string
	}{
		{
			want: true,
			src:  []int{1, 4, 6, 2, 6},
			dst:  []int{1, 4, 6, 2},
			name: "src exist one not in dst",
		},
		{
			want: false,
			src:  []int{1, 4, 6, 2, 6},
			dst:  []int{1, 4, 6, 2, 6, 7},
			name: "src not include the whole ele",
		},
		{
			want: false,
			src:  []int{},
			dst:  []int{1},
			name: "length of src is 0",
		},
		{
			want: true,
			src:  nil,
			dst:  []int{},
			name: "src nil dst empty",
		},
		{
			want: true,
			src:  nil,
			name: "src and dst nil",
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.want, ContainsAllFunc[int](test.src, test.dst, func(src, dst int) bool {
				return src == dst
			}))
		})
	}
}
