package op

import (
	"net/url"
	"strconv"
	"time"

	"github.com/mariomang/hitokoto-go/constants"
	"github.com/valyala/fastjson"
)

type HitokotoRequest struct {
	API       string
	Type      []constants.HitokotoType
	Encode    constants.HitokotoEncode
	Charset   constants.HitokotoCharset
	Callback  string
	Select    string
	MinLength int
	MaxLength int
}

func (r *HitokotoRequest) FormatToValues() url.Values {
	var values = url.Values{}

	for _, t := range r.Type {
		values.Add("c", string(t))
	}
	var encode = string(r.Encode)
	if encode != "" {
		values.Add("encode", encode)
	}
	var charset = string(r.Charset)
	if charset != "" {
		values.Add("charset", charset)
	}
	var callback = r.Callback
	if callback != "" {
		values.Add("callback", callback)
	}
	var selectStr = r.Select
	if selectStr != "" {
		values.Add("select", selectStr)
	}
	var minLength = r.MinLength
	if minLength > 0 {
		values.Add("minlength", strconv.Itoa(minLength))
	}
	var maxLength = r.MaxLength
	if maxLength > 0 {
		values.Add("maxlength", strconv.Itoa(maxLength))
	}

	return values
}

type HitokotoResponse struct {
	ID         int    `json:"id"`
	Hitokoto   string `json:"hitokoto"`
	Type       string `json:"type"`
	From       string `json:"from"`
	FromWho    string `json:"from_who"`
	Creator    string `json:"creator"`
	CreatorUID int    `json:"creator_uid"`
	Reviewer   int    `json:"reviewer"`
	UUID       string `json:"uuid"`
	CommitFrom string `json:"commit_from"`
	Length     int    `json:"length"`

	CreatedAt time.Time `json:"created_at"`
}

func (r *HitokotoResponse) Parse(data []byte) error {
	v, err := fastjson.ParseBytes(data)
	if err != nil {
		return err
	}

	r.ID = v.GetInt("id")
	r.UUID = string(v.GetStringBytes("uuid"))
	r.Hitokoto = string(v.GetStringBytes("hitokoto"))
	r.Type = string(v.GetStringBytes("type"))
	r.From = string(v.GetStringBytes("from"))
	r.FromWho = string(v.GetStringBytes("from_who"))
	r.Creator = string(v.GetStringBytes("creator"))
	r.CreatorUID = v.GetInt("creator_uid")
	r.CommitFrom = string(v.GetStringBytes("commit_from"))
	r.Length = v.GetInt("length")

	createdAt, err := strconv.ParseInt(string(v.GetStringBytes("created_at")), 10, 64)
	if err != nil {
		createdAt = 0
	}
	r.CreatedAt = time.Unix(createdAt, 0)

	return nil
}

type HitokotoAppendRequest struct {
	Token    string
	From     string
	FromWho  string
	Hitokoto string
	Type     string
}

func (r *HitokotoAppendRequest) FormatToValues() url.Values {
	var values = url.Values{}

	values.Add("token", r.Token)
	values.Add("from", r.From)
	values.Add("from_who", r.FromWho)
	values.Add("hitokoto", r.Hitokoto)
	values.Add("type", r.Type)

	return values
}

type HitokotoAppendResponse struct {
	ID         int    `json:"id"`
	UUID       string `json:"uuid"`
	Hitokoto   string `json:"hitokoto"`
	Type       string `json:"type"`
	From       string `json:"from"`
	FromWho    string `json:"from_who"`
	Creator    string `json:"creator"`
	CreatorUID int    `json:"creator_uid"`
	CommitFrom string `json:"commit_from"`

	CreatedAt time.Time `json:"created_at"`
}

func (r *HitokotoAppendResponse) Parse(data []byte) error {
	v, err := fastjson.ParseBytes(data)
	if err != nil {
		return err
	}

	list, err := v.Array()
	if err != nil {
		return err
	}
	if len(list) > 0 {

		var d = list[0]

		r.ID = d.GetInt("id")
		r.UUID = string(d.GetStringBytes("uuid"))
		r.Hitokoto = string(d.GetStringBytes("hitokoto"))
		r.Type = string(d.GetStringBytes("type"))
		r.From = string(d.GetStringBytes("from"))
		r.FromWho = string(d.GetStringBytes("from_who"))
		r.Creator = string(d.GetStringBytes("creator"))
		r.CreatorUID = d.GetInt("creator_uid")
		r.CommitFrom = string(d.GetStringBytes("commit_from"))

		createdAt, err := strconv.ParseInt(string(d.GetStringBytes("created_at")), 10, 64)
		if err != nil {
			createdAt = 0
		}
		r.CreatedAt = time.Unix(createdAt, 0)
	}
	return nil
}

type HitokotoUUIDRequest struct {
	Token string
	UUID  string
}

func (r *HitokotoUUIDRequest) FormatToValues() url.Values {
	var values = url.Values{}

	values.Add("token", r.Token)
	values.Add("uuid", r.UUID)

	return values
}

type HitokotoUUIDResponse struct {
	Hitokoto   string `json:"hitokoto"`
	UUID       string `json:"uuid"`
	Type       string `json:"type"`
	From       string `json:"from"`
	FromWho    string `json:"from_who"`
	Creator    string `json:"creator"`
	CreatorUID int    `json:"creator_uid"`
	Reviewer   int    `json:"reviewer"`
	CommitFrom string `json:"commit_from"`
	Status     string `json:"status"`

	CreatedAt time.Time `json:"created_at"`
}

func (r *HitokotoUUIDResponse) Parse(data []byte) error {
	v, err := fastjson.ParseBytes(data)
	if err != nil {
		return err
	}

	list, err := v.Array()
	if err != nil {
		return err
	}
	if len(list) > 0 {

		var d = list[0]

		r.UUID = string(d.GetStringBytes("uuid"))
		r.Hitokoto = string(d.GetStringBytes("hitokoto"))
		r.Type = string(d.GetStringBytes("type"))
		r.From = string(d.GetStringBytes("from"))
		r.FromWho = string(d.GetStringBytes("from_who"))
		r.Creator = string(d.GetStringBytes("creator"))
		r.CreatorUID = d.GetInt("creator_uid")
		r.CommitFrom = string(d.GetStringBytes("commit_from"))
		r.Reviewer = d.GetInt("reviewer")
		r.Status = string(d.GetStringBytes("status"))

		createdAt, err := strconv.ParseInt(string(d.GetStringBytes("created_at")), 10, 64)
		if err != nil {
			createdAt = 0
		}
		r.CreatedAt = time.Unix(createdAt, 0)
	}
	return nil
}

type HitokotoUUIDMarkRequest struct {
	Token string
	UUID  string
}

func (r *HitokotoUUIDMarkRequest) FormatToValues() url.Values {
	var values = url.Values{}

	values.Add("token", r.Token)
	values.Add("uuid", r.UUID)

	return values
}

type HitokotoUUIDMarkResponse struct {
	Data []int `json:"data"`
}

func (r *HitokotoUUIDMarkResponse) Parse(data []byte) error {
	v, err := fastjson.ParseBytes(data)
	if err != nil {
		return err
	}

	r.Data = make([]int, 0)
	list, err := v.Array()
	if err != nil {
		return err
	}

	for _, d := range list {
		i, err := d.Int()
		if err != nil {
			i = 0
		}

		r.Data = append(r.Data, i)
	}
	return nil
}

type SentenceScore struct {
	Total        int `json:"total"`
	Participants int `json:"participants"`
	Average      int `json:"average"`
}

type HitokotoScorePostRequest struct {
	Token   string
	UUID    string
	Score   int
	Comment string
}

func (r *HitokotoScorePostRequest) FormatToValues() url.Values {
	var values = url.Values{}

	values.Add("token", r.Token)
	values.Add("uuid", r.UUID)
	values.Add("score", strconv.Itoa(r.Score))
	values.Add("comment", r.Comment)

	return values
}

type HitokotoScorePostResponse struct {
	Score         string         `json:"score"`
	Comment       string         `json:"comment"`
	SentenceUUID  string         `json:"sentence_uuid"`
	SentenceScore *SentenceScore `json:"sentence_score"`
}

func (r *HitokotoScorePostResponse) Parse(data []byte) error {
	v, err := fastjson.ParseBytes(data)
	if err != nil {
		return err
	}

	r.SentenceScore = new(SentenceScore)

	list, err := v.Array()
	if err != nil {
		return err
	}
	if len(list) > 0 {

		var d = list[0]

		r.Score = string(d.GetStringBytes("score"))
		r.Comment = string(d.GetStringBytes("comment"))
		r.SentenceUUID = string(d.GetStringBytes("sentence_uuid"))
		var sentenceScore = d.Get("sentence_score")
		r.SentenceScore.Average = sentenceScore.GetInt("average")
		r.SentenceScore.Participants = sentenceScore.GetInt("participants")
		r.SentenceScore.Total = sentenceScore.GetInt("total")

	}
	return nil
}

type Log struct {
	ID           int    `json:"id"`
	SentenceUUID string `json:"sentence_uuid"`
	UserID       int    `json:"user_id"`
	Score        int    `json:"score"`
	Comment      string `json:"comment"`
	UpdatedAt    string `json:"updated_at"`
	CreatedAt    string `json:"created_at"`
}

type HitokotoScoreGetRequest struct {
	Token string
	UUID  string
}

func (r *HitokotoScoreGetRequest) FormatToValues() url.Values {
	var values = url.Values{}

	values.Add("token", r.Token)
	values.Add("uuid", r.UUID)

	return values
}

type HitokotoScoreGetResponse struct {
	ID           int            `json:"id"`
	SentenceUUID string         `json:"sentence_uuid"`
	Score        *SentenceScore `json:"score"`
	Logs         []*Log         `json:"logs"`
	UpdatedAt    time.Time      `json:"updated_at"`
	CreatedAt    time.Time      `json:"created_at"`
}

func (r *HitokotoScoreGetResponse) Parse(data []byte) error {
	v, err := fastjson.ParseBytes(data)
	if err != nil {
		return err
	}

	r.Score = new(SentenceScore)
	list, err := v.Array()
	if err != nil {
		return err
	}
	if len(list) > 0 {

		var d = list[0]

		r.ID = d.GetInt("id")
		r.SentenceUUID = string(d.GetStringBytes("sentence_uuid"))

		var score = d.Get("score")
		r.Score.Average = score.GetInt("average")
		r.Score.Participants = score.GetInt("participants")
		r.Score.Total = score.GetInt("total")

		var logs = d.GetArray("logs")

		for _, log := range logs {
			var l = new(Log)
			l.ID = log.GetInt("id")
			l.SentenceUUID = string(log.GetStringBytes("sentence_uuid"))
			l.UserID = log.GetInt("user_id")
			l.Score = log.GetInt("score")
			l.Comment = string(log.GetStringBytes("comment"))
			l.UpdatedAt = string(log.GetStringBytes("updated_at"))
			l.CreatedAt = string(log.GetStringBytes("created_at"))
			r.Logs = append(r.Logs, l)
		}

		updatedAt, err := time.Parse(time.RFC3339Nano, string(d.GetStringBytes("updated_at")))
		if err != nil {
			updatedAt = time.Time{}
		}
		r.UpdatedAt = updatedAt

		createdAt, err := time.Parse(time.RFC3339Nano, string(d.GetStringBytes("created_at")))
		if err != nil {
			createdAt = time.Time{}
		}
		r.CreatedAt = createdAt

	}
	return nil
}

type HitokotoReportRequest struct {
	Token   string
	UUID    string
	Comment string
}

func (r *HitokotoReportRequest) FormatToValues() url.Values {
	var values = url.Values{}

	values.Add("token", r.Token)
	values.Add("uuid", r.UUID)
	values.Add("comment", r.Comment)

	return values
}

type HitokotoReportResponse struct {
	SentenceUUID string    `json:"sentence_uuid"`
	UserID       int       `json:"user_id"`
	Comment      string    `json:"comment"`
	UpdatedAt    time.Time `json:"updated_at"`
	CreatedAt    time.Time `json:"created_at"`
	ID           int       `json:"id"`
}

func (r *HitokotoReportResponse) Parse(data []byte) error {
	v, err := fastjson.ParseBytes(data)
	if err != nil {
		return err
	}

	list, err := v.Array()
	if err != nil {
		return err
	}
	if len(list) > 0 {

		var d = list[0]

		r.ID = d.GetInt("id")
		r.SentenceUUID = string(d.GetStringBytes("sentence_uuid"))
		r.UserID = d.GetInt("user_id")
		r.Comment = string(d.GetStringBytes("comment"))

		updatedAt, err := time.Parse(time.RFC3339Nano, string(d.GetStringBytes("updated_at")))
		if err != nil {
			updatedAt = time.Time{}
		}
		r.UpdatedAt = updatedAt

		createdAt, err := time.Parse(time.RFC3339Nano, string(d.GetStringBytes("created_at")))
		if err != nil {
			createdAt = time.Time{}
		}
		r.CreatedAt = createdAt

	}
	return nil
}
