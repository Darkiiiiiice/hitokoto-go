package op

import (
	"net/url"
	"time"

	"github.com/valyala/fastjson"
)

type UserRequest struct {
	Token string
}

func (r *UserRequest) FormatToValues() url.Values {
	var values = url.Values{}

	values.Add("token", r.Token)
	return values
}

type UserResponse struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Email       string `json:"email"`
	IsSuspended int    `json:"is_suspended"`
	IsAdmin     int    `json:"is_admin"`
	IsReviewer  int    `json:"is_reviewer"`
	CreatedFrom string `json:"created_from"`

	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
	EmailVerifiedAt time.Time `json:"email_verified_at"`
}

func (r *UserResponse) Parse(data []byte) error {
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
		r.ID = d.GetInt("id")
		r.Name = string(d.GetStringBytes("name"))
		r.Email = string(d.GetStringBytes("email"))
		r.IsSuspended = d.GetInt("is_suspended")
		r.IsAdmin = d.GetInt("is_admin")
		r.IsReviewer = d.GetInt("is_reviewer")
		r.CreatedFrom = string(d.GetStringBytes("created_from"))

		emailVerifiedAt, err := time.Parse(time.RFC3339Nano, string(d.GetStringBytes("email_verified_at")))
		if err != nil {
			emailVerifiedAt = time.Time{}
		}
		r.EmailVerifiedAt = emailVerifiedAt

		creatAt, err := time.Parse(time.RFC3339Nano, string(d.GetStringBytes("created_at")))
		if err != nil {
			creatAt = time.Time{}
		}
		r.CreatedAt = creatAt

		updatedAt, err := time.Parse(time.RFC3339Nano, string(d.GetStringBytes("updated_at")))
		if err != nil {
			updatedAt = time.Time{}
		}
		r.UpdatedAt = updatedAt
	}

	return nil
}
