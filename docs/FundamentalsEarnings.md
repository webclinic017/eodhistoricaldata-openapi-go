# FundamentalsEarnings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**History** | Pointer to [**map[string]FundamentalsEarningsHistory**](FundamentalsEarningsHistory.md) |  | [optional] 
**Trend** | Pointer to [**map[string]FundamentalsEarningsTrend**](FundamentalsEarningsTrend.md) |  | [optional] 
**Annual** | Pointer to [**map[string]FundamentalsEarningsAnnual**](FundamentalsEarningsAnnual.md) |  | [optional] 

## Methods

### NewFundamentalsEarnings

`func NewFundamentalsEarnings() *FundamentalsEarnings`

NewFundamentalsEarnings instantiates a new FundamentalsEarnings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFundamentalsEarningsWithDefaults

`func NewFundamentalsEarningsWithDefaults() *FundamentalsEarnings`

NewFundamentalsEarningsWithDefaults instantiates a new FundamentalsEarnings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHistory

`func (o *FundamentalsEarnings) GetHistory() map[string]FundamentalsEarningsHistory`

GetHistory returns the History field if non-nil, zero value otherwise.

### GetHistoryOk

`func (o *FundamentalsEarnings) GetHistoryOk() (*map[string]FundamentalsEarningsHistory, bool)`

GetHistoryOk returns a tuple with the History field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHistory

`func (o *FundamentalsEarnings) SetHistory(v map[string]FundamentalsEarningsHistory)`

SetHistory sets History field to given value.

### HasHistory

`func (o *FundamentalsEarnings) HasHistory() bool`

HasHistory returns a boolean if a field has been set.

### GetTrend

`func (o *FundamentalsEarnings) GetTrend() map[string]FundamentalsEarningsTrend`

GetTrend returns the Trend field if non-nil, zero value otherwise.

### GetTrendOk

`func (o *FundamentalsEarnings) GetTrendOk() (*map[string]FundamentalsEarningsTrend, bool)`

GetTrendOk returns a tuple with the Trend field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrend

`func (o *FundamentalsEarnings) SetTrend(v map[string]FundamentalsEarningsTrend)`

SetTrend sets Trend field to given value.

### HasTrend

`func (o *FundamentalsEarnings) HasTrend() bool`

HasTrend returns a boolean if a field has been set.

### GetAnnual

`func (o *FundamentalsEarnings) GetAnnual() map[string]FundamentalsEarningsAnnual`

GetAnnual returns the Annual field if non-nil, zero value otherwise.

### GetAnnualOk

`func (o *FundamentalsEarnings) GetAnnualOk() (*map[string]FundamentalsEarningsAnnual, bool)`

GetAnnualOk returns a tuple with the Annual field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnnual

`func (o *FundamentalsEarnings) SetAnnual(v map[string]FundamentalsEarningsAnnual)`

SetAnnual sets Annual field to given value.

### HasAnnual

`func (o *FundamentalsEarnings) HasAnnual() bool`

HasAnnual returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


