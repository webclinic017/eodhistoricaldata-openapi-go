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

// CompaniesApiService CompaniesApi service
type CompaniesApiService service

type ApiListBulkFundamentalsRequest struct {
	ctx        _context.Context
	ApiService *CompaniesApiService
	exchange   string
	fmt        *string
	symbols    *string
	offset     *string
	limit      *string
}

// string fmt (name or id) of the bulk-fundamentals
func (r ApiListBulkFundamentalsRequest) Fmt(fmt string) ApiListBulkFundamentalsRequest {
	r.fmt = &fmt
	return r
}

// string symbols (name or id) of the bulk-fundamentals
func (r ApiListBulkFundamentalsRequest) Symbols(symbols string) ApiListBulkFundamentalsRequest {
	r.symbols = &symbols
	return r
}

// string offset (name or id) of the bulk-fundamentals
func (r ApiListBulkFundamentalsRequest) Offset(offset string) ApiListBulkFundamentalsRequest {
	r.offset = &offset
	return r
}

// string limit (name or id) of the bulk-fundamentals
func (r ApiListBulkFundamentalsRequest) Limit(limit string) ApiListBulkFundamentalsRequest {
	r.limit = &limit
	return r
}

func (r ApiListBulkFundamentalsRequest) Execute() ([]Fundamentals, *_nethttp.Response, error) {
	return r.ApiService.ListBulkFundamentalsExecute(r)
}

/*
ListBulkFundamentals Method for ListBulkFundamentals

List properties of bulk-fundamentals

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param exchange string exchange (name or id) of the bulk-fundamentals
 @return ApiListBulkFundamentalsRequest
*/
func (a *CompaniesApiService) ListBulkFundamentals(ctx _context.Context, exchange string) ApiListBulkFundamentalsRequest {
	return ApiListBulkFundamentalsRequest{
		ApiService: a,
		ctx:        ctx,
		exchange:   exchange,
	}
}

// Execute executes the request
//  @return []Fundamentals
func (a *CompaniesApiService) ListBulkFundamentalsExecute(r ApiListBulkFundamentalsRequest) ([]Fundamentals, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  []Fundamentals
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CompaniesApiService.ListBulkFundamentals")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/bulk-fundamentals/{exchange}"
	localVarPath = strings.Replace(localVarPath, "{"+"exchange"+"}", _neturl.PathEscape(parameterToString(r.exchange, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.fmt == nil {
		return localVarReturnValue, nil, reportError("fmt is required and must be specified")
	}
	if r.symbols == nil {
		return localVarReturnValue, nil, reportError("symbols is required and must be specified")
	}

	localVarQueryParams.Add("fmt", parameterToString(*r.fmt, ""))
	localVarQueryParams.Add("symbols", parameterToString(*r.symbols, ""))
	if r.offset != nil {
		localVarQueryParams.Add("offset", parameterToString(*r.offset, ""))
	}
	if r.limit != nil {
		localVarQueryParams.Add("limit", parameterToString(*r.limit, ""))
	}
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

type ApiReadFundamentalsRequest struct {
	ctx        _context.Context
	ApiService *CompaniesApiService
	ticker     string
	filter     *string
}

// string filter (name or id) of the fundamentals
func (r ApiReadFundamentalsRequest) Filter(filter string) ApiReadFundamentalsRequest {
	r.filter = &filter
	return r
}

func (r ApiReadFundamentalsRequest) Execute() (Fundamentals, *_nethttp.Response, error) {
	return r.ApiService.ReadFundamentalsExecute(r)
}

/*
ReadFundamentals Method for ReadFundamentals

Read properties of fundamentals

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param ticker string ticker (name or id) of the fundamentals
 @return ApiReadFundamentalsRequest
*/
func (a *CompaniesApiService) ReadFundamentals(ctx _context.Context, ticker string) ApiReadFundamentalsRequest {
	return ApiReadFundamentalsRequest{
		ApiService: a,
		ctx:        ctx,
		ticker:     ticker,
	}
}

// Execute executes the request
//  @return Fundamentals
func (a *CompaniesApiService) ReadFundamentalsExecute(r ApiReadFundamentalsRequest) (Fundamentals, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  Fundamentals
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CompaniesApiService.ReadFundamentals")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/fundamentals/{ticker}"
	localVarPath = strings.Replace(localVarPath, "{"+"ticker"+"}", _neturl.PathEscape(parameterToString(r.ticker, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	if r.filter != nil {
		localVarQueryParams.Add("filter", parameterToString(*r.filter, ""))
	}
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
