/**
 * Description：
 * FileName：error.go
 * Author：CJiaの用心
 * Create：2025/9/5 00:32:50
 * Remark：
 */

package errs

import (
	"fmt"
	"time"
)

// NewErrIndexOutOfRange 创建一个代表下标超出范围的错误
func NewErrIndexOutOfRange(length int, index int) error {
	return fmt.Errorf("echo: 下标超出范围，长度 %d, 下标 %d", length, index)
}

// NewErrInvalidType 创建一个代表类型转换失败的错误
func NewErrInvalidType(want string, got any) error {
	return fmt.Errorf("echo: 类型转换失败，预期类型:%s, 实际值:%#v", want, got)
}

// NewErrInvalidIntervalValue 创建一个无效间隔值的错误
func NewErrInvalidIntervalValue(interval time.Duration) error {
	return fmt.Errorf("echo: 无效的间隔时间 %d, 预期值应大于 0", interval)
}

// NewErrInvalidMaxIntervalValue 创建一个无效最大间隔值的错误
func NewErrInvalidMaxIntervalValue(maxInterval, initialInterval time.Duration) error {
	return fmt.Errorf("echo: 最大重试间隔的时间 [%d] 应大于等于初始重试的间隔时间 [%d] ", maxInterval, initialInterval)
}

// NewErrRetryExhausted 创建一个超过最大重试次数的错误
func NewErrRetryExhausted(lastErr error) error {
	return fmt.Errorf("echo: 超过最大重试次数，业务返回的最后一个 error %w", lastErr)
}
