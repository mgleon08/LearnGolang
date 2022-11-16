# Golang 介紹和安裝

- [Golang 介紹和安裝](#golang-介紹和安裝)
  - [介紹](#介紹)
    - [撰寫風格](#撰寫風格)
    - [特色](#特色)
  - [安裝](#安裝)
    - [檢查版號](#檢查版號)
    - [設定 $GOPATH](#設定-gopath)
  - [專案類別](#專案類別)
  - [命名規則](#命名規則)
  - [特殊關鍵字 && 內建類別 && 內建函示](#特殊關鍵字--內建類別--內建函示)
  - [簡單範例 Hello World](#簡單範例-hello-world)
    - [建立 Go Module](#建立-go-module)
    - [建立 main.go](#建立-maingo)
    - [執行檔案](#執行檔案)

## 介紹

Golang 是 Google 開發的一種 `靜態` `強型別` `編譯` 程式語言，支援垃圾回收(garbage collection)與併發 (concurrency)。

### 撰寫風格

* 每行程式結束後不需要撰寫分號 `;`
* 大括號 `{` 不能夠換行放置
* `if` 判斷式和 `for` 迴圈不需要以小括號包覆起來
* 使用 tab 做排版

### 特色

* 開放原始碼 (open source)
* 跨平台(cross-platform)
* 語言層原生支援併發(concurrency)
* 內建垃圾回收 (garbage collection)，可手動調整觸發時機
* 編譯速度快，執行效能高，部屬快速容易
* 程式風格強制統一
* 內建開發相關工具
* 豐富的標準函式庫
* 代碼風格清晰、簡單，並且有強制性

## 安裝

https://go.dev/dl/

mac 可以用 [Homebrew](https://brew.sh/index_zh-tw)

```go
brew update && brew upgrade
brew install go
```

### 檢查版號

```go
go version
go version go1.19.2 darwin/amd64
```

### 設定 $GOPATH

GOPATH 就是 golang 的 Workspaces

可以改設定在 `.bashrc` or `.zshrc`

```go
// golang
export GOPATH="$HOME/go"
export GOBIN="$GOPATH/bin"
export PATH="$PATH:$GOBIN"
```

GOPATH 中會在細分三個資料夾

```go
src - 放 Go 程式碼的地方
pkg - 放 Go package 的地方
bin - 編譯好的執行檔會放在這裡
```

設定好可以打 `go env`

```go
GO111MODULE=""
GOARCH="amd64"
GOBIN="/Users/username/go/bin"
GOCACHE="/Users/username/Library/Caches/go-build"
GOENV="/Users/username/Library/Application Support/go/env"
GOEXE=""
GOEXPERIMENT=""
GOFLAGS=""
GOHOSTARCH="amd64"
GOHOSTOS="darwin"
GOINSECURE=""
GOMODCACHE="/Users/username/go/pkg/mod"
GONOPROXY=""
GONOSUMDB=""
GOOS="darwin"
GOPATH="/Users/username/go"
GOPRIVATE=""
GOPROXY="https://proxy.golang.org,direct"
GOROOT="/usr/local/Cellar/go/1.19.2/libexec"
GOSUMDB="sum.golang.org"
GOTMPDIR=""
GOTOOLDIR="/usr/local/Cellar/go/1.19.2/libexec/pkg/tool/darwin_amd64"
GOVCS=""
GOVERSION="go1.19.2"
GCCGO="gccgo"
GOAMD64="v1"
AR="ar"
CC="clang"
CXX="clang++"
CGO_ENABLED="1"
GOMOD="/dev/null"
GOWORK=""
CGO_CFLAGS="-g -O2"
CGO_CPPFLAGS=""
CGO_CXXFLAGS="-g -O2"
CGO_FFLAGS="-g -O2"
CGO_LDFLAGS="-g -O2"
PKG_CONFIG="pkg-config"
GOGCCFLAGS="-fPIC -arch x86_64 -m64 -pthread -fno-caret-diagnostics -Qunused-arguments -fmessage-length=0 -fdebug-prefix-map=/var/folders/dh/br43j__55rz00yzmjlyqb59r0000gp/T/go-build3370784997=/tmp/go-build -gno-record-gcc-switches -fno-common"
```

## 專案類別

go 又分兩種專案

執行檔 (executable)
  * created for running
  * name should be main
  * always func main

函式庫 (library)
  * created for reusability
  * can have any name
  * no function main

[Program_execution](https://go.dev/ref/spec#Program_execution)

## 命名規則

1. golang use camelcase
2. 字首大寫代表可讓其他 package 使用，可理解為大寫 public 小寫 private
3. 在 `func` 外面一定要加上 `var` OR `func` 才可以去做定義
4. 常數const(constants) 通常也會第一個字大寫
5. 沒用到的參數可用 `blank identifier (_)` 來代替

## 特殊關鍵字 && 內建類別 && 內建函示

```go
// 特殊關鍵字
break  default  func    interface   select
case   defer    go   map struct
chan  else goto package switch
const fallthrough  if   range    type
continue   for  import  return   var

// 內建類別
- string
- bool
- Numeric Types
  - int  int8  int16  int32(rune)  int64
  - uint uint8(byte) uint16 uint32 uint64
  - float32 float64
  - complex64 complex128
  - byte // alias for uint8
  - rune // alias for int32，represents a Unicode code point

// 內建函示
make len cap new append copy close delete
complex real imag panic recover
```

## 簡單範例 Hello World

先建立一個資料夾

```go
mkdir $HOME/HelloWorld
```

### 建立 Go Module

`Go Module` 用來追蹤專案內的相關套件及版號。

```go
// cd $HOME/HelloWorld

go mod init HelloWorld
```

會產生 `go.mod`

```go
// module 名稱
module HelloWorld

// 使用的 go 版號
go 1.19
```

另外可以用

```go
go mod tidy
```

來加入 `缺少的模組` 和刪除 `未使用的模組`

### 建立 main.go

`touch main.go`

```go
// 宣告程式屬於哪個 package，所有的 go 檔案都必須聲明
package main

// 引入套件，多個可以加括號 ()
import "fmt"

// 程式執行入口，main 在 golang 中是特殊的 function，每個執行檔只能有一個，告訴 Golang 要從哪裡開始執行
func main() {
    // 使用 fmt 套件印出字串 hello world
    fmt.Println("Hello World")
}
```

[Go Playground](https://go.dev/play/p/FY3IAusUD60)

### 執行檔案

執行有三種方式

1. `go install`

執行 `go install` 後，會在產生執行檔在 `$GOPATH/bin`

```go
// $GOPATH/bin/HelloWorld

Hello World
```

另外在上面有設定 `export PATH="$PATH:$GOBIN"` 所以也可以直接執行 `HelloWorld` 不需要前面的 PATH

```go
// HelloWorld

Hello World
```

2. `go build`

會直接在當下目錄產生執行檔 `HelloWorld`

3. `go run`

```go
// go run main.go

Hello World
```

跟 `go build` 有點類似，差別是要給 file 參數

另外可以加上 `--work` 能看到會將執行檔放在一個暫存的地方，再去執行

```go
// go run --work main.go

WORK=/var/folders/dh/br43j__55rz00yzmjlyqb59r0000gp/T/go-build3630724493
Hello World
```
