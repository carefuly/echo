/**
 * Description：
 * FileName：contains.go
 * Author：CJiaの用心
 * Create：2025/9/9 00:11:54
 * Remark：
 */

package slice

// Contains 判断 src 里面是否存在 dst
// 参数：
// 待搜索的切片
// 需要查找的元素
// 返回值：
// 如果切片中包含该元素，则返回 true；否则返回 false
func Contains[T comparable](src []T, dst T) bool {
	return ContainsFunc[T](src, func(src T) bool {
		return src == dst
	})
}

// ContainsFunc 判断 src 里面是否存在 dst
// 参数：
// 待搜索的切片
// 匹配函数，接受一个元素并返回是否匹配
// 返回值：
// 如果切片中存在至少一个元素满足匹配条件，则返回 true；否则返回 false
func ContainsFunc[T any](src []T, equal func(src T) bool) bool {
	// 遍历调用equal函数进行判断
	for _, v := range src {
		if equal(v) {
			return true
		}
	}
	return false
}

// ContainsAny 判断 src 里面是否存在 dst 中的任何一个元素
// 参数：
// 源切片，用于搜索的切片
// 目标切片，包含需要查找的元素
// 返回值：
// 如果源切片中包含目标切片中的至少一个元素，则返回 true；否则返回 false
func ContainsAny[T comparable](src, dst []T) bool {
	// 处理空切片情况
	if len(src) == 0 || len(dst) == 0 {
		return false
	}

	srcMap := toMap[T](src)
	for _, v := range dst {
		if _, exist := srcMap[v]; exist {
			return true
		}
	}
	return false
}

// ContainsAnyFunc 判断 src 里面是否存在 dst 中的任何一个元素
// 参数：
// 源切片，用于搜索的切片
// 目标切片，包含需要查找的元素
// 比较函数，用于判断两个元素是否相等
// 返回值：
// 如果源切片中包含目标切片中的至少一个元素（根据比较函数判断），则返回 true；否则返回 false
func ContainsAnyFunc[T any](src, dst []T, equal equalFunc[T]) bool {
	// 处理空切片情况
	if len(src) == 0 || len(dst) == 0 {
		return false
	}

	for _, valDst := range dst {
		for _, valSrc := range src {
			if equal(valSrc, valDst) {
				return true
			}
		}
	}
	return false
}

// ContainsAll 判断 src 里面是否存在 dst 中的所有元素
// 参数：
// 源切片，用于搜索的切片
// 目标切片，包含需要查找的元素
// 返回值：
// 如果源切片包含目标切片中的所有元素，则返回 true；否则返回 false
func ContainsAll[T comparable](src, dst []T) bool {
	// 空目标切片总是返回 true
	if len(dst) == 0 {
		return true
	}

	// 源切片为空但目标切片非空时返回 false
	if len(src) == 0 {
		return false
	}

	srcMap := toMap[T](src)
	for _, v := range dst {
		if _, exist := srcMap[v]; !exist {
			return false
		}
	}
	return true
}

// ContainsAllFunc 判断 src 里面是否存在 dst 中的所有元素
// 参数：
// 源切片，用于搜索的切片
// 目标切片，包含需要查找的元素
// 比较函数，用于判断两个元素是否相等
// 返回值：
// 如果源切片包含目标切片中的所有元素（根据比较函数判断），则返回 true；否则返回 false
func ContainsAllFunc[T any](src, dst []T, equal equalFunc[T]) bool {
	// 空目标切片总是返回 true
	if len(dst) == 0 {
		return true
	}

	// 源切片为空但目标切片非空时返回 false
	if len(src) == 0 {
		return false
	}

	for _, valDst := range dst {
		if !ContainsFunc[T](src, func(src T) bool {
			return equal(src, valDst)
		}) {
			return false
		}
	}
	return true
}
