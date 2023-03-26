# Hitokoto-GO-SDK(一言网GO-SDK非官方版本)

[官网](https://hitokoto.cn) | [开发中心](https://developer.hitokoto.cn)

本项目使用 [Hitokoto](https://hitokoto.cn)(一言网) 官方文档中的API，对官方文档中的API使用Golang语言封装的SDK。所有API均可在官方开发者文档中查询，不包含未开放API

## 本项目

Hitokoto GO SDK是一个为Hitokoto平台设计的软件开发工具包，旨在让开发者更方便地使用Hitokoto平台的API接口，快速地获取短句子数据。

## 一言网(Hitokoto)

Hitokoto是一个开源的短句子分享平台，旨在让用户分享和发现有趣、有意义、有启发性的句子。在Hitokoto网站上，每次访问页面时会随机展示一句短语，这些短语来自全球用户提交的数据，并经过审核和分类。

一言网（hitokoto.cn）创立于 2016 年，隶属于萌创团队，目前网站主要提供一句话服务。

动漫也好、小说也好、网络也好，不论在哪里，我们总会看到有那么一两个句子能穿透你的心。我们把这些句子汇聚起来，形成一言网络，以传递更多的感动。如果可以，我们希望我们没有停止服务的那一天。

简单来说，一言指的就是一句话，可以是动漫中的台词，也可以是网络上的各种小段子。 或是感动，或是开心，有或是单纯的回忆。来到这里，留下你所喜欢的那一句句话，与大家分享，这就是一言存在的目的。

## 安装

使用Hitokoto GO SDK前，需要先安装golang。安装完成后，使用go get工具可以轻松地安装Hitokoto GO SDK：

``` bash
go get github.com/MarioMang/hitokoto@v1.0.2
```

## 使用方法

### 引入 SDK

使用 go get 或者在 go.mod 中手动添加

``` text
 required github.com/mariomang/hitokoto-go v1.0.2
```

### 初始化

在使用Hitokoto GO SDK之前，需要先进行初始化。初始化方法需要传入Hitokoto API接口的相关参数。在初始化之后，就可以使用SDK提供的方法获取短句子数据了。
> ps: SDK已经内置了API请求地址等一些参数，必须要自己设置

不使用自定义 User-Agent:

``` golang
var executor = NewExecutor()
```

使用自定义 User-Agent:

``` golang
var executor = NewExecutor().WithUserAgent(userAgent)
```

### 部分示例

#### 语句接口

| 地址 | 协议 | 方法 | QPS 限制 | 线路 |
|:---|:---|:---|:---:|:---|
|v1.hitokoto.cn| HTTPS| Any| 2| 全球|
|international.v1.hitokoto.cn| HTTPS| Any| 10| 海外|

使用默认配置
> 默认请求 (international.v1.hitokoto.cn) 地址

``` golang
e := NewExecutor()

req := &op.HitokotoRequest{
    Type:    []constants.HitokotoType{constants.Animation},
    Encode:  constants.EncodeJson,
    Charset: constants.CharsetUTF8,
}
resp := &op.HitokotoResponse{}
err := e.Do(&constants.APIHitokoto, req, resp)
if err != nil {
    t.Errorf("Error executing request: %v", err)
}
```

使用Global地址:
> 请求 (v1.hitokoto.cn) 地址

``` golang
e := NewExecutor().WithGlobalAPI()

req := &op.HitokotoRequest{
    Type:    []constants.HitokotoType{constants.Animation},
    Encode:  constants.EncodeJson,
    Charset: constants.CharsetUTF8,
}
resp := &op.HitokotoResponse{}
err := e.Do(&constants.APIHitokoto, req, resp)
if err != nil {
    t.Errorf("Error executing request: %v", err)
}
```

句子类型(HitokotoType):

|参数| 说明| 类型 |
|:---:|:---:|:---|
|a |动画| constants.Animation |
|b| 漫画| constants.Cartoon |
|c| 游戏| constants.Game |
|d| 文学| constants.Literature |
|e| 原创| constants.Original |
|f| 来自网络| constants.Internet |
|g| 其他| constants.Other |
|h| 影视| constants.MoviesAndTv |
|i| 诗词| constants.Poetry |
|j| 网易云| constants.Netease |
|k| 哲学| constants.Philosophy |
|l| 抖机灵| constants.PettyTrick |

返回编码(HitokotoEncode):
|参数| 说明| 类型 |
|:---:|:---|:---|
|text |返回纯洁文本| constants.EncodeText |
|json| 返回格式化后的 JSON 文本| constants.EncodeJson |
|js| 返回指定选择器的同步调用函数。默认选择器为：.hitokoto| constants.EncodeJs |

字符集(HitokotoCharset):
|参数| 说明| 类型 |
|:---:|:---|:---|
|utf-8 |返回 utf-8 编码的内容 | constants.CharsetUTF8 |
|gbk| 返回 gbk 编码的内容。不支持与异步函数同用 | constants.CharsetGBK |

#### 身份认证(/auth/login)

登录接口，成功返回用户信息（包含令牌）

``` golang
e := NewExecutor()

req := &op.LoginRequest{
    Email:    "user@hitokoto.cn",
    Password: "xxxxxxxxx",
}
resp := &op.LoginResponse{}
err := e.Do(&constants.APILogin, req, resp)
if err != nil {
    e, ok := err.(*HitokotoError)
    if !ok {
        t.Errorf("Error executing request: %v", err)
    }

    if e.Status != 200 {
        t.Errorf("Status is not correct: %v", e.Status)
    }
}
```

#### 注册接口(/auth/register)

注册接口，成功返回用户信息。

``` golang
e := NewExecutor()

req := &op.RegisterRequest{
    Name:     "皮皮喵",
    Email:    "pipi@hitokoto.cn",
Password: "gugugu",
}
resp := &op.RegisterResponse{}
err := e.Do(&constants.APIRegister, req, resp)
if err != nil {
    e, ok := err.(*HitokotoError)
    if !ok {
        t.Errorf("Error executing request: %v", err)
    }

    if e.Status != 200 {
        t.Errorf("Status is not correct: %v", e.Status)
    }
}
```

#### 获取赞过的句子

获得用户赞的句子

``` golang
e := NewExecutor()

req := &op.UserHitokotoLikeRequest{
    Token:  "xxxxxx",
    Offset: 0,
    Limit:  10,
}
resp := &op.UserHitokotoLikeResponse{}
err := e.Do(&constants.APIUserHitokotoLike, req, resp)
if err != nil {
    e, ok := err.(*HitokotoError)
    if !ok {
    t.Errorf("Error executing request: %v", err)
    }

    if e.Status != 200 {
    t.Errorf("Status is not correct: %v", e.Status)
    }
}
```

### 映射关系

|api | method | structure | request | response |
|:---|:---|:---|:---|:---|
|/auth/login|POST|constants.APILogin|op.LoginRequest|op.LoginResponse|
|/auth/register|POST|constants.APIRegister|op.RegisterRequest|op.RegisterResponse|
|/auth/password/reset|POST|constants.APIPasswordReset|op.PasswordResetRequest|op.PasswordResetResponse|
|/like|GET|constants.APILikeGet|op.LikeGetRequest|op.LikeGetResponse|
|/like|POST|constants.APILikePost|op.LikePostRequest|op.LikePostResponse|
|/like/cancel|POST|constants.APILikeCancel|op.LikeCancelRequest|op.LikeCancelResponse|
|/mark|GET|constants.APIMark|op.MarkRequest|op.MarkResponse|
|/user|GET|constants.APIUser|op.UserRequest|op.UserResponse|
|/user/email/verify|PUT|constants.APIUserEmailVerify|op.UserEmailVerifyRequest|op.UserEmailVerifyResponse|
|/user/token|GET|constants.APIUserToken|op.UserTokenRequest|op.UserTokenResponse|
|/user/token/refresh|PUT|constants.APIUserTokenRefresh|op.UserTokenRefreshRequest|op.UserTokenRefreshResponse|
|/user/password|PUT|constants.APIUserPassword|op.UserPasswordRequest|op.UserPasswordResponse|
|/user/email|PUT|constants.APIUserEmail|op.UserEmailRequest|op.UserEmailResponse|
|/user/notification/settings|GET|constants.APIUserNotificationSettingsGet|op.UserNotificationSettingsGetRequest|op.UserNotificationSettingsGetResponse|
|/user/notification/settings|PUT|constants.APIUserNotificationSettingsPut|op.UserNotificationSettingsPutRequest|op.UserNotificationSettingsPutResponse|
|/user/hitokoto/like|GET|constants.APIUserHitokotoLike|op.UserHitokotoLikeRequest|op.UserHitokotoLikeResponse|
|/user/hitokoto/summary|GET|constants.APIUserHitokotoSummary|op.UserHitokotoSummaryRequest|op.UserHitokotoSummaryResponse|
|/user/hitokoto/history|GET|constants.APIUserHitokotoHistory|op.UserHitokotoHistoryRequest|op.UserHitokotoHistoryResponse|
|/user/hitokoto/history/pending|GET|constants.APIUserHitokotoHistoryPending|op.UserHitokotoHistoryPendingRequest|op.UserHitokotoHistoryPendingResponse|
|/user/hitokoto/history/refuse|GET|constants.APIUserHitokotoHistoryRefuse|op.UserHitokotoHistoryRefuseRequest|op.UserHitokotoHistoryRefuseResponse|
|/user/hitokoto/history/accept|GET|constants.APIUserHitokotoHistoryAccept|op.UserHitokotoHistoryAcceptRequest|op.UserHitokotoHistoryAcceptResponse|
|/hitokoto/append|POST|constants.APIHitokotoAppend|op.HitokotoAppendRequest|op.HitokotoAppendResponse|
|/hitokoto/{uuid}|GET|constants.APIHitokotoUUID|op.HitokotoUUIDRequest|op.HitokotoUUIDResponse|
|/hitokoto/{uuid}/mark|GET|constants.APIHitokotoUUIDMark|op.HitokotoUUIDMarkRequest|op.HitokotoUUIDMarkResponse|
|/hitokoto/{uuid}/score|POST|constants.APIHitokotoScorePost|op.HitokotoScorePostRequest|op.HitokotoScorePostResponse|
|/hitokoto/{uuid}/score|GET|constants.APIHitokotoScoreGet|op.HitokotoScoreGetRequest|op.HitokotoScoreGetResponse|
|/hitokoto/{uuid}/report|POST|constants.APIHitokotoReport|op.HitokotoReportRequest|op.HitokotoReportResponse|
