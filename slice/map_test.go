/**
 * Description：
 * FileName：map_test.go.go
 * Author：CJiaの用心
 * Create：2025/9/16 13:42:56
 * Remark：
 */

package slice

import (
	"github.com/stretchr/testify/assert"
	"strconv"
	"testing"
)

func TestFilterMap(t *testing.T) {
	testCases := []struct {
		name string
		src  []int
		want []string
	}{
		{
			name: "src nil",
			want: []string{},
		},
		{
			name: "src empty",
			src:  []int{},
			want: []string{},
		},
		{
			name: "src has element",
			src:  []int{1, -2, 3},
			want: []string{"1", "3"},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			res := FilterMap(tc.src, func(idx int, src int) (string, bool) {
				return strconv.Itoa(src), src >= 0
			})
			assert.Equal(t, res, tc.want)
		})
	}
}

func TestMap(t *testing.T) {
	testCases := []struct {
		name string
		src  []int
		want []string
	}{
		{
			name: "src nil",
			want: []string{},
		},
		{
			name: "src empty",
			src:  []int{},
			want: []string{},
		},
		{
			name: "src has element",
			src:  []int{1, 2, 3},
			want: []string{"1", "2", "3"},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			res := Map(tc.src, func(idx int, src int) string {
				return strconv.Itoa(src)
			})
			assert.Equal(t, res, tc.want)
		})
	}
}

func TestToMap(t *testing.T) {
	t.Run("integer-string to map[int]string", func(t *testing.T) {
		elements := []string{"1", "2", "3", "4", "5"}
		resMap := ToMap(elements, func(str string) int {
			num, _ := strconv.Atoi(str)
			return num
		})
		epeMap := map[int]string{
			1: "1",
			2: "2",
			3: "3",
			4: "4",
			5: "5",
		}
		assert.Equal(t, epeMap, resMap)
	})
	t.Run("struct<string, string, int> to map[string]struct<string, string, int>", func(t *testing.T) {
		type eleType struct {
			A string
			B string
			C int
		}
		elements := []eleType{
			{
				A: "a",
				B: "b",
				C: 1,
			},
			{
				A: "c",
				B: "d",
				C: 2,
			},
		}
		resMap := ToMap(elements, func(ele eleType) string {
			return ele.A
		})
		epeMap := map[string]eleType{
			"a": {
				A: "a",
				B: "b",
				C: 1,
			},
			"c": {
				A: "c",
				B: "d",
				C: 2,
			},
		}
		assert.Equal(t, epeMap, resMap)
	})

	t.Run("struct<string, string, int> to map[string]struct<string, string, int>, 重复的key", func(t *testing.T) {
		type eleType struct {
			A string
			B string
			C int
		}
		elements := []eleType{
			{
				A: "a",
				B: "b",
				C: 1,
			},
			{
				A: "c",
				B: "d",
				C: 2,
			},
		}
		resMap := ToMap(elements, func(ele eleType) string {
			return ele.A
		})
		epeMap := map[string]eleType{
			"a": {
				A: "a",
				B: "b",
				C: 1,
			},
			"c": {
				A: "c",
				B: "d",
				C: 2,
			},
		}
		assert.Equal(t, epeMap, resMap)
	})

	t.Run("传入nil slice,返回空map", func(t *testing.T) {
		var elements []string = nil
		resMap := ToMap(elements, func(str string) int {
			num, _ := strconv.Atoi(str)
			return num
		})
		epeMap := make(map[int]string)
		assert.Equal(t, epeMap, resMap)
	})
}

func TestToMapV(t *testing.T) {
	t.Run("integer-string to map[int]int", func(t *testing.T) {
		elements := []string{"1", "2", "3", "4", "5"}
		resMap := ToMapV(elements, func(str string) (int, int) {
			num, _ := strconv.Atoi(str)
			return num, num
		})
		epeMap := map[int]int{
			1: 1,
			2: 2,
			3: 3,
			4: 4,
			5: 5,
		}
		assert.Equal(t, epeMap, resMap)
	})
	t.Run("struct<string, string, int> to map[string]struct<string, string, int>", func(t *testing.T) {
		type eleType struct {
			A string
			B string
			C int
		}
		elements := []eleType{
			{
				A: "a",
				B: "b",
				C: 1,
			},
			{
				A: "c",
				B: "d",
				C: 2,
			},
		}
		resMap := ToMapV(elements, func(ele eleType) (string, eleType) {
			return ele.A, ele
		})
		epeMap := map[string]eleType{
			"a": {
				A: "a",
				B: "b",
				C: 1,
			},
			"c": {
				A: "c",
				B: "d",
				C: 2,
			},
		}
		assert.Equal(t, epeMap, resMap)
	})

	t.Run("struct<string, string, int> to map[string]struct<string, string, int>, 重复的key", func(t *testing.T) {
		type eleType struct {
			A string
			B string
			C int
		}
		elements := []eleType{
			{
				A: "a",
				B: "b",
				C: 1,
			},
			{
				A: "c",
				B: "d",
				C: 2,
			},
			{
				A: "a",
				B: "d",
				C: 3,
			},
		}
		resMap := ToMapV(elements, func(ele eleType) (string, eleType) {
			return ele.A, ele
		})
		epeMap := map[string]eleType{
			"a": {
				A: "a",
				B: "d",
				C: 3,
			},
			"c": {
				A: "c",
				B: "d",
				C: 2,
			},
		}
		assert.Equal(t, epeMap, resMap)
	})

	t.Run("传入nil slice,返回空map", func(t *testing.T) {
		var elements []string = nil
		resMap := ToMapV(elements, func(str string) (int, int) {
			num, _ := strconv.Atoi(str)
			return num, num
		})
		epeMap := make(map[int]int)
		assert.Equal(t, epeMap, resMap)
	})
}
