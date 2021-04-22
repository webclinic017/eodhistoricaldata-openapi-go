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
)

// Linger please
var (
	_ _context.Context
)

// NewsApiService NewsApi service
type NewsApiService service

type ApiListNewsRequest struct {
	ctx        _context.Context
	ApiService *NewsApiService
	s          *string
	from       *string
	to         *string
	limit      *string
	offset     *string
}

func (r ApiListNewsRequest) S(s string) ApiListNewsRequest {
	r.s = &s
	return r
}
func (r ApiListNewsRequest) From(from string) ApiListNewsRequest {
	r.from = &from
	return r
}
func (r ApiListNewsRequest) To(to string) ApiListNewsRequest {
	r.to = &to
	return r
}
func (r ApiListNewsRequest) Limit(limit string) ApiListNewsRequest {
	r.limit = &limit
	return r
}
func (r ApiListNewsRequest) Offset(offset string) ApiListNewsRequest {
	r.offset = &offset
	return r
}

func (r ApiListNewsRequest) Execute() ([]Quote, *_nethttp.Response, error) {
	return r.ApiService.ListNewsExecute(r)
}

/*
 * ListNews Method for ListNews
 * List properties of news
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @return ApiListNewsRequest
 */
func (a *NewsApiService) ListNews(ctx _context.Context) ApiListNewsRequest {
	return ApiListNewsRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

/*
 * Execute executes the request
 * @return []Quote
 */
func (a *NewsApiService) ListNewsExecute(r ApiListNewsRequest) ([]Quote, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  []Quote
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "NewsApiService.ListNews")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/news"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.s == nil {
		return localVarReturnValue, nil, reportError("s is required and must be specified")
	}
	if r.from == nil {
		return localVarReturnValue, nil, reportError("from is required and must be specified")
	}
	if r.to == nil {
		return localVarReturnValue, nil, reportError("to is required and must be specified")
	}
	if r.limit == nil {
		return localVarReturnValue, nil, reportError("limit is required and must be specified")
	}
	if r.offset == nil {
		return localVarReturnValue, nil, reportError("offset is required and must be specified")
	}

	localVarQueryParams.Add("s", parameterToString(*r.s, ""))
	localVarQueryParams.Add("from", parameterToString(*r.from, ""))
	localVarQueryParams.Add("to", parameterToString(*r.to, ""))
	localVarQueryParams.Add("limit", parameterToString(*r.limit, ""))
	localVarQueryParams.Add("offset", parameterToString(*r.offset, ""))
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
