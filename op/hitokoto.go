package op

import (
	"net/url"
	"strconv"
	"time"

	"github.com/valyala/fastjson"
)

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
	values.Add(":uuid", r.UUID)

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
	values.Add(":uuid", r.UUID)

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
	values.Add(":uuid", r.UUID)
	values.Add("score", strconv.Itoa(r.Score))
	values.Add("comment", r.Comment)

	return values
}

type HitokotoScorePostResponse struct {
	Score         string        `json:"score"`
	Comment       string        `json:"comment"`
	SentenceUUID  string        `json:"sentence_uuid"`
	SentenceScore SentenceScore `json:"sentence_score"`
}

func (r *HitokotoScorePostResponse) Parse(data []byte) error {
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
