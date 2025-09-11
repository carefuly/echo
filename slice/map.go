/**
 * Description：
 * FileName：map.go
 * Author：CJiaの用心
 * Create：2025/9/9 00:24:41
 * Remark：
 */

package slice

// toMap 将切片转换为 map 用于高效查找
// 参数：
// 需要转换的切片
// 返回值：
// 包含切片元素的 map，值使用空结构体以减少内存占用
func toMap[T comparable](src []T) map[T]struct{} {
	var dataMap = make(map[T]struct{}, len(src))
	for _, v := range src {
		// 使用空结构体,减少内存消耗
		dataMap[v] = struct{}{}
	}
	return dataMap
}

// deduplicateFunc 使用自定义相等函数对切片进行去重
// 参数：
// 需要去重的切片
// 自定义相等判断函数
// 返回值：
// 去重后的切片，保留原始顺序中第一次出现的元素
func deduplicateFunc[T any](data []T, equal equalFunc[T]) []T {
	if len(data) == 0 {
		return data
	}

	// 使用一个切片记录唯一元素
	unique := make([]T, 0, len(data))

	// 使用双重循环但优化比较次数
	for i, v := range data {
		found := false

		// 只检查前面的元素
		for j := 0; j < i; j++ {
			if equal(data[j], v) {
				found = true
				break
			}
		}

		if !found {
			unique = append(unique, v)
		}
	}

	return unique
}
