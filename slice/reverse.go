/**
 * Description：
 * FileName：reverse.go
 * Author：CJiaの用心
 * Create：2025/9/16 11:36:59
 * Remark：
 */

package slice

// Reverse 创建一个与原切片顺序相反的新切片
// 参数:
// 原始切片，类型为任意类型的切片 []T
// 返回值:
// 倒序后的新切片，长度和容量与原始切片相同
func Reverse[T any](src []T) []T {
	var ret = make([]T, 0, len(src))

	// 从后向前遍历原切片，填充新切片
	for i := len(src) - 1; i >= 0; i-- {
		ret = append(ret, src[i])
	}
	return ret
}

// ReverseSelf 原地反转切片元素顺序。
// 参数:
// 需要反转的切片，类型为任意类型的切片 []T
func ReverseSelf[T any](src []T) {
	// 使用双指针从两端向中间交换元素
	left, right := 0, len(src)-1
	for left < right {
		// 同时交换头尾元素
		src[left], src[right] = src[right], src[left]
		// 指针向中间移动
		left++
		right--
	}
}
