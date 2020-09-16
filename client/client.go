package client
import (
"github.com/go-resty/resty"
)
		

type Client struct {
	ApiKey string
	ApiHost string
	ApiClient *resty.Client
	Host string
	Port string
	User string
}

func NewClient(apiKey, apiHost string) *Client {
	return &Client{
		ApiKey: apiKey,
		ApiHost: apiHost,
		ApiClient: resty.New(),
	}
}

func (c *Client) R() *resty.R {
	return c.ApiClient.R().
		SetHeader("X-Api-Key", c.ApiKey).
		SetQueryParams(map[string]string{
				"host": c.Host,
				"port": c.Port,
				"user":c.User,
		})
}

func (c *Client) GetUsers() []passwd.User, error {
	request := {}
      SetBody(api.GetUsersRequest{

			}).
      SetResult(&AuthSuccess{}).    // or SetResult(AuthSuccess{}).
      Post(c.ApiHost + "/" + url)
}
