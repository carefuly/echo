/**
 * Description：
 * FileName：add.go
 * Author：CJiaの用心
 * Create：2025/9/6 00:51:03
 * Remark：
 */

package slice

import "github.com/carefuly/carefuly-echo/internal/slice"

// Add 在切片的指定索引位置插入元素
// 参数：
// src: 原始切片
// element: 要插入的元素
// index: 插入位置索引
// 返回值：
// 插入元素后的新切片
// 错误信息（索引越界时返回错误）
func Add[T any](src []T, element T, index int) ([]T, error) {
	return slice.Add(src, element, index)
}
