package wechatopenplatform

import "net/url"

const (
	baseURL         = "https://api.weixin.qq.com"
	userInfoAPI     = "/sns/userinfo"
	tokenAPI        = "/sns/oauth2/access_token"
	refreshTokenAPI = "/sns/oauth2/refresh_token"
)

// https://open.weixin.qq.com/cgi-bin/showdocument?action=dir_list&t=resource/res_list&verify=1&id=open1419316505&token=&lang=zh_CN
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

func TokenURL(appId, appSecret, code string) (s string, err error) {
	u, err := url.Parse(baseURL + tokenAPI)
	if err != nil {
		return s, err
	}

	query := u.Query()

	query.Set("appid", appId)
	query.Set("secret", appSecret)
	query.Set("code", code)
	query.Set("grant_type", "authorization_code")

	u.RawQuery = query.Encode()

	return u.String(), nil
}

func RefreshTokenURL(appId, refreshToken string) (s string, err error) {
	u, err := url.Parse(baseURL + refreshTokenAPI)
	if err != nil {
		return s, err
	}

	query := u.Query()

	query.Set("appid", appId)
	query.Set("refresh_token", refreshToken)
	query.Set("grant_type", "refresh_token")

	u.RawQuery = query.Encode()

	return u.String(), nil
}
