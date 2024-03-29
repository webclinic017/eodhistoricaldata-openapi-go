/*
eodhistoricaldata

eodhistoricaldata API

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"bytes"
	_context "context"
	_ioutil "io/ioutil"
	_nethttp "net/http"
	_neturl "net/url"
	"strings"
)

// Linger please
var (
	_ _context.Context
)

// DividendsApiService DividendsApi service
type DividendsApiService service

type ApiListDividendsRequest struct {
	ctx        _context.Context
	ApiService *DividendsApiService
	ticker     string
	fmt        *string
	from       *string
	to         *string
}

// string fmt (name or id) of the dividends
func (r ApiListDividendsRequest) Fmt(fmt string) ApiListDividendsRequest {
	r.fmt = &fmt
	return r
}

// string from (name or id) of the dividends
func (r ApiListDividendsRequest) From(from string) ApiListDividendsRequest {
	r.from = &from
	return r
}

// string to (name or id) of the dividends
func (r ApiListDividendsRequest) To(to string) ApiListDividendsRequest {
	r.to = &to
	return r
}

func (r ApiListDividendsRequest) Execute() ([]Dividend, *_nethttp.Response, error) {
	return r.ApiService.ListDividendsExecute(r)
}

/*
ListDividends Method for ListDividends

List properties of dividends

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param ticker string ticker (name or id) of the dividends
 @return ApiListDividendsRequest
*/
func (a *DividendsApiService) ListDividends(ctx _context.Context, ticker string) ApiListDividendsRequest {
	return ApiListDividendsRequest{
		ApiService: a,
		ctx:        ctx,
		ticker:     ticker,
	}
}

// Execute executes the request
//  @return []Dividend
func (a *DividendsApiService) ListDividendsExecute(r ApiListDividendsRequest) ([]Dividend, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  []Dividend
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DividendsApiService.ListDividends")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/div/{ticker}"
	localVarPath = strings.Replace(localVarPath, "{"+"ticker"+"}", _neturl.PathEscape(parameterToString(r.ticker, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.fmt == nil {
		return localVarReturnValue, nil, reportError("fmt is required and must be specified")
	}
	if r.from == nil {
		return localVarReturnValue, nil, reportError("from is required and must be specified")
	}
	if r.to == nil {
		return localVarReturnValue, nil, reportError("to is required and must be specified")
	}

	localVarQueryParams.Add("fmt", parameterToString(*r.fmt, ""))
	localVarQueryParams.Add("from", parameterToString(*r.from, ""))
	localVarQueryParams.Add("to", parameterToString(*r.to, ""))
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["ApiKeyAuth"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarQueryParams.Add("api_token", key)
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
