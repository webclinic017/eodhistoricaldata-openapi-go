# \MacroIndicatorApi

All URIs are relative to *https://eodhistoricaldata.com/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListMacroIndicator**](MacroIndicatorApi.md#ListMacroIndicator) | **Get** /macro-indicator/{country} | 



## ListMacroIndicator

> MacroIndicator ListMacroIndicator(ctx, country).Fmt(fmt).Indicator(indicator).Execute()





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
    country := "country_example" // string | string country (name or id) of the macroindicator
    fmt := "["json","csv"]" // string | string fmt (name or id) of the macroindicator
    indicator := "["population_total"]" // string | string indicator (name or id) of the macroindicator

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MacroIndicatorApi.ListMacroIndicator(context.Background(), country).Fmt(fmt).Indicator(indicator).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MacroIndicatorApi.ListMacroIndicator``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListMacroIndicator`: MacroIndicator
    fmt.Fprintf(os.Stdout, "Response from `MacroIndicatorApi.ListMacroIndicator`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**country** | **string** | string country (name or id) of the macroindicator | 

### Other Parameters

Other parameters are passed through a pointer to a apiListMacroIndicatorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fmt** | **string** | string fmt (name or id) of the macroindicator | 
 **indicator** | **string** | string indicator (name or id) of the macroindicator | 

### Return type

[**MacroIndicator**](MacroIndicator.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

