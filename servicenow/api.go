package servicenow

// SnowClient is a struct to wrap HTTPClient and with aditional service now parameters
type SnowClient struct {
	access       AccessConfig
	defaultTable string
}

// AccessConfig is a struct to store sensitive snow data
type AccessConfig struct {
	Credential string
	Endpoint   string
}

// opsSnowConfig is type to create new options
type opsSnowConfig func(s *SnowClient) error

// withDefaultTable is Generic Helper so set a default table
func withDefaultTable(table string) opsSnowConfig {
	return func(s *SnowClient) error {
		s.defaultTable = table
		return nil
	}
}

// NewSnowClient create a new nowClient setted
func NewSnowClient(access AccessConfig, opts ...opsSnowConfig) (*SnowClient, error) {
	snow := &SnowClient{
		access: access,
	}

	for _, opts := range opts {
		if err := opts(snow); err != nil {
			return nil, err
		}
	}

	return snow, nil
}