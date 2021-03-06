package weapp

import (
	"strings"

	"github.com/goodwong/go-x/dingtalk/client"
)

// New 创建
func New(cfg *Config) *Weapp {
	r := strings.NewReplacer("APPID", cfg.AppID, "APPSECRET", cfg.AppSecret)
	tokenAPI := "https://api.weixin.qq.com/cgi-bin/token?grant_type=client_credential&appid=APPID&secret=APPSECRET"
	tokenAPI = r.Replace(tokenAPI)

	client := client.New(&client.Config{
		TokenAPI: tokenAPI,
	})
	return &Weapp{
		AppID:  cfg.AppID,
		config: cfg,
		Client: client,
	}
}

// Config 配置类
type Config struct {
	AppID     string
	AppSecret string
}

// Weapp 功能类
type Weapp struct {
	AppID  string
	config *Config
	Client *client.APIClient
}

// UserSession 用户会话信息
type UserSession struct {
	OpenID     string `json:"openid"`      //用户唯一标识
	SessionKey string `json:"session_key"` //会话密钥
	UnionID    string `json:"unionid"`     //用户在开放平台的唯一标识符，在满足 UnionID 下发条件的情况下会返回，详见 UnionID 机制说明。
}

// Code2session 根据免登授权码获取用户信息
func (weapp *Weapp) Code2session(code string) (*UserSession, error) {
	url := "https://api.weixin.qq.com/sns/jscode2session?appid=APPID&secret=SECRET&js_code=JSCODE&grant_type=authorization_code"
	r := strings.NewReplacer(
		"APPID", weapp.config.AppID,
		"SECRET", weapp.config.AppSecret,
		"JSCODE", code,
	)

	var result UserSession
	err := weapp.Client.Get(r.Replace(url), &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// UnlimitedWacode 获取无限制小程序码
// Usage:
// ```go
// 	respBytes, err := wechatWeapp.UnlimitedWacode("pages/gift/index/main", "gift=123")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	output, err := os.Create("create_unlimited_wxcode_output.png")
// 	if err != nil {
// 		log.Fatal("创建output文件失败：", err)
// 	}
// 	defer output.Close()
// 	output.Write(respBytes)
// ```
func (weapp *Weapp) UnlimitedWacode(page, scene string, widths ...int) (respBytes []byte, err error) {
	url := "https://api.weixin.qq.com/wxa/getwxacodeunlimit?access_token=ACCESS_TOKEN"
	params := map[string]interface{}{
		"page":       page,
		"scene":      scene,
		"is_hyaline": true,
	}
	if len(widths) > 0 {
		params["width"] = widths[0]
	}
	err = weapp.Client.PostJSON(url, params, &respBytes)
	return
}
