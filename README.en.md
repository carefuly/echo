## carefuly-echo

A Go utility library. This document focuses on the `list` package (a generic List interface and three implementations):
- `ArrayList[T]`: slice-backed dynamic array with optimized memory management and shrink policy
- `ConcurrentList[T]`: thread-safe wrapper with RWMutex for any `List[T]`
- `CopyOnWriteArrayList[T]`: copy-on-write semantics, lock-free reads; ideal for read-heavy workloads

### Install

```bash
go get github.com/carefuly/carefuly-echo
```

### Import

```go
import (
	"github.com/carefuly/carefuly-echo/list"
)
```

### List interface

Unified API exposed by `list` (simplified description below):

```go
// Notes: bounds-checked; AsSlice returns a fresh slice every call and never nil
Get(index int) (T, error)
Append(ts ...T) error
Add(index int, t T) error
Set(index int, t T) error
Delete(index int) (T, error)
Len() int
Cap() int
Range(fn func(index int, t T) error) error
AsSlice() []T
```

### Implementations & traits

- ArrayList[T]
  - Wrapper over a Go slice; append/insert/delete with bounds checking
  - Shrink policy after deletions to avoid over-allocations:
    - If capacity > 2048 and length < 1/2 capacity: shrink to 5/8 of capacity (floor)
    - If 64 < capacity ≤ 2048 and length ≤ 1/4 capacity: shrink to 1/2 of capacity
    - If capacity ≤ 64: no shrinking
  - `AsSlice` returns a newly allocated slice; never returns `nil`

- ConcurrentList[T]
  - Adds thread-safety to any `List[T]`
  - Read methods (`Get/Len/Cap/Range/AsSlice`) use RLock; write methods (`Append/Add/Set/Delete`) use Lock

- CopyOnWriteArrayList[T]
  - Copy-on-write on mutations; reads are lock-free
  - Best for read-heavy, infrequent writes and iteration intensive scenarios
  - `AsSlice` returns a copy; `Delete` always reallocates to exact length (no shrink policy)

### Quick start

```go
package main

import (
	"fmt"
	"github.com/carefuly/carefuly-echo/list"
)

func main() {
	// 1) ArrayList: general-purpose dynamic array
	a := list.NewArrayList[int](0)
	_ = a.Append(1, 2, 3)          // [1 2 3]
	_ = a.Add(1, 9)                // [1 9 2 3]
	_ = a.Set(2, 5)                // [1 9 5 3]
	v, _ := a.Delete(1)            // v=9, a=[1 5 3]; may shrink capacity
	fmt.Println(v, a.AsSlice())

	// 2) ConcurrentList: make it thread-safe
	cl := &list.ConcurrentList[int]{List: a}
	_ = cl.Append(100)
	fmt.Println(cl.Len(), cl.AsSlice())

	// 3) CopyOnWriteArrayList: read-heavy workloads
	cow := list.NewCopyOnWriteArrayList[int]()
	_ = cow.Append(7, 8, 9)
	_ = cow.Add(1, 88)
	fmt.Println(cow.AsSlice())
}
```

### Errors & edge cases

- Indexing methods validate bounds and return errors (e.g., `echo: 下标超出范围，长度 %d, 下标 %d`)
- `Range` stops early if the callback returns an error and bubbles it up
- `AsSlice` returns an empty (length 0, capacity 0) slice when no elements are present, never `nil`

### Guidance

- Single-threaded or light contention: use `ArrayList[T]`
- Concurrent read/write: wrap with `ConcurrentList[T]{List: yourList}`
- Read-mostly workloads: prefer `CopyOnWriteArrayList[T]`

### Tests

- Cover `Get/Append/Add/Set/Delete/Len/Cap/Range/AsSlice`
- Shrink rules of `ArrayList` are validated with both logical and boundary cases (thresholds: 64, 2048, 2049, etc.)

### License

MIT


