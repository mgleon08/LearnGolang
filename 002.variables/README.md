# Variables 變數

- [Variables 變數](#variables-變數)
	- [Short hand declaration](#short-hand-declaration)
	- [可以 assign runtime 計算的值](#可以-assign-runtime-計算的值)
	- [Declared and Assign Error](#declared-and-assign-error)

golang 宣告變數會透過 `var` 並且能給定特定的型別，如果沒給予則會自動判定，也可以同時宣告並給予初始值

```go
package main

import "fmt"

func main() {
	// 宣告成 int
	var a1 int

	// 初始化同時宣告
	var a2 int = 10

	// a3 跟 b3 都是 int，沒有給值 int 預設是 0
	var a3, b3 int

	// 同時宣告一樣的 type 並給值
	var a4, b4 int = 100, 50

	// 自動推斷型別
	var a5 = 10

	// 同時宣告自動判斷 type (必須是同時給值)
	var a6, b6 = 0, "test"

	// 多個同時宣告和給值，可以用括號包再一起 (不能同行，不然會錯)
	var (
		a7 bool = false
		b7 int
		c7 = "hello"
	)

	fmt.Println("a1:", a1)
	fmt.Println("a2:", a2)
	fmt.Println("a3, b3:", a3, b3)
	fmt.Println("a4, b4:", a4, b4)
	fmt.Println("a5:", a5)
	fmt.Println("a6, b6:", a6, b6)
	fmt.Println("a7, b7, c7:", a7, b7, c7)

}
```

```go
a1: 0
a2: 10
a3, b3: 0 0
a4, b4: 100 50
a5: 10
a6, b6: 0 test
a7, b7, c7: false 0 hello
```

[Go Playground](https://go.dev/play/p/Tgb_x22FJaz)

## Short hand declaration

`:=` 是聲明並賦值，並且系統自動推斷類型，只能在 `main()` 裡面

```go
package main

import (
	"fmt"
)

func main() {
	// 單一宣告
	a1 := 0

	// 多個宣告
	a2, b2, c2 := 0, true, "test"

	a3, b3 := 1, 2
	b3, c3 := 3, 4 // 重複宣告，因為 c 是新的變數，因此可以通過

	a4, b4 := 1, 2
	a4, b4 = 3, 4 // 這邊是用 = 不是 := 因此是 assign 新的值

	a5 := float64(0) // 宣告並給值

	fmt.Println("a1:", a1)
	fmt.Println("a2:", a2, b2, c2)
	fmt.Println("a3, b3, c3:", a3, b3, c3)
	fmt.Println("a4, b4:", a4, b4)
	fmt.Println("a5", a5)

}
```

```go
a1: 0
a2: 0 true test
a3, b3, c3: 1 3 4
a4, b4: 3 4
a5 0
```

[Go Playground](https://go.dev/play/p/FLNU243YO9N)

## 可以 assign runtime 計算的值

```go
package main

import (
	"fmt"
	"math"
)

func main() {
	a, b := 1.1, 2.2
	c := math.Min(a, b)
	fmt.Println("Minimum value is", c)
}
```

```
Minimum value is 1.1
```

[Go Playground](https://go.dev/play/p/bJEV7W-lcW3)

## Declared and Assign Error

宣告後未使用會出現 error

```go
a declared but not used
```

golang 是強型別，宣告後就不能 assign 不同的 type，會出現 error

```go
var a int = 123
a = "test"

// cannot use "test" (untyped string constant) as int value in assignment

a5 := 123
a5 = "test"
// cannot use "test" (untyped string constant) as int value in assignment
```

重複宣告，因為沒有新的參數，因此會報錯

```go
package main

import "fmt"

func main() {
	a, b := 1, 2
	a, b := 3, 4 // 重複宣告，因為沒有新的參數，因此會報錯

	fmt.Println(a, b)
	// no new variables on left side of :=
}
```

[Go Playground](https://go.dev/play/p/__AbPvDJ1ks)
