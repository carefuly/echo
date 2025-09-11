/**
 * Description：
 * FileName：find.go
 * Author：CJiaの用心
 * Create：2025/9/12 00:14:21
 * Remark：
 */

package slice

// Find 在切片中查找第一个满足条件的元素
// 参数：
// 待搜索的切片
// 匹配函数，接受元素并返回是否匹配
// 返回值：
// 第一个匹配的元素（如果找到）
// 布尔值表示是否找到
func Find[T any](src []T, match matchFunc[T]) (T, bool) {
	for i := range src {
		if match(src[i]) {
			return src[i], true
		}
	}
	var zero T
	return zero, false
}

// FindAll 查找所有满足条件的元素
// 永远不会返回 nil
// 自动优化容量分配
// 使用两阶段收集提高性能
// 参数：
// 待搜索的切片
// 匹配函数，接受元素并返回是否匹配
// 返回值：
// 包含所有匹配元素的切片（可能为空）
func FindAll[T any](src []T, match matchFunc[T]) []T {
	indices := make([]int, 0, m(len(src)/8, 64))
	for i, val := range src {
		if match(val) {
			indices = append(indices, i)
		}
	}

	// 直接填充结果
	res := make([]T, len(indices))
	for j, idx := range indices {
		res[j] = src[idx]
	}
	return res
}

// m 返回两个整数中较小的一个
func m(a, b int) int {
	if a < b {
		return a
	}
	return b
}
