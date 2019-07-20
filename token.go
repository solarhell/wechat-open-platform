package wechatopenplatform

import (
	"errors"
	"github.com/imroc/req"
)

func GetAccessToken(appId, appSecret, code string) (ak AccessToken, err error) {
	api, err := TokenURL(appId, appSecret, code)
	if err != nil {
		return ak, err
	}

	r, err := req.Get(api)
	if err != nil {
		return ak, err
	}

	if r.Response().StatusCode != 200 {
		return ak, ErrConnectWechatServer
	}

	err = r.ToJSON(&ak)
	if err != nil {
		return ak, err
	}

	if ak.Errcode != 0 {
		return ak, errors.New(ak.Errmsg)
	}

	return ak, nil
}

func RefreshToken(appId, refreshToken string) (ak AccessToken, err error) {
	api, err := RefreshTokenURL(appId, refreshToken)
	if err != nil {
		return ak, err
	}

	r, err := req.Get(api)
	if err != nil {
		return ak, err
	}

	if r.Response().StatusCode != 200 {
		return ak, ErrConnectWechatServer
	}

	err = r.ToJSON(&ak)
	if err != nil {
		return ak, err
	}

	if ak.Errcode != 0 {
		return ak, errors.New(ak.Errmsg)
	}

	return ak, nil
}
