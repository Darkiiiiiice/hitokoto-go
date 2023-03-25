package op

import (
	"net/url"
	"time"

	"github.com/valyala/fastjson"
)

type RegisterRequest struct {
	Name     string
	Email    string
	Password string
}

func (r *RegisterRequest) FormatToValues() url.Values {
	var values = url.Values{}

	values.Add("name", r.Name)
	values.Add("email", r.Email)
	values.Add("password", r.Password)

	return values
}

type RegisterResponse struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Token string `json:"token"`

	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
	EmailVerifiedAt time.Time `json:"email_verified_at"`
}

func (r *RegisterResponse) Parse(data []byte) error {
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
		r.Token = string(d.GetStringBytes("token"))

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
