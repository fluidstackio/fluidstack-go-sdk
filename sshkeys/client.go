// This file was auto-generated by Fern from our API Definition.

package sshkeys

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

// Fetch a list of SSH key names associated with the authenticated user.
func (c *Client) List(
	ctx context.Context,
	request *fluidstackgosdk.SshKeysListRequest,
	opts ...option.RequestOption,
) ([]*fluidstackgosdk.SshKeyResponse, error) {
	options := core.NewRequestOptions(opts...)
	baseURL := internal.ResolveBaseURL(
		options.BaseURL,
		c.baseURL,
		"https://platform.fluidstack.io",
	)
	endpointURL := baseURL + "/ssh_keys"
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

	var response []*fluidstackgosdk.SshKeyResponse
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

// Create a new SSH key for the authenticated user.
//
// A unique name must be provided for the SSH key, along with a public key. The public key you provide is stored on your FluidStack account for use in SSH authentication.
//
// Supported public key formats: ssh-rsa, ssh-dss (DSA), ssh-ed25519, and ecdsa keys with NIST curves.
func (c *Client) Create(
	ctx context.Context,
	request *fluidstackgosdk.CreateSshKeyRequest,
	opts ...option.RequestOption,
) (*fluidstackgosdk.SshKeyResponse, error) {
	options := core.NewRequestOptions(opts...)
	baseURL := internal.ResolveBaseURL(
		options.BaseURL,
		c.baseURL,
		"https://platform.fluidstack.io",
	)
	endpointURL := baseURL + "/ssh_keys"
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

	var response *fluidstackgosdk.SshKeyResponse
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

// Delete an existing SSH key by its name.
func (c *Client) Delete(
	ctx context.Context,
	sshKeyName string,
	request *fluidstackgosdk.SshKeysDeleteRequest,
	opts ...option.RequestOption,
) (interface{}, error) {
	options := core.NewRequestOptions(opts...)
	baseURL := internal.ResolveBaseURL(
		options.BaseURL,
		c.baseURL,
		"https://platform.fluidstack.io",
	)
	endpointURL := internal.EncodeURL(
		baseURL+"/ssh_keys/%v",
		sshKeyName,
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
