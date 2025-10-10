/**
 * Description：
 * FileName：string_test.go.go
 * Author：CJiaの用心
 * Create：2025/10/10 16:34:42
 * Remark：
 */

package stringx

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"os"
	"testing"
)

func TestUnsafeToBytes(t *testing.T) {
	testCase := []struct {
		name string
		val  string
		want []byte
	}{
		{
			name: "normal conversion",
			val:  "hello",
			want: []byte("hello"),
		},
		{
			name: "emoji conversion",
			val:  "😀!hello world",
			want: []byte("😀!hello world"),
		},
		{
			name: "chinese conversion",
			val:  "你好 世界！",
			want: []byte("你好 世界！"),
		},
	}

	for _, tt := range testCase {
		t.Run(tt.name, func(t *testing.T) {
			val := UnsafeToBytes(tt.val)
			assert.Equal(t, tt.want, val)
		})
	}
}

func TestUnsafeToString(t *testing.T) {
	testCase := []struct {
		name   string
		before func(t *testing.T)
		after  func(t *testing.T)
		val    func(t *testing.T) []byte
		want   string
	}{
		{
			name:   "normal conversion",
			before: func(t *testing.T) {},
			after:  func(t *testing.T) {},
			val: func(t *testing.T) []byte {
				return []byte("hello")
			},
			want: "hello",
		},
		{
			name:   "emoji conversion",
			before: func(t *testing.T) {},
			after:  func(t *testing.T) {},
			val: func(t *testing.T) []byte {
				return []byte("😀!hello world")
			},
			want: "😀!hello world",
		},
		{
			name:   "chinese conversion",
			before: func(t *testing.T) {},
			after:  func(t *testing.T) {},
			val: func(t *testing.T) []byte {
				return []byte("你好 世界！")
			},
			want: "你好 世界！",
		},
		{
			// 通过读取 file 文件 模拟 io.Reader 中存在的字节流 并将其转换为 string 检查他的正确性
			// 当然他必须是可控制的
			name: "file(io.Reader) read bytes stream conversion string",
			before: func(t *testing.T) {
				create, err := os.Create("/tmp/test_put.txt")
				require.NoError(t, err)
				defer func(create *os.File) {
					err := create.Close()
					if err != nil {
						return
					}
				}(create)
				_, err = create.WriteString("the test file...")
				require.NoError(t, err)
			},
			after: func(t *testing.T) {
				require.NoError(t, os.Remove("/tmp/test_put.txt"))
			},
			val: func(t *testing.T) []byte {
				open, err := os.Open("/tmp/test_put.txt")
				require.NoError(t, err)
				defer func(open *os.File) {
					err := open.Close()
					if err != nil {
						return
					}
				}(open)
				buf := bytes.Buffer{}
				_, err = buf.ReadFrom(open)
				require.NoError(t, err)
				return buf.Bytes()
			},
			want: "the test file...",
		},
	}

	for _, tt := range testCase {
		t.Run(tt.name, func(t *testing.T) {
			defer tt.after(t)
			tt.before(t)
			b := tt.val(t)
			val := UnsafeToString(b)
			assert.Equal(t, tt.want, val)
		})
	}
}
