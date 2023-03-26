package op

import (
	"net/url"
	"time"

	"github.com/valyala/fastjson"
)

type Data struct {
	ID        int       `json:"id"`
	Text      string    `json:"text"`
	Level     string    `json:"level"`
	Property  int       `json:"property"`
	UpdatedAt time.Time `json:"updated_at"`
	CreatedAt time.Time `json:"created_at"`
}

type MarkRequest struct {
	Token string
}

func (r *MarkRequest) FormatToValues() url.Values {
	var values = url.Values{}

	values.Add("token", r.Token)

	return values
}

type MarkResponse struct {
	Data []*Data `json:"data"`
}

func (r *MarkResponse) Parse(data []byte) error {
	v, err := fastjson.ParseBytes(data)
	if err != nil {
		return err
	}

	r.Data = make([]*Data, 0)

	list, err := v.Array()
	if err != nil {
		return err
	}
	for _, item := range list {
		var d = new(Data)

		d.ID = item.GetInt("id")
		d.Text = string(item.GetStringBytes("text"))
		d.Level = string(item.GetStringBytes("level"))
		d.Property = item.GetInt("property")

		updatedAt, err := time.Parse(time.RFC3339Nano, string(item.GetStringBytes("updated_at")))
		if err != nil {
			updatedAt = time.Time{}
		}
		d.UpdatedAt = updatedAt

		createdAt, err := time.Parse(time.RFC3339Nano, string(item.GetStringBytes("created_at")))
		if err != nil {
			createdAt = time.Time{}
		}
		d.CreatedAt = createdAt
		r.Data = append(r.Data, d)
	}
	return nil
}
