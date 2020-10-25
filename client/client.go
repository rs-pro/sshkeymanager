package client

import (
	"github.com/go-resty/resty/v2"
)

type Client struct {
	ApiKey    string
	ApiHost   string
	ApiClient *resty.Client
	Host      string
	Port      string
	User      string
}

func NewClient(apiHost, apiKey string) *Client {
	return &Client{
		ApiKey:    apiKey,
		ApiHost:   apiHost,
		ApiClient: resty.New(),
	}
}

func (c *Client) R() *resty.Request {
	return c.ApiClient.R().
		SetHeader("X-Api-Key", c.ApiKey).
		SetQueryParams(map[string]string{
			"host": c.Host,
			"port": c.Port,
			"user": c.User,
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

func (c *Client) Execute(method string, request, result interface{}) (*resty.Response, error) {
	return c.R().
		SetBody(request).
		SetResult(result).
		Post(c.GetURL(method))
}
