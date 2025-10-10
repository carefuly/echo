/**
 * Description：
 * FileName：rand_code_test.go.go
 * Author：CJiaの用心
 * Create：2025/10/10 16:44:12
 * Remark：
 */

package randx_test

import (
	"errors"
	"github.com/carefuly/careful-echo/randx"
	"github.com/stretchr/testify/assert"
	"regexp"
	"strings"
	"testing"
)

var (
	errTypeNotSupported   = errors.New("echo:不支持的类型")
	errLengthLessThanZero = errors.New("echo:长度必须大于等于0")
)

func TestRandCode(t *testing.T) {
	testCases := []struct {
		name      string
		length    int
		typ       randx.Type
		wantMatch string
		wantErr   error
	}{
		{
			name:      "数字验证码",
			length:    100,
			typ:       randx.TypeDigit,
			wantMatch: "^[0-9]+$",
			wantErr:   nil,
		},
		{
			name:      "小写字母验证码",
			length:    100,
			typ:       randx.TypeLowerCase,
			wantMatch: "^[a-z]+$",
			wantErr:   nil,
		},
		{
			name:      "数字+小写字母验证码",
			length:    100,
			typ:       randx.TypeDigit | randx.TypeLowerCase,
			wantMatch: "^[a-z0-9]+$",
			wantErr:   nil,
		},
		{
			name:      "数字+大写字母验证码",
			length:    100,
			typ:       randx.TypeDigit | randx.TypeUpperCase,
			wantMatch: "^[A-Z0-9]+$",
			wantErr:   nil,
		},
		{
			name:      "大写字母验证码",
			length:    100,
			typ:       randx.TypeUpperCase,
			wantMatch: "^[A-Z]+$",
			wantErr:   nil,
		},
		{
			name:      "大小写字母验证码",
			length:    100,
			typ:       randx.TypeUpperCase | randx.TypeLowerCase,
			wantMatch: "^[a-zA-Z]+$",
			wantErr:   nil,
		},
		{
			name:      "数字+大小写字母验证码",
			length:    100,
			typ:       randx.TypeDigit | randx.TypeUpperCase | randx.TypeLowerCase,
			wantMatch: "^[0-9a-zA-Z]+$",
			wantErr:   nil,
		},
		{
			name:      "所有类型验证",
			length:    100,
			typ:       randx.TypeMixed,
			wantMatch: "^[\\S\\s]+$",
			wantErr:   nil,
		},
		{
			name:      "特殊字符类型验证",
			length:    100,
			typ:       randx.TypeSpecial,
			wantMatch: "^[^0-9a-zA-Z]+$",
			wantErr:   nil,
		},
		{
			name:      "未定义类型(超过范围)",
			length:    100,
			typ:       randx.TypeMixed + 1,
			wantMatch: "",
			wantErr:   errTypeNotSupported,
		},
		{
			name:      "未定义类型(0)",
			length:    100,
			typ:       0,
			wantMatch: "",
			wantErr:   errTypeNotSupported,
		},
		{
			name:      "长度小于0",
			length:    -1,
			typ:       0,
			wantMatch: "",
			wantErr:   errLengthLessThanZero,
		},
		{
			name:      "长度等于0",
			length:    0,
			typ:       randx.TypeMixed,
			wantMatch: "",
			wantErr:   nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			code, err := randx.RandCode(tc.length, tc.typ)
			if tc.wantErr != nil {
				assert.Equal(t, tc.wantErr, err)
				return
			}
			assert.Len(t, code, tc.length)
			if tc.length > 0 {
				matched, err := regexp.MatchString(tc.wantMatch, code)
				assert.Nil(t, err)
				assert.Truef(t, matched, "expected %s but got %s", tc.wantMatch, code)
			}

		})
	}
}

func TestRandStrByCharset(t *testing.T) {
	matchFunc := func(str, charset string) bool {
		for _, c := range str {
			if !strings.Contains(charset, string(c)) {
				return false
			}
		}
		return true
	}
	testCases := []struct {
		name    string
		length  int
		charset string
		wantErr error
	}{
		{
			name:    "长度小于0",
			length:  -1,
			charset: "123",
			wantErr: errLengthLessThanZero,
		},
		{
			name:    "长度等于0",
			length:  0,
			charset: "123",
			wantErr: nil,
		},
		{
			name:    "随机字符串测试",
			length:  100,
			charset: "2rg248ry227t@@",
			wantErr: nil,
		},
		{
			name:    "随机字符串测试",
			length:  100,
			charset: "2rg248ry227t@&*($.!",
			wantErr: nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			code, err := randx.RandStrByCharset(tc.length, tc.charset)
			if tc.wantErr != nil {
				assert.Equal(t, tc.wantErr, err)
				return
			}

			assert.Len(t, code, tc.length)
			if tc.length > 0 {
				assert.True(t, matchFunc(code, tc.charset))
			}

		})
	}
}
