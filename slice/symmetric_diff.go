/**
 * Description：
 * FileName：symmetric_diff.go
 * Author：CJiaの用心
 * Create：2025/9/16 11:51:44
 * Remark：
 */

package slice

// SymmetricDiffSet 计算两个切片的对称差集（已去重）
// 使用内置 comparable 约束，适用于可直接比较的元素类型
// 返回结果顺序不固定
// 参数:
// 第一个切片
// 第二个切片
// 返回值:
// 对称差集切片，包含所有只存在于一个切片中的元素
func SymmetricDiffSet[T comparable](src, dst []T) []T {
	srcMap, dstMap := toMap[T](src), toMap[T](dst)
	for k := range dstMap {
		if _, ok := srcMap[k]; ok {
			delete(srcMap, k)
		} else {
			srcMap[k] = struct{}{}
		}
	}

	res := make([]T, 0, len(srcMap))
	for k := range srcMap {
		res = append(res, k)
	}

	return res
}

// SymmetricDiffSetFunc 计算两个切片的对称差集（已去重）
// 使用自定义相等函数，适用于无法直接比较的元素类型
// 你应该优先使用 SymmetricDiffSet（当元素类型满足 comparable 时）
// 参数:
// 第一个切片
// 第二个切片
// 判断两个元素是否相等的函数
// 返回值:
// 对称差集切片，包含所有只存在于一个切片中的元素
func SymmetricDiffSetFunc[T any](src, dst []T, equal equalFunc[T]) []T {
	var res []T

	// 找出在src不在dst的元素
	for _, v := range src {
		if !ContainsFunc[T](dst, func(t T) bool {
			return equal(t, v)
		}) {
			res = append(res, v)
		}
	}

	// 找出在dst不在src的元素
	for _, v := range dst {
		if !ContainsFunc[T](src, func(t T) bool {
			return equal(t, v)
		}) {
			res = append(res, v)
		}
	}

	return deduplicateFunc[T](res, equal)
}
