# Functions

- [Functions](#functions)
	- [函式 function](#函式-function)
	- [不定參數函式 Variadic Function](#不定參數函式-variadic-function)
	- [匿名函式 Anonymous Function](#匿名函式-anonymous-function)
	- [Defined Function Types](#defined-function-types)
	- [高階函式 Higher Order Function](#高階函式-higher-order-function)
	- [閉包函式 Closure Function](#閉包函式-closure-function)
	- [初始化函式 Init Function](#初始化函式-init-function)
- [斐波那契数 fibonacci](#斐波那契数-fibonacci)

## 函式 function

```go
package main

import "fmt"

// 基本 function
func foo1(name string) {
	fmt.Println("1. Hi " + name)
}

// 多個傳入值
func foo2(name1 string, name2 string) {
	fmt.Println("2. Hi " + name1 + ", " + name2)
}

// 傳入值同個形態，可只寫一個
func foo3(name1, name2 string) {
	fmt.Println("3. Hi " + name1 + ", " + name2)
}

// return 值的型態定義在後面
func foo4(name string) string {
	var str = "Hi " + name
	return str
}

// 可以直接在 func 的回傳區塊命名回傳變數
func foo5(name string) (str string) {
	str = "Hi " + name
	return
}

// 多重return
func foo6(x, y int) (int, int) {
	return x + y, x - y
}

// 多個傳入值 ... (當會有不確定的個數傳入值，有兩個變數，不確定的要放置在最後)
func foo7(x ...int) int {
	// 收到 Type 為 []int
	var t int
	for _, n := range x {
		t += n
	}
	return t
}

// 傳入 slice, slice 是類似矩陣的東西
func foo8(x []int) int {
	var t int
	for _, n := range x {
		t += n
	}
	return t
}

// function 回傳 function，function 可以當變數，也可以用來回傳
func foo9() func() string {
	return func() string {
		return "foo9"
	}
}

func foo10() (string, string) {
	return "foo10", "bar10"
}

func main() {
	// 自動判別型別，必須宣告再 main 裡面
	foo11 := func(name string) {
		fmt.Println("11. Hi " + name)
	}

	foo1("foo1")
	foo2("foo2", "bar2")
	foo3("foo3", "bar3")
	fmt.Println("4.", foo4("foo4"))
	fmt.Println("5.", foo5("foo5"))
	a6, b6 := foo6(1, 2)
	fmt.Println("6.", a6, b6)
	fmt.Println("7.", foo7(1, 2, 3))
	nums := []int{1, 2, 3, 4}
	fmt.Println("8.", foo8(nums))
	fmt.Println("9.", foo9()())
	// function 回傳兩個值，不需要可以用 _ 取代
	a10, _ := foo10()
	fmt.Println("10.", a10)
	foo11("foo11")
}
```

```go
1. Hi foo1
2. Hi foo2, bar2
3. Hi foo3, bar3
4. Hi foo4
5. Hi foo5
6. 3 -1
7. 6
8. 10
9. foo9
10. foo10
11. Hi foo11
```

[Go Playground](https://go.dev/play/p/7vZO_sGWOmn)

## 不定參數函式 Variadic Function

當有不確定的參數時，可用 `...` 來代替 (會自動將後面的值，轉成 slice，並且根據 type)

> 只能夠放在最後一個參數

```go
package main

import (
	"fmt"
)

func find(num int, nums ...int) {
	fmt.Printf("type of nums is %T\n", nums)
	fmt.Println("nums is", nums)

}
func main() {
	find(1, 2, 3, 4)
	nums := []int{5, 6, 7}
	find(8, nums...) // 必須加上 ... 否則會造成 type error
}
```

```go
type of nums is []int
nums is [2 3 4]
type of nums is []int
nums is [5 6 7]
```

[Go Playground](https://go.dev/play/p/6aVta_0xfWd)

## 匿名函式 Anonymous Function

常用在只使用一次的情況

```go
package main

import (
	"fmt"
)

func main() {
	a := func() {
		fmt.Println("Anonymous function 1")
	}
	a()

	func() {
		fmt.Println("Anonymous function 2")
	}()

	func(n string) {
		fmt.Println("Anonymous", n)
	}("function 3")
}
```

```go
Anonymous function 1
Anonymous function 2
Anonymous function 3
```

[Go Playground](https://go.dev/play/p/5uITx5Fy-gz)

## Defined Function Types

像定義 struct type 一樣，也可以定義 func type

```go
type add func(a int, b int) int
```

```go
package main

import (
	"fmt"
)

// 定義一個 type 是 func，並且裡面要給的參數和 return 的值都已經確認
type add func(a int, b int) int

func main() {
	var a add = func(a int, b int) int {
		return a + b
	}
	s := a(5, 6)
	fmt.Println("Sum", s)
}
```

```go
Sum 11
```

[Go Playground](https://go.dev/play/p/GglgP2cZxf0)

## 高階函式 Higher Order Function

滿足其中之一即可

* 能夠接收函式作為參數（這個被接收的函式稱為回調函式）
* 將函式作為返回值輸出的函式

接收函式

```go
package main

import (
	"fmt"
)

func simple(a func(a, b int) int) {
	fmt.Println(a(2, 6))
}

func main() {
	f := func(a, b int) int {
		return a + b
	}
	simple(f)
}
```

```go
8
```

[Go Playground](https://go.dev/play/p/seSsLCA3vdt)

返回函式

```go
package main

import (
	"fmt"
)

func simple() func(a, b int) int {
	f := func(a, b int) int {
		return a + b
	}
	return f
}

func main() {
	s := simple()
	fmt.Println(s(2, 6))
}
```

```go
8
```

[Go Playground](https://go.dev/play/p/usDtowfxwNT)

## 閉包函式 Closure Function

即使函式已經執行結束，其內部的變數卻並未跟著消除，還能繼續被呼叫；這種能將外層變數「包」在內層暫存、使用的方式，就是所謂的「閉包」。

```go
package main

import (
	"fmt"
)

// return func(int) int
func closure(x int) func(int) int {
	fmt.Println("in closure", x, &x)
	return func(y int) int {
		fmt.Println("closure return function 1", x, &x)
		fmt.Println("closure return function 2", y, &y)
		return x + y
	}
}

func main() {
	f := closure(10)

	fmt.Println(f(1))
	fmt.Println(f(2))
}
```

```go
in closure 10 0xc00001c030
closure return function 1 10 0xc00001c030
closure return function 2 1 0xc00001c038
11
closure return function 1 10 0xc00001c030
closure return function 2 2 0xc00001c050
12
```

[Go Playground](https://go.dev/play/p/FafyLrL9njo)

```go
package main

import (
	"fmt"
)

func appendStr() func(string) string {
	text := "Hello"
	returnFunction := func(name string) string {
		text = text + " " + name
		return text
	}
	return returnFunction
}

func main() {
	a := appendStr()
	b := appendStr()

	fmt.Println(a("foo"))
	fmt.Println(b("bar"))

	fmt.Println(a("bar"))
	fmt.Println(b("foo"))
}
```

```go
Hello foo
Hello bar
Hello foo bar
Hello bar foo
```

[Go Playground](https://go.dev/play/p/qWmpqVxVvit)

## 初始化函式 Init Function

* import Package 時 `init()` 會在一開始的時候被呼叫
* 只有 `init()` 可以多次宣告
* `init` 不會有任何參數和回傳值
* 當 package 只要執行 `init()` 可以使用 `_`

```go
import _ "fmt"
```

執行順序會是

1. Package Level Variables
2. Init Function
3. Main Function

```go
package main

import "fmt"

var a, b int = 1, 2

func init() {
	fmt.Println("Init", a)
}

func init() {
	fmt.Println("Init", b)
}

func main() {
	fmt.Println("Main")
}
```

```go
Init 1
Init 2
Main
```

[Go Playground](https://go.dev/play/p/A0vzgoJuYR9)

# 斐波那契数 fibonacci

```go
package main

import (
	"fmt"
)

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	x, y := 0, 1
	fmt.Println("test1 x=", x, "y=", y)
	return func() int {
		fmt.Println("test2 x=", x, "y=", y)
		x, y = y, x+y
		fmt.Println("test3 x=", x, "y=", y)
		return x
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
```

```go
test1 x= 0 y= 1
test2 x= 0 y= 1
test3 x= 1 y= 1
1
test2 x= 1 y= 1
test3 x= 1 y= 2
1
test2 x= 1 y= 2
test3 x= 2 y= 3
2
test2 x= 2 y= 3
test3 x= 3 y= 5
3
test2 x= 3 y= 5
test3 x= 5 y= 8
5
test2 x= 5 y= 8
test3 x= 8 y= 13
8
test2 x= 8 y= 13
test3 x= 13 y= 21
13
test2 x= 13 y= 21
test3 x= 21 y= 34
21
test2 x= 21 y= 34
test3 x= 34 y= 55
34
test2 x= 34 y= 55
test3 x= 55 y= 89
55
```

[Go Playground](https://go.dev/play/p/54BE3GHNjcF)
