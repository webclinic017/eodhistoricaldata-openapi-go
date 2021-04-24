# \DividendsApi

All URIs are relative to *https://eodhistoricaldata.com/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListDividends**](DividendsApi.md#ListDividends) | **Get** /div/{ticker} | 



## ListDividends

> []Dividend ListDividends(ctx, ticker).Fmt(fmt).From(from).To(to).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    ticker := "ticker_example" // string | string ticker (name or id) of the dividends
    fmt := "["json","csv"]" // string | string fmt (name or id) of the dividends
    from := "["2021-03-01"]" // string | string from (name or id) of the dividends
    to := "["2021-03-10"]" // string | string to (name or id) of the dividends

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DividendsApi.ListDividends(context.Background(), ticker).Fmt(fmt).From(from).To(to).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DividendsApi.ListDividends``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListDividends`: []Dividend
    fmt.Fprintf(os.Stdout, "Response from `DividendsApi.ListDividends`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ticker** | **string** | string ticker (name or id) of the dividends | 

### Other Parameters

Other parameters are passed through a pointer to a apiListDividendsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fmt** | **string** | string fmt (name or id) of the dividends | 
 **from** | **string** | string from (name or id) of the dividends | 
 **to** | **string** | string to (name or id) of the dividends | 

### Return type

[**[]Dividend**](Dividend.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

