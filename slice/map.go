/**
 * Description：
 * FileName：map.go
 * Author：CJiaの用心
 * Create：2025/9/9 00:24:41
 * Remark：
 */

package slice

// FilterMap 对切片元素进行过滤和转换
// 对每个元素调用映射函数 m，仅当 m 返回 true 时，将转换结果加入返回切片
// 即使某些元素被过滤，也会遍历所有元素
// 参数:
// 输入切片
// 映射函数，接收索引和元素，返回转换后的元素和是否保留的布尔值
// 返回值:
// 过滤并转换后的新切片
func FilterMap[Src any, Dst any](src []Src, m func(idx int, src Src) (Dst, bool)) []Dst {
	// 预分配足够容量（最大可能长度）
	res := make([]Dst, 0, len(src))
	// 使用索引变量避免每次迭代计算索引
	for i := 0; i < len(src); i++ {
		if dst, ok := m(i, src[i]); ok {
			res = append(res, dst)
		}
	}
	return res
}

// Map 将切片元素转换为新类型
// 对每个元素调用映射函数 m，将转换结果组成新切片返回
// 参数:
// 输入切片
// 映射函数，接收索引和元素，返回转换后的元素
// 返回值:
// 转换后的新切片，长度与输入相同
func Map[Src any, Dst any](src []Src, m func(idx int, src Src) Dst) []Dst {
	dst := make([]Dst, len(src))
	// 使用索引变量避免每次迭代计算索引
	for i := 0; i < len(src); i++ {
		dst[i] = m(i, src[i])
	}
	return dst
}

// ToMap 将切片转换为映射 [Key]Ele
// 使用提供的函数从元素中提取键
// 注意:
// - 如果多个元素产生相同的键，后面的元素会覆盖前面的
// - 即使输入为 nil，也保证返回非 nil 的空映射
// 参数:
// 输入切片
// 从元素中提取键的函数
// 返回值:
// 结果映射，键为 fn 返回的值，值为原始元素
func ToMap[Ele any, Key comparable](
	elements []Ele,
	fn func(element Ele) Key,
) map[Key]Ele {
	return ToMapV(
		elements,
		func(element Ele) (Key, Ele) {
			return fn(element), element
		})
}

// ToMapV 将切片转换为映射 [Key]Val
// 使用提供的函数从元素中提取键和值
//
// 注意:
// - 如果多个元素产生相同的键，后面的元素会覆盖前面的
// - 即使输入为 nil，也保证返回非 nil 的空映射
//
// 参数:
//
//	elements: 输入切片
//	fn: 从元素中提取键和值的函数
//
// 返回:
//
//	结果映射，键和值由 fn 返回
func ToMapV[Ele any, Key comparable, Val any](
	elements []Ele,
	fn func(element Ele) (Key, Val),
) (resultMap map[Key]Val) {
	resultMap = make(map[Key]Val, len(elements))
	for _, element := range elements {
		k, v := fn(element)
		resultMap[k] = v
	}
	return
}

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

// deduplicate 对切片进行去重处理，返回不包含重复元素的新切片
// 参数：
// 待去重的原始切片
// 返回值：
// 去重后的新切片（元素顺序不保证与原始切片一致）
func deduplicate[T comparable](data []T) []T {
	dataMap := toMap[T](data)

	var newData = make([]T, 0, len(dataMap))

	for key := range dataMap {
		newData = append(newData, key)
	}

	return newData
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
