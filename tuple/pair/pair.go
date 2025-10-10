/**
 * Description：
 * FileName：pair.go
 * Author：CJiaの用心
 * Create：2025/10/10 16:41:57
 * Remark：
 */

package pair

import "fmt"

type Pair[K any, V any] struct {
	Key   K
	Value V
}

func (pair *Pair[K, V]) String() string {
	return fmt.Sprintf("<%#v, %#v>", pair.Key, pair.Value)
}

// Split 方法将Key, Value作为返回参数传出。
func (pair *Pair[K, V]) Split() (K, V) {
	return pair.Key, pair.Value
}

func NewPair[K any, V any](
	key K,
	value V,
) Pair[K, V] {
	return Pair[K, V]{
		Key:   key,
		Value: value,
	}
}
