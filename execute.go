package hitokoto

import (
	"net/http"

	"github.com/go-resty/resty/v2"
	"github.com/mariomang/hitokoto-go/constants"
	"github.com/valyala/fastjson"
)

// Executor is used to execute commands on the API gateway
type Executor struct {
	// APIGateway is the URL of the API gateway
	APIGateway string
	// UserAgent is the user agent used to execute the command
	UserAgent string

	// GlobalAPI is used to set the global API
	GlobalAPI bool
}

// NewExecutor initializes an Executor with the default APIGateway URL.
func NewExecutor() *Executor {
	return &Executor{
		APIGateway: constants.Scheme + constants.APIGatewayV1,
	}
}

// WithUserAgent sets the User-Agent header to the given value. If the
// value is empty, the User-Agent header will not be sent.
func (e *Executor) WithUserAgent(userAgent string) *Executor {
	e.UserAgent = userAgent
	return e
}

func (e *Executor) WithGlobalAPI() *Executor {
	e.GlobalAPI = true
	return e
}

// Do executes the API request and response
func (e *Executor) Do(api *constants.API, req Request, resp Response) (err error) {

	var requestSentence = false
	if api.Api == "" {
		requestSentence = true
	}
	// Create a Resty Client
	client := resty.New()

	var values = req.FormatToValues()

	var userAgent = constants.UserAgent
	if e.UserAgent != "" {
		userAgent = e.UserAgent
	}

	var request = client.R().
		SetQueryParamsFromValues(values).
		SetHeader("Content-Type", api.ContentType).
		SetHeader("Host", constants.Host).
		SetHeader("User-Agent", userAgent)

	var path = e.APIGateway + api.Api
	if requestSentence {
		if e.GlobalAPI {
			path = constants.Scheme + constants.SentenceGlobalAPI
			request = request.SetHeader("Host", constants.SentenceGlobalAPI)
		} else {
			path = constants.Scheme + constants.SentenceAbroadAPI
			request = request.SetHeader("Host", constants.SentenceAbroadAPI)
		}
	}

	if api.PathVar != "" {
		request = request.SetPathParam("uuid", values.Get("uuid"))
		values.Del(api.PathVar)
	}

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

	if requestSentence {
		return resp.Parse(body)
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

// get sends a GET request to path, and returns the response body.
func (e *Executor) get(request *resty.Request, path string) ([]byte, error) {

	resp, err := request.Get(path)
	if err != nil {
		return nil, err
	}
	return resp.Body(), nil
}

// post sends a POST request to path, and returns the response body.
func (e *Executor) post(request *resty.Request, path string) ([]byte, error) {

	resp, err := request.Post(path)
	if err != nil {
		return nil, err
	}
	return resp.Body(), nil
}

// put sends a PUT request to path, and returns the response body.
func (e *Executor) put(request *resty.Request, path string) ([]byte, error) {

	resp, err := request.Put(path)
	if err != nil {
		return nil, err
	}
	return resp.Body(), nil
}
