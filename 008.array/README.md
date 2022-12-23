# Array 陣列

- [Array 陣列](#array-陣列)
	- [Array](#array)
		- [Arrays are value types](#arrays-are-value-types)
		- [Length of an array](#length-of-an-array)
	- [Iterating](#iterating)
		- [for](#for)
		- [range](#range)
	- [Slices](#slices)
		- [建立 slice](#建立-slice)
		- [nil with slice](#nil-with-slice)
		- [length and capacity](#length-and-capacity)
		- [Slices of slices 二維切片](#slices-of-slices-二維切片)
		- [Append](#append)
		- [透過 make 來建立 slice (dynamically-sized arrays)](#透過-make-來建立-slice-dynamically-sized-arrays)
		- [Memory Optimisation](#memory-optimisation)
		- [copy](#copy)
	- [Variadic Functions](#variadic-functions)
		- [example 1](#example-1)
		- [example 2](#example-2)

## Array

* Array 的長度是其類型的一部分，因此陣列不能改變大小。
* 當定義為 int Array，之後就只能放 int，不能放其他 type。
* int default 值是 0, string 是 nil。
* 當設定一定長度時，內容不一定要填滿，但不能超過當初定義的長度。
* size 不一樣的 Array 是不同的 Type。`[1]int` != `[2]int`
* 透過 `...` 可以在編譯的時候決定長度。

```go
package main

import "fmt"

func main() {
	// 宣告一個變數 a 為一個 int type 的 Array 並且長度只有10.(int 預設是 0, string 則為 nil)
	var a1 [10]int

	// 宣告變數 a2 長度為 3，並且第一個為 1
	a2 := [3]int{1}

	// 宣告一個變數 a3 為 string type, 並且裡面的值是 A, B 並且長度只有 2
	a3 := [2]string{"A", "B"}
	a4 := [...]string{"A", "B"} // ... 編譯的時候才決定長度

	// 二維陣列
	a5 := [3][2]string{
		{"lion", "tiger"},
		{"cat", "dog"},
		{"pigeon", "peacock"}, // 最後一行必須要有 ,
	}

	// value 塞 struct
	a6 := [6]struct {
		i int
		b bool
	}{
		{2, true},
		{3, false},
		{5, true},
		{7, true},
		{11, false},
		{13, true},
	}

	// 透過 0:, 1: 可指定 value 要塞的位置
	a7 := [3]int{
		0: 111,
		1: 222,
		2: 333,
	}

	a8 := [...]int{
		3: 333,
	}

	a9 := [...]int{
		5: 111,
		222,
		0: 333,
	}

	fmt.Println("a1:", a1)
	fmt.Println("a2:", a2)
	fmt.Println("a3:", a3)
	fmt.Println("a4", a4)
	fmt.Println("a5:", a5)
	fmt.Println("a6:", a6)
	fmt.Println("a7:", a7)
	fmt.Println("a8:", a8)
	fmt.Println("a9:", a9)
}
```

```go
a1: [0 0 0 0 0 0 0 0 0 0]
a2: [1 0 0]
a3: [A B]
a4 [A B]
a5: [[lion tiger] [cat dog] [pigeon peacock]]
a6: [{2 true} {3 false} {5 true} {7 true} {11 false} {13 true}]
a7: [111 222 333]
a8: [0 0 0 333]
a9: [333 0 0 0 0 111 222]
```

[Go Playground](https://go.dev/play/p/D4lw2MUDigI)

### Arrays are value types

Array assign 給另一個變數時，會是建立新的 value，不是 reference

```go
package main

import "fmt"

func main() {
	a := [...]string{"FooA", "barA"}
	b := a // a copy of a is assigned to b
	b[0] = "FooB"
	fmt.Println("a is", a)
	fmt.Println("b is", b)
}
```

```go
a is [FooA barA]
b is [FooB barA]
```

[Go Playground](https://go.dev/play/p/-9mbyS2b5Rv)

將 array 丟入 function 時，也是一個新的 array，不會影響原本的

```go
package main

import "fmt"

func changeLocal(num [2]int) {
	num[0] = 3
	fmt.Println("inside", num)

}
func main() {
	num := [...]int{1, 2}
	fmt.Println("before", num)
	changeLocal(num)
	fmt.Println("after", num)
}
```

```go
before [1 2]
inside [3 2]
after [1 2]
```

[Go Playground](https://go.dev/play/p/X_yMgLYeI5q)

### Length of an array

透過 `len` 取得 array 長度

```go
package main

import "fmt"

func main() {
	a := [...]int{1, 2, 3}
	fmt.Println("length of a is", len(a))

}
```

```go
length of a is 3
```

[Go Playground](https://go.dev/play/p/0Sri_MYTel0)

## Iterating

### for

```go
package main

import "fmt"

func main() {
	a := [...]int{9, 8, 7}
	for i := 0; i < len(a); i++ { //looping from 0 to the length of the array
		fmt.Printf("%d th element of a is %d\n", i, a[i])
	}
}
```

```go
0 th element of a is 9
1 th element of a is 8
2 th element of a is 7
```

[Go Playground](https://go.dev/play/p/yGSsOnatsHO)

### range

可以達成 foreach 的方式

```go
package main

import (
	"fmt"
)

func main() {
	data := []string{"a", "b", "c"}

	for index, value := range data {
		fmt.Printf("%d%s|", index, value) // 輸出：0a|1b|2c|
	}

	for index := range data {
		fmt.Printf("%d|", index) // 輸出：0|1|2|
	}

	for _, value := range data {
		fmt.Printf("%s|", value) // 輸出：a|b|c|
	}
}
```

```go
0a|1b|2c|0|1|2|a|b|c|
```

[Go Playground](https://go.dev/play/p/xMfePDdjjlj)

如果是在 map，就會是 key, value

```go
package main

import "fmt"

func main() {
	kvs := map[string]string{"a": "apple", "b": "banana"}
	for k, v := range kvs {
		fmt.Printf("%s -> %s\n", k, v)
	}
}
```

```go
a -> apple
b -> banana
```

[Go Playground](https://go.dev/play/p/H7F7UzJElQp)

## Slices

[slices-intro](https://go.dev/blog/slices-intro)

不用定義其最大長度，而且可以直接賦予值

* An array has a fixed size. A slice, on the other hand, is a dynamically-sizeds
* slice 的零值是 nil，一個 nil 的 slice 的長度和容量是 0，使用 make 就不會是 nil
* Slice 是用 reference

nil slice 和 empty slice 是不一樣的，但 length 都一樣是 0
slice 只能跟 nil 去做比較，兩個 slice 要比較，必須要用 loop 去一個一個比對

```go
a[start:end]
```

### 建立 slice

```go
package main

import (
	"fmt"
)

func main() {
	a := [5]int{1, 2, 3, 4, 5}
	var b []int = a[1:4] //creates a slice from a[1] to a[3]
	fmt.Println(a)
	fmt.Println(b)

	b[0] = 123
	a[len(a)-1] = 321
	fmt.Println(a)
	fmt.Println(b)
}
```

```go
[1 2 3 4 5]
[2 3 4]
[1 123 3 4 321]
[123 3 4]
```

[Go Playground](https://go.dev/play/p/7MSQchDMTW7)

實際上是建立了一個長度 3 的 array，並回傳 slice reference 到變數中
slices 都是用 references 的方式，因此改值也會影響到原本的值

```go
package main

import (
	"fmt"
)

func main() {
	a := []int{1, 2, 3}
	b := a[:] // 建立一個包含全部的 slice
	c := a[:] // 建立一個包含全部的 slice

	fmt.Println("a", a)
	fmt.Println("b", b)
	fmt.Println("c", c)
	b[0] = 123
	c[1] = 321
	fmt.Println("a", a)
	fmt.Println("b", b)
	fmt.Println("c", c)
}
```

```go
a [1 2 3]
b [1 2 3]
c [1 2 3]
a [123 321 3]
b [123 321 3]
c [123 321 3]
```

[Go Playground](https://go.dev/play/p/qZ12J2NfQfo)

### nil with slice

空值的時候，slice 可以用 nil 判斷

```go
package main

import (
	"fmt"
)

func main() {
	var nums []int

	if nums == nil {
		fmt.Printf("%d: %T\n", nums, nums)
		nums = append(nums, 1, 2, 3)
		fmt.Printf("%d: %T\n", nums, nums)
	}
}
```

```go
[]: []int
[1 2 3]: []int
```

[Go Playground](https://go.dev/play/p/LadU78aov9-)

### length and capacity

* len is the number of elements in the slice.
* cap is the number of elements in the underlying array starting from the index from which the slice is created.

從頭開始切，cap 就會縮小

```go
// 設定 slice 長度，容量
array[low : high : cap]
```

```go
package main

import (
	"fmt"
)

func main() {
	s := []int{1, 2, 3, 4, 5, 6}
	printSlice(s)
	s = s[:0]
	printSlice(s)
	s = s[:4] // Extend its length.
	printSlice(s)
	s = s[:cap(s)] // re-sliced
	printSlice(s)

	s = s[2:] // 切掉頭兩個
	printSlice(s)
	s = s[:cap(s)] // re-sliced
	printSlice(s)
	s = s[0:2:3] // 設定 capacity
	printSlice(s)
	s = s[:cap(s)] // re-sliced
	printSlice(s)
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
```

```go
len=6 cap=6 [1 2 3 4 5 6]
len=0 cap=6 []
len=4 cap=6 [1 2 3 4]
len=6 cap=6 [1 2 3 4 5 6]
len=4 cap=4 [3 4 5 6]
len=4 cap=4 [3 4 5 6]
len=2 cap=3 [3 4]
len=3 cap=3 [3 4 5]
```

[Go Playground](https://go.dev/play/p/jxrZkHCfG1v)

### Slices of slices 二維切片

```go
package main

import (
	"fmt"
	"strings"
)

func main() {
	// Create a tic-tac-toe board.
	board := [][]string{
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
	}

	// The players take turns.
	board[0][0] = "X"
	board[2][2] = "O"
	board[1][2] = "X"
	board[1][0] = "O"
	board[0][2] = "X"

	for i := 0; i < len(board); i++ {
		fmt.Printf("%s\n", strings.Join(board[i], " "))
	}
}
```

```go
X _ X
O _ X
_ _ O
```

[Go Playground](https://go.dev/play/p/73TnhLf50PU)
[strings.Join](https://pkg.go.dev/strings#Join)

### Append

當空間不夠時，會自動擴充 capacity(*2)，實際上是建立一個新的 array，將原有的值 copy 過去，新的 slice 在 reference 到新的 array 上面

```go
func append(s []T, x ...T) []T
```

```go
package main

import (
	"fmt"
)

func main() {
	var s []int // slice 沒給值的話，預設會是 nil
	printSlice(s)

	s = append(s, 0)
	printSlice(s)

	s = append(s, 1)
	printSlice(s)

	s = append(s, 2)
	printSlice(s)
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
```

```go
len=0 cap=0 []
len=1 cap=1 [0]
len=2 cap=2 [0 1]
len=3 cap=4 [0 1 2]
```

[Go Playground](https://go.dev/play/p/TscPwGGgq4D)

### 透過 make 來建立 slice (dynamically-sized arrays)

[make](https://pkg.go.dev/builtin#make)

可以透過 make 來指定 slice type, length, capacity

```go
func make([]T, len, cap)

// make(類型, 長度, 容量)
// cap option 預設會跟長度一樣
```

```go
package main

import (
	"fmt"
)

func main() {
	a := make([]int, 5) // a := [5]int{}
	printSlice("a", a)
	b := make([]int, 0, 5)
	printSlice("b", b)
	c := b[:2]
	printSlice("c", c)
	d := c[2:5]
	printSlice("d", d)
}

func printSlice(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n",
		s, len(x), cap(x), x)
}
```

```go
a len=5 cap=5 [0 0 0 0 0]
b len=0 cap=5 []
c len=2 cap=5 [0 0]
d len=3 cap=3 [0 0 0]
```

[Go Playground](https://go.dev/play/p/9GXim_7_th7)

### Memory Optimisation

因為 slice 是 reference array，因此 array 就無法被 garbage collected(垃圾回收)

用 slices 時，會 reference 到 array，此時會造成 array 無法被 garbage collected，因此會造成 memory 的浪費

因此要解決的話可以利用 copy，copy 會產生新的 slices，而原本的 array 就可以被 garbage collected

### copy

[copy](https://pkg.go.dev/builtin#copy)

```go
func copy(dst, src []T) int
```

```go
package main

import (
	"fmt"
)

func nums() []int {
	nums := []int{1, 2, 3, 4, 5}
	neededNums := nums[:len(nums)-2]
	numsCpy := make([]int, len(neededNums))
	copy(numsCpy, neededNums)
	return numsCpy
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

func main() {
	numsSlice := nums()
	fmt.Println(numsSlice)
	printSlice(numsSlice)
}
```

```go
[1 2 3]
len=3 cap=3 [1 2 3]
```

[Go Playground](https://go.dev/play/p/yGh24nPRPgu)

## Variadic Functions

`Variadic Functions` 是一個可接受任意的參數，一定要放最後面一個

```go
// ...Type 代表可以接收任意參數
func variadic(a int, b ...int) {
}
```

### example 1

原理是會將，後面的參數轉成 slice，所以 type 會變成 []int，但是因為 find 後面接收的參數 type 是 int，也不能直接將 []int 帶進去，因此 []int 透過 ... 語法糖，帶到 Variadic Functions 裡

```go
package main

import (
	"fmt"
)

func find(num int, nums ...int) {
	found := false
	for i, v := range nums {
		if v == num {
			fmt.Println(num, "found at index", i, "in", nums)
			found = true
		}
	}
	if !found {
		fmt.Println(num, "not found in ", nums)
	}
	fmt.Printf("\n")
}

func main() {
	find(89, 89, 90, 95)
	find(45, 56, 67, 45, 90, 109)
	find(78, 38, 56, 98)
	find(87)
	// 相當於
	nums := []int{89, 90, 95}
	find(89, nums...) // 要加 ... 否則 type 會是 []int，造成 error
}
```

```go
89 found at index 0 in [89 90 95]
45 found at index 2 in [56 67 45 90 109]
78 not found in  [38 56 98]
87 not found in  []
89 found at index 0 in [89 90 95]
```

[Go Playground](https://go.dev/play/p/ZcnRGDydcNA)

### example 2

一開始是用 slice(reference) 傳入，因此改變值，外面也會跟著改變
append 則是空間不夠時會建立一個新的 array， 因此不會影響到原來的值

```go
package main

import (
	"fmt"
)

func change(s ...string) {
	s[0] = "Go"
	s = append(s, "playground")
	fmt.Println(s) // [Go world playground]
}

func main() {
	welcome := []string{"hello", "world"}
	change(welcome...)
	fmt.Println(welcome) // [Go world]
}
```

```go
[Go world playground]
[Go world]
```

[Go Playground](https://go.dev/play/p/yHMnWzPs3Ct)
