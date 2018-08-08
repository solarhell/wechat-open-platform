package wechatopenplatform

import "net/url"

const (
	baseURL     = "https://api.weixin.qq.com"
	userInfoAPI = "/sns/userinfo"
	tokenAPI    = "/sns/oauth2/access_token"
)

func UserInfoURL(accessToken, openId string) (s string, err error) {
	u, err := url.Parse(baseURL + userInfoAPI)
	if err != nil {
		return s, err
	}

	query := u.Query()

	query.Set("access_token", accessToken)
	query.Set("openid", openId)
	query.Set("lang", "zh-CN")

	u.RawQuery = query.Encode()

	return u.String(), nil
}

func TokenURL(appId, appSecret string) (s string, err error) {
	u, err := url.Parse(baseURL + tokenAPI)
	if err != nil {
		return s, err
	}

	query := u.Query()

	query.Set("appid", appId)
	query.Set("secret", appSecret)
	query.Set("grant_type", "authorization_code")

	u.RawQuery = query.Encode()

	return u.String(), nil
}
