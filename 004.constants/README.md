# Constants 常數

- [Constants 常數](#constants-常數)
	- [宣告](#宣告)
	- [constants 宣告後就不能 reassign (immutable)](#constants-宣告後就不能-reassign-immutable)
	- [不能 assign variable 給 constants](#不能-assign-variable-給-constants)
	- [const 沒有特別聲明是哪種類型時，是屬於 Untyped](#const-沒有特別聲明是哪種類型時是屬於-untyped)
	- [多重宣告 const，後面沒有給值，會跟前面的一樣](#多重宣告-const後面沒有給值會跟前面的一樣)
		- [IOTA 常量的計數器 const number generator](#iota-常量的計數器-const-number-generator)
	- [constants 在編譯的時候就要定義，不能在執行時才給值](#constants-在編譯的時候就要定義不能在執行時才給值)

## 宣告

constants 常用於用固定常數，必須一開始就 initialized

```go
package main

import "fmt"

func main() {
	const a1 = 123
	const a2 string = "Test"
	const (
		a3 bool = false
		b3 int  = 321
		c3      = "hello"
	)

	fmt.Println("a1:", a1)
	fmt.Println("a2:", a2)
	fmt.Println("a3, b3, c3:", a3, b3, c3)

}
```

```go
a1: 123
a2: Test
a3, b3, c3: false 321 hello
```

[Go Playground](https://go.dev/play/p/hz6OEB2IpX3)

## constants 宣告後就不能 reassign (immutable)

```go
package main

func main() {
	const a = 1
	a = 2
}
```

```go
cannot assign to a (untyped int constant 1)
```

[Go Playground](https://go.dev/play/p/bze4zg_i7R7)

## 不能 assign variable 給 constants

```go
package main

import "fmt"

func main() {
	a := 1
	const b = a

	fmt.Println(b)
}
```

```go
a (variable of type int) is not constant
```

[Go Playground](https://go.dev/play/p/l94Ee6NyYGI)

## const 沒有特別聲明是哪種類型時，是屬於 Untyped

當 const assign 給任意 type 的值時，會由 Untyped 會自動轉變為該 type

而 var 一開始就會給定 type，因此會 error

```go
package main

import (
	"fmt"
)

func main() {
	const a = 100 // Untyped

	var intVar int = a             // var a = int(min)
	var int32Var int32 = a         // var a = int32(min)
	var byteVar byte = a           // var a = byte(min)
	var runeVar rune = a           // var a = rune(min)
	var float64Var float64 = a     // var a = float64(min)
	var complex64Var complex64 = a // var a = complex64(min)

	fmt.Printf("intVar: %v, %T \n", intVar, intVar)
	fmt.Printf("int32Var: %v, %T \n", int32Var, int32Var)
	fmt.Printf("byteVar: %v, %T \n", byteVar, byteVar)
	fmt.Printf("runeVar: %v, %T \n", runeVar, runeVar)
	fmt.Printf("float64Var: %v, %T \n", float64Var, float64Var)
	fmt.Printf("complex64Var: %v, %T \n", complex64Var, complex64Var)
}
```

```go
intVar: 100, int
int32Var: 100, int32
byteVar: 100, uint8
runeVar: 100, int32
float64Var: 100, float64
complex64Var: (100+0i), complex64
```

[Go Playground](https://go.dev/play/p/sqGuzpir-iX)

## 多重宣告 const，後面沒有給值，會跟前面的一樣

```go
package main

import (
	"fmt"
)

func main() {
	const (
		a int = 123
		b
	)
	fmt.Println(a, b)
}
```

```go
123 123
```

[Go Playground](https://go.dev/play/p/9nI8b6GN6bz)

### IOTA 常量的計數器 const number generator

遞增

```go
package main

import (
	"fmt"
)

func main() {
	const (
		a = iota
		b
		_
		c
		d
	)
	fmt.Println(a, b, c, d)
}
```

```go
0 1 3 4
```

[Go Playground](https://go.dev/play/p/awChwvv8QCd)

遞減

```go
package main

import (
	"fmt"
)

func main() {
	const (
		a = 9 - iota
		b
		_
		c
		d
	)
	fmt.Println(a, b, c, d)
}
```

```go
9 8 6 5
```

[Go Playground](https://go.dev/play/p/EJZy-iOUXAo)

Reference: [Iota](https://github.com/golang/go/wiki/Iota)

## constants 在編譯的時候就要定義，不能在執行時才給值

```go
package main

import (
	"fmt"
	"math"
)

func main() {
	// 因為 math.Sqrt(4) 是在 runtime 才會知道值，所以會 error
	const a = math.Sqrt(4)

	fmt.Println(a)
}
```

```go
math.Sqrt(4) (value of type float64) is not constant
```

[Go Playground](https://go.dev/play/p/xM5ZP4EXrcg)

const 可以在 compile time 時檢查出來

```go
package main

import (
	"fmt"
)

func main() {
	const (
		max int = 5
		min int = 0
	)
	fmt.Println(max / min)
}
```

```go
invalid operation: division by zero
```

[Go Playground](https://go.dev/play/p/PigqrP8fDjU)

runtime 的 error 無法在 compile-time 時檢查出來

```go
package main

import (
	"fmt"
)

func main() {
	max, min := 5, 0

	fmt.Println(max / min)
}
```

```go
panic: runtime error: integer divide by zero
```

[Go Playground](https://go.dev/play/p/jII2Y5ySONb)
