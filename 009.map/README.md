# Map

- [Map](#map)
  - [Example](#example)
  - [Map of structs](#map-of-structs)
  - [Maps are reference types](#maps-are-reference-types)
  - [Maps equality](#maps-equality)


A map maps keys to values. (相當於其當語言的 hash, json)

- Similar to slices, maps are reference types
- Only be compared to nil
- The zero value of a map is nil
- nil 的 map 無法賦值，必須用 make

```go
make(map[type_of_key]type_of_value)
value, ok := map[key] // 判斷 key 的 value 在不在
delete(map, key)
```

## Example

```go
package main

import (
	"fmt"
)

func main() {
	// 聲明 map 的 key 和 value 的 type
	var m1 map[string]int
	fmt.Println("m1 == nil:", m1 == nil, m1) // true

	// 使用make函式建立一個非 nil 的 map，nil map 不能賦值
	m1 = make(map[string]int)
	fmt.Println("m1 == nil:", m1 == nil, m1) // false

	// 最後給已聲明的 map 賦值
	m1["a"] = 1
	fmt.Println("m1 =", m1)
	fmt.Println("m1[a] =", m1["a"])

	// 直接建立
	// map[keyType]valueTypes
	m2 := make(map[string]string)

	// 然後賦值
	m2["a"] = "aa"
	fmt.Println("m2 =", m2)

	// 初始化 + 賦值一體化
	m3 := map[string]string{
		"a": "aa",
		"b": "bb",
		"c": "cc",
	}
	fmt.Println("m3 =", m3)

	// 查找鍵值是否存在，ok 為 true / false, v 為 value
	if v, ok := m1["c"]; ok {
		fmt.Println(v)
	} else {
		fmt.Println("Key Not Found")
	}

	// for loop map
	for k, v := range m3 {
		fmt.Println(k, v)
	}

	// delete
	delete(m3, "a")
	fmt.Println(m3)

	//Length of the map
	fmt.Println(len(m3))
}
```

```go
m1 == nil: true map[]
m1 == nil: false map[]
m1 = map[a:1]
m1[a] = 1
m2 = map[a:aa]
m3 = map[a:aa b:bb c:cc]
Key Not Found
b bb
c cc
a aa
map[b:bb c:cc]
2
```

[Go Playground](https://go.dev/play/p/5kW2maLVaSS)

## Map of structs

```go
package main

import (
	"fmt"
)

type fruit struct {
	price int
	name  string
}

func main() {
	fruit1 := fruit{
		price: 10,
	}
	fruit2 := fruit{
		price: 20,
	}
	fruitInfo := map[string]fruit{
		"Apple":  fruit1,
		"Banana": fruit2,
	}

	for name, info := range fruitInfo {
		fmt.Printf("fruit: %s price: $%d\n", name, info.price)
	}
}
```

```go
fruit: Apple price: $10
fruit: Banana price: $20
```

[Go Playground](https://go.dev/play/p/LHdcB5U3whD)

## Maps are reference types

```go
package main

import (
	"fmt"
)

type fruit struct {
	price int
	name  string
}

func main() {
	fruit1 := fruit{
		price: 10,
	}
	fruitInfo := map[string]fruit{
		"Apple": fruit1,
	}

	for name, info := range fruitInfo {
		fmt.Printf("fruit: %s price: $%d\n", name, info.price)
	}
	fruitInfo2 := fruitInfo
	fruitInfo2["Apple"] = fruit{
		price: 30,
	}
	for name, info := range fruitInfo {
		fmt.Printf("fruit: %s price: $%d\n", name, info.price)
	}
}
```

```go
fruit: Apple price: $10
fruit: Apple price: $30
```

[Go Playground](https://go.dev/play/p/MWayqPwMibH)


## Maps equality

map 無法比較，只能比較 `nil`

```go
package main

func main() {
	map1 := map[string]int{
		"one": 1,
		"two": 2,
	}

	map2 := map1

	if map1 == map2 {
	}
}
```

```go
invalid operation: map1 == map2 (map can only be compared to nil)
```

[Go Playground](https://go.dev/play/p/mkfkORUV5pi)

```go
package main

import "fmt"

func main() {
	// 這樣就不是 nil，有賦與 {}
	// map1 := map[string]int{}
	var map1 map[string]int
	map2 := map1
	fmt.Println("map1 is nil", map1)

	if map1 == nil {
		fmt.Println("map1 is nil")
	}

	if map2 == nil {
		fmt.Println("map2 is nil")
	}

}
```

```go
map1 is nil map[]
map1 is nil
map2 is nil
```

[Go Playground](https://go.dev/play/p/vapejuNrbR7)
