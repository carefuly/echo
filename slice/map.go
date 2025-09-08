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
