# Fluidstack Go Library

[![fern shield](https://img.shields.io/badge/%F0%9F%8C%BF-SDK%20generated%20by%20Fern-brightgreen)](https://github.com/fern-api/fern)

The Fluidstack Go library provides convenient access to the Fluidstack API from Go.

## Requirements

This module requires Go version >= 1.13.

# Installation

Run the following command to use the Fluidstack Go library in your module:

```sh
go get github.com/fluidstackio/fluidstack-go-sdk
```

## Usage

```go
import fluidstackclient "github.com/fluidstackio/fluidstack-go-sdk/client"
import option "github.com/fluidstackio/fluidstack-go-sdk/option"

c := fluidstackclient.NewClient(
    option.WithApiKey(API_KEY) // defaults to FLUIDSTACK_API_KEY
)
```

## Timeouts

Setting a timeout for each individual request is as simple as using the standard
`context` library. Setting a one second timeout for an individual API call looks
like the following:

```go
ctx, cancel := context.WithTimeout(context.TODO(), time.Second)
defer cancel()

response, err := client.SshKeys.List(
  ctx,
  &fluidstack.SshKeysListRequest{
			ShowAll: true,
	}
)
```

## Request Options

A variety of request options are included to adapt the behavior of the library, which includes
configuring authorization tokens, or providing your own instrumented `*http.Client`. Both of
these options are shown below:

```go
client := fluidstackclient.NewClient(
  option.WithApiKey("API_KEY"),
  option.WithHTTPClient(
    &http.Client{
      Timeout: 5 * time.Second,
    },
  ),
)
```

> Providing your own `*http.Client` is recommended. Otherwise, the `http.DefaultClient` will be used,
> and your client will wait indefinitely for a response (unless the per-request, context-based timeout
> is used).

## Errors

Structured error types are returned from API calls that return non-success status codes. For example,
you can check if the error was due to an unauthorized request (i.e. status code 401) with the following:

```go
response, err := client.SshKeys.List(...)
if err != nil {
  if unauthorizedErr, ok := err.(*fluidstack.UnauthorizedError);
    // Do something with the unauthorized request ...
  }
  return err
}
```

These errors are also compatible with the `errors.Is` and `errors.As` APIs, so you can access the error
like so:

```go
response, err := client.SshKeys.List(...)
if err != nil {
  var unauthorizedErr *fluidstack.UnauthorizedError
  if errors.As(err, unauthorizedErr) {
    // Do something with the unauthorized request ...
  }
  return err
}
```

If you'd like to wrap the errors with additional information and still retain the ability
to access the type with `errors.Is` and `errors.As`, you can use the `%w` directive:

```go
response, err := client.SshKeys.List(...)
if err != nil {
  return fmt.Errorf("failed to identify user: %w", err)
}
```

## Contributing

While we value open-source contributions to this SDK, this library is generated programmatically.
Additions made directly to this library would have to be moved over to our generation code,
otherwise they would be overwritten upon the next generated release. Feel free to open a PR as
a proof of concept, but know that we will not be able to merge it as-is. We suggest opening
an issue first to discuss with us!

On the other hand, contributions to the `README.md` are always very welcome!
