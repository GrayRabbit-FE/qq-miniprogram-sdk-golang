package types

type LoginRequest struct {
	AppId     string `json:"appid"`
	Secret    string `json:"secret"`
	JsCode    string `json:"js_code"`
	GrantType string `json:"grant_type"` // 默认值：authorization_code
}
