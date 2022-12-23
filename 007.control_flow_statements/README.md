# Control Flow Statements (if/for/while/range/switch)

- [Control Flow Statements (if/for/while/range/switch)](#control-flow-statements-ifforwhilerangeswitch)
	- [If/Else](#ifelse)
		- [Syntax](#syntax)
		- [Example](#example)
	- [Loop](#loop)
		- [For](#for)
			- [Syntax](#syntax-1)
			- [Example](#example-1)
		- [infinite loop](#infinite-loop)
			- [Syntax](#syntax-2)
			- [Example](#example-2)
		- [range (each) 迴圈](#range-each-迴圈)
		- [Labels](#labels)
	- [Switch](#switch)
		- [Syntax](#syntax-3)
		- [Example](#example-3)
		- [Fallthrough](#fallthrough)

## If/Else

### Syntax

注意 `{` 必須要在 `else` 後面，不能是換行

```go
if condition {
} else if condition {
} else {
}
```

```go
if statement; condition {
}
```


### Example

```go
package main

import (
	"fmt"
)

func main() {
	num := 1
	if num == 1 {
		fmt.Println("num1", num)
	}
	if num%2 == 0 {
		fmt.Println("num2", num)
	} else {
		fmt.Println("num3", num)
	}
	if num += 1; num == 2 {
		fmt.Println("num4", num)
	}

}
```

```go
num1 1
num3 1
num4 2
```

[Go Playground](https://go.dev/play/p/CiI4_PWApCQ)

## Loop

### For

Golang 中只有 for 一種迴圈，但能夠達成 for、while、foreach 多種用法

#### Syntax

```go
for init; condition; post {
}
```

#### Example

```go
package main

import (
	"fmt"
)

func main() {
	a1 := 1
	for a1 <= 3 {
		fmt.Println("a1: ", a1)
		a1 = a1 + 1
	}

	for a2 := 7; a2 <= 9; a2++ {
		fmt.Println("a2: ", a2)
	}

	for a4, b4 := 1, 2; b4 <= 10 && a4 <= 5; b4, a4 = b4+1, a4+1 {
		fmt.Printf("a4: %d * %d = %d\n", a4, b4, a4*b4)
	}

	for a5 := 0; a5 <= 5; a5++ {
		if a5%2 == 0 {
			continue
		}
		fmt.Println("a5: ", a5)
	}

	for a6 := 0; a6 <= 5; a6++ {
		if a6%2 == 0 {
			break
		}
		fmt.Println("a6: ")
	}
}
```

```go
a1:  1
a1:  2
a1:  3
a2:  7
a2:  8
a2:  9
a4: 1 * 2 = 2
a4: 2 * 3 = 6
a4: 3 * 4 = 12
a4: 4 * 5 = 20
a4: 5 * 6 = 30
a5:  1
a5:  3
a5:  5
```

[Go Playground](https://go.dev/play/p/dQJcj_gB-6_F)

### infinite loop

會一直跑，直到 `break` or `return`，也能夠透過 `continue` 繼續跑下去

#### Syntax

```go
for {
}
```


#### Example

```go
package main

import (
	"fmt"
)

func main() {
	for {
		fmt.Println("loop")
		break
	}
}
```

```go
loop
```

[Go Playground](https://go.dev/play/p/K_qJ1uJ8pFk)

### range (each) 迴圈

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

### Labels

當在中間用 break，只會中斷裡面的迴圈，繼續外面的迴圈

```go
package main

import (
	"fmt"
)

func main() {
	for i := 0; i < 3; i++ {
		for j := 1; j < 4; j++ {
			fmt.Printf("i = %d , j = %d\n", i, j)
			if i == j {
				break
			}
		}
	}
}
```

```go
i = 0 , j = 1
i = 0 , j = 2
i = 0 , j = 3
i = 1 , j = 1
i = 2 , j = 1
i = 2 , j = 2
```

[Go Playground](https://go.dev/play/p/_ymnujBU1IE)

加上 label 後，就可以直接跳到外面

```go
package main

import (
	"fmt"
)

func main() {
outer:
	for i := 0; i < 3; i++ {
		for j := 1; j < 4; j++ {
			fmt.Printf("i = %d , j = %d\n", i, j)
			if i == j {
				break outer
			}
		}
	}
}
```

```go
i = 0 , j = 1
i = 0 , j = 2
i = 0 , j = 3
i = 1 , j = 1
```

[Go Playground](https://go.dev/play/p/_-dsKTsGvtN)

## Switch

### Syntax

`condition` 不能重複, `default` 不一定要

```go
switch value {
case condition:
case condition:
default:
}
```

### Example

```go
package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("When's Saturday?")
	today := time.Now().Weekday()
	switch time.Saturday {
	case today + 0:
		fmt.Println("Today.")
	case today + 1:
		fmt.Println("Tomorrow.")
	case today + 2:
		fmt.Println("In two days.")
	default:
		fmt.Println("Too far away.")
	}

	switch 1 {
	case 1, 3, 5, 7, 9:
		fmt.Println("odd")
	case 2, 4, 6, 8, 10:
		fmt.Println("even")
	default:
		fmt.Println("-")
	}
}
```

```go
When's Saturday?
Too far away.
odd
```

[Go Playground](https://go.dev/play/p/OgA8H0I6r-b)

### Fallthrough

可以遇到條件符合之後，繼續往下走一個，後面就不管條件直接往下走，可以透過 `break` 來中止

```go
package main

import (
	"fmt"
)

func main() {
	switch num := 1; {
	case num < 10:
		fmt.Printf("%d is lesser than 10\n", num)
		fallthrough
	case num > 20:
		fmt.Printf("%d is greater than 20\n", num)
		fallthrough
	case num > 30:
		fmt.Printf("%d is greater than 30\n", num)
	case num > 40:
		fmt.Printf("%d is greater than 40\n", num)
	}

}
```

```go
1 is lesser than 10
1 is greater than 20
1 is greater than 30
```

[Go Playground](https://go.dev/play/p/Qr3_g2nog4k)
