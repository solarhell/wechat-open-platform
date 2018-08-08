package wechatopenplatform

type CommonResponse struct {
	Errcode int    `json:"errcode,omitempty"`
	Errmsg  string `json:"errmsg,omitempty"`
}

type AccessToken struct {
	CommonResponse
	AccessToken  string `json:"access_token"`
	ExpiresIn    int    `json:"expires_in"`
	RefreshToken string `json:"refresh_token"`
	Openid       string `json:"openid"`
	Scope        string `json:"scope"`
	Unionid      string `json:"unionid,omitempty"`
}

type UserInfo struct {
	CommonResponse
	Openid     string   `json:"openid"`
	Nickname   string   `json:"nickname"`
	Sex        int      `json:"sex"` // 普通用户性别，1为男性，2为女性
	Language   string   `json:"language"`
	City       string   `json:"city"`
	Province   string   `json:"province"`
	Country    string   `json:"country"`
	Headimgurl string   `json:"headimgurl,omitempty"`
	Privilege  []string `json:"privilege"`
	Unionid    string   `json:"unionid"`
}
