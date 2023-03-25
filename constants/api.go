package constants

import "net/http"

const (
	Scheme       = "https://"                   // Scheme is the scheme of the API
	APIGatewayV1 = "hitokoto.cn/api/restful/v1" // APIGatewayV1 is the API gateway of v1
)

var (
	APILogin = API{
		Api:         "/auth/login",
		Method:      http.MethodPost,
		ContentType: "application/json",
		PathVar:     "",
		Desc:        "登录接口，成功返回用户信息（包含令牌）",
	}
)

type API struct {
	Api         string
	Method      string
	PathVar     string
	ContentType string
	Desc        string
}
