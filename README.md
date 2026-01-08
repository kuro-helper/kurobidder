# kurobidder

## 安裝

### 1. Clone 專案

由於此專案不會被 `go get` 快取，需要先自行 clone 專案到本地：

```bash
git clone https://github.com/kuro-helper/kurobidder.git
cd kurobidder
```

### 2. 安裝依賴

```bash
go get github.com/PuerkitoBio/goquery
go get github.com/kuro-helper/kurohelper-proxy
```

> [!NOTICE]
> 本專案有使用到 kurohelper-proxy 模組來做代理伺服器

### 3. 在目標專案中使用 replace

在你的目標專案中的 `go.mod` 檔案裡，使用 `replace` 指令來引用本地 clone 的專案：

```go
module your-project

go 1.24.0

require (
    github.com/kuro-helper/kurobidder v0.0.0
)

replace github.com/kuro-helper/kurobidder => <local-path>
```

然後執行：

```bash
go mod tidy
```
