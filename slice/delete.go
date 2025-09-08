/**
 * Description：
 * FileName：delete.go
 * Author：CJiaの用心
 * Create：2025/9/6 00:55:37
 * Remark：
 */

package slice

import "github.com/carefuly/carefuly-echo/internal/slice"

// Delete 从切片中删除指定索引位置的元素
// 参数：
// src: 原始切片
// index: 要删除的元素索引位置
// 返回值：
// 删除元素后的新切片
// 被删除的元素值
// 错误信息（索引越界时返回错误）
func Delete[T any](src []T, index int) ([]T, T, error) {
	return slice.Delete[T](src, index)
}

// FilterDelete 删除切片中满足条件的元素
// 参数：
// src: 原始切片（函数会直接修改此切片）
// m: 条件判断函数，返回 true 表示删除该元素
// 返回值：
// 删除元素后的新切片（与原切片共享底层数组）
func FilterDelete[T any](src []T, m func(idx int, src T) bool) []T {
	// emptyPos 指向下一个保留元素应该放置的位置
	emptyPos := 0
	length := len(src)

	// 检查是否没有任何元素被删除
	allKept := true
	for i := 0; i < length; i++ {
		if m(i, src[i]) {
			allKept = false
			break
		}
	}
	if allKept {
		return src
	}

	// 遍历切片，移动需要保留的元素
	for i := 0; i < length; i++ {
		if m(i, src[i]) {
			// 跳过需要删除的元素
			continue
		}

		// 仅当位置不同时才移动元素
		if emptyPos != i {
			src[emptyPos] = src[i]
		}
		emptyPos++
	}

	return src[:emptyPos]
}
