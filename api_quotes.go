/*
 * eodhistoricaldata
 *
 * eodhistoricaldata API
 *
 * API version: 1.0.0
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

// QuotesApiService QuotesApi service
type QuotesApiService service

type ApiListHistoryIntradayQuotesRequest struct {
	ctx        _context.Context
	ApiService *QuotesApiService
	ticker     string
	fmt        *string
	interval   *string
	from       *string
	to         *string
}

func (r ApiListHistoryIntradayQuotesRequest) Fmt(fmt string) ApiListHistoryIntradayQuotesRequest {
	r.fmt = &fmt
	return r
}
func (r ApiListHistoryIntradayQuotesRequest) Interval(interval string) ApiListHistoryIntradayQuotesRequest {
	r.interval = &interval
	return r
}
func (r ApiListHistoryIntradayQuotesRequest) From(from string) ApiListHistoryIntradayQuotesRequest {
	r.from = &from
	return r
}
func (r ApiListHistoryIntradayQuotesRequest) To(to string) ApiListHistoryIntradayQuotesRequest {
	r.to = &to
	return r
}

func (r ApiListHistoryIntradayQuotesRequest) Execute() ([]Quote, *_nethttp.Response, error) {
	return r.ApiService.ListHistoryIntradayQuotesExecute(r)
}

/*
 * ListHistoryIntradayQuotes Method for ListHistoryIntradayQuotes
 * List properties of historyintradayquotes
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param ticker string ticker (name or id) of the historyintradayquotes
 * @return ApiListHistoryIntradayQuotesRequest
 */
func (a *QuotesApiService) ListHistoryIntradayQuotes(ctx _context.Context, ticker string) ApiListHistoryIntradayQuotesRequest {
	return ApiListHistoryIntradayQuotesRequest{
		ApiService: a,
		ctx:        ctx,
		ticker:     ticker,
	}
}

/*
 * Execute executes the request
 * @return []Quote
 */
func (a *QuotesApiService) ListHistoryIntradayQuotesExecute(r ApiListHistoryIntradayQuotesRequest) ([]Quote, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  []Quote
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "QuotesApiService.ListHistoryIntradayQuotes")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/intraday/{ticker}"
	localVarPath = strings.Replace(localVarPath, "{"+"ticker"+"}", _neturl.PathEscape(parameterToString(r.ticker, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.fmt == nil {
		return localVarReturnValue, nil, reportError("fmt is required and must be specified")
	}
	if r.interval == nil {
		return localVarReturnValue, nil, reportError("interval is required and must be specified")
	}
	if r.from == nil {
		return localVarReturnValue, nil, reportError("from is required and must be specified")
	}
	if r.to == nil {
		return localVarReturnValue, nil, reportError("to is required and must be specified")
	}

	localVarQueryParams.Add("fmt", parameterToString(*r.fmt, ""))
	localVarQueryParams.Add("interval", parameterToString(*r.interval, ""))
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

type ApiListHistoryQuotesRequest struct {
	ctx        _context.Context
	ApiService *QuotesApiService
	ticker     string
	fmt        *string
	g          *string
	filter     *string
	order      *string
	from       *string
	to         *string
}

func (r ApiListHistoryQuotesRequest) Fmt(fmt string) ApiListHistoryQuotesRequest {
	r.fmt = &fmt
	return r
}
func (r ApiListHistoryQuotesRequest) G(g string) ApiListHistoryQuotesRequest {
	r.g = &g
	return r
}
func (r ApiListHistoryQuotesRequest) Filter(filter string) ApiListHistoryQuotesRequest {
	r.filter = &filter
	return r
}
func (r ApiListHistoryQuotesRequest) Order(order string) ApiListHistoryQuotesRequest {
	r.order = &order
	return r
}
func (r ApiListHistoryQuotesRequest) From(from string) ApiListHistoryQuotesRequest {
	r.from = &from
	return r
}
func (r ApiListHistoryQuotesRequest) To(to string) ApiListHistoryQuotesRequest {
	r.to = &to
	return r
}

func (r ApiListHistoryQuotesRequest) Execute() ([]Quote, *_nethttp.Response, error) {
	return r.ApiService.ListHistoryQuotesExecute(r)
}

/*
 * ListHistoryQuotes Method for ListHistoryQuotes
 * List properties of historyquotes
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param ticker string ticker (name or id) of the historyquotes
 * @return ApiListHistoryQuotesRequest
 */
func (a *QuotesApiService) ListHistoryQuotes(ctx _context.Context, ticker string) ApiListHistoryQuotesRequest {
	return ApiListHistoryQuotesRequest{
		ApiService: a,
		ctx:        ctx,
		ticker:     ticker,
	}
}

/*
 * Execute executes the request
 * @return []Quote
 */
func (a *QuotesApiService) ListHistoryQuotesExecute(r ApiListHistoryQuotesRequest) ([]Quote, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  []Quote
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "QuotesApiService.ListHistoryQuotes")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/eod/{ticker}"
	localVarPath = strings.Replace(localVarPath, "{"+"ticker"+"}", _neturl.PathEscape(parameterToString(r.ticker, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.fmt == nil {
		return localVarReturnValue, nil, reportError("fmt is required and must be specified")
	}

	if r.g != nil {
		localVarQueryParams.Add("g", parameterToString(*r.g, ""))
	}
	localVarQueryParams.Add("fmt", parameterToString(*r.fmt, ""))
	if r.filter != nil {
		localVarQueryParams.Add("filter", parameterToString(*r.filter, ""))
	}
	if r.order != nil {
		localVarQueryParams.Add("order", parameterToString(*r.order, ""))
	}
	if r.from != nil {
		localVarQueryParams.Add("from", parameterToString(*r.from, ""))
	}
	if r.to != nil {
		localVarQueryParams.Add("to", parameterToString(*r.to, ""))
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

type ApiListRealtimeQuotesRequest struct {
	ctx        _context.Context
	ApiService *QuotesApiService
	ticker     string
	fmt        *string
	interval   *string
	filter     *string
	s          *string
}

func (r ApiListRealtimeQuotesRequest) Fmt(fmt string) ApiListRealtimeQuotesRequest {
	r.fmt = &fmt
	return r
}
func (r ApiListRealtimeQuotesRequest) Interval(interval string) ApiListRealtimeQuotesRequest {
	r.interval = &interval
	return r
}
func (r ApiListRealtimeQuotesRequest) Filter(filter string) ApiListRealtimeQuotesRequest {
	r.filter = &filter
	return r
}
func (r ApiListRealtimeQuotesRequest) S(s string) ApiListRealtimeQuotesRequest {
	r.s = &s
	return r
}

func (r ApiListRealtimeQuotesRequest) Execute() ([]Quote, *_nethttp.Response, error) {
	return r.ApiService.ListRealtimeQuotesExecute(r)
}

/*
 * ListRealtimeQuotes Method for ListRealtimeQuotes
 * List properties of realtimequotes
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param ticker string ticker (name or id) of the realtimequotes
 * @return ApiListRealtimeQuotesRequest
 */
func (a *QuotesApiService) ListRealtimeQuotes(ctx _context.Context, ticker string) ApiListRealtimeQuotesRequest {
	return ApiListRealtimeQuotesRequest{
		ApiService: a,
		ctx:        ctx,
		ticker:     ticker,
	}
}

/*
 * Execute executes the request
 * @return []Quote
 */
func (a *QuotesApiService) ListRealtimeQuotesExecute(r ApiListRealtimeQuotesRequest) ([]Quote, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  []Quote
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "QuotesApiService.ListRealtimeQuotes")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/real-time/{ticker}"
	localVarPath = strings.Replace(localVarPath, "{"+"ticker"+"}", _neturl.PathEscape(parameterToString(r.ticker, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.fmt == nil {
		return localVarReturnValue, nil, reportError("fmt is required and must be specified")
	}
	if r.interval == nil {
		return localVarReturnValue, nil, reportError("interval is required and must be specified")
	}
	if r.filter == nil {
		return localVarReturnValue, nil, reportError("filter is required and must be specified")
	}
	if r.s == nil {
		return localVarReturnValue, nil, reportError("s is required and must be specified")
	}

	localVarQueryParams.Add("fmt", parameterToString(*r.fmt, ""))
	localVarQueryParams.Add("interval", parameterToString(*r.interval, ""))
	localVarQueryParams.Add("filter", parameterToString(*r.filter, ""))
	localVarQueryParams.Add("s", parameterToString(*r.s, ""))
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

type ApiListShortsQuotesRequest struct {
	ctx        _context.Context
	ApiService *QuotesApiService
	ticker     string
	fmt        *string
}

func (r ApiListShortsQuotesRequest) Fmt(fmt string) ApiListShortsQuotesRequest {
	r.fmt = &fmt
	return r
}

func (r ApiListShortsQuotesRequest) Execute() ([]SplitsQuote, *_nethttp.Response, error) {
	return r.ApiService.ListShortsQuotesExecute(r)
}

/*
 * ListShortsQuotes Method for ListShortsQuotes
 * List properties of shortsquotes
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param ticker string ticker (name or id) of the shortsquotes
 * @return ApiListShortsQuotesRequest
 */
func (a *QuotesApiService) ListShortsQuotes(ctx _context.Context, ticker string) ApiListShortsQuotesRequest {
	return ApiListShortsQuotesRequest{
		ApiService: a,
		ctx:        ctx,
		ticker:     ticker,
	}
}

/*
 * Execute executes the request
 * @return []SplitsQuote
 */
func (a *QuotesApiService) ListShortsQuotesExecute(r ApiListShortsQuotesRequest) ([]SplitsQuote, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  []SplitsQuote
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "QuotesApiService.ListShortsQuotes")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/shorts/{ticker}"
	localVarPath = strings.Replace(localVarPath, "{"+"ticker"+"}", _neturl.PathEscape(parameterToString(r.ticker, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.fmt == nil {
		return localVarReturnValue, nil, reportError("fmt is required and must be specified")
	}

	localVarQueryParams.Add("fmt", parameterToString(*r.fmt, ""))
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

type ApiListSplitsQuotesRequest struct {
	ctx        _context.Context
	ApiService *QuotesApiService
	ticker     string
	fmt        *string
}

func (r ApiListSplitsQuotesRequest) Fmt(fmt string) ApiListSplitsQuotesRequest {
	r.fmt = &fmt
	return r
}

func (r ApiListSplitsQuotesRequest) Execute() ([]SplitsQuote, *_nethttp.Response, error) {
	return r.ApiService.ListSplitsQuotesExecute(r)
}

/*
 * ListSplitsQuotes Method for ListSplitsQuotes
 * List properties of splitsquotes
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param ticker string ticker (name or id) of the splitsquotes
 * @return ApiListSplitsQuotesRequest
 */
func (a *QuotesApiService) ListSplitsQuotes(ctx _context.Context, ticker string) ApiListSplitsQuotesRequest {
	return ApiListSplitsQuotesRequest{
		ApiService: a,
		ctx:        ctx,
		ticker:     ticker,
	}
}

/*
 * Execute executes the request
 * @return []SplitsQuote
 */
func (a *QuotesApiService) ListSplitsQuotesExecute(r ApiListSplitsQuotesRequest) ([]SplitsQuote, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  []SplitsQuote
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "QuotesApiService.ListSplitsQuotes")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/splits/{ticker}"
	localVarPath = strings.Replace(localVarPath, "{"+"ticker"+"}", _neturl.PathEscape(parameterToString(r.ticker, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.fmt == nil {
		return localVarReturnValue, nil, reportError("fmt is required and must be specified")
	}

	localVarQueryParams.Add("fmt", parameterToString(*r.fmt, ""))
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
