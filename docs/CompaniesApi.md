# \CompaniesApi

All URIs are relative to *https://eodhistoricaldata.com/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListBulkFundamentals**](CompaniesApi.md#ListBulkFundamentals) | **Get** /bulk-fundamentals/{exchange} | 
[**ReadFundamentals**](CompaniesApi.md#ReadFundamentals) | **Get** /fundamentals/{ticker} | 



## ListBulkFundamentals

> []Fundamentals ListBulkFundamentals(ctx, exchange).Fmt(fmt).Symbols(symbols).Offset(offset).Limit(limit).Execute()





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
    exchange := "exchange_example" // string | string exchange (name or id) of the bulk-fundamentals
    fmt := "["json","csv"]" // string | string fmt (name or id) of the bulk-fundamentals
    symbols := "symbols_example" // string | string symbols (name or id) of the bulk-fundamentals
    offset := "offset_example" // string | string offset (name or id) of the bulk-fundamentals (optional)
    limit := "limit_example" // string | string limit (name or id) of the bulk-fundamentals (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CompaniesApi.ListBulkFundamentals(context.Background(), exchange).Fmt(fmt).Symbols(symbols).Offset(offset).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CompaniesApi.ListBulkFundamentals``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListBulkFundamentals`: []Fundamentals
    fmt.Fprintf(os.Stdout, "Response from `CompaniesApi.ListBulkFundamentals`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**exchange** | **string** | string exchange (name or id) of the bulk-fundamentals | 

### Other Parameters

Other parameters are passed through a pointer to a apiListBulkFundamentalsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fmt** | **string** | string fmt (name or id) of the bulk-fundamentals | 
 **symbols** | **string** | string symbols (name or id) of the bulk-fundamentals | 
 **offset** | **string** | string offset (name or id) of the bulk-fundamentals | 
 **limit** | **string** | string limit (name or id) of the bulk-fundamentals | 

### Return type

[**[]Fundamentals**](Fundamentals.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReadFundamentals

> Fundamentals ReadFundamentals(ctx, ticker).Filter(filter).Execute()





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
    ticker := "AAPL.US" // string | string ticker (name or id) of the fundamentals
    filter := "filter_example" // string | string filter (name or id) of the fundamentals (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CompaniesApi.ReadFundamentals(context.Background(), ticker).Filter(filter).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CompaniesApi.ReadFundamentals``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReadFundamentals`: Fundamentals
    fmt.Fprintf(os.Stdout, "Response from `CompaniesApi.ReadFundamentals`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ticker** | **string** | string ticker (name or id) of the fundamentals | 

### Other Parameters

Other parameters are passed through a pointer to a apiReadFundamentalsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **filter** | **string** | string filter (name or id) of the fundamentals | 

### Return type

[**Fundamentals**](Fundamentals.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

