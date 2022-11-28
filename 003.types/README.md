
# Types 類型

- [Types 類型](#types-類型)
	- [基本類型](#基本類型)
	- [Type Alias](#type-alias)
	- [數字儲存範圍 Signed & Unsigned integers](#數字儲存範圍-signed--unsigned-integers)
	- [轉型](#轉型)
		- [透過 strconv 套件來轉成不同的 type](#透過-strconv-套件來轉成不同的-type)


## 基本類型

golang 內建的類別

> `*` 為比較常使用

```go
- *string
- *bool
- Numeric Types
  - *int  int8  int16  *int32(rune)  int64
  - uint *uint8(byte) uint16 uint32 uint64
  - float32 *float64
  - complex64 complex128
  - byte // alias for uint8
  - rune // alias for int32，represents a Unicode code point
```

## Type Alias

可以建立 alias type，並且建立後的 alias 不能跟原本的 type 去 assign


```go
package main

import "fmt"

func main() {
	var a = "test"

	type newString string
	var b newString = "test"
	b = a

	fmt.Println(b)
}
```

```
cannot use a (variable of type string) as type newString in assignment
```

[Go Playground](https://go.dev/play/p/OVPIr2GfI8-)

## 數字儲存範圍 Signed & Unsigned integers

`int` 會根據底層是 32 bit(4 bytes) or 64 bit(8 bytes) 來表現，以下面範例來說就是 64 bit

```go
package main

import (
	"fmt"
	"math"
	"reflect"
)

func main() {
	fmt.Printf("uint8  : 0 ~ %d\n", math.MaxUint8)
	fmt.Printf("uint16 : 0 ~ %d\n", math.MaxUint16)
	fmt.Printf("uint32 : 0 ~ %d\n", math.MaxUint32)
	fmt.Printf("uint64 : 0 ~ %d\n", uint64(math.MaxUint64))
	fmt.Printf("int    : %d ~ %d\n", math.MinInt, math.MaxInt)
	fmt.Printf("int8   : %d ~ %d\n", math.MinInt8, math.MaxInt8)
	fmt.Printf("int16  : %d ~ %d\n", math.MinInt16, math.MaxInt16)
	fmt.Printf("int32  : %d ~ %d\n", math.MinInt32, math.MaxInt32)
	fmt.Printf("int64  : %d ~ %d\n", math.MinInt64, math.MaxInt64)
	fmt.Printf("整數預設型態: %s\n", reflect.TypeOf(1))
	fmt.Printf("浮點數預設型態: %s\n", reflect.TypeOf(1.8))
}
```

```go
uint8  : 0 ~ 255
uint16 : 0 ~ 65535
uint32 : 0 ~ 4294967295
uint64 : 0 ~ 18446744073709551615
int    : -9223372036854775808 ~ 9223372036854775807
int8   : -128 ~ 127
int16  : -32768 ~ 32767
int32  : -2147483648 ~ 2147483647
int64  : -9223372036854775808 ~ 9223372036854775807
整數預設型態: int
浮點數預設型態: float64
```

[Go Playground](https://go.dev/play/p/tiXTXJoLwdw)

## 轉型

因為 golang 是強型別，因此在做 assign 或是兩個數字相加時，必須先轉型

```go
package main

import (
	"fmt"
)

func main() {
	a := 1
	b := 1.8
	sum := a + int(b)
	// 如果沒透過 int()，會出現 error
	// invalid operation: a + b (mismatched types int and float64)

	fmt.Println(sum)
}
```

```
2
```

[Go Playground](https://go.dev/play/p/zgGCAdFCPLW)

### 透過 strconv 套件來轉成不同的 type

```go
package main

import (
	"fmt"
	"strconv"
)

func main() {
	b, errB := strconv.ParseBool("true")
	f, errF := strconv.ParseFloat("3.1415", 64)
	i, errI := strconv.ParseInt("-42", 10, 64)
	u, errU := strconv.ParseUint("42", 10, 64)

	fmt.Printf("b Type: %T, Value: %v, Err: %v\n", b, b, errB)
	fmt.Printf("f Type: %T, Value: %v, Err: %v\n", f, f, errF)
	fmt.Printf("i Type: %T, Value: %v, Err: %v\n", i, i, errI)
	fmt.Printf("u Type: %T, Value: %v, Err: %v\n", u, u, errU)
}
```

```go
b Type: bool, Value: true, Err: <nil>
f Type: float64, Value: 3.1415, Err: <nil>
i Type: int64, Value: -42, Err: <nil>
u Type: uint64, Value: 42, Err: <nil>
```

[Go Playground](https://go.dev/play/p/l2AkGuB4y0p)

Reference: [strconv](https://pkg.go.dev/strconv)
