# FundamentalsFinancialsCashFlow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CurrencySymbol** | Pointer to **string** |  | [optional] 
**Quarterly** | Pointer to [**map[string]CashFlow**](CashFlow.md) |  | [optional] 
**Yearly** | Pointer to [**map[string]CashFlow**](CashFlow.md) |  | [optional] 

## Methods

### NewFundamentalsFinancialsCashFlow

`func NewFundamentalsFinancialsCashFlow() *FundamentalsFinancialsCashFlow`

NewFundamentalsFinancialsCashFlow instantiates a new FundamentalsFinancialsCashFlow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFundamentalsFinancialsCashFlowWithDefaults

`func NewFundamentalsFinancialsCashFlowWithDefaults() *FundamentalsFinancialsCashFlow`

NewFundamentalsFinancialsCashFlowWithDefaults instantiates a new FundamentalsFinancialsCashFlow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCurrencySymbol

`func (o *FundamentalsFinancialsCashFlow) GetCurrencySymbol() string`

GetCurrencySymbol returns the CurrencySymbol field if non-nil, zero value otherwise.

### GetCurrencySymbolOk

`func (o *FundamentalsFinancialsCashFlow) GetCurrencySymbolOk() (*string, bool)`

GetCurrencySymbolOk returns a tuple with the CurrencySymbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencySymbol

`func (o *FundamentalsFinancialsCashFlow) SetCurrencySymbol(v string)`

SetCurrencySymbol sets CurrencySymbol field to given value.

### HasCurrencySymbol

`func (o *FundamentalsFinancialsCashFlow) HasCurrencySymbol() bool`

HasCurrencySymbol returns a boolean if a field has been set.

### GetQuarterly

`func (o *FundamentalsFinancialsCashFlow) GetQuarterly() map[string]CashFlow`

GetQuarterly returns the Quarterly field if non-nil, zero value otherwise.

### GetQuarterlyOk

`func (o *FundamentalsFinancialsCashFlow) GetQuarterlyOk() (*map[string]CashFlow, bool)`

GetQuarterlyOk returns a tuple with the Quarterly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuarterly

`func (o *FundamentalsFinancialsCashFlow) SetQuarterly(v map[string]CashFlow)`

SetQuarterly sets Quarterly field to given value.

### HasQuarterly

`func (o *FundamentalsFinancialsCashFlow) HasQuarterly() bool`

HasQuarterly returns a boolean if a field has been set.

### GetYearly

`func (o *FundamentalsFinancialsCashFlow) GetYearly() map[string]CashFlow`

GetYearly returns the Yearly field if non-nil, zero value otherwise.

### GetYearlyOk

`func (o *FundamentalsFinancialsCashFlow) GetYearlyOk() (*map[string]CashFlow, bool)`

GetYearlyOk returns a tuple with the Yearly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYearly

`func (o *FundamentalsFinancialsCashFlow) SetYearly(v map[string]CashFlow)`

SetYearly sets Yearly field to given value.

### HasYearly

`func (o *FundamentalsFinancialsCashFlow) HasYearly() bool`

HasYearly returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


