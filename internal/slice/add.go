/**
 * Description：
 * FileName：add.go
 * Author：CJiaの用心
 * Create：2025/9/5 00:37:26
 * Remark：
 */

package slice

import "github.com/carefuly/careful-echo/internal/errs"

// Add 在切片的指定索引位置插入元素
// 参数：
// src: 原始切片
// element: 要插入的元素
// index: 插入位置索引
// 返回值：
// 插入元素后的新切片
// 错误信息（索引越界时返回错误）
func Add[T any](src []T, element T, index int) ([]T, error) {
	length := len(src)

	if index < 0 || index > length {
		return nil, errs.NewErrIndexOutOfRange(length, index)
	}

	// 扩展切片：追加零值元素
	// 容量不足时会自动扩容，确保后续操作不会触发多次扩容
	src = append(src, *new(T)) // 追加类型零值

	// 移动元素：使用 copy 高效移动元素段
	// 将 [index:] 区间的元素整体后移一位
	// copy 参数说明：
	//   dst: src[index+1:] -> 从 index+1 开始的切片
	//   src: src[index:]   -> 从 index 开始到倒数第二个元素
	copy(src[index+1:], src[index:])

	// 插入新元素
	src[index] = element

	return src, nil
}
