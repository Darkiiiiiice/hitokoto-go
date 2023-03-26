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
