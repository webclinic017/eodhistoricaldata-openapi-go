# FundamentalsSplitsDividends

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ForwardAnnualDividendRate** | Pointer to **float64** |  | [optional] 
**ForwardAnnualDividendYield** | Pointer to **float64** |  | [optional] 
**PayoutRatio** | Pointer to **float64** |  | [optional] 
**DividendDate** | Pointer to **string** |  | [optional] 
**ExDividendDate** | Pointer to **string** |  | [optional] 
**LastSplitFactor** | Pointer to **string** |  | [optional] 
**LastSplitDate** | Pointer to **string** |  | [optional] 
**NumberDividendsByYear** | Pointer to [**[]FundamentalsSplitsDividendsNumberDividendsByYear**](FundamentalsSplitsDividendsNumberDividendsByYear.md) |  | [optional] 

## Methods

### NewFundamentalsSplitsDividends

`func NewFundamentalsSplitsDividends() *FundamentalsSplitsDividends`

NewFundamentalsSplitsDividends instantiates a new FundamentalsSplitsDividends object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFundamentalsSplitsDividendsWithDefaults

`func NewFundamentalsSplitsDividendsWithDefaults() *FundamentalsSplitsDividends`

NewFundamentalsSplitsDividendsWithDefaults instantiates a new FundamentalsSplitsDividends object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetForwardAnnualDividendRate

`func (o *FundamentalsSplitsDividends) GetForwardAnnualDividendRate() float64`

GetForwardAnnualDividendRate returns the ForwardAnnualDividendRate field if non-nil, zero value otherwise.

### GetForwardAnnualDividendRateOk

`func (o *FundamentalsSplitsDividends) GetForwardAnnualDividendRateOk() (*float64, bool)`

GetForwardAnnualDividendRateOk returns a tuple with the ForwardAnnualDividendRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForwardAnnualDividendRate

`func (o *FundamentalsSplitsDividends) SetForwardAnnualDividendRate(v float64)`

SetForwardAnnualDividendRate sets ForwardAnnualDividendRate field to given value.

### HasForwardAnnualDividendRate

`func (o *FundamentalsSplitsDividends) HasForwardAnnualDividendRate() bool`

HasForwardAnnualDividendRate returns a boolean if a field has been set.

### GetForwardAnnualDividendYield

`func (o *FundamentalsSplitsDividends) GetForwardAnnualDividendYield() float64`

GetForwardAnnualDividendYield returns the ForwardAnnualDividendYield field if non-nil, zero value otherwise.

### GetForwardAnnualDividendYieldOk

`func (o *FundamentalsSplitsDividends) GetForwardAnnualDividendYieldOk() (*float64, bool)`

GetForwardAnnualDividendYieldOk returns a tuple with the ForwardAnnualDividendYield field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForwardAnnualDividendYield

`func (o *FundamentalsSplitsDividends) SetForwardAnnualDividendYield(v float64)`

SetForwardAnnualDividendYield sets ForwardAnnualDividendYield field to given value.

### HasForwardAnnualDividendYield

`func (o *FundamentalsSplitsDividends) HasForwardAnnualDividendYield() bool`

HasForwardAnnualDividendYield returns a boolean if a field has been set.

### GetPayoutRatio

`func (o *FundamentalsSplitsDividends) GetPayoutRatio() float64`

GetPayoutRatio returns the PayoutRatio field if non-nil, zero value otherwise.

### GetPayoutRatioOk

`func (o *FundamentalsSplitsDividends) GetPayoutRatioOk() (*float64, bool)`

GetPayoutRatioOk returns a tuple with the PayoutRatio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayoutRatio

`func (o *FundamentalsSplitsDividends) SetPayoutRatio(v float64)`

SetPayoutRatio sets PayoutRatio field to given value.

### HasPayoutRatio

`func (o *FundamentalsSplitsDividends) HasPayoutRatio() bool`

HasPayoutRatio returns a boolean if a field has been set.

### GetDividendDate

`func (o *FundamentalsSplitsDividends) GetDividendDate() string`

GetDividendDate returns the DividendDate field if non-nil, zero value otherwise.

### GetDividendDateOk

`func (o *FundamentalsSplitsDividends) GetDividendDateOk() (*string, bool)`

GetDividendDateOk returns a tuple with the DividendDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDividendDate

`func (o *FundamentalsSplitsDividends) SetDividendDate(v string)`

SetDividendDate sets DividendDate field to given value.

### HasDividendDate

`func (o *FundamentalsSplitsDividends) HasDividendDate() bool`

HasDividendDate returns a boolean if a field has been set.

### GetExDividendDate

`func (o *FundamentalsSplitsDividends) GetExDividendDate() string`

GetExDividendDate returns the ExDividendDate field if non-nil, zero value otherwise.

### GetExDividendDateOk

`func (o *FundamentalsSplitsDividends) GetExDividendDateOk() (*string, bool)`

GetExDividendDateOk returns a tuple with the ExDividendDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExDividendDate

`func (o *FundamentalsSplitsDividends) SetExDividendDate(v string)`

SetExDividendDate sets ExDividendDate field to given value.

### HasExDividendDate

`func (o *FundamentalsSplitsDividends) HasExDividendDate() bool`

HasExDividendDate returns a boolean if a field has been set.

### GetLastSplitFactor

`func (o *FundamentalsSplitsDividends) GetLastSplitFactor() string`

GetLastSplitFactor returns the LastSplitFactor field if non-nil, zero value otherwise.

### GetLastSplitFactorOk

`func (o *FundamentalsSplitsDividends) GetLastSplitFactorOk() (*string, bool)`

GetLastSplitFactorOk returns a tuple with the LastSplitFactor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSplitFactor

`func (o *FundamentalsSplitsDividends) SetLastSplitFactor(v string)`

SetLastSplitFactor sets LastSplitFactor field to given value.

### HasLastSplitFactor

`func (o *FundamentalsSplitsDividends) HasLastSplitFactor() bool`

HasLastSplitFactor returns a boolean if a field has been set.

### GetLastSplitDate

`func (o *FundamentalsSplitsDividends) GetLastSplitDate() string`

GetLastSplitDate returns the LastSplitDate field if non-nil, zero value otherwise.

### GetLastSplitDateOk

`func (o *FundamentalsSplitsDividends) GetLastSplitDateOk() (*string, bool)`

GetLastSplitDateOk returns a tuple with the LastSplitDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSplitDate

`func (o *FundamentalsSplitsDividends) SetLastSplitDate(v string)`

SetLastSplitDate sets LastSplitDate field to given value.

### HasLastSplitDate

`func (o *FundamentalsSplitsDividends) HasLastSplitDate() bool`

HasLastSplitDate returns a boolean if a field has been set.

### GetNumberDividendsByYear

`func (o *FundamentalsSplitsDividends) GetNumberDividendsByYear() []FundamentalsSplitsDividendsNumberDividendsByYear`

GetNumberDividendsByYear returns the NumberDividendsByYear field if non-nil, zero value otherwise.

### GetNumberDividendsByYearOk

`func (o *FundamentalsSplitsDividends) GetNumberDividendsByYearOk() (*[]FundamentalsSplitsDividendsNumberDividendsByYear, bool)`

GetNumberDividendsByYearOk returns a tuple with the NumberDividendsByYear field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberDividendsByYear

`func (o *FundamentalsSplitsDividends) SetNumberDividendsByYear(v []FundamentalsSplitsDividendsNumberDividendsByYear)`

SetNumberDividendsByYear sets NumberDividendsByYear field to given value.

### HasNumberDividendsByYear

`func (o *FundamentalsSplitsDividends) HasNumberDividendsByYear() bool`

HasNumberDividendsByYear returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


