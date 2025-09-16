## carefuly-echo

Go 通用工具库，包含对切片的高效与易用的操作封装。本文聚焦 `slice` 包的能力与用法。

### 安装

```bash
go get github.com/carefuly/carefuly-echo
```

### 引用

```base
import (
    s "github.com/carefuly/carefuly-echo/slice"
)
```

### 功能概览（slice）

- 添加/删除
  - `Add[T any]([]T, T, index) ([]T, error)`: 在指定位置插入元素（索引越界返回错误）
  - `Delete[T any]([]T, index) ([]T, T, error)`: 删除指定位置元素（索引越界返回错误）
  - `FilterDelete[T any]([]T, func(idx int, v T) bool) []T`: 删除满足条件的元素

- 查询/包含判断
  - `Contains[T comparable]([]T, T) bool`
  - `ContainsFunc[T any]([]T, func(v T) bool) bool`
  - `ContainsAny[T comparable](src, dst []T) bool`
  - `ContainsAnyFunc[T any](src, dst []T, equal func(a, b T) bool) bool`
  - `ContainsAll[T comparable](src, dst []T) bool`
  - `ContainsAllFunc[T any](src, dst []T, equal func(a, b T) bool) bool`

- 查找/索引
  - `Find[T any]([]T, func(v T) bool) (T, bool)`: 返回首个匹配元素
  - `FindAll[T any]([]T, func(v T) bool) []T`: 返回所有匹配元素
  - `Index[T comparable]([]T, T) int`
  - `IndexFunc[T any]([]T, func(v T) bool) int`
  - `LastIndex[T comparable]([]T, T) int`
  - `LastIndexFunc[T any]([]T, func(v T) bool) int`
  - `IndexAll[T comparable]([]T, T) []int`
  - `IndexAllFunc[T any]([]T, func(v T) bool) []int`

- 映射/转换
  - `Map[Src, Dst any]([]Src, func(i int, v Src) Dst) []Dst`
  - `FilterMap[Src, Dst any]([]Src, func(i int, v Src) (Dst, bool)) []Dst`
  - `ToMap[Ele any, Key comparable]([]Ele, func(Ele) Key) map[Key]Ele`
  - `ToMapV[Ele any, Key comparable, Val any]([]Ele, func(Ele) (Key, Val)) map[Key]Val`

- 聚合（数值）
  - `Max[T echo.RealNumber]([]T) T`
  - `Min[T echo.RealNumber]([]T) T`
  - `Sum[T echo.Number]([]T) T`
  - 注意：`Max`/`Min` 空切片会 `panic`（至少需要一个元素）

- 集合运算（去重）
  - 差集：`DiffSet[T comparable](src, dst []T) []T`，`DiffSetFunc[T any](src, dst []T, equal func(a, b T) bool) []T`
  - 交集：`IntersectSet[T comparable](src, dst []T) []T`，`IntersectSetFunc[T any](src, dst []T, equal func(a, b T) bool) []T`
  - 对称差：`SymmetricDiffSet[T comparable](src, dst []T) []T`，`SymmetricDiffSetFunc[T any](src, dst []T, equal func(a, b T) bool) []T`
  - 并集：`UnionSet[T comparable](src, dst []T) []T`，`UnionSetFunc[T any](src, dst []T, equal func(a, b T) bool) []T`
  - 说明：`Set` 版本要求元素 `comparable`，返回结果顺序不固定；`Func` 版本支持任意类型并保留首次出现顺序。

- 反转
  - `Reverse[T any]([]T) []T`: 返回倒序的新切片
  - `ReverseSelf[T any]([]T)`: 原地倒序

### 快速上手示例

```go
package main

import (
    "fmt"
    s "github.com/carefuly/carefuly-echo/slice"
)

func main() {
    // 添加/删除
    nums := []int{1, 2, 3}
    nums, _ = s.Add(nums, 9, 1)       // [1 9 2 3]
    nums, del, _ := s.Delete(nums, 2) // nums=[1 9 3], del=2
    fmt.Println(nums, del)

    // 查询/包含
    fmt.Println(s.Contains([]string{"a", "b"}, "b")) // true
    any := s.ContainsAny([]int{1, 2, 3}, []int{8, 3})    // true
    all := s.ContainsAll([]int{1, 2, 3}, []int{1, 3})    // true
    fmt.Println(any, all)

    // 查找/索引
    v, ok := s.Find([]int{3, 6, 9}, func(x int) bool { return x%3 == 0 && x > 5 })
    fmt.Println(v, ok)                    // 6 true
    fmt.Println(s.Index([]int{7, 8, 7}, 7))     // 0
    fmt.Println(s.LastIndex([]int{7, 8, 7}, 7)) // 2

    // 映射/转换
    squares := s.Map([]int{1, 2, 3}, func(_ int, x int) int { return x * x })
    evens := s.FilterMap([]int{1, 2, 3, 4}, func(_ int, x int) (int, bool) {
        return x, x%2 == 0
    })
    fmt.Println(squares, evens) // [1 4 9] [2 4]

    // 集合运算
    a, b := []int{1, 2, 2, 3}, []int{2, 4}
    fmt.Println(s.DiffSet(a, b))        // 可能输出 [1 3]
    fmt.Println(s.IntersectSet(a, b))   // 可能输出 [2]
    fmt.Println(s.SymmetricDiffSet(a,b))// 可能输出 [1 3 4]
    fmt.Println(s.UnionSet(a, b))       // 可能输出 [1 2 3 4]

    // 反转
    fmt.Println(s.Reverse([]int{1,2,3})) // [3 2 1]
}
```

### 复杂类型的等值比较（Func 变体）

```go
type User struct { ID int; Name string }

equalUser := func(a, b User) bool { return a.ID == b.ID }

u1 := []User{{1, "A"}, {2, "B"}, {2, "B2"}}
u2 := []User{{2, "B"}, {3, "C"}}

onlyInU1 := s.DiffSetFunc(u1, u2, equalUser)      // 基于 ID 的差集
both := s.IntersectSetFunc(u1, u2, equalUser)     // 基于 ID 的交集
all := s.UnionSetFunc(u1, u2, equalUser)          // 基于 ID 的并集（去重，保留首次出现）
xor := s.SymmetricDiffSetFunc(u1, u2, equalUser)  // 基于 ID 的对称差
_ = []any{onlyInU1, both, all, xor}
```

### 约束说明

- `echo.RealNumber`: `uint/uint8/uint16/uint32/uint64/int/int8/int16/int32/int64/float32/float64`
- `echo.Number`: `RealNumber` 联合集合再加 `complex64/complex128`

### 注意事项

- `Set` 后缀的方法会去重且返回顺序不固定；如需保序，使用对应的 `Func` 变体。
- `Max`/`Min` 需至少一个元素；空切片调用将 `panic`。
- `Add`/`Delete` 的索引越界会返回错误。

### 许可证

MIT


