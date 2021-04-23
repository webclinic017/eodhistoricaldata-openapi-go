# \QuotesApi

All URIs are relative to *https://eodhistoricaldata.com/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListHistoryIntradayQuotes**](QuotesApi.md#ListHistoryIntradayQuotes) | **Get** /intraday/{ticker} | 
[**ListHistoryQuotes**](QuotesApi.md#ListHistoryQuotes) | **Get** /eod/{ticker} | 
[**ListRealtimeQuotes**](QuotesApi.md#ListRealtimeQuotes) | **Get** /real-time/{ticker} | 
[**ListShortsQuotes**](QuotesApi.md#ListShortsQuotes) | **Get** /shorts/{ticker} | 
[**ListSplitsQuotes**](QuotesApi.md#ListSplitsQuotes) | **Get** /splits/{ticker} | 



## ListHistoryIntradayQuotes

> []Quote ListHistoryIntradayQuotes(ctx, ticker).Fmt(fmt).Interval(interval).From(from).To(to).Execute()





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
    ticker := "ticker_example" // string | string ticker (name or id) of the historyintradayquotes
    fmt := "["json","csv"]" // string | string fmt (name or id) of the historyintradayquotes
    interval := "["1m","5m"]" // string | string interval (name or id) of the historyintradayquotes
    from := "[10000,40000]" // string | string from (name or id) of the historyintradayquotes
    to := "[10000,40000]" // string | string to (name or id) of the historyintradayquotes

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.QuotesApi.ListHistoryIntradayQuotes(context.Background(), ticker).Fmt(fmt).Interval(interval).From(from).To(to).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `QuotesApi.ListHistoryIntradayQuotes``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListHistoryIntradayQuotes`: []Quote
    fmt.Fprintf(os.Stdout, "Response from `QuotesApi.ListHistoryIntradayQuotes`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ticker** | **string** | string ticker (name or id) of the historyintradayquotes | 

### Other Parameters

Other parameters are passed through a pointer to a apiListHistoryIntradayQuotesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fmt** | **string** | string fmt (name or id) of the historyintradayquotes | 
 **interval** | **string** | string interval (name or id) of the historyintradayquotes | 
 **from** | **string** | string from (name or id) of the historyintradayquotes | 
 **to** | **string** | string to (name or id) of the historyintradayquotes | 

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


## ListHistoryQuotes

> []Quote ListHistoryQuotes(ctx, ticker).Fmt(fmt).Order(order).Execute()





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
    ticker := "["EUR.FOREX","BZ","CL","CT","GC","HG","KC","MALTR","MZN","NG","NICKEL","PA","PL","SB","SI","ZC"]" // string | string ticker (name or id) of the historyquotes
    fmt := "["json","csv"]" // string | string fmt (name or id) of the historyquotes
    order := "["d"]" // string | string order (name or id) of the historyquotes

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.QuotesApi.ListHistoryQuotes(context.Background(), ticker).Fmt(fmt).Order(order).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `QuotesApi.ListHistoryQuotes``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListHistoryQuotes`: []Quote
    fmt.Fprintf(os.Stdout, "Response from `QuotesApi.ListHistoryQuotes`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ticker** | **string** | string ticker (name or id) of the historyquotes | 

### Other Parameters

Other parameters are passed through a pointer to a apiListHistoryQuotesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fmt** | **string** | string fmt (name or id) of the historyquotes | 
 **order** | **string** | string order (name or id) of the historyquotes | 

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


## ListRealtimeQuotes

> []Quote ListRealtimeQuotes(ctx, ticker).Fmt(fmt).Interval(interval).Filter(filter).S(s).Execute()





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
    ticker := "ticker_example" // string | string ticker (name or id) of the realtimequotes
    fmt := "["json","csv"]" // string | string fmt (name or id) of the realtimequotes
    interval := "["1m","5m"]" // string | string interval (name or id) of the realtimequotes
    filter := "["close"]" // string | string filter (name or id) of the realtimequotes
    s := "["AAPL.US"]" // string | string s (name or id) of the realtimequotes

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.QuotesApi.ListRealtimeQuotes(context.Background(), ticker).Fmt(fmt).Interval(interval).Filter(filter).S(s).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `QuotesApi.ListRealtimeQuotes``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListRealtimeQuotes`: []Quote
    fmt.Fprintf(os.Stdout, "Response from `QuotesApi.ListRealtimeQuotes`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ticker** | **string** | string ticker (name or id) of the realtimequotes | 

### Other Parameters

Other parameters are passed through a pointer to a apiListRealtimeQuotesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fmt** | **string** | string fmt (name or id) of the realtimequotes | 
 **interval** | **string** | string interval (name or id) of the realtimequotes | 
 **filter** | **string** | string filter (name or id) of the realtimequotes | 
 **s** | **string** | string s (name or id) of the realtimequotes | 

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


## ListShortsQuotes

> []SplitsQuote ListShortsQuotes(ctx, ticker).Fmt(fmt).Execute()





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
    ticker := "ticker_example" // string | string ticker (name or id) of the shortsquotes
    fmt := "["json","csv"]" // string | string fmt (name or id) of the shortsquotes

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.QuotesApi.ListShortsQuotes(context.Background(), ticker).Fmt(fmt).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `QuotesApi.ListShortsQuotes``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListShortsQuotes`: []SplitsQuote
    fmt.Fprintf(os.Stdout, "Response from `QuotesApi.ListShortsQuotes`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ticker** | **string** | string ticker (name or id) of the shortsquotes | 

### Other Parameters

Other parameters are passed through a pointer to a apiListShortsQuotesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fmt** | **string** | string fmt (name or id) of the shortsquotes | 

### Return type

[**[]SplitsQuote**](SplitsQuote.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListSplitsQuotes

> []SplitsQuote ListSplitsQuotes(ctx, ticker).Fmt(fmt).Execute()





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
    ticker := "ticker_example" // string | string ticker (name or id) of the splitsquotes
    fmt := "["json","csv"]" // string | string fmt (name or id) of the splitsquotes

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.QuotesApi.ListSplitsQuotes(context.Background(), ticker).Fmt(fmt).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `QuotesApi.ListSplitsQuotes``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListSplitsQuotes`: []SplitsQuote
    fmt.Fprintf(os.Stdout, "Response from `QuotesApi.ListSplitsQuotes`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ticker** | **string** | string ticker (name or id) of the splitsquotes | 

### Other Parameters

Other parameters are passed through a pointer to a apiListSplitsQuotesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fmt** | **string** | string fmt (name or id) of the splitsquotes | 

### Return type

[**[]SplitsQuote**](SplitsQuote.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

