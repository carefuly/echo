/**
 * Description：
 * FileName：aggregate_test.go.go
 * Author：CJiaの用心
 * Create：2025/9/9 00:06:03
 * Remark：
 */

package slice

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type Integer int

func TestMax(t *testing.T) {
	testCases := []struct {
		name  string
		input []Integer
		want  Integer
	}{
		{
			name:  "value",
			input: []Integer{1},
			want:  1,
		},
		{
			name:  "values",
			input: []Integer{2, 3, 1},
			want:  3,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			res := Max[Integer](tc.input)
			assert.Equal(t, tc.want, res)
		})
	}
}

func TestMin(t *testing.T) {
	testCases := []struct {
		name  string
		input []Integer
		want  Integer
	}{
		{
			name:  "value",
			input: []Integer{3},
			want:  3,
		},
		{
			name:  "values",
			input: []Integer{3, 1, 2},
			want:  1,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			res := Min[Integer](tc.input)
			assert.Equal(t, tc.want, res)
		})
	}
}

func TestSum(t *testing.T) {
	testCases := []struct {
		name  string
		input []Integer
		want  Integer
	}{
		{
			name: "nil",
		},
		{
			name:  "empty",
			input: []Integer{},
		},
		{
			name:  "value",
			input: []Integer{1},
			want:  1,
		},
		{
			name:  "values",
			input: []Integer{1, 2, 3},
			want:  6,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			res := Sum[Integer](tc.input)
			assert.Equal(t, tc.want, res)
		})
	}
}
