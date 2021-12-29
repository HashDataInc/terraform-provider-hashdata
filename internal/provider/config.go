package provider

import (
	"context"
	"github.com/HashDataInc/terraform-provider-hashdata/internal/provider/cloudmgr"
	"golang.org/x/oauth2"
	"net/url"
)

type Config struct {
	Username     string
	Password     string
	ClientId     string
	ClientSecret string
	EndPoint     string
}

func (c *Config) Client() (interface{}, error) {
	cfg := cloudmgr.NewConfiguration()
	configOauth2 := oauth2.Config{
		ClientID:     c.ClientId,
		ClientSecret: c.ClientSecret,
		Endpoint: oauth2.Endpoint{
			TokenURL: c.EndPoint + TOKEN_PATH,
		},
	}
	token, err := configOauth2.PasswordCredentialsToken(context.Background(), c.Username, c.Password)
	if err != nil {
		return nil, err
	}
	u, err := url.Parse(c.EndPoint)
	if err != nil {
		panic(err)
	}
	cfg.Scheme = u.Scheme
	var header = make(map[string]string)
	header[DEFAULT_HEADER_KEY] = AUTH_PREFIX + token.AccessToken
	cfg.DefaultHeader = header
	cfg.Servers = cloudmgr.ServerConfigurations{
		{
			URL:         u.Host,
			Description: "Default server",
		},
	}
	return cloudmgr.NewAPIClient(cfg), nil
}
