/**
 * Description：
 * FileName：add_test.go.go
 * Author：CJiaの用心
 * Create：2025/9/5 00:47:37
 * Remark：
 */

package slice

import (
	"fmt"
	"github.com/carefuly/careful-echo/internal/errs"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAdd(t *testing.T) {
	testCases := []struct {
		name      string
		slice     []int
		addVal    int
		index     int
		wantSlice []int
		wantErr   error
	}{
		{
			name:      "index 0",
			slice:     []int{1, 2, 3},
			addVal:    0,
			index:     0,
			wantSlice: []int{0, 1, 2, 3},
		},
		{
			name:      "index middle",
			slice:     []int{123, 124, 125},
			addVal:    233,
			index:     1,
			wantSlice: []int{123, 233, 124, 125},
		},
		{
			name:    "index out of range",
			slice:   []int{123, 1000},
			addVal:  233,
			index:   12,
			wantErr: errs.NewErrIndexOutOfRange(2, 12),
		},
		{
			name:    "index less than 0",
			slice:   []int{123, 1000},
			addVal:  233,
			index:   -1,
			wantErr: errs.NewErrIndexOutOfRange(2, -1),
		},
		{
			name:      "index last",
			slice:     []int{11, 12, 13, 14, 15, 16},
			addVal:    233,
			index:     5,
			wantSlice: []int{11, 12, 13, 14, 15, 233, 16},
		},
		{
			name:      "append on last",
			slice:     []int{23, 24, 25, 26, 27, 28},
			addVal:    233,
			index:     6,
			wantSlice: []int{23, 24, 25, 26, 27, 28, 233},
		},
		{
			name:    "index out of range",
			slice:   []int{78, 79, 80, 81, 82, 83},
			addVal:  233,
			index:   7,
			wantErr: errs.NewErrIndexOutOfRange(6, 7),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			fmt.Println("oldVal >>> ", tc.name, tc.slice)
			res, err := Add(tc.slice, tc.addVal, tc.index)
			fmt.Println("newRes >>> ", tc.name, res)
			assert.Equal(t, tc.wantErr, err)
			if err != nil {
				println("err >>> ", tc.name, err.Error())
				return
			}
			assert.Equal(t, tc.wantSlice, res)
		})
	}
}
