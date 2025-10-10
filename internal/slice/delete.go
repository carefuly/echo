/**
 * Description：
 * FileName：delete.go
 * Author：CJiaの用心
 * Create：2025/9/6 00:10:24
 * Remark：
 */

package slice

import "github.com/carefuly/careful-echo/internal/errs"

// Delete 从切片中删除指定索引位置的元素
// 参数：
// src: 原始切片
// index: 要删除的元素索引位置
// 返回值：
// 删除元素后的新切片
// 被删除的元素值
// 错误信息（索引越界时返回错误）
func Delete[T any](src []T, index int) ([]T, T, error) {
	length := len(src)

	if index < 0 || index >= length {
		var zero T
		return nil, zero, errs.NewErrIndexOutOfRange(length, index)
	}

	// 保存被删除的元素值
	deleted := src[index]

	// 移动元素：使用 copy 高效移动元素段
	// 将 [index+1:] 区间的元素整体前移一位
	// copy 参数说明：
	//   dst: src[index:]   -> 从 index 开始的切片
	//   src: src[index+1:] -> 从 index+1 开始到末尾
	copy(src[index:], src[index+1:])

	// 截断切片：移除最后一个多余元素
	// 注意：底层数组不变，保留容量避免频繁内存分配
	src = src[:length-1]

	return src, deleted, nil
}
