package constants

import "net/http"

const (
	Scheme       = "https://" // Scheme is the scheme of the API
	Host         = "hitokoto.cn"
	APIGatewayV1 = Host + "/api/restful/v1" // APIGatewayV1 is the API gateway of v1
	UserAgent    = "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/16.3 Safari/605.1.15"

	SentenceGlobalAPI = "v1.hitokoto.cn"
	SentenceAbroadAPI = "international.v1.hitokoto.cn"
)

var (
	APILogin = API{
		Api:         "/auth/login",
		Method:      http.MethodPost,
		ContentType: "application/json",
		PathVar:     "",
		Desc:        "登录接口，成功返回用户信息（包含令牌）",
	}
	APIRegister = API{
		Api:         "/auth/register",
		Method:      http.MethodPost,
		ContentType: "application/json",
		PathVar:     "",
		Desc:        "注册接口，成功返回用户信息。",
	}
	APIPasswordReset = API{
		Api:         "/auth/password/reset",
		Method:      http.MethodPost,
		ContentType: "application/json",
		PathVar:     "",
		Desc:        "重置密码接口",
	}
	APILikeGet = API{
		Api:         "/like",
		Method:      http.MethodGet,
		ContentType: "application/json",
		PathVar:     "",
		Desc:        "返回句子赞的相关信息",
	}
	APILikePost = API{
		Api:         "/like",
		Method:      http.MethodPost,
		ContentType: "application/json",
		PathVar:     "",
		Desc:        "提交赞，成功返回提交者 IP",
	}
	APILikeCancel = API{
		Api:         "/like/cancel",
		Method:      http.MethodPost,
		ContentType: "application/json",
		PathVar:     "",
		Desc:        "撤回已经发出的喜爱",
	}
	APIMark = API{
		Api:         "/mark",
		Method:      http.MethodGet,
		ContentType: "application/json",
		PathVar:     "",
		Desc:        "获取所有的审核标记",
	}
	APIUser = API{
		Api:         "/user",
		Method:      http.MethodGet,
		ContentType: "application/json",
		PathVar:     "",
		Desc:        "获取用户信息",
	}
	APIUserEmailVerify = API{
		Api:         "/user/email/verify",
		Method:      http.MethodPut,
		ContentType: "application/x-www-form-urlencoded",
		PathVar:     "",
		Desc:        "申请验证邮箱",
	}
	APIUserToken = API{
		Api:         "/user/token",
		Method:      http.MethodGet,
		ContentType: "application/json",
		PathVar:     "",
		Desc:        "返回用户令牌的相关信息",
	}
	APIUserTokenRefresh = API{
		Api:         "/user/token/refresh",
		Method:      http.MethodPut,
		ContentType: "application/x-www-form-urlencoded",
		PathVar:     "",
		Desc:        "返重置令牌，返回新令牌的相关信息",
	}
	APIUserPassword = API{
		Api:         "/user/password",
		Method:      http.MethodPut,
		ContentType: "application/x-www-form-urlencoded",
		PathVar:     "",
		Desc:        "修改密码",
	}
	APIUserEmail = API{
		Api:         "/user/email",
		Method:      http.MethodPut,
		ContentType: "application/x-www-form-urlencoded",
		PathVar:     "",
		Desc:        "修改邮箱（未来行为可能发生变化）",
	}
	APIUserNotificationSettingsGet = API{
		Api:         "/user/notification/settings",
		Method:      http.MethodGet,
		ContentType: "application/json",
		PathVar:     "",
		Desc:        "获取用户通知设置",
	}
	APIUserNotificationSettingsPut = API{
		Api:         "/user/notification/settings",
		Method:      http.MethodPut,
		ContentType: "application/x-www-form-urlencoded",
		PathVar:     "",
		Desc:        "设定用户通知设置，返回新设置",
	}
	APIUserHitokotoSummary = API{
		Api:         "/user/hitokoto/summary",
		Method:      http.MethodGet,
		ContentType: "application/json",
		PathVar:     "",
		Desc:        "获得用户一言提交数据的概览",
	}
	APIUserHitokotoLike = API{
		Api:         "/user/hitokoto/like",
		Method:      http.MethodGet,
		ContentType: "application/json",
		PathVar:     "",
		Desc:        "获得用户赞的句子",
	}
	APIUserHitokotoHistory = API{
		Api:         "/user/hitokoto/history",
		Method:      http.MethodGet,
		ContentType: "application/json",
		PathVar:     "",
		Desc:        "获得用户历史的一言提交",
	}
	APIUserHitokotoHistoryPending = API{
		Api:         "/user/hitokoto/history/pending",
		Method:      http.MethodGet,
		ContentType: "application/json",
		PathVar:     "",
		Desc:        "获得用户历史的一言提交（审核中部分）",
	}
	APIUserHitokotoHistoryRefuse = API{
		Api:         "/user/hitokoto/history/refuse",
		Method:      http.MethodGet,
		ContentType: "application/json",
		PathVar:     "",
		Desc:        "获得用户历史的一言提交（已拒绝部分）",
	}
	APIUserHitokotoHistoryAccept = API{
		Api:         "/user/hitokoto/history/accept",
		Method:      http.MethodGet,
		ContentType: "application/json",
		PathVar:     "",
		Desc:        "获得用户历史的一言提交（已上线部分）",
	}
	APIHitokotoAppend = API{
		Api:         "/hitokoto/append",
		Method:      http.MethodPost,
		ContentType: "application/json",
		PathVar:     "",
		Desc:        "添加一言，返回审核队列中新句子的信息",
	}
	APIHitokotoUUID = API{
		Api:         "/hitokoto/{uuid}",
		Method:      http.MethodGet,
		ContentType: "application/json",
		PathVar:     "uuid",
		Desc:        "查看指定一言的信息（通过 UUID）",
	}
	APIHitokotoUUIDMark = API{
		Api:         "/hitokoto/{uuid}/mark",
		Method:      http.MethodGet,
		ContentType: "application/json",
		PathVar:     "uuid",
		Desc:        "查看指定一言的审核标记（通过 UUID）",
	}
	APIHitokotoScorePost = API{
		Api:         "/hitokoto/{uuid}/score",
		Method:      http.MethodPost,
		ContentType: "application/json",
		PathVar:     "uuid",
		Desc:        "为已上线的句子评分，返回评分相关信息",
	}
	APIHitokotoScoreGet = API{
		Api:         "/hitokoto/{uuid}/score",
		Method:      http.MethodGet,
		ContentType: "application/json",
		PathVar:     "uuid",
		Desc:        "获得句子的评分信息",
	}
	APIHitokotoReport = API{
		Api:         "/hitokoto/{uuid}/report",
		Method:      http.MethodPost,
		ContentType: "application/json",
		PathVar:     "uuid",
		Desc:        "举报一言存在问题，返回提交举报的相关信息",
	}
	APIHitokoto = API{
		Api:         "",
		Method:      http.MethodGet,
		ContentType: "application/json",
		PathVar:     "",
		Desc:        "语句接口",
	}
)

type API struct {
	Api         string
	Method      string
	PathVar     string
	ContentType string
	Desc        string
}
