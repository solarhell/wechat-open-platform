# wechat-open-platform
[![Build Status](https://travis-ci.org/solarhell/wechat-open-platform.svg?branch=master)](https://travis-ci.org/solarhell/wechat-open-platform)


Golang 微信开放平台 SDK

## done
登录和解密用户信息
AccessToken(需持久化 防止超过请求限制)
生成二维码

## todo


## usage

### 登录
```go
package mina

import (
	"github.com/solarhell/wechat-open-platform"
	"net/http"
)

func main() {
	c := wechatopenplatform.NewClient(&http.Client{
		Transport: &wechatopenplatform.DebugRequestTransport{
			RequestHeader:  true,
			RequestBody:    true,
			ResponseHeader: true,
			ResponseBody:   true,
		},
	})

	ui, err := c.Login("appid", "secret", "code")
	...
}
```

### 生成不限量二维码
```go
recordId := 23333
scene := fmt.Sprintf("target_id=%d", recordId)
if len(scene) > 32 {
	err
	...
}
param := mina.QRCoder{
	Scene:     scene,
	Page:      "pages/index/index",
	Width:     430,
	AutoColor: true,
	IsHyaline: true,
}
...
```
