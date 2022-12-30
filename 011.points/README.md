# 指標 Pointer

將變數直接指向記憶體位置就叫做 Pointer，要修改內容就直接到該記憶體位置修改

```go
*T
```

* The type *T is a pointer to a T value. Its zero value is nil.
* The & operator generates a pointer to its operand.
* Do not pass a pointer to an array as a argument to a function. Use slice instead.
* Go does not support pointer arithmetic(pointer 位置運算)
* 使用 value receivers，function 會用 copy 的方式，並且改不到原本的值，pointer receivers 則會使用 reference 且會改到原本的值，比較不會浪費記憶體

```go
package main

import (
	"fmt"
)

func main() {
	var i = 8  // i 佔用了一個記憶體空間
	var p *int // 宣告 p 是一個 int 的指標，但此時要指向哪還不知道
	fmt.Printf("Type of p is %T\n", p)
	fmt.Printf("Type of &i is %T\n", &i)
	fmt.Println("p =", p) // The zero value of a pointer is nil

	p = &i                  // 將 p 指到 i 的記憶體位置
	fmt.Println("p =", p)   // p 所指到的記憶體位置，就是i
	fmt.Println("&p =", &p) // p 的記憶體位置
	fmt.Println("*p =", *p) // '*' 代表透過 pointer 顯示該記憶體位置的值

	*p = 20 // 透過 pointer 寫入 i 的值
	fmt.Println("*p =", *p)
	fmt.Println("i =", i)

	// 透過 new 建立指標
	size := new(int)
	fmt.Printf("Size value is %d, type is %T, address is %v\n", *size, size, size)
	*size = 85
	fmt.Println("New size value is", *size)
}
```

```go
Type of p is *int
Type of &i is *int
p = <nil>
p = 0xc0000b4000
&p = 0xc0000b0018
*p = 8
*p = 20
i = 20
Size value is 0, type is *int, address is 0xc0000b4008
New size value is 85
```

[Go Playground](https://go.dev/play/p/y2MXwNTRYFW)

## Passing pointer to a function

```go
package main

import (
	"fmt"
)

func change(val *int) {
	*val = 55
}
func main() {
	a := 58
	fmt.Println("value of a before function call is", a)
	b := &a
	change(b)
	fmt.Println("value of a after function call is", a)
}
```

```go
value of a before function call is 58
value of a after function call is 55
```

[Go Playground](https://go.dev/play/p/P2kOH-Yv662)

## Returning pointer from a function

```go
package main

import (
	"fmt"
)

func hello() *int {
	i := 5
	return &i
}
func main() {
	d := hello()
	fmt.Println("Value of d", *d)
}
```

```go
Value of d 5
```

[Go Playground](https://go.dev/play/p/I6r-fRx2qML)

## Do not pass a pointer to an array as an argument to a function. Use slice instead.


```go
package main

import (
	"fmt"
)

func modifyA(arr *[3]int) {
	// 相當於 arr[0]
	(*arr)[0] = 90
}

func modifyB(arr *[3]int) {
	// 相當於 (*arr)[0]
	arr[0] = 90
}

func modifyC(sls []int) {
	sls[0] = 90
}

func main() {
	a := [3]int{89, 90, 91}
	modifyA(&a)
	fmt.Println(a)

	b := [3]int{89, 90, 91}
	modifyB(&b)
	fmt.Println(b)

	c := [3]int{89, 90, 91}
	modifyC(c[:])
	fmt.Println(c)

}
```

```go
[90 90 91]
[90 90 91]
[90 90 91]
```

[Go Playground](https://go.dev/play/p/nGsMF2epbH3)
