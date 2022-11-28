# Packages

- [Packages](#packages)
  - [宣告](#宣告)
  - [Init Function](#init-function)
  - [Import](#import)
  - [Importing Custom Package](#importing-custom-package)
  - [Exported Names](#exported-names)

## 宣告

宣告程式屬於哪個 package

```go
package test
```

所有的 go 檔案都必須聲明，要 import 這個檔案時，就必須使用這個名稱。

```go
import test
```

而 go 又分兩種專案

* 執行檔 (executable)
* 函式庫 (library)

 執行檔一定要宣告為 `main` 套件

```go
// 沒聲明 main 會顯示
go run: cannot run non-main package
```

[Program_execution](https://go.dev/ref/spec#Program_execution)

> Package main is special. It defines a standalone executable program, not a library. Within package main the function main is also special—it’s where execution of the program begins. Whatever main does is what the program does.


## Init Function

當 Import Package 只要執行 `init()` 可以使用 `_`

```go
import _ "fmt"
```

## Import

```go
// 引入套件，多個可以加括號 ()
import "fmt"

import (
	"fmt"
	"math"
)

// 希望使用匯入的套件，是為了要觸發那個套件的 init func 而引用的話，可以在前面加上一個底線 _
import _ math

// 如果名字有衝突可以加上 neko
import (
    "github.com/test1/foo"
    neko "github.com/test2/foo"
)

// 當前文件同一目錄的 model 目錄，但是不建議這種方式 import
import "./test"

// 載入 GOPATH/src/test1/foo 模組
import "github.com/test1/foo"

// 點操作
import(
    . "fmt" // 可以使 fmt.Println("Hi") 省略為 Println("Hi")
)

// 別名操作
import(
    f "fmt" // 就可以改為用 f 來呼叫，f.Println("Hi")
)
```

## Importing Custom Package

建立一個 custom package

```go
package custompackage

func SayHi(name string) string {
    text := "Hello "
    text += name
    text += " I'm Custom Package"
    return text
}
```

import

```go
package main

import (
    "fmt"
    "learnpackage/custompackage"
)

func main() {
    text := custompackage.SayHi("Foo")
    fmt.Println(text)
}
```

結構

```go
learnpackage
  ├── go.mod
  ├── main.go
  └── custompackage
      └── custompackage.go
```

`go install` 也必須在 learnpackage 目錄底下，且必須要有 `go.mod`

```go
go mod init learnpackage
```

## Exported Names

如果 package 的 function 是要給外面使用時，必須要大寫開頭，像是上面的 `SayHi` 一樣，如果是小寫則是認定為是 private 而報錯

```go
undefined: custompackage.sayHi
```
