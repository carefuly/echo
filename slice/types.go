/**
 * Description：
 * FileName：types.go
 * Author：CJiaの用心
 * Create：2025/9/9 00:26:37
 * Remark：
 */

package slice

// equalFunc 比较两个元素是否相等
type equalFunc[T any] func(src, dst T) bool

type matchFunc[T any] func(src T) bool
