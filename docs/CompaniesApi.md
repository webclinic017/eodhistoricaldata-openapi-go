# \CompaniesApi

All URIs are relative to *https://eodhistoricaldata.com/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListBulkFundamentals**](CompaniesApi.md#ListBulkFundamentals) | **Get** /bulk-fundamentals/{ticker} | 
[**ListFundamentals**](CompaniesApi.md#ListFundamentals) | **Get** /fundamentals/{ticker} | 



## ListBulkFundamentals

> []Fundamentals ListBulkFundamentals(ctx, ticker).Fmt(fmt).Offset(offset).Limit(limit).Execute()





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
    ticker := "AAPL.US" // string | string ticker (name or id) of the bulk-fundamentals
    fmt := "fmt_example" // string | string fmt (name or id) of the bulk-fundamentals
    offset := "offset_example" // string | string offset (name or id) of the bulk-fundamentals
    limit := "limit_example" // string | string limit (name or id) of the bulk-fundamentals

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CompaniesApi.ListBulkFundamentals(context.Background(), ticker).Fmt(fmt).Offset(offset).Limit(limit).Execute()
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
**ticker** | **string** | string ticker (name or id) of the bulk-fundamentals | 

### Other Parameters

Other parameters are passed through a pointer to a apiListBulkFundamentalsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fmt** | **string** | string fmt (name or id) of the bulk-fundamentals | 
 **offset** | **string** | string offset (name or id) of the bulk-fundamentals | 
 **limit** | **string** | string limit (name or id) of the bulk-fundamentals | 

### Return type

[**[]Fundamentals**](Fundamentals.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListFundamentals

> Fundamentals ListFundamentals(ctx, ticker).Filter(filter).Execute()





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
    filter := "filter_example" // string | string filter (name or id) of the fundamentals

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CompaniesApi.ListFundamentals(context.Background(), ticker).Filter(filter).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CompaniesApi.ListFundamentals``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListFundamentals`: Fundamentals
    fmt.Fprintf(os.Stdout, "Response from `CompaniesApi.ListFundamentals`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ticker** | **string** | string ticker (name or id) of the fundamentals | 

### Other Parameters

Other parameters are passed through a pointer to a apiListFundamentalsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **filter** | **string** | string filter (name or id) of the fundamentals | 

### Return type

[**Fundamentals**](Fundamentals.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

