## carefuly-echo

A Go utility library. This document focuses on the `slice` package: practical and efficient helpers for working with slices.

### Install

```bash
go get github.com/carefuly/carefuly-echo
```

### Import

```go
import (
    s "github.com/carefuly/carefuly-echo/slice"
)
```

### Feature overview (slice)

- Add/Delete
  - `Add[T any]([]T, T, index) ([]T, error)`: insert an element at position (bounds-checked)
  - `Delete[T any]([]T, index) ([]T, T, error)`: delete element at position (bounds-checked)
  - `FilterDelete[T any]([]T, func(idx int, v T) bool) []T`: delete elements that match predicate

- Containment
  - `Contains[T comparable]([]T, T) bool`
  - `ContainsFunc[T any]([]T, func(v T) bool) bool`
  - `ContainsAny[T comparable](src, dst []T) bool`
  - `ContainsAnyFunc[T any](src, dst []T, equal func(a, b T) bool) bool`
  - `ContainsAll[T comparable](src, dst []T) bool`
  - `ContainsAllFunc[T any](src, dst []T, equal func(a, b T) bool) bool`

- Find/Index
  - `Find[T any]([]T, func(v T) bool) (T, bool)`: first matching element
  - `FindAll[T any]([]T, func(v T) bool) []T`: all matching elements
  - `Index[T comparable]([]T, T) int`
  - `IndexFunc[T any]([]T, func(v T) bool) int`
  - `LastIndex[T comparable]([]T, T) int`
  - `LastIndexFunc[T any]([]T, func(v T) bool) int`
  - `IndexAll[T comparable]([]T, T) []int`
  - `IndexAllFunc[T any]([]T, func(v T) bool) []int`

- Map/Transform
  - `Map[Src, Dst any]([]Src, func(i int, v Src) Dst) []Dst`
  - `FilterMap[Src, Dst any]([]Src, func(i int, v Src) (Dst, bool)) []Dst`
  - `ToMap[Ele any, Key comparable]([]Ele, func(Ele) Key) map[Key]Ele`
  - `ToMapV[Ele any, Key comparable, Val any]([]Ele, func(Ele) (Key, Val)) map[Key]Val`

- Aggregation (numeric)
  - `Max[T echo.RealNumber]([]T) T`
  - `Min[T echo.RealNumber]([]T) T`
  - `Sum[T echo.Number]([]T) T`
  - Note: `Max`/`Min` panic on empty slices (require at least one element)

- Set operations (deduplicated)
  - Difference: `DiffSet[T comparable](src, dst []T) []T`, `DiffSetFunc[T any](src, dst []T, equal func(a, b T) bool) []T`
  - Intersection: `IntersectSet[T comparable](src, dst []T) []T`, `IntersectSetFunc[T any](src, dst []T, equal func(a, b T) bool) []T`
  - Symmetric difference: `SymmetricDiffSet[T comparable](src, dst []T) []T`, `SymmetricDiffSetFunc[T any](src, dst []T, equal func(a, b T) bool) []T`
  - Union: `UnionSet[T comparable](src, dst []T) []T`, `UnionSetFunc[T any](src, dst []T, equal func(a, b T) bool) []T`
  - Notes: `Set` variants require `comparable` elements and return in unspecified order; `Func` variants support any type and preserve first-occurrence order.

- Reverse
  - `Reverse[T any]([]T) []T`: returns a new reversed slice
  - `ReverseSelf[T any]([]T)`: in-place reverse

### Quick start

```go
package main

import (
    "fmt"
    s "github.com/carefuly/carefuly-echo/slice"
)

func main() {
    // Add/Delete
    nums := []int{1, 2, 3}
    nums, _ = s.Add(nums, 9, 1)       // [1 9 2 3]
    nums, del, _ := s.Delete(nums, 2) // nums=[1 9 3], del=2
    fmt.Println(nums, del)

    // Containment
    fmt.Println(s.Contains([]string{"a", "b"}, "b")) // true
    any := s.ContainsAny([]int{1, 2, 3}, []int{8, 3})    // true
    all := s.ContainsAll([]int{1, 2, 3}, []int{1, 3})    // true
    fmt.Println(any, all)

    // Find/Index
    v, ok := s.Find([]int{3, 6, 9}, func(x int) bool { return x%3 == 0 && x > 5 })
    fmt.Println(v, ok)                    // 6 true
    fmt.Println(s.Index([]int{7, 8, 7}, 7))     // 0
    fmt.Println(s.LastIndex([]int{7, 8, 7}, 7)) // 2

    // Map/Transform
    squares := s.Map([]int{1, 2, 3}, func(_ int, x int) int { return x * x })
    evens := s.FilterMap([]int{1, 2, 3, 4}, func(_ int, x int) (int, bool) {
        return x, x%2 == 0
    })
    fmt.Println(squares, evens) // [1 4 9] [2 4]

    // Set operations
    a, b := []int{1, 2, 2, 3}, []int{2, 4}
    fmt.Println(s.DiffSet(a, b))         // e.g. [1 3]
    fmt.Println(s.IntersectSet(a, b))    // e.g. [2]
    fmt.Println(s.SymmetricDiffSet(a,b)) // e.g. [1 3 4]
    fmt.Println(s.UnionSet(a, b))        // e.g. [1 2 3 4]

    // Reverse
    fmt.Println(s.Reverse([]int{1,2,3})) // [3 2 1]
}
```

### Equality for complex types (Func variants)

```go
type User struct { ID int; Name string }

equalUser := func(a, b User) bool { return a.ID == b.ID }

u1 := []User{{1, "A"}, {2, "B"}, {2, "B2"}}
u2 := []User{{2, "B"}, {3, "C"}}

onlyInU1 := s.DiffSetFunc(u1, u2, equalUser)
both := s.IntersectSetFunc(u1, u2, equalUser)
all := s.UnionSetFunc(u1, u2, equalUser)
xor := s.SymmetricDiffSetFunc(u1, u2, equalUser)
_ = []any{onlyInU1, both, all, xor}
```

### Constraints

- `echo.RealNumber`: `uint/uint8/uint16/uint32/uint64/int/int8/int16/int32/int64/float32/float64`
- `echo.Number`: `RealNumber` plus `complex64/complex128`

### Notes

- Methods with `Set` suffix deduplicate and return in unspecified order; use `Func` variants to preserve first occurrence order.
- `Max`/`Min` require at least one element; empty slice will panic.
- `Add`/`Delete` return errors on index out of range.

### License

MIT


