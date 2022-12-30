# String

String 在 GO 裡面是一個 slice 裡面都是 byte

## 單引號 / 雙引號 / 反引號

### 雙引號

用來建立可解析的字串字面量(支援轉義，但不能用來引用多行)

```go
fmt.Printf("\u65e5\u672c\u8a9e") // 日本語
```

### 反引號 Raw string literal

用來建立原生的字串字面量，這些字串可能由多行組成(不支援任何轉義序列)，原生的字串字面量多用於書寫多行訊息、HTML以及正則表達式

```go
fmt.Printf(`\u65e5\u672c\u8a9e`) // \u65e5\u672c\u8a9e
```

### 單引號

表示 Golang 的一個特殊類型：rune(int32的別稱)，類似其他語言的 byte 但又不完全一樣，是指：碼點字面量（Unicode code point），不做任何轉義的原始內容。

* [Unicode Character Set and UTF-8, UTF-16, UTF-32 Encoding](https://naveenr.net/unicode-character-set-and-utf-8-utf-16-utf-32-encoding/)
* [UTF-8 encoder/decoder](https://mothereff.in/utf-8#%C3%B1)


## 列出每個 string 的 byte

```go
package main

import (
	"fmt"
)

func printBytes(s string) {
	fmt.Printf("Bytes: ")
	for i := 0; i < len(s); i++ {
		fmt.Printf("%x ", s[i])
	}
}

func main() {
	name := "Hello World"
	fmt.Printf("String: %s\n", name)
	printBytes(name)
}
```

```go
String: Hello World
Bytes: 48 65 6c 6c 6f 20 57 6f 72 6c 64
```

[Go Playground](https://go.dev/play/p/yWIGvZenKQY)

但在有些字元會有問題，因為字元是兩個 bytes 去組成

```go
package main

import (
	"fmt"
)

func printBytes(s string) {
	fmt.Printf("Bytes: ")
	for i := 0; i < len(s); i++ {
		fmt.Printf("%x ", s[i])
	}
}

func printChars(s string) {
	fmt.Printf("Characters: ")
	for i := 0; i < len(s); i++ {
		fmt.Printf("%c ", s[i])
	}
}

func main() {
	name := "Hello World"
	fmt.Printf("String: %s\n", name)
	printChars(name)
	fmt.Printf("\n")
	printBytes(name)
	fmt.Printf("\n\n")
	name = "Señor"
	fmt.Printf("String: %s\n", name)
	printChars(name)
	fmt.Printf("\n")
	printBytes(name)
}
```

```go
String: Hello World
Characters: H e l l o   W o r l d
Bytes: 48 65 6c 6c 6f 20 57 6f 72 6c 64

String: Señor
Characters: S e Ã ± o r
Bytes: 53 65 c3 b1 6f 72
```

[Go Playground](https://go.dev/play/p/2hyVf8l9fiO)

## Rune

rune 的 alice int32

Rune 表示 Go 中的一個 Unicode 代碼點
不管 code point 佔用多少字節，都可以用一個符文來表示


改用 rune 顯示正確字元

```go
package main

import (
	"fmt"
)

func printBytes(s string) {
	fmt.Printf("Bytes: ")
	for i := 0; i < len(s); i++ {
		fmt.Printf("%x ", s[i])
	}
}

func printChars(s string) {
	fmt.Printf("Characters: ")
	runes := []rune(s)
	for i := 0; i < len(runes); i++ {
		fmt.Printf("%c ", runes[i])
	}
}

func main() {
	name := "Hello World"
	fmt.Printf("String: %s\n", name)
	printChars(name)
	fmt.Printf("\n")
	printBytes(name)
	fmt.Printf("\n\n")
	name = "Señor"
	fmt.Printf("String: %s\n", name)
	printChars(name)
	fmt.Printf("\n")
	printBytes(name)
}
```

```go
String: Hello World
Characters: H e l l o   W o r l d
Bytes: 48 65 6c 6c 6f 20 57 6f 72 6c 64

String: Señor
Characters: S e ñ o r
Bytes: 53 65 c3 b1 6f 72
```

[Go Playground](https://go.dev/play/p/n8rsfagm2SJ)

透過 range 會自動轉 rune

```go
package main

import (
	"fmt"
)

func charsAndBytePosition(s string) {
	for index, rune := range s {
		fmt.Printf("%c starts at byte %d\n", rune, index)
	}
}

func main() {
	name := "Señor"
	charsAndBytePosition(name)
}
```

```go
S starts at byte 0
e starts at byte 1
ñ starts at byte 2
o starts at byte 4
r starts at byte 5
```

[Go Playground](https://go.dev/play/p/0ldNBeffjYI)

## 透過 slice of byte 顯示 string

```go
package main

import (
	"fmt"
)

func main() {
	byteSlice := []byte{67, 97, 102, 195, 169} //decimal equivalent of {'\x43', '\x61', '\x66', '\xC3', '\xA9'}
	str := string(byteSlice)
	fmt.Println(str)
}
```

```go
Café
```

[Go Playground](https://go.dev/play/p/LlCGNjsWFaZ)

## 透過 slice of rune 顯示 string

```go
package main

import (
	"fmt"
)

func main() {
	runeSlice := []rune{0x0053, 0x0065, 0x00f1, 0x006f, 0x0072}
	str := string(runeSlice)
	fmt.Println(str)
}
```

```go
Señor
```

[Go Playground](https://go.dev/play/p/dw_rPMLE1te)


## String 長度

透過 `len` 顯示的會是 string 轉 byte 後的長度

```go
package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	word1 := "Señor"
	fmt.Printf("String: %s\n", word1)
	fmt.Printf("Length: %d\n", utf8.RuneCountInString(word1))
	fmt.Printf("Number of bytes: %d\n", len(word1))

	fmt.Printf("\n")
	word2 := "Pets"
	fmt.Printf("String: %s\n", word2)
	fmt.Printf("Length: %d\n", utf8.RuneCountInString(word2))
	fmt.Printf("Number of bytes: %d\n", len(word2))
}
```

```go
String: Señor
Length: 5
Number of bytes: 6

String: Pets
Length: 4
Number of bytes: 4
```

[Go Playground](https://go.dev/play/p/KBQg1qagnfC)

## Strings are immutable

string 不能直接改字元，必須轉成 rune 更改後再轉回來

```go
package main

import (
	"fmt"
)

func mutate(s []rune) string {
	s[0] = 'a'
	return string(s)
}
func main() {
	h := "hello"
	fmt.Println(mutate([]rune(h)))
}
```

```go
aello
```

[Go Playground](https://go.dev/play/p/Ym_D7yJhKco)
