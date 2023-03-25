package hitokoto

import (
	"net/http"
	"strings"

	"github.com/go-resty/resty/v2"
	"github.com/mariomang/hitokoto/constants"
	"github.com/valyala/fastjson"
)

type Executor struct {
	APIGateway string
	UserAgent  string
}

func NewExecutor(token string) *Executor {
	return &Executor{
		APIGateway: constants.Scheme + constants.APIGatewayV1,
	}
}

func (e *Executor) WithUserAgent(userAgent string) *Executor {
	e.UserAgent = userAgent
	return e
}

func (e *Executor) Do(api *constants.API, req Request, resp Response) (err error) {

	// Create a Resty Client
	client := resty.New()

	var values = req.FormatToValues()

	var path = e.APIGateway + api.Api
	if api.PathVar != "" {
		var pathVar = values.Get(api.PathVar)
		values.Del(api.PathVar)
		path = strings.ReplaceAll(path, api.PathVar, pathVar)
	}

	var userAgent = constants.UserAgent
	if userAgent != "" {
		userAgent = e.UserAgent
	}

	var request = client.R().
		SetQueryParamsFromValues(values).
		SetHeader("Content-Type", api.ContentType).
		SetHeader("Host", constants.Host).
		SetHeader("User-Agent", userAgent)

	var body []byte
	switch api.Method {
	case http.MethodGet:
		body, err = e.get(request, path)
		if err != nil {
			return err
		}
	case http.MethodPost:
		body, err = e.post(request, path)
		if err != nil {
			return err
		}
	case http.MethodPut:
		body, err = e.put(request, path)
		if err != nil {
			return err
		}
	default:
	}

	v, err := fastjson.ParseBytes(body)
	if err != nil {
		return err
	}

	var status = v.GetInt("status")
	if status != 200 {
		return NewHitokotoError(status, string(v.GetStringBytes("message")), v.GetInt64("ts"))
	}

	var buf = v.Get("data").MarshalTo(nil)

	if err := resp.Parse(buf); err != nil {
		return err
	}

	return nil
}

func (e *Executor) get(request *resty.Request, path string) ([]byte, error) {

	resp, err := request.Get(path)
	if err != nil {
		return nil, err
	}
	return resp.Body(), nil
}

func (e *Executor) post(request *resty.Request, path string) ([]byte, error) {

	resp, err := request.Post(path)
	if err != nil {
		return nil, err
	}
	return resp.Body(), nil
}

func (e *Executor) put(request *resty.Request, path string) ([]byte, error) {

	resp, err := request.Put(path)
	if err != nil {
		return nil, err
	}
	return resp.Body(), nil
}
