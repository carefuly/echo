/**
 * Description：
 * FileName：aggregate.go
 * Author：CJiaの用心
 * Create：2025/9/8 23:51:46
 * Remark：
 */

package slice

import echo "github.com/carefuly/carefuly-echo"

// Max 返回切片中的最大值
// 参数：
// 包含数值的切片
// 返回值：
// 切片中的最大值
func Max[T echo.RealNumber](values []T) T {
	if len(values) == 0 {
		panic("切片至少包含一个元素")
	}
	res := values[0]
	for _, v := range values[1:] {
		if v > res {
			res = v
		}
	}
	return res
}

// Min 返回切片中的最小值
// 参数：
// 包含数值的切片
// 返回值：
// 切片中的最小值
func Min[T echo.RealNumber](values []T) T {
	if len(values) == 0 {
		panic("切片至少包含一个元素")
	}
	res := values[0]
	for _, v := range values[1:] {
		if v < res {
			res = v
		}
	}
	return res
}

// Sum 计算切片所有元素的总和
// 参数：
// 包含数值的切片
// 返回值：
// 所有元素的总和
func Sum[T echo.Number](values []T) T {
	var res T
	for _, v := range values {
		res += v
	}
	return res
}
