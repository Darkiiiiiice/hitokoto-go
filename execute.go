package hitokoto

import (
	"net/http"
	"strings"

	"github.com/go-resty/resty/v2"
	"github.com/mariomang/hitokoto/op"
)

type Executor struct {
	APIGateway string
	Token      string
}

func (e *Executor) Do(command op.Command) (err error) {

	// Create a Resty Client
	client := resty.New()

	var api = command.API()
	var values = command.Values()
	if api.Auth {
		values.Add("token", e.Token)
	}

	var path = api.Api
	if api.PathVar != "" {
		var pathVar = values.Get(api.PathVar)
		values.Del(api.PathVar)
		path = strings.ReplaceAll(path, ":"+api.PathVar, pathVar)
	}

	var request = client.R().
		SetQueryParamsFromValues(values).
		SetHeader("Content-Type", api.ContentType)

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

	if err := command.Parse(body); err != nil {
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
