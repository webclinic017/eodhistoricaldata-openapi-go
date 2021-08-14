# FundamentalsFinancialsIncomeStatement

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CurrencySymbol** | Pointer to **string** |  | [optional] 
**Quarterly** | Pointer to [**map[string]IncomeStatement**](IncomeStatement.md) |  | [optional] 
**Yearly** | Pointer to [**map[string]IncomeStatement**](IncomeStatement.md) |  | [optional] 

## Methods

### NewFundamentalsFinancialsIncomeStatement

`func NewFundamentalsFinancialsIncomeStatement() *FundamentalsFinancialsIncomeStatement`

NewFundamentalsFinancialsIncomeStatement instantiates a new FundamentalsFinancialsIncomeStatement object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFundamentalsFinancialsIncomeStatementWithDefaults

`func NewFundamentalsFinancialsIncomeStatementWithDefaults() *FundamentalsFinancialsIncomeStatement`

NewFundamentalsFinancialsIncomeStatementWithDefaults instantiates a new FundamentalsFinancialsIncomeStatement object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCurrencySymbol

`func (o *FundamentalsFinancialsIncomeStatement) GetCurrencySymbol() string`

GetCurrencySymbol returns the CurrencySymbol field if non-nil, zero value otherwise.

### GetCurrencySymbolOk

`func (o *FundamentalsFinancialsIncomeStatement) GetCurrencySymbolOk() (*string, bool)`

GetCurrencySymbolOk returns a tuple with the CurrencySymbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencySymbol

`func (o *FundamentalsFinancialsIncomeStatement) SetCurrencySymbol(v string)`

SetCurrencySymbol sets CurrencySymbol field to given value.

### HasCurrencySymbol

`func (o *FundamentalsFinancialsIncomeStatement) HasCurrencySymbol() bool`

HasCurrencySymbol returns a boolean if a field has been set.

### GetQuarterly

`func (o *FundamentalsFinancialsIncomeStatement) GetQuarterly() map[string]IncomeStatement`

GetQuarterly returns the Quarterly field if non-nil, zero value otherwise.

### GetQuarterlyOk

`func (o *FundamentalsFinancialsIncomeStatement) GetQuarterlyOk() (*map[string]IncomeStatement, bool)`

GetQuarterlyOk returns a tuple with the Quarterly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuarterly

`func (o *FundamentalsFinancialsIncomeStatement) SetQuarterly(v map[string]IncomeStatement)`

SetQuarterly sets Quarterly field to given value.

### HasQuarterly

`func (o *FundamentalsFinancialsIncomeStatement) HasQuarterly() bool`

HasQuarterly returns a boolean if a field has been set.

### GetYearly

`func (o *FundamentalsFinancialsIncomeStatement) GetYearly() map[string]IncomeStatement`

GetYearly returns the Yearly field if non-nil, zero value otherwise.

### GetYearlyOk

`func (o *FundamentalsFinancialsIncomeStatement) GetYearlyOk() (*map[string]IncomeStatement, bool)`

GetYearlyOk returns a tuple with the Yearly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYearly

`func (o *FundamentalsFinancialsIncomeStatement) SetYearly(v map[string]IncomeStatement)`

SetYearly sets Yearly field to given value.

### HasYearly

`func (o *FundamentalsFinancialsIncomeStatement) HasYearly() bool`

HasYearly returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


