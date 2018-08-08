package wechatopenplatform

import (
	"encoding/json"
	"errors"
)

func (c *Client) GetUserInfo(accessToken AccessToken, openId string) (ui UserInfo, err error) {
	api, err := UserInfoURL(accessToken.AccessToken, openId)
	if err != nil {
		return ui, err
	}

	res, err := c.client.Get(api)

	if err != nil {
		return ui, err
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		return ui, ErrConnectWechatServer
	}

	err = json.NewDecoder(res.Body).Decode(&ui)
	if err != nil {
		return ui, err
	}

	if ui.Errcode != 0 {
		return ui, errors.New(ui.Errmsg)
	}

	return ui, nil
}
