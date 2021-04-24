# \NewsApi

All URIs are relative to *https://eodhistoricaldata.com/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListNews**](NewsApi.md#ListNews) | **Get** /news | 



## ListNews

> []Quote ListNews(ctx).S(s).From(from).To(to).Limit(limit).Offset(offset).Execute()





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
    s := "["AAPL.US"]" // string | string s (name or id) of the news
    from := "["2021-03-01"]" // string | string from (name or id) of the news
    to := "["2021-03-10"]" // string | string to (name or id) of the news
    limit := "[1000]" // string | string limit (name or id) of the news
    offset := "[200]" // string | string offset (name or id) of the news

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NewsApi.ListNews(context.Background()).S(s).From(from).To(to).Limit(limit).Offset(offset).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NewsApi.ListNews``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListNews`: []Quote
    fmt.Fprintf(os.Stdout, "Response from `NewsApi.ListNews`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListNewsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **s** | **string** | string s (name or id) of the news | 
 **from** | **string** | string from (name or id) of the news | 
 **to** | **string** | string to (name or id) of the news | 
 **limit** | **string** | string limit (name or id) of the news | 
 **offset** | **string** | string offset (name or id) of the news | 

### Return type

[**[]Quote**](Quote.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

