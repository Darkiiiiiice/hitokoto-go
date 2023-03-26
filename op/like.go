package op

import (
	"net/url"
	"time"

	"github.com/valyala/fastjson"
)

type Set struct {
	UserId      int       `json:"user_id"`
	CreatedTime time.Time `json:"created_time"`
}

type LikeGetRequest struct {
	SentenceUuid string
}

func (r *LikeGetRequest) FormatToValues() url.Values {
	var values = url.Values{}

	values.Add("sentence_uuid", r.SentenceUuid)

	return values
}

type LikeGetResponse struct {
	Sets  []*Set `json:"sets"`
	Total int    `json:"total"`
}

func (r *LikeGetResponse) Parse(data []byte) error {
	v, err := fastjson.ParseBytes(data)
	if err != nil {
		return err
	}

	list, err := v.Array()
	if err != nil {
		return err
	}
	if len(list) > 0 {
		d := list[0]
		r.Total = d.GetInt("total")
		r.Sets = make([]*Set, 0)

		for _, item := range d.GetArray("sets") {
			var hitokoto = new(Set)
			hitokoto.UserId = item.GetInt("user_id")

			createdAt, err := time.Parse(time.RFC3339Nano, string(item.GetStringBytes("created_time")))
			if err != nil {
				createdAt = time.Time{}
			}
			hitokoto.CreatedTime = createdAt

			r.Sets = append(r.Sets, hitokoto)
		}
	}
	return nil
}

type LikePostRequest struct {
	SentenceUuid string
}

func (r *LikePostRequest) FormatToValues() url.Values {
	var values = url.Values{}

	values.Add("sentence_uuid", r.SentenceUuid)

	return values
}

type LikePostResponse struct {
	IP     string `json:"ip"`
	UserId int    `json:"user_id"`
}

func (r *LikePostResponse) Parse(data []byte) error {
	v, err := fastjson.ParseBytes(data)
	if err != nil {
		return err
	}

	list, err := v.Array()
	if err != nil {
		return err
	}
	if len(list) > 0 {
		d := list[0]
		r.IP = string(d.GetStringBytes("ip"))
		r.UserId = d.GetInt("user_id")
	}
	return nil
}
