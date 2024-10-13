rmdir /s/q vendor

set GOPROXY=https://goproxy.cn

set GOPRIVATE=golang.corp.yxkj.com

:: 强制更新代码
:: go get -u golang.corp.yxkj.com/orange/common
:: go get -u golang.corp.yxkj.com/orange/pkg
:: go get -u golang.corp.yxkj.com/orange/yhLineServer/internal/common
:: go get -u golang.corp.yxkj.com/orange/plugins
:: go get -u golang.corp.yxkj.com/orange/convert
:: 更新依赖
go mod tidy
:: 更新本地依赖
go mod vendor