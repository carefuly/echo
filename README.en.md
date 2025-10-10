## carefuly-echo

A Go utility library. This document focuses on these four packages:
- `tuple`: generic tuples (currently `Pair[K,V]`)
- `stringx`: zero-copy conversions between string and []byte (unsafe)
- `randx`: high-performance random string generation
- `bean`: generic utilities for the Option pattern

### Install

```bash
go get github.com/carefuly/carefuly-echo
```

### tuple

Import path: `github.com/carefuly/carefuly-echo/tuple/pair`

- **Type**: `Pair[K any, V any]`
  - Fields: `Key K`, `Value V`
  - Methods:
    - `func (p *Pair[K,V]) String() string`
    - `func (p *Pair[K,V]) Split() (K, V)`
  - Constructor: `func NewPair[K any, V any](key K, value V) Pair[K,V]`

Example:

```go
package main

import (
	"fmt"
	"github.com/carefuly/carefuly-echo/tuple/pair"
)

func main() {
	p := pair.NewPair("id", 123)
	fmt.Println(p.String()) // e.g.: <"id", 123>
	k, v := p.Split()
	fmt.Println(k, v)
}
```

### stringx (unsafe zero-copy)

Import path: `github.com/carefuly/carefuly-echo/stringx`

- `func UnsafeToBytes(val string) []byte`
  - Converts `string` to `[]byte` without allocation (zero-copy)
- `func UnsafeToString(val []byte) string`
  - Converts `[]byte` to `string` without allocation

Important notes:
- Ensure the source `string`/`[]byte` outlives the converted view; do not free or mutate the underlying storage
- Keep length/capacity assumptions consistent to avoid out-of-bounds access
- Do not mutate the converted result; backing storage may be shared

Example (use only if you fully understand the risks):

```go
package main

import (
	"fmt"
	"github.com/carefuly/carefuly-echo/stringx"
)

func main() {
	s := "hello"
	bs := stringx.UnsafeToBytes(s) // zero-copy view
	fmt.Println(len(bs), cap(bs))

	b := []byte{'a', 'b'}
	s2 := stringx.UnsafeToString(b) // zero-copy view
	fmt.Println(s2)
}
```

### randx (random strings)

Import path: `github.com/carefuly/carefuly-echo/randx`

- Types and charsets:
  - `type Type int`
  - Flags:
    - `TypeDigit` (digits)
    - `TypeLowerCase` (lowercase letters)
    - `TypeUpperCase` (uppercase letters)
    - `TypeSpecial` (special symbols)
    - `TypeMixed = TypeDigit | TypeUpperCase | TypeLowerCase | TypeSpecial`
  - Preset charsets: `CharsetDigit/LowerCase/UpperCase/Special`

- Generators:
  - `func RandCode(length int, typ Type) (string, error)`
    - Generates by type combination; `length < 0` errors; `length == 0` returns empty string
    - Errors if `typ` exceeds `TypeMixed`
  - `func RandStrByCharset(length int, charset string) (string, error)`
    - Custom charset; errors if empty

Implementation detail:
- Bit-cache extraction from `rand.Int63()` using masks to reduce RNG calls

Example:

```go
package main

import (
	"fmt"
	"github.com/carefuly/carefuly-echo/randx"
)

func main() {
	code, _ := randx.RandCode(8, randx.TypeMixed)
	fmt.Println(code)

	// Custom charset
	onlyHex, _ := randx.RandStrByCharset(16, "0123456789abcdef")
	fmt.Println(onlyHex)
}
```

Errors:
- `echo:长度必须大于等于0` (length < 0)
- `echo:不支持的类型` (unsupported type or empty charset)

### bean (Option pattern)

Import path: `github.com/carefuly/carefuly-echo/bean/option`

- `type Option[T any] func(*T)`
- `func Apply[T any](t *T, opts ...Option[T])`
- `type OptionErr[T any] func(*T) error`
- `func ApplyErr[T any](t *T, opts ...OptionErr[T]) error`

Example:

```go
package main

import (
	"errors"
	"fmt"
	"github.com/carefuly/carefuly-echo/bean/option"
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
	// Without errors
	s := Server{}
	option.Apply(&s, WithAddr(":8080"), WithTLS(true))
	fmt.Printf("%+v\n", s)

	// With errors
	s2 := Server{}
	_ = option.ApplyErr(&s2, WithAddrNonEmpty(":9090"))
	fmt.Printf("%+v\n", s2)
}
```

### License

MIT


