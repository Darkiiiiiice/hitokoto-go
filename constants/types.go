package constants

type HitokotoType string

const (
	Animation   HitokotoType = "a" // 动画
	Cartoon     HitokotoType = "b" // 漫画
	Game        HitokotoType = "c" // 游戏
	Literature  HitokotoType = "d" // 文学
	Original    HitokotoType = "e" // 原创
	Internet    HitokotoType = "f" // 来自网络
	Other       HitokotoType = "g" // 其他
	MoviesAndTv HitokotoType = "h" // 影视
	Poetry      HitokotoType = "i" // 诗词
	Netease     HitokotoType = "j" // 网易云
	Philosophy  HitokotoType = "k" // 哲学
	PettyTrick  HitokotoType = "l" // 抖机灵
)

type HitokotoEncode string

const (
	EncodeText  HitokotoEncode = "text"  // 纯文本
	EncodeJson  HitokotoEncode = "json"  // JSON
	EncodeJs    HitokotoEncode = "js"    // JS
	EncodeOther HitokotoEncode = "other" // 其他
)

type HitokotoCharset string

const (
	CharsetUTF8  HitokotoCharset = "utf-8" // UTF-8
	CharsetGBK   HitokotoCharset = "gbk"   // GBK
	CharsetOther HitokotoCharset = "other" // 其他
)
