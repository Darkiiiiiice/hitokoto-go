package op

import (
	"net/url"
	"strconv"
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

type UserTokenRequest struct {
	Token string
}

func (r *UserTokenRequest) FormatToValues() url.Values {
	var values = url.Values{}

	values.Add("token", r.Token)
	return values
}

type UserTokenResponse struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Token string `json:"token"`
}

func (r *UserTokenResponse) Parse(data []byte) error {
	v, err := fastjson.ParseBytes(data)
	if err != nil {
		return err
	}

	list, err := v.Array()
	if err != nil {
		return err
	}
	if len(list) > 0 {
		d := list[0].Get("user")
		r.ID = d.GetInt("id")
		r.Name = string(d.GetStringBytes("name"))
		r.Email = string(d.GetStringBytes("email"))
		r.Token = string(d.GetStringBytes("token"))
	}

	return nil
}

type UserTokenRefreshRequest struct {
	Token string
}

func (r *UserTokenRefreshRequest) FormatToValues() url.Values {
	var values = url.Values{}

	values.Add("token", r.Token)
	return values
}

type UserTokenRefreshResponse struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Token string `json:"token"`
}

func (r *UserTokenRefreshResponse) Parse(data []byte) error {
	v, err := fastjson.ParseBytes(data)
	if err != nil {
		return err
	}

	list, err := v.Array()
	if err != nil {
		return err
	}
	if len(list) > 0 {
		d := list[0].Get("user")
		r.ID = d.GetInt("id")
		r.Name = string(d.GetStringBytes("name"))
		r.Email = string(d.GetStringBytes("email"))
		r.Token = string(d.GetStringBytes("token"))
	}

	return nil
}

type UserEmailVerifyRequest struct {
	Token string
}

func (r *UserEmailVerifyRequest) FormatToValues() url.Values {
	var values = url.Values{}

	values.Add("token", r.Token)
	return values
}

type UserEmailVerifyResponse struct {
}

func (r *UserEmailVerifyResponse) Parse(data []byte) error {
	return nil
}

type UserPasswordRequest struct {
	Token       string
	Password    string
	NewPassword string
}

func (r *UserPasswordRequest) FormatToValues() url.Values {
	var values = url.Values{}

	values.Add("token", r.Token)
	values.Add("password", r.Password)
	values.Add("new_password", r.NewPassword)
	return values
}

type UserPasswordResponse struct {
}

func (r *UserPasswordResponse) Parse(data []byte) error {
	return nil
}

type UserEmailRequest struct {
	Token    string
	Password string
	Email    string
}

func (r *UserEmailRequest) FormatToValues() url.Values {
	var values = url.Values{}

	values.Add("token", r.Token)
	values.Add("password", r.Password)
	values.Add("email", r.Email)
	return values
}

type UserEmailResponse struct {
}

func (r *UserEmailResponse) Parse(data []byte) error {
	return nil
}

type UserNotificationSettingsGetRequest struct {
	Token string
}

func (r *UserNotificationSettingsGetRequest) FormatToValues() url.Values {
	var values = url.Values{}

	values.Add("token", r.Token)
	return values
}

type UserNotificationSettingsGetResponse struct {
	ID     int `json:"id"`
	UserID int `json:"user_id"`

	EmailNotificationGlobal           int `json:"email_notification_global"`
	EmailNotificationHitokotoAppended int `json:"email_notification_hitokoto_appended"`
	EmailNotificationHitokotoReviewed int `json:"email_notification_hitokoto_reviewed"`
	EmailNotificationPollCreated      int `json:"email_notification_poll_created"`
	EmailNotificationPollResult       int `json:"email_notification_poll_result"`
	EmailNotificationPollDailyReport  int `json:"email_notification_poll_daily_report"`

	UpdatedAt time.Time `json:"updated_at"`
	CreatedAt time.Time `json:"created_at"`
}

func (r *UserNotificationSettingsGetResponse) Parse(data []byte) error {
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
		r.UserID = d.GetInt("user_id")
		r.EmailNotificationGlobal = d.GetInt("email_notification_global")
		r.EmailNotificationHitokotoAppended = d.GetInt("email_notification_hitokoto_appended")
		r.EmailNotificationHitokotoReviewed = d.GetInt("email_notification_hitokoto_reviewed")
		r.EmailNotificationPollCreated = d.GetInt("email_notification_poll_created")
		r.EmailNotificationPollResult = d.GetInt("email_notification_poll_result")
		r.EmailNotificationPollDailyReport = d.GetInt("email_notification_poll_daily_report")

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

type UserNotificationSettingsPutRequest struct {
	Token string

	EmailGlobal           bool `json:"email_global,omitempty"`
	EmailHitokotoAppended bool `json:"email_hitokoto_appended,omitempty"`
	EmailHitokotoReviewed bool `json:"email_hitokoto_reviewed,omitempty"`
	EmailPollCreated      bool `json:"email_poll_created,omitempty"`
	EmailPollResult       bool `json:"email_poll_result,omitempty"`
	EmailPollReportDaily  bool `json:"email_poll_report_daily,omitempty"`
}

func (r *UserNotificationSettingsPutRequest) FormatToValues() url.Values {
	var values = url.Values{}

	values.Add("token", r.Token)
	values.Add("email_global", strconv.FormatBool(r.EmailGlobal))
	values.Add("email_hitokoto_appended", strconv.FormatBool(r.EmailHitokotoAppended))
	values.Add("email_hitokoto_reviewed", strconv.FormatBool(r.EmailHitokotoReviewed))
	values.Add("email_poll_created", strconv.FormatBool(r.EmailPollCreated))
	values.Add("email_poll_result", strconv.FormatBool(r.EmailPollResult))
	values.Add("email_poll_report_daily", strconv.FormatBool(r.EmailPollReportDaily))

	return values
}

type UserNotificationSettingsPutResponse struct {
	ID     int `json:"id"`
	UserID int `json:"user_id"`

	EmailNotificationGlobal           int `json:"email_notification_global"`
	EmailNotificationHitokotoAppended int `json:"email_notification_hitokoto_appended"`
	EmailNotificationHitokotoReviewed int `json:"email_notification_hitokoto_reviewed"`
	EmailNotificationPollCreated      int `json:"email_notification_poll_created"`
	EmailNotificationPollResult       int `json:"email_notification_poll_result"`
	EmailNotificationPollDailyReport  int `json:"email_notification_poll_daily_report"`

	UpdatedAt time.Time `json:"updated_at"`
	CreatedAt time.Time `json:"created_at"`
}

func (r *UserNotificationSettingsPutResponse) Parse(data []byte) error {
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
		r.UserID = d.GetInt("user_id")
		r.EmailNotificationGlobal = d.GetInt("email_notification_global")
		r.EmailNotificationHitokotoAppended = d.GetInt("email_notification_hitokoto_appended")
		r.EmailNotificationHitokotoReviewed = d.GetInt("email_notification_hitokoto_reviewed")
		r.EmailNotificationPollCreated = d.GetInt("email_notification_poll_created")
		r.EmailNotificationPollResult = d.GetInt("email_notification_poll_result")
		r.EmailNotificationPollDailyReport = d.GetInt("email_notification_poll_daily_report")

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
