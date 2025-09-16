/**
 * Description：
 * FileName：union.go
 * Author：CJiaの用心
 * Create：2025/9/16 11:59:43
 * Remark：
 */

package slice

// UnionSet 计算两个切片的并集（已去重）
// 使用内置 comparable 约束，适用于可直接比较的元素类型
// 返回结果顺序不固定
// 参数:
// 第一个切片
// 第二个切片
// 返回值:
// 并集切片，包含所有出现在任一输入切片中的唯一元素
func UnionSet[T comparable](src, dst []T) []T {
	srcMap, dstMap := toMap[T](src), toMap[T](dst)
	for key := range srcMap {
		dstMap[key] = struct{}{}
	}

	var ret = make([]T, 0, len(dstMap))
	for key := range dstMap {
		ret = append(ret, key)
	}

	return ret
}

// UnionSetFunc 计算两个切片的并集（已去重）
// 使用自定义相等函数，适用于无法直接比较的元素类型
// 你应该优先使用 UnionSet（当元素类型满足 comparable 时）
// 参数:
// 第一个切片
// 第二个切片
// 判断两个元素是否相等的函数
// 返回值:
// 并集切片，包含所有出现在任一输入切片中的唯一元素
func UnionSetFunc[T any](src, dst []T, equal equalFunc[T]) []T {
	var ret = make([]T, 0, len(src)+len(dst))

	ret = append(ret, dst...)
	ret = append(ret, src...)

	return deduplicateFunc[T](ret, equal)
}
