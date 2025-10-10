## careful-echo

Go 通用工具库。本文档聚焦以下四个包：
- `tuple`：泛型元组（当前提供 `Pair[K,V]`）
- `stringx`：字符串/字节零拷贝转换（不安全）
- `randx`：高性能随机字符串生成
- `bean`：Option 模式的泛型工具

### 安装

```bash
go get github.com/carefuly/careful-echo
```

### tuple（元组）

包路径：`github.com/carefuly/careful-echo/tuple/pair`

- **类型**：`Pair[K any, V any]`
  - 字段：`Key K`，`Value V`
  - 方法：
    - `func (p *Pair[K,V]) String() string`
    - `func (p *Pair[K,V]) Split() (K, V)`
  - 构造：`func NewPair[K any, V any](key K, value V) Pair[K,V]`

示例：

```go
package main

import (
	"fmt"
	"github.com/carefuly/careful-echo/tuple/pair"
)

func main() {
	p := pair.NewPair("id", 123)
	fmt.Println(p.String()) // 形如：<"id", 123>
	k, v := p.Split()
	fmt.Println(k, v)
}
```

### stringx（不安全零拷贝转换）

包路径：`github.com/carefuly/careful-echo/stringx`

- `func UnsafeToBytes(val string) []byte`
  - 将 `string` 转为 `[]byte`，零拷贝，不会分配新内存
- `func UnsafeToString(val []byte) string`
  - 将 `[]byte` 转为 `string`，零拷贝

重要注意：
- 确保参与转换的 `string` 与 `[]byte` 生命周期足够长；不要在转换后释放或修改其底层存储
- 确保长度/容量一致，避免越界
- 不要修改转换后的视图（它们可能共享同一块内存）

示例（仅在明确理解风险时使用）：

```go
package main

import (
	"fmt"
	"github.com/carefuly/careful-echo/stringx"
)

func main() {
	s := "hello"
	bs := stringx.UnsafeToBytes(s) // 零拷贝视图
	fmt.Println(len(bs), cap(bs))

	b := []byte{'a', 'b'}
	s2 := stringx.UnsafeToString(b) // 零拷贝视图
	fmt.Println(s2)
}
```

### randx（随机字符串）

包路径：`github.com/carefuly/careful-echo/randx`

- 类型与字符集：
  - `type Type int`
  - 标志位：
    - `TypeDigit`（数字）
    - `TypeLowerCase`（小写字母）
    - `TypeUpperCase`（大写字母）
    - `TypeSpecial`（特殊符号）
    - `TypeMixed = TypeDigit | TypeUpperCase | TypeLowerCase | TypeSpecial`
  - 预置字符集：`CharsetDigit/LowerCase/UpperCase/Special`

- 生成函数：
  - `func RandCode(length int, typ Type) (string, error)`
    - 基于类型组合快速生成；`length < 0` 返回错误；`length == 0` 返回空串
    - `typ` 超过 `TypeMixed` 返回错误
  - `func RandStrByCharset(length int, charset string) (string, error)`
    - 自定义字符集；`charset` 为空返回错误

实现细节：
- 使用位段缓存（按位掩码从 `rand.Int63()` 中多次取用），减少随机源调用次数

示例：

```go
package main

import (
	"fmt"
	"github.com/carefuly/careful-echo/randx"
)

func main() {
	code, _ := randx.RandCode(8, randx.TypeMixed)
	fmt.Println(code)

	// 自定义字符集
	onlyHex, _ := randx.RandStrByCharset(16, "0123456789abcdef")
	fmt.Println(onlyHex)
}
```

错误：
- `echo:长度必须大于等于0`（length < 0）
- `echo:不支持的类型`（`typ` 超范围或字符集为空）

### bean（Option 模式）

包路径：`github.com/carefuly/careful-echo/bean/option`

- `type Option[T any] func(*T)`：无错误的配置项
- `func Apply[T any](t *T, opts ...Option[T])`：依次应用所有 `Option`
- `type OptionErr[T any] func(*T) error`：带错误的配置项
- `func ApplyErr[T any](t *T, opts ...OptionErr[T]) error`：遇错即停并返回

示例：

```go
package main

import (
	"errors"
	"fmt"
	"github.com/carefuly/careful-echo/bean/option"
)

type Server struct {
	Addr string
	TLS  bool
}

func WithAddr(addr string) option.Option[Server] {
	return func(s *Server) { s.Addr = addr }
}

func WithTLS(enabled bool) option.Option[Server] {
	return func(s *Server) { s.TLS = enabled }
}

func WithAddrNonEmpty(addr string) option.OptionErr[Server] {
	return func(s *Server) error {
		if addr == "" { return errors.New("addr empty") }
		s.Addr = addr
		return nil
	}
}

func main() {
	// 无错版本
	s := Server{}
	option.Apply(&s, WithAddr(":8080"), WithTLS(true))
	fmt.Printf("%+v\n", s)

	// 带错版本
	s2 := Server{}
	_ = option.ApplyErr(&s2, WithAddrNonEmpty(":9090"))
	fmt.Printf("%+v\n", s2)
}
```

### 许可证

MIT


