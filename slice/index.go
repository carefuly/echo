/**
 * Description：
 * FileName：index.go
 * Author：CJiaの用心
 * Create：2025/9/12 14:28:54
 * Remark：
 */

package slice

// Index 返回和 dst 相等的第一个元素下标
// 参数：
// 源切片
// 目标元素
// 返回值：
// 元素下标（未找到返回 -1）
func Index[T comparable](src []T, dst T) int {
	return IndexFunc[T](src, func(elem T) bool {
		return elem == dst
	})
}

// IndexFunc 返回 match 返回 true 的第一个下标
// 参数：
// 源切片
// 元素匹配函数
// 返回值：
// 元素下标（未找到返回 -1）
func IndexFunc[T any](src []T, match matchFunc[T]) int {
	for i, v := range src {
		if match(v) {
			return i
		}
	}
	return -1
}

// LastIndex 返回和 dst 相等的最后一个元素下标
// 参数：
// 源切片
// 目标元素
// 返回值：
// 元素下标（未找到返回 -1）
func LastIndex[T comparable](src []T, dst T) int {
	return LastIndexFunc[T](src, func(elem T) bool {
		return elem == dst
	})
}

// LastIndexFunc 返回和 dst 相等的最后一个元素下标
// 参数：
// 源切片
// 元素匹配函数
// 返回值：
// 元素下标（未找到返回 -1）
func LastIndexFunc[T any](src []T, match matchFunc[T]) int {
	for i := len(src) - 1; i >= 0; i-- {
		if match(src[i]) {
			return i
		}
	}
	return -1
}

// IndexAll 返回和 dst 相等的所有元素的下标
// 参数：
// 源切片
// 目标元素
// 返回值：
// 所有匹配元素的下标切片（未找到返回空切片）
func IndexAll[T comparable](src []T, dst T) []int {
	return IndexAllFunc[T](src, func(elem T) bool {
		return elem == dst
	})
}

// IndexAllFunc 返回和 match 返回 true 的所有元素的下标
// 参数：
// 源切片
// 元素匹配函数
// 返回值：
// 所有匹配元素的下标切片（未找到返回空切片）
func IndexAllFunc[T any](src []T, match matchFunc[T]) []int {
	var indexes = make([]int, 0, len(src))
	for k, v := range src {
		if match(v) {
			indexes = append(indexes, k)
		}
	}
	return indexes
}
