package snow

import (
	"context"
	"fmt"
	"time"

	"github.com/matheuspolitano/servicenow-go-sdk/myrequest"
)

// SnowClient is a struct to wrap HTTPClient and with aditional service now parameters
type SnowClient struct {
	access       AccessConfig
	httpClient *myrequest.HTTPClient
	timeout time.Duration
}

// AccessConfig is a struct to store sensitive snow data
type AccessConfig struct {
	endpoint   string
	username string
	password string
}

// NewAccessConfig create a new access config
func NewAccessConfig(username, password, endpoint string) *AccessConfig{
	return &AccessConfig{
		endpoint: endpoint,
		username: username,
		password: password,
	}
}

// opsSnowConfig is type to create new options
type opsSnowConfig func(s *SnowClient) error


// NewSnowClient create a new nowClient
func NewSnowClient(access AccessConfig, opts ...opsSnowConfig) (*SnowClient, error) {
	snow := &SnowClient{
		access: access,
		timeout: 30 * time.Second,
	}

	for _, opts := range opts {
		if err := opts(snow); err != nil {
			return nil, err
		}
	}
	client := myrequest.NewHTTPClient(snow.access.endpoint, snow.timeout, nil, access.username, access.password)
	snow.httpClient = client
	return snow, nil
}

func (sn *SnowClient) ExecuteQuery(ctx context.Context, table string,query QueryElement) (map[string]any, error){
	url := fmt.Sprintf("/%s?sysparm_query=%s", table, query.String())
	return sn.httpClient.Get(ctx, url)
}
