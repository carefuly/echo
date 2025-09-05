/**
 * Description：
 * FileName：shrink.go
 * Author：CJiaの用心
 * Create：2025/9/6 00:27:04
 * Remark：
 */

package slice

const (
	smallCapacityThreshold = 64
	largeCapacityThreshold = 2048
	largeShrinkRatio       = 0.625 // 5/8
	minShrinkLengthRatio   = 2     // 容量/长度 ≥ 2 时才考虑缩容
	smallShrinkLengthRatio = 4     // 小切片需要更高的空间浪费比例才缩容
)

// Shrink 对切片进行缩容以减少内存占用
// 参数：
// src: 需要缩容的切片
// 返回值：
// 缩容后的新切片
func Shrink[T any](src []T) []T {
	c, l := cap(src), len(src)

	// 特殊情况处理：空切片或无需缩容
	if l == 0 {
		return make([]T, 0)
	}

	newCap, shouldShrink := calCapacity(c, l)
	if !shouldShrink {
		return src
	}

	// 创建新切片并拷贝数据
	// 使用精确容量分配避免额外内存浪费
	newSlice := make([]T, l, newCap)
	copy(newSlice, src)
	return newSlice
}

// calCapacity 计算缩容后的新容量
// 参数：
// cap: 当前切片容量
// len: 当前切片长度
// 返回值：
// newCap: 计算后的新容量
// shrink: 是否需要执行缩容操作
func calCapacity(cap, len int) (newCap int, shrink bool) {
	// 小切片不缩容
	if cap < smallCapacityThreshold {
		return cap, false
	}

	// 计算空间浪费比例
	spaceRatio := float64(cap) / float64(len)

	// 大切片处理逻辑
	if cap > largeCapacityThreshold {
		if spaceRatio >= minShrinkLengthRatio {
			// 使用浮点计算保证精度，然后四舍五入
			return int(float64(cap)*largeShrinkRatio + 0.5), true
		}
		return cap, false
	}

	// 中等切片处理逻辑
	if spaceRatio >= smallShrinkLengthRatio {
		return cap / 2, true
	}

	return cap, false
}
