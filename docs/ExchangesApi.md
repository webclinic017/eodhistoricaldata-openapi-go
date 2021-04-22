# \ExchangesApi

All URIs are relative to *https://eodhistoricaldata.com/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListExchangeTickers**](ExchangesApi.md#ListExchangeTickers) | **Get** /exchange-symbol-list | 
[**ListExchanges**](ExchangesApi.md#ListExchanges) | **Get** /exchanges-list | 



## ListExchangeTickers

> []ExchangeTicker ListExchangeTickers(ctx, exchange).Fmt(fmt).Execute()





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
    exchange := "exchange_example" // string | string exchange (name or id) of the exchangetickers
    fmt := "["json","csv"]" // string | string fmt (name or id) of the exchangetickers

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ExchangesApi.ListExchangeTickers(context.Background(), exchange).Fmt(fmt).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ExchangesApi.ListExchangeTickers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListExchangeTickers`: []ExchangeTicker
    fmt.Fprintf(os.Stdout, "Response from `ExchangesApi.ListExchangeTickers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**exchange** | **string** | string exchange (name or id) of the exchangetickers | 

### Other Parameters

Other parameters are passed through a pointer to a apiListExchangeTickersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fmt** | **string** | string fmt (name or id) of the exchangetickers | 

### Return type

[**[]ExchangeTicker**](ExchangeTicker.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListExchanges

> []Exchange ListExchanges(ctx).Fmt(fmt).Execute()





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
    fmt := "["json","csv"]" // string | string fmt (name or id) of the exchanges

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ExchangesApi.ListExchanges(context.Background()).Fmt(fmt).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ExchangesApi.ListExchanges``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListExchanges`: []Exchange
    fmt.Fprintf(os.Stdout, "Response from `ExchangesApi.ListExchanges`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListExchangesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fmt** | **string** | string fmt (name or id) of the exchanges | 

### Return type

[**[]Exchange**](Exchange.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

