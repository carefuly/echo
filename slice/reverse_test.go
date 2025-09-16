/**
 * Description：
 * FileName：reverse_test.go.go
 * Author：CJiaの用心
 * Create：2025/9/16 11:39:28
 * Remark：
 */

package slice

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestReverseInt(t *testing.T) {
	testCases := []struct {
		name string
		src  []int
		want []int
	}{
		{
			name: "normal test",
			src:  []int{1, 3, 5, 7},
			want: []int{7, 5, 3, 1},
		},
		{
			name: "length of src is 0",
			src:  []int{},
			want: []int{},
		},
		{
			name: "length of src is nil",
			src:  nil,
			want: []int{},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			res := Reverse[int](tc.src)
			assert.ElementsMatch(t, tc.want, res)
		})
	}
}

func TestReverseStruct(t *testing.T) {
	type testStruct struct {
		A int
		B []int
	}
	testCases := []struct {
		name string
		src  []testStruct
		want []testStruct
	}{
		{
			name: "normal test",
			src:  []testStruct{{5, []int{7, 8, 9}}, {3, []int{4, 5, 6}}, {1, []int{1, 2, 3}}},
			want: []testStruct{{1, []int{1, 2, 3}}, {3, []int{4, 5, 6}}, {5, []int{7, 8, 9}}},
		},
		{
			src:  []testStruct{},
			want: []testStruct{},
			name: "length of src is 0",
		},
		{
			src:  nil,
			want: []testStruct{},
			name: "length of src is nil",
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			res := Reverse[testStruct](tc.src)
			assert.ElementsMatch(t, tc.want, res)
			// assert.NotSame(t, tc.src, res)
		})
	}
}

func TestReverseSelfInt(t *testing.T) {
	testCases := []struct {
		name string
		src  []int
		want []int
	}{
		{
			name: "normal test",
			src:  []int{1, 3, 5, 7},
			want: []int{7, 5, 3, 1},
		},
		{
			name: "length of src is 0",
			src:  []int{},
			want: []int{},
		},
		{
			name: "length of src is nil",
			src:  nil,
			want: []int{},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ReverseSelf[int](tc.src)
			assert.ElementsMatch(t, tc.want, tc.src)
		})
	}
}

func TestReverseSelfStruct(t *testing.T) {
	type testStruct struct {
		A int
		B []int
	}
	testCases := []struct {
		name string
		src  []testStruct
		want []testStruct
	}{
		{
			name: "normal test",
			src:  []testStruct{{5, []int{7, 8, 9}}, {3, []int{4, 5, 6}}, {1, []int{1, 2, 3}}},
			want: []testStruct{{1, []int{1, 2, 3}}, {3, []int{4, 5, 6}}, {5, []int{7, 8, 9}}},
		},
		{
			name: "length of src is 0",
			src:  []testStruct{},
			want: []testStruct{},
		},
		{
			name: "length of src is nil",
			src:  nil,
			want: []testStruct{},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ReverseSelf[testStruct](tc.src)
			assert.ElementsMatch(t, tc.want, tc.src)
		})
	}
}
