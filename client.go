package rippling

import "github.com/deepmap/oapi-codegen/pkg/securityprovider"

var RipplingAPIEndpoint string = "https://api.rippling.com"

func New(key string) (*Rippling, error) {
	auth, err := securityprovider.NewSecurityProviderBearerToken(key)
	if err != nil {
		return nil, err
	}
	client, err := NewClient(RipplingAPIEndpoint, WithRequestEditorFn(auth.Intercept))
	if err != nil {
		return nil, err
	}
	return client, nil
}
