package wechatopenplatform

import "errors"

var (
	ErrSignature           = errors.New("signature error")
	ErrPaddingSize         = errors.New("padding size error")
	ErrAppidMismatch       = errors.New("app id mismtch error")
	ErrDataExpire          = errors.New("data expire error")
	ErrConnectWechatServer = errors.New("err connecet Tencent Wechat server")
)
