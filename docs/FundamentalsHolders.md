# FundamentalsHolders

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Institutions** | Pointer to [**map[string]FundamentalsHoldersData**](FundamentalsHoldersData.md) |  | [optional] 
**Funds** | Pointer to [**map[string]FundamentalsHoldersData**](FundamentalsHoldersData.md) |  | [optional] 

## Methods

### NewFundamentalsHolders

`func NewFundamentalsHolders() *FundamentalsHolders`

NewFundamentalsHolders instantiates a new FundamentalsHolders object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFundamentalsHoldersWithDefaults

`func NewFundamentalsHoldersWithDefaults() *FundamentalsHolders`

NewFundamentalsHoldersWithDefaults instantiates a new FundamentalsHolders object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInstitutions

`func (o *FundamentalsHolders) GetInstitutions() map[string]FundamentalsHoldersData`

GetInstitutions returns the Institutions field if non-nil, zero value otherwise.

### GetInstitutionsOk

`func (o *FundamentalsHolders) GetInstitutionsOk() (*map[string]FundamentalsHoldersData, bool)`

GetInstitutionsOk returns a tuple with the Institutions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstitutions

`func (o *FundamentalsHolders) SetInstitutions(v map[string]FundamentalsHoldersData)`

SetInstitutions sets Institutions field to given value.

### HasInstitutions

`func (o *FundamentalsHolders) HasInstitutions() bool`

HasInstitutions returns a boolean if a field has been set.

### GetFunds

`func (o *FundamentalsHolders) GetFunds() map[string]FundamentalsHoldersData`

GetFunds returns the Funds field if non-nil, zero value otherwise.

### GetFundsOk

`func (o *FundamentalsHolders) GetFundsOk() (*map[string]FundamentalsHoldersData, bool)`

GetFundsOk returns a tuple with the Funds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunds

`func (o *FundamentalsHolders) SetFunds(v map[string]FundamentalsHoldersData)`

SetFunds sets Funds field to given value.

### HasFunds

`func (o *FundamentalsHolders) HasFunds() bool`

HasFunds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


