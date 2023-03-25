package op

import (
	"net/url"
	"time"

	"github.com/valyala/fastjson"
)

type LoginRequest struct {
	Email    string
	Password string
}

func (lr *LoginRequest) FormatToValues() url.Values {
	var values = url.Values{}

	values.Add("email", lr.Email)
	values.Add("password", lr.Password)

	return values
}

type LoginResponse struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Email       string `json:"email"`
	IsSuspended int    `json:"is_suspended"`
	IsAdmin     int    `json:"is_admin"`
	IsReviewer  int    `json:"is_reviewer"`
	CreatedFrom string `json:"created_from"`
	Token       string `json:"token"`

	UpdatedAt       time.Time `json:"updated_at"`
	CreatedAt       time.Time `json:"created_at"`
	EmailVerifiedAt time.Time `json:"email_verified_at"`
}

func (lr *LoginResponse) Parse(data []byte) error {
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
		lr.ID = d.GetInt("id")
		lr.Name = string(d.GetStringBytes("name"))
		lr.Email = string(d.GetStringBytes("email"))
		lr.IsSuspended = d.GetInt("is_suspended")
		lr.IsAdmin = d.GetInt("is_admin")
		lr.IsReviewer = d.GetInt("is_reviewer")
		lr.CreatedFrom = string(d.GetStringBytes("created_from"))
		lr.Token = string(d.GetStringBytes("token"))

		emailVerifiedAt, err := time.Parse(time.RFC3339Nano, string(d.GetStringBytes("email_verified_at")))
		if err != nil {
			emailVerifiedAt = time.Time{}
		}
		lr.EmailVerifiedAt = emailVerifiedAt

		creatAt, err := time.Parse(time.RFC3339Nano, string(d.GetStringBytes("created_at")))
		if err != nil {
			creatAt = time.Time{}
		}
		lr.CreatedAt = creatAt

		updatedAt, err := time.Parse(time.RFC3339Nano, string(d.GetStringBytes("updated_at")))
		if err != nil {
			updatedAt = time.Time{}
		}
		lr.UpdatedAt = updatedAt

	}

	return nil
}
