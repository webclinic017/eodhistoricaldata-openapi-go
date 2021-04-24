# \BondsApi

All URIs are relative to *https://eodhistoricaldata.com/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ReadBondFundamentals**](BondsApi.md#ReadBondFundamentals) | **Get** /bond-fundamentals/{bond} | 



## ReadBondFundamentals

> BondFundamentals ReadBondFundamentals(ctx, bond).Fmt(fmt).Execute()





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
    bond := "bond_example" // string | string bond (name or id) of the bondfundamentals
    fmt := "["json","csv"]" // string | string fmt (name or id) of the bondfundamentals

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.BondsApi.ReadBondFundamentals(context.Background(), bond).Fmt(fmt).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BondsApi.ReadBondFundamentals``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReadBondFundamentals`: BondFundamentals
    fmt.Fprintf(os.Stdout, "Response from `BondsApi.ReadBondFundamentals`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bond** | **string** | string bond (name or id) of the bondfundamentals | 

### Other Parameters

Other parameters are passed through a pointer to a apiReadBondFundamentalsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fmt** | **string** | string fmt (name or id) of the bondfundamentals | 

### Return type

[**BondFundamentals**](BondFundamentals.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

