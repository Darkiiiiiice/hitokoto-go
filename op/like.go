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

type LikeRequest struct {
	SentenceUuid string
}

func (r *LikeRequest) FormatToValues() url.Values {
	var values = url.Values{}

	values.Add("sentence_uuid", r.SentenceUuid)

	return values
}

type LikeResponse struct {
	Sets  []*Set `json:"sets"`
	Total int    `json:"total"`
}

func (r *LikeResponse) Parse(data []byte) error {
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
