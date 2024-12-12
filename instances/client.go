// This file was auto-generated by Fern from our API Definition.

package instances

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

// This endpoint is used to retrieve a list of all instances associated with the authenticated user.
func (c *Client) List(
	ctx context.Context,
	request *fluidstackgosdk.InstancesListRequest,
	opts ...option.RequestOption,
) ([]*fluidstackgosdk.ListInstanceResponse, error) {
	options := core.NewRequestOptions(opts...)
	baseURL := internal.ResolveBaseURL(
		options.BaseURL,
		c.baseURL,
		"https://platform.fluidstack.io",
	)
	endpointURL := baseURL + "/instances"
	queryParams, err := internal.QueryValues(request)
	if err != nil {
		return nil, err
	}
	if len(queryParams) > 0 {
		endpointURL += "?" + queryParams.Encode()
	}
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

	var response []*fluidstackgosdk.ListInstanceResponse
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

// This endpoint is used to create a new instance. You must provide a custom `name` for the instance, its `gpu_type`, and the name of its `ssh_key`.
//
// If no values are provided for the `gpu_count` and `operating_system_label`, the default values of `1` and `ubuntu_20_04_lts_nvidia` are used respectively.
func (c *Client) Create(
	ctx context.Context,
	request *fluidstackgosdk.CreateInstanceRequest,
	opts ...option.RequestOption,
) (*fluidstackgosdk.CreateInstanceResponse, error) {
	options := core.NewRequestOptions(opts...)
	baseURL := internal.ResolveBaseURL(
		options.BaseURL,
		c.baseURL,
		"https://platform.fluidstack.io",
	)
	endpointURL := baseURL + "/instances"
	headers := internal.MergeHeaders(
		c.header.Clone(),
		options.ToHeader(),
	)
	if request.ApiKey != nil {
		headers.Add("api-key", fmt.Sprintf("%v", *request.ApiKey))
	}
	headers.Set("Content-Type", "application/json")
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

	var response *fluidstackgosdk.CreateInstanceResponse
	if err := c.caller.Call(
		ctx,
		&internal.CallParams{
			URL:             endpointURL,
			Method:          http.MethodPost,
			Headers:         headers,
			MaxAttempts:     options.MaxAttempts,
			BodyProperties:  options.BodyProperties,
			QueryParameters: options.QueryParameters,
			Client:          options.HTTPClient,
			Request:         request,
			Response:        &response,
			ErrorDecoder:    internal.NewErrorDecoder(errorCodes),
		},
	); err != nil {
		return nil, err
	}
	return response, nil
}

// This endpoint is used to retrieve a single instance associated with the authenticated user by its ID.
//
//	This endpoint returns HTTP 202 Accepted code if the instance is still pending. Otherwise, it returns HTTP 200 OK code.
func (c *Client) Get(
	ctx context.Context,
	instanceId string,
	request *fluidstackgosdk.InstancesGetRequest,
	opts ...option.RequestOption,
) (*fluidstackgosdk.InstanceResponse, error) {
	options := core.NewRequestOptions(opts...)
	baseURL := internal.ResolveBaseURL(
		options.BaseURL,
		c.baseURL,
		"https://platform.fluidstack.io",
	)
	endpointURL := internal.EncodeURL(
		baseURL+"/instances/%v",
		instanceId,
	)
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

	var response *fluidstackgosdk.InstanceResponse
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

// This endpoint is used to terminate an existing instance by its ID.
func (c *Client) Delete(
	ctx context.Context,
	instanceId string,
	request *fluidstackgosdk.InstancesDeleteRequest,
	opts ...option.RequestOption,
) (interface{}, error) {
	options := core.NewRequestOptions(opts...)
	baseURL := internal.ResolveBaseURL(
		options.BaseURL,
		c.baseURL,
		"https://platform.fluidstack.io",
	)
	endpointURL := internal.EncodeURL(
		baseURL+"/instances/%v",
		instanceId,
	)
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

	var response interface{}
	if err := c.caller.Call(
		ctx,
		&internal.CallParams{
			URL:             endpointURL,
			Method:          http.MethodDelete,
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

// This endpoint is used to stop an existing instance by its ID.
func (c *Client) Stop(
	ctx context.Context,
	instanceId string,
	request *fluidstackgosdk.InstancesStopRequest,
	opts ...option.RequestOption,
) (*fluidstackgosdk.ListInstanceResponse, error) {
	options := core.NewRequestOptions(opts...)
	baseURL := internal.ResolveBaseURL(
		options.BaseURL,
		c.baseURL,
		"https://platform.fluidstack.io",
	)
	endpointURL := internal.EncodeURL(
		baseURL+"/instances/%v/stop",
		instanceId,
	)
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

	var response *fluidstackgosdk.ListInstanceResponse
	if err := c.caller.Call(
		ctx,
		&internal.CallParams{
			URL:             endpointURL,
			Method:          http.MethodPut,
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

// This endpoint is used to start an existing instance by its ID.
func (c *Client) Start(
	ctx context.Context,
	instanceId string,
	request *fluidstackgosdk.InstancesStartRequest,
	opts ...option.RequestOption,
) (*fluidstackgosdk.ListInstanceResponse, error) {
	options := core.NewRequestOptions(opts...)
	baseURL := internal.ResolveBaseURL(
		options.BaseURL,
		c.baseURL,
		"https://platform.fluidstack.io",
	)
	endpointURL := internal.EncodeURL(
		baseURL+"/instances/%v/start",
		instanceId,
	)
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

	var response *fluidstackgosdk.ListInstanceResponse
	if err := c.caller.Call(
		ctx,
		&internal.CallParams{
			URL:             endpointURL,
			Method:          http.MethodPut,
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
