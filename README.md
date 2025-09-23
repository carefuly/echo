## carefuly-echo

Go 通用工具库。本文档聚焦 `list` 包（通用 List 接口与三种实现）：
- `ArrayList[T]`：基于切片的高性能可变数组，支持自动缩容
- `ConcurrentList[T]`：为任意 `List[T]` 提供读写锁包装，线程安全
- `CopyOnWriteArrayList[T]`：写时复制，读不加锁，适合读多写少

### 安装

```bash
go get github.com/carefuly/carefuly-echo
```

### 引用

```go
import (
	"github.com/carefuly/carefuly-echo/list"
)
```

### List 接口

`list` 包对外暴露统一接口 `List[T]`（下述为精简说明）：

```go
// 仅示意：索引越界返回错误，AsSlice 每次返回新切片
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

### 实现与特性

- ArrayList[T]
  - 基于切片的封装，追加/插入/删除均保持与原生切片一致的语义（越界报错）
  - 删除后自动缩容（避免长期占用过大容量）：
    - 容量 > 2048 且长度 < 容量一半：缩容至原容量的 5/8（向下取整）
    - 64 < 容量 ≤ 2048 且长度 ≤ 容量的 1/4：缩容至原容量的一半
    - 容量 ≤ 64：不缩容
  - `AsSlice` 每次返回全新切片，永不返回 `nil`

- ConcurrentList[T]
  - 为任意 `List[T]` 提供线程安全包装
  - 读操作（`Get/Len/Cap/Range/AsSlice`）加读锁，写操作（`Append/Add/Set/Delete`）加写锁

- CopyOnWriteArrayList[T]
  - 写时复制：写操作复制底层数组再修改，读操作不加锁
  - 适合读多写少、遍历频繁的场景
  - `AsSlice` 返回副本；`Delete` 不做“缩容”，每次分配恰当长度

### 快速上手

```go
package main

import (
	"fmt"
	"github.com/carefuly/carefuly-echo/list"
)

func main() {
	// 1) ArrayList：常规可变数组
	a := list.NewArrayList[int](0)
	_ = a.Append(1, 2, 3)          // [1 2 3]
	_ = a.Add(1, 9)                // [1 9 2 3]
	_ = a.Set(2, 5)                // [1 9 5 3]
	v, _ := a.Delete(1)            // v=9, a=[1 5 3]，必要时触发缩容
	fmt.Println(v, a.AsSlice())

	// 2) ConcurrentList：线程安全包装
	cl := &list.ConcurrentList[int]{List: a}
	_ = cl.Append(100)
	fmt.Println(cl.Len(), cl.AsSlice())

	// 3) CopyOnWriteArrayList：读多写少
	cow := list.NewCopyOnWriteArrayList[int]()
	_ = cow.Append(7, 8, 9)
	_ = cow.Add(1, 88)
	fmt.Println(cow.AsSlice())
}
```

### 错误与边界

- 所有按下标访问/修改/删除的方法在越界时返回错误（错误文案示例：`echo: 下标超出范围，长度 %d, 下标 %d`）
- `Range` 迭代时：当回调返回错误即中断并返回该错误
- `AsSlice`：无元素时返回长度、容量均为 0 的新切片（非 `nil`）

### 选型建议

- 纯本地单线程或轻并发读写：优先 `ArrayList[T]`
- 多并发读写：用 `ConcurrentList[T]{List: yourList}` 包装
- 读多写少、写入频率低：考虑 `CopyOnWriteArrayList[T]`

### 测试概览

- 覆盖 `Get/Append/Add/Set/Delete/Len/Cap/Range/AsSlice` 行为
- `ArrayList` 缩容规则包含逻辑用例与边界用例（如 64、2048、2049 等容量阈值）

### 许可证

MIT


