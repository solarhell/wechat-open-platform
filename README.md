# wechat-open-platform
[![Build Status](https://travis-ci.org/solarhell/wechat-open-platform.svg?branch=master)](https://travis-ci.org/solarhell/wechat-open-platform)


Golang 微信开放平台 SDK

## done
获取用户信息
AccessToken(需持久化 防止超过请求限制)
RefreshToken

## todo


## usage

### 获取用户信息
```go
package mina

import (
	WOP "github.com/solarhell/wechat-open-platform"
)

func main() {
	ak, err := WOP.GetAccessToken("appId", "appSecret", "code")
	if err != nil {
		// err handle
		...
	}
	
	ui, err := WOP.GetUserInfo(ak.AccessToken, ak.Openid)
	if err != nil {
		// err handle
		...
	}
	
	
}
```
