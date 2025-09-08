/**
 * Description：
 * FileName：constrain.go
 * Author：CJiaの用心
 * Create：2025/9/8 23:54:11
 * Remark：
 */

package echo

// RealNumber 实数
// 绝大多数情况下，你都应该用这个来表达数字的含义
type RealNumber interface {
	~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 |
		~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~float32 | ~float64
}

type Number interface {
	RealNumber | ~complex64 | ~complex128
}
