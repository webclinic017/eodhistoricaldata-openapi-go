# Fundamentals

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**General** | Pointer to [**FundamentalsGeneral**](FundamentalsGeneral.md) |  | [optional] 
**Highlights** | Pointer to [**FundamentalsHighlights**](FundamentalsHighlights.md) |  | [optional] 
**Valuation** | Pointer to [**FundamentalsValuation**](FundamentalsValuation.md) |  | [optional] 
**SharesStats** | Pointer to [**FundamentalsSharesStats**](FundamentalsSharesStats.md) |  | [optional] 
**Technicals** | Pointer to [**FundamentalsTechnicals**](FundamentalsTechnicals.md) |  | [optional] 
**SplitsDividends** | Pointer to [**FundamentalsSplitsDividends**](FundamentalsSplitsDividends.md) |  | [optional] 
**AnalystRatings** | Pointer to [**FundamentalsAnalystRatings**](FundamentalsAnalystRatings.md) |  | [optional] 
**Holders** | Pointer to [**FundamentalsHolders**](FundamentalsHolders.md) |  | [optional] 
**InsiderTransactions** | Pointer to [**[]FundamentalsInsiderTransactionsData**](FundamentalsInsiderTransactionsData.md) |  | [optional] 
**ESGScores** | Pointer to [**FundamentalsESGScores**](FundamentalsESGScores.md) |  | [optional] 
**OutstandingShares** | Pointer to [**FundamentalsOutstandingShares**](FundamentalsOutstandingShares.md) |  | [optional] 
**Earnings** | Pointer to [**FundamentalsEarnings**](FundamentalsEarnings.md) |  | [optional] 
**Financials** | Pointer to [**FundamentalsFinancials**](FundamentalsFinancials.md) |  | [optional] 

## Methods

### NewFundamentals

`func NewFundamentals() *Fundamentals`

NewFundamentals instantiates a new Fundamentals object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFundamentalsWithDefaults

`func NewFundamentalsWithDefaults() *Fundamentals`

NewFundamentalsWithDefaults instantiates a new Fundamentals object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGeneral

`func (o *Fundamentals) GetGeneral() FundamentalsGeneral`

GetGeneral returns the General field if non-nil, zero value otherwise.

### GetGeneralOk

`func (o *Fundamentals) GetGeneralOk() (*FundamentalsGeneral, bool)`

GetGeneralOk returns a tuple with the General field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeneral

`func (o *Fundamentals) SetGeneral(v FundamentalsGeneral)`

SetGeneral sets General field to given value.

### HasGeneral

`func (o *Fundamentals) HasGeneral() bool`

HasGeneral returns a boolean if a field has been set.

### GetHighlights

`func (o *Fundamentals) GetHighlights() FundamentalsHighlights`

GetHighlights returns the Highlights field if non-nil, zero value otherwise.

### GetHighlightsOk

`func (o *Fundamentals) GetHighlightsOk() (*FundamentalsHighlights, bool)`

GetHighlightsOk returns a tuple with the Highlights field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHighlights

`func (o *Fundamentals) SetHighlights(v FundamentalsHighlights)`

SetHighlights sets Highlights field to given value.

### HasHighlights

`func (o *Fundamentals) HasHighlights() bool`

HasHighlights returns a boolean if a field has been set.

### GetValuation

`func (o *Fundamentals) GetValuation() FundamentalsValuation`

GetValuation returns the Valuation field if non-nil, zero value otherwise.

### GetValuationOk

`func (o *Fundamentals) GetValuationOk() (*FundamentalsValuation, bool)`

GetValuationOk returns a tuple with the Valuation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValuation

`func (o *Fundamentals) SetValuation(v FundamentalsValuation)`

SetValuation sets Valuation field to given value.

### HasValuation

`func (o *Fundamentals) HasValuation() bool`

HasValuation returns a boolean if a field has been set.

### GetSharesStats

`func (o *Fundamentals) GetSharesStats() FundamentalsSharesStats`

GetSharesStats returns the SharesStats field if non-nil, zero value otherwise.

### GetSharesStatsOk

`func (o *Fundamentals) GetSharesStatsOk() (*FundamentalsSharesStats, bool)`

GetSharesStatsOk returns a tuple with the SharesStats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharesStats

`func (o *Fundamentals) SetSharesStats(v FundamentalsSharesStats)`

SetSharesStats sets SharesStats field to given value.

### HasSharesStats

`func (o *Fundamentals) HasSharesStats() bool`

HasSharesStats returns a boolean if a field has been set.

### GetTechnicals

`func (o *Fundamentals) GetTechnicals() FundamentalsTechnicals`

GetTechnicals returns the Technicals field if non-nil, zero value otherwise.

### GetTechnicalsOk

`func (o *Fundamentals) GetTechnicalsOk() (*FundamentalsTechnicals, bool)`

GetTechnicalsOk returns a tuple with the Technicals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTechnicals

`func (o *Fundamentals) SetTechnicals(v FundamentalsTechnicals)`

SetTechnicals sets Technicals field to given value.

### HasTechnicals

`func (o *Fundamentals) HasTechnicals() bool`

HasTechnicals returns a boolean if a field has been set.

### GetSplitsDividends

`func (o *Fundamentals) GetSplitsDividends() FundamentalsSplitsDividends`

GetSplitsDividends returns the SplitsDividends field if non-nil, zero value otherwise.

### GetSplitsDividendsOk

`func (o *Fundamentals) GetSplitsDividendsOk() (*FundamentalsSplitsDividends, bool)`

GetSplitsDividendsOk returns a tuple with the SplitsDividends field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSplitsDividends

`func (o *Fundamentals) SetSplitsDividends(v FundamentalsSplitsDividends)`

SetSplitsDividends sets SplitsDividends field to given value.

### HasSplitsDividends

`func (o *Fundamentals) HasSplitsDividends() bool`

HasSplitsDividends returns a boolean if a field has been set.

### GetAnalystRatings

`func (o *Fundamentals) GetAnalystRatings() FundamentalsAnalystRatings`

GetAnalystRatings returns the AnalystRatings field if non-nil, zero value otherwise.

### GetAnalystRatingsOk

`func (o *Fundamentals) GetAnalystRatingsOk() (*FundamentalsAnalystRatings, bool)`

GetAnalystRatingsOk returns a tuple with the AnalystRatings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnalystRatings

`func (o *Fundamentals) SetAnalystRatings(v FundamentalsAnalystRatings)`

SetAnalystRatings sets AnalystRatings field to given value.

### HasAnalystRatings

`func (o *Fundamentals) HasAnalystRatings() bool`

HasAnalystRatings returns a boolean if a field has been set.

### GetHolders

`func (o *Fundamentals) GetHolders() FundamentalsHolders`

GetHolders returns the Holders field if non-nil, zero value otherwise.

### GetHoldersOk

`func (o *Fundamentals) GetHoldersOk() (*FundamentalsHolders, bool)`

GetHoldersOk returns a tuple with the Holders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHolders

`func (o *Fundamentals) SetHolders(v FundamentalsHolders)`

SetHolders sets Holders field to given value.

### HasHolders

`func (o *Fundamentals) HasHolders() bool`

HasHolders returns a boolean if a field has been set.

### GetInsiderTransactions

`func (o *Fundamentals) GetInsiderTransactions() []FundamentalsInsiderTransactionsData`

GetInsiderTransactions returns the InsiderTransactions field if non-nil, zero value otherwise.

### GetInsiderTransactionsOk

`func (o *Fundamentals) GetInsiderTransactionsOk() (*[]FundamentalsInsiderTransactionsData, bool)`

GetInsiderTransactionsOk returns a tuple with the InsiderTransactions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInsiderTransactions

`func (o *Fundamentals) SetInsiderTransactions(v []FundamentalsInsiderTransactionsData)`

SetInsiderTransactions sets InsiderTransactions field to given value.

### HasInsiderTransactions

`func (o *Fundamentals) HasInsiderTransactions() bool`

HasInsiderTransactions returns a boolean if a field has been set.

### GetESGScores

`func (o *Fundamentals) GetESGScores() FundamentalsESGScores`

GetESGScores returns the ESGScores field if non-nil, zero value otherwise.

### GetESGScoresOk

`func (o *Fundamentals) GetESGScoresOk() (*FundamentalsESGScores, bool)`

GetESGScoresOk returns a tuple with the ESGScores field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetESGScores

`func (o *Fundamentals) SetESGScores(v FundamentalsESGScores)`

SetESGScores sets ESGScores field to given value.

### HasESGScores

`func (o *Fundamentals) HasESGScores() bool`

HasESGScores returns a boolean if a field has been set.

### GetOutstandingShares

`func (o *Fundamentals) GetOutstandingShares() FundamentalsOutstandingShares`

GetOutstandingShares returns the OutstandingShares field if non-nil, zero value otherwise.

### GetOutstandingSharesOk

`func (o *Fundamentals) GetOutstandingSharesOk() (*FundamentalsOutstandingShares, bool)`

GetOutstandingSharesOk returns a tuple with the OutstandingShares field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutstandingShares

`func (o *Fundamentals) SetOutstandingShares(v FundamentalsOutstandingShares)`

SetOutstandingShares sets OutstandingShares field to given value.

### HasOutstandingShares

`func (o *Fundamentals) HasOutstandingShares() bool`

HasOutstandingShares returns a boolean if a field has been set.

### GetEarnings

`func (o *Fundamentals) GetEarnings() FundamentalsEarnings`

GetEarnings returns the Earnings field if non-nil, zero value otherwise.

### GetEarningsOk

`func (o *Fundamentals) GetEarningsOk() (*FundamentalsEarnings, bool)`

GetEarningsOk returns a tuple with the Earnings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEarnings

`func (o *Fundamentals) SetEarnings(v FundamentalsEarnings)`

SetEarnings sets Earnings field to given value.

### HasEarnings

`func (o *Fundamentals) HasEarnings() bool`

HasEarnings returns a boolean if a field has been set.

### GetFinancials

`func (o *Fundamentals) GetFinancials() FundamentalsFinancials`

GetFinancials returns the Financials field if non-nil, zero value otherwise.

### GetFinancialsOk

`func (o *Fundamentals) GetFinancialsOk() (*FundamentalsFinancials, bool)`

GetFinancialsOk returns a tuple with the Financials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinancials

`func (o *Fundamentals) SetFinancials(v FundamentalsFinancials)`

SetFinancials sets Financials field to given value.

### HasFinancials

`func (o *Fundamentals) HasFinancials() bool`

HasFinancials returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


