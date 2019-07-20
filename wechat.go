package wechatopenplatform

import (
	"errors"
	"github.com/imroc/req"
)

func GetUserInfo(accessToken, openId string) (ui UserInfo, err error) {
	api, err := UserInfoURL(accessToken, openId)
	if err != nil {
		return ui, err
	}

	r, err := req.Get(api)
	if err != nil {
		return ui, err
	}

	if r.Response().StatusCode != 200 {
		return ui, ErrConnectWechatServer
	}

	err = r.ToJSON(&ui)
	if err != nil {
		return ui, err
	}

	if ui.Errcode != 0 {
		return ui, errors.New(ui.Errmsg)
	}

	return ui, nil
}
