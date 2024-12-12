// This file was auto-generated by Fern from our API Definition.

package configurations

import (
	context "context"
	fmt "fmt"
	fluidstackgosdk "github.com/fluidstackio/fluidstack-go-sdk"
	core "github.com/fluidstackio/fluidstack-go-sdk/core"
	internal "github.com/fluidstackio/fluidstack-go-sdk/internal"
	option "github.com/fluidstackio/fluidstack-go-sdk/option"
	http "net/http"
	os "os"
)

type Client struct {
	baseURL string
	caller  *internal.Caller
	header  http.Header
}

func NewClient(opts ...option.RequestOption) *Client {
	options := core.NewRequestOptions(opts...)
	if options.ApiKey == "" {
		options.ApiKey = os.Getenv("FLUIDSTACK_API_KEY")
	}
	return &Client{
		baseURL: options.BaseURL,
		caller: internal.NewCaller(
			&internal.CallerParams{
				Client:      options.HTTPClient,
				MaxAttempts: options.MaxAttempts,
			},
		),
		header: options.ToHeader(),
	}
}

// List available configurations including GPU type, GPU count, RAM size, and disk size.
func (c *Client) List(
	ctx context.Context,
	request *fluidstackgosdk.ConfigurationsListRequest,
	opts ...option.RequestOption,
) ([]*fluidstackgosdk.ConfigurationResponse, error) {
	options := core.NewRequestOptions(opts...)
	baseURL := internal.ResolveBaseURL(
		options.BaseURL,
		c.baseURL,
		"https://platform.fluidstack.io",
	)
	endpointURL := baseURL + "/list_available_configurations"
	headers := internal.MergeHeaders(
		c.header.Clone(),
		options.ToHeader(),
	)
	if request.ApiKey != nil {
		headers.Add("api-key", fmt.Sprintf("%v", *request.ApiKey))
	}
	errorCodes := internal.ErrorCodes{
		401: func(apiError *core.APIError) error {
			return &fluidstackgosdk.UnauthorizedError{
				APIError: apiError,
			}
		},
		422: func(apiError *core.APIError) error {
			return &fluidstackgosdk.UnprocessableEntityError{
				APIError: apiError,
			}
		},
	}

	var response []*fluidstackgosdk.ConfigurationResponse
	if err := c.caller.Call(
		ctx,
		&internal.CallParams{
			URL:             endpointURL,
			Method:          http.MethodGet,
			Headers:         headers,
			MaxAttempts:     options.MaxAttempts,
			BodyProperties:  options.BodyProperties,
			QueryParameters: options.QueryParameters,
			Client:          options.HTTPClient,
			Response:        &response,
			ErrorDecoder:    internal.NewErrorDecoder(errorCodes),
		},
	); err != nil {
		return nil, err
	}
	return response, nil
}
