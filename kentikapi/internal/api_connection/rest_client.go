package api_connection

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/hashicorp/go-retryablehttp"
	"github.com/kentik/community_sdk_golang/apiv6/kentikapi/httputil"
)

//nolint:gosec
const (
	authEmailKey    = "X-CH-Auth-Email"
	authAPITokenKey = "X-CH-Auth-API-Token"
	defaultTimeout  = 100 * time.Second
)

type RestClient struct {
	config     RestClientConfig
	httpClient *retryablehttp.Client
	// Timeout specifies a time limit for requests made by this Client. The timeout includes connection time, any
	// redirects, and reading the response body. The timer remains running after Get, Head, Post, or Do return and will
	// interrupt reading of the Response.Body. A Timeout of zero means no timeout.
	Timeout *time.Duration
}

type RestClientConfig struct {
	APIURL    string
	AuthEmail string
	AuthToken string
	RetryCfg  httputil.RetryConfig
	Timeout   *time.Duration
}

func NewRestClient(c RestClientConfig) *RestClient {
	if c.Timeout != nil {
		return &RestClient{
			config: c,
			httpClient: httputil.
				NewRetryingClient(
					httputil.ClientConfig{
						HTTPClient: nil,
						RetryCfg:   c.RetryCfg,
					},
				),
			Timeout: c.Timeout,
		}
	}
	return &RestClient{
		config: c,
		httpClient: httputil.
			NewRetryingClient(
				httputil.ClientConfig{
					HTTPClient: nil,
					RetryCfg:   c.RetryCfg,
				},
			),
		Timeout: durationPtr(defaultTimeout),
	}
}

func durationPtr(duration time.Duration) *time.Duration {
	return &duration
}

// Get sends GET request to the API and returns raw response body.
//nolint:dupl
func (c *RestClient) Get(ctx context.Context, path string) (responseBody json.RawMessage, err error) {
	ctx, cancel := context.WithTimeout(ctx, *c.Timeout)
	defer cancel()

	request, err := c.newRequest(ctx, http.MethodGet, path, json.RawMessage{})
	if err != nil {
		return nil, fmt.Errorf("new request: %v", err)
	}

	response, err := c.httpClient.Do(request)
	if err != nil {
		return nil, fmt.Errorf("do request: %v", err)
	}

	defer func() {
		cErr := response.Body.Close()
		if err == nil && cErr != nil {
			err = fmt.Errorf("close response body: %v", cErr)
		}
	}()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf("read response body: %v", err)
	}

	return body, errorFromResponseStatus(response, string(body))
}

// Post sends POST request to the API and returns raw response body
//nolint:dupl
func (c *RestClient) Post(ctx context.Context, path string, payload json.RawMessage,
) (responseBody json.RawMessage, err error) {
	ctx, cancel := context.WithTimeout(ctx, *c.Timeout)
	defer cancel()

	request, err := c.newRequest(ctx, http.MethodPost, path, payload)
	if err != nil {
		return nil, fmt.Errorf("new request: %v", err)
	}

	response, err := c.httpClient.Do(request)
	if err != nil {
		return nil, fmt.Errorf("do request: %v", err)
	}
	defer func() {
		cErr := response.Body.Close()
		if err == nil && cErr != nil {
			err = fmt.Errorf("close response body: %v", cErr)
		}
	}()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf("read response body: %v", err)
	}

	return body, errorFromResponseStatus(response, string(body))
}

// Put sends PUT request to the API and returns raw response body
//nolint:dupl
func (c *RestClient) Put(ctx context.Context, path string, payload json.RawMessage,
) (responseBody json.RawMessage, err error) {
	ctx, cancel := context.WithTimeout(ctx, *c.Timeout)
	defer cancel()

	request, err := c.newRequest(ctx, http.MethodPut, path, payload)
	if err != nil {
		return nil, fmt.Errorf("new request: %v", err)
	}

	response, err := c.httpClient.Do(request)
	if err != nil {
		return nil, fmt.Errorf("do request: %v", err)
	}
	defer func() {
		cErr := response.Body.Close()
		if err == nil && cErr != nil {
			err = fmt.Errorf("close response body: %v", cErr)
		}
	}()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf("read response body: %v", err)
	}

	return body, errorFromResponseStatus(response, string(body))
}

// Delete sends DELETE request to the API and returns raw response body
//nolint:dupl
func (c *RestClient) Delete(ctx context.Context, path string) (responseBody json.RawMessage, err error) {
	ctx, cancel := context.WithTimeout(ctx, *c.Timeout)
	defer cancel()

	request, err := c.newRequest(ctx, http.MethodDelete, path, json.RawMessage{})
	if err != nil {
		return nil, fmt.Errorf("new request: %v", err)
	}

	response, err := c.httpClient.Do(request)
	if err != nil {
		return nil, fmt.Errorf("do request: %v", err)
	}

	defer func() {
		cErr := response.Body.Close()
		if err == nil && cErr != nil {
			err = fmt.Errorf("close response body: %v", cErr)
		}
	}()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf("read response body: %v", err)
	}

	return body, errorFromResponseStatus(response, string(body))
}

func errorFromResponseStatus(r *http.Response, responseBody string) error {
	// TODO(dfurman): return more specific errors
	if r.StatusCode >= http.StatusBadRequest {
		return fmt.Errorf("API response error, status: %v, response body: %v", r.Status, responseBody)
	}
	return nil
}

func (c *RestClient) newRequest(ctx context.Context, method string, path string, payload json.RawMessage,
) (*retryablehttp.Request, error) {
	request, err := http.NewRequestWithContext(ctx, method, c.makeFullURL(path), bytes.NewReader(payload))
	if err != nil {
		return nil, err
	}

	request.Header.Set(authEmailKey, c.config.AuthEmail)
	request.Header.Set(authAPITokenKey, c.config.AuthToken)
	request.Header.Set("Content-Type", "application/json")

	return retryablehttp.FromRequest(request)
}

func (c *RestClient) makeFullURL(path string) string {
	return c.config.APIURL + path
}
