package client

import (
	"errors"
	"reflect"
	"strconv"

	"github.com/go-resty/resty/v2"
	"github.com/rs-pro/sshkeymanager/api"
)

type Client struct {
	ApiKey    string
	ApiHost   string
	ApiClient *resty.Client
	// ApiComment is used to add extra info (like user that performed the action) to keyserver log
	ApiComment string
	Host       string
	Port       string
	User       string
}

func NewClient(apiHost, apiKey string) *Client {
	r := resty.New()
	r.SetDebug(true)

	return &Client{
		ApiKey:    apiKey,
		ApiHost:   apiHost,
		ApiClient: r,
	}
}

func (c *Client) R() *resty.Request {
	return c.ApiClient.R().
		SetHeader("X-Api-Key", c.ApiKey).
		SetQueryParams(map[string]string{
			"host":    c.Host,
			"port":    c.Port,
			"user":    c.User,
			"comment": c.ApiComment,
		})
}

func (c *Client) WithConfig(host, port, user string) *Client {
	return &Client{
		ApiKey:    c.ApiKey,
		ApiHost:   c.ApiHost,
		ApiClient: resty.New(),
		Host:      host,
		Port:      port,
		User:      user,
	}
}

func (c *Client) GetURL(method string) string {
	return c.ApiHost + "/" + method
}

func (c *Client) Execute(method string, request, result interface{}) (interface{}, error) {
	r, err := c.R().
		SetBody(request).
		SetResult(result).
		SetError(&api.BasicError{}).
		Post(c.GetURL(method))

	if err != nil {
		// err is returned in case of network or other non-keymanager error
		return nil, err
	}

	if r.IsSuccess() {
		return r.Result(), nil
	}

	if r.IsError() {
		// in case of sshkeymanager error (not network etc) we do not return err, insted we set Err field in the response
		out := result
		ec := r.Error().(*api.BasicError)
		var e error
		if ec.Err != nil {
			e = errors.New(*ec.Err)
		} else {
			e = errors.New("bad status code " + strconv.Itoa(r.StatusCode()))
		}

		reflect.Indirect(reflect.ValueOf(&out).Elem().Elem()).FieldByName("Err").Set(reflect.ValueOf(api.MakeKmError(e)))
		return out, nil
	}

	return nil, errors.New("unexpected result (not success or error)")
}
