package pve

import (
	"bytes"
	"context"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/mitchellh/mapstructure"
)

type Options struct {
	// The HTTP client to invoke API calls with. Defaults to client's default HTTP
	// implementation if nil.
	HTTPClient HTTPClient
	TLS        *tls.Config
}

// Copy creates a clone where the APIOptions list is deep copied.
func (o Options) Copy() Options {
	to := o
	return to
}

func WithTLSOptions(c *tls.Config) func(*Options) {
	return func(o *Options) {
		o.TLS = c
	}
}

func WithHTTPClientOptions(hc HTTPClient) func(*Options) {
	return func(o *Options) {
		o.HTTPClient = hc
	}
}

type HTTPClient interface {
	Do(*http.Request) (*http.Response, error)
}

// Client provides the API client to make API call for proxmox pve
type Client struct {
	options Options

	auth AuthOptions
}

func New(auth *AuthOptions, optFns ...func(*Options)) *Client {
	options := Options{}

	for _, fn := range optFns {
		fn(&options)
	}

	if options.HTTPClient == nil {
		tr := &http.Transport{}
		if options.TLS != nil {
			tr.TLSClientConfig = options.TLS
		}
		options.HTTPClient = &http.Client{
			Transport: tr,
		}
	}

	client := &Client{
		options: options,
		auth:    *auth,
	}

	return client
}

func (c *Client) doRequest(
	ctx context.Context,
	method string,
	uri string,
	params *url.Values,
	headers *http.Header,
	body io.Reader,
) (resp *http.Response, err error) {

	// add params to url here
	// reqUrl := url.JoinPath(c.options.APIEndpoint, uri)
	reqUrl := c.auth.APIEndpoint + uri
	if params != nil && len(*params) > 0 {
		reqUrl = reqUrl + "?" + params.Encode()
	}

	request, err := http.NewRequestWithContext(ctx, method, reqUrl, body)
	if err != nil {
		return nil, err
	}

	if headers != nil {
		request.Header = *headers
	}

	// set token
	authHeader := fmt.Sprintf("PVEAPIToken=%s", c.auth.ApiToken)
	request.Header.Set("Authorization", authHeader)

	return c.options.HTTPClient.Do(request)
}

func (c *Client) apiCall(
	ctx context.Context,
	method string,
	path string,
	request interface{},
	response interface{},
) error {

	uri := path

	params, err := constructRequestParmas(request)
	if err != nil {
		return err
	}
	headers, err := constructRequestHeaders(request)
	if err != nil {
		return err
	}
	body, err := constructRequestBody(request)
	if err != nil {
		return err
	}

	result, err := c.doRequest(ctx, method, uri, params, headers, body)
	if err != nil {
		return err
	}

	bodyBytes, err := io.ReadAll(result.Body)
	if err != nil {
		return err
	}

	responseBody := ApiResponseBody{
		Data: response,
	}

	if len(bodyBytes) > 0 {
		err := json.Unmarshal(bodyBytes, &responseBody)
		if err != nil {
			return err
		}
	}

	if result.StatusCode < 200 || result.StatusCode > 299 {
		if responseBody.Errors != nil {
			errorDetial, err := json.Marshal(responseBody.Errors)
			if err != nil {
				panic(err)
			}
			return fmt.Errorf("%s | details: %s", result.Status, string(errorDetial))
		}
		return fmt.Errorf("%s", result.Status)
	}

	return nil
}

type ApiResponseBody struct {
	Data   interface{}            `json:"data"`
	Errors map[string]interface{} `json:"errors"`
}

type ApiResponseErrors struct {
	Name string
}

const (
	RequestBodyTagName    = "json"
	RequestParamsTagName  = "query"
	RequestHeadersTagName = "header"
	RequestPathTagName    = "path"
)

func constructRequestParmas(request interface{}) (*url.Values, error) {
	var params map[string]interface{}

	decoder, err := mapstructure.NewDecoder(&mapstructure.DecoderConfig{
		Metadata:             nil,
		Result:               &params,
		TagName:              RequestParamsTagName,
		IgnoreUntaggedFields: true,
	})
	if err != nil {
		return nil, err
	}

	err = decoder.Decode(request)
	if err != nil {
		return nil, err
	}

	result := url.Values{}
	for key, val := range params {
		switch v := val.(type) {
		case string:
			result.Add(key, v)
		case *string:
			result.Add(key, *v)
		case []string:
			for _, vi := range v {
				result.Add(key, vi)
			}
		case []*string:
			for _, vi := range v {
				if vi != nil {
					result.Add(key, *vi)
				}
			}
		}
	}

	return &result, nil
}

func constructRequestHeaders(request interface{}) (*http.Header, error) {

	var headers map[string]interface{}

	decoder, err := mapstructure.NewDecoder(&mapstructure.DecoderConfig{
		Metadata:             nil,
		Result:               &headers,
		TagName:              RequestHeadersTagName,
		IgnoreUntaggedFields: true,
	})
	if err != nil {
		return nil, err
	}

	err = decoder.Decode(request)
	if err != nil {
		return nil, err
	}

	result := http.Header{}
	for key, val := range headers {
		switch v := val.(type) {
		case string:
			result.Add(key, v)
		case []string:
			for _, vi := range v {
				result.Add(key, vi)
			}
		}
	}

	return &result, nil
}

func constructRequestBody(request interface{}) (io.Reader, error) {

	var body map[string]interface{}

	decoder, err := mapstructure.NewDecoder(&mapstructure.DecoderConfig{
		Metadata:             nil,
		Result:               &body,
		TagName:              RequestBodyTagName,
		IgnoreUntaggedFields: true,
	})
	if err != nil {
		return nil, err
	}

	err = decoder.Decode(request)
	if err != nil {
		return nil, err
	}

	// nil empty body
	if len(body) == 0 {
		return nil, nil
	}

	bodyBytes, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}

	buf := bytes.NewReader(bodyBytes)

	return buf, err
}
