/**
 * Description：
 * FileName：intersect.go
 * Author：CJiaの用心
 * Create：2025/9/12 15:58:49
 * Remark：
 */

package slice

// IntersectSet 计算两个切片的交集（只支持 comparable 类型）
// 参数：
// 第一个切片
// 第二个切片
// 返回值：
// 交集切片（已去重）
func IntersectSet[T comparable](src []T, dst []T) []T {
	srcMap := toMap(src)
	var ret = make([]T, 0, len(src))
	// 交集小于等于两个集合中的任意一个
	for _, val := range dst {
		if _, exist := srcMap[val]; exist {
			ret = append(ret, val)
		}
	}
	return deduplicate[T](ret)
}

// IntersectSetFunc 使用自定义相等函数计算两个切片的交集（支持任意类型）
// 参数：
// 第一个切片
// 第二个切片
// 元素相等性判断函数
// 返回值：
// 交集切片（已去重）
func IntersectSetFunc[T any](src []T, dst []T, equal equalFunc[T]) []T {
	var ret = make([]T, 0, len(src))
	for _, v := range dst {
		if ContainsFunc[T](src, func(t T) bool {
			return equal(t, v)
		}) {
			ret = append(ret, v)
		}
	}
	return deduplicateFunc[T](ret, equal)
}
