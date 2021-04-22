# \QuotesApi

All URIs are relative to *https://eodhistoricaldata.com/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ReadQuotes**](QuotesApi.md#ReadQuotes) | **Get** /eod/{ticker} | 



## ReadQuotes

> []Quote ReadQuotes(ctx, ticker).Fmt(fmt).Order(order).Execute()





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
    ticker := "["EUR.FOREX","BZ","CL","CT","GC","HG","KC","MALTR","MZN","NG","NICKEL","PA","PL","SB","SI","ZC"]" // string | string ticker (name or id) of the quotes
    fmt := "["json","csv"]" // string | string fmt (name or id) of the quotes
    order := "["d"]" // string | string order (name or id) of the quotes

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.QuotesApi.ReadQuotes(context.Background(), ticker).Fmt(fmt).Order(order).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `QuotesApi.ReadQuotes``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReadQuotes`: []Quote
    fmt.Fprintf(os.Stdout, "Response from `QuotesApi.ReadQuotes`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ticker** | **string** | string ticker (name or id) of the quotes | 

### Other Parameters

Other parameters are passed through a pointer to a apiReadQuotesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fmt** | **string** | string fmt (name or id) of the quotes | 
 **order** | **string** | string order (name or id) of the quotes | 

### Return type

[**[]Quote**](Quote.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

