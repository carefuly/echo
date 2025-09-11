/**
 * Description：
 * FileName：diff.go
 * Author：CJiaの用心
 * Create：2025/9/11 23:38:20
 * Remark：
 */

package slice

// DiffSet 计算两个切片的差集（src - dst）
// 结果已去重
// 返回顺序不确定
// 只支持 comparable 类型
// 参数：
// 源切片，作为被减数
// 目标切片，作为减数
// 返回值：
// 差集切片，包含所有在 src 中但不在 dst 中的元素
// 当 src 为空时返回空切片
func DiffSet[T comparable](src, dst []T) []T {
	srcMap := toMap[T](src)
	dstMap := toMap(dst)

	// 只遍历去重后的目标元素
	for key := range dstMap {
		delete(srcMap, key)
	}

	// 构建结果切片
	ret := make([]T, 0, len(srcMap))
	for key := range srcMap {
		ret = append(ret, key)
	}
	return ret
}

// DiffSetFunc 计算两个切片的差集（src - dst），使用自定义相等函数
// 结果已去重
// 返回顺序为源切片中第一次出现的顺序
// 适用于任意类型（any）
// 参数：
// 源切片，作为被减数
// 目标切片，作为减数
// 自定义相等判断函数，用于比较两个元素是否相等
// 返回值：
// 差集切片，包含所有在 src 中但不在 dst 中的元素
// 当 src 为空时返回空切片
func DiffSetFunc[T any](src, dst []T, equal equalFunc[T]) []T {
	var ret = make([]T, 0, len(src))
	for _, val := range src {
		if !ContainsFunc[T](dst, func(src T) bool {
			return equal(src, val)
		}) {
			ret = append(ret, val)
		}
	}

	return deduplicateFunc[T](ret, equal)
}
