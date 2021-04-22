# FundamentalsESGScores

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Disclaimer** | Pointer to **string** |  | [optional] 
**RatingDate** | Pointer to **string** |  | [optional] 
**TotalEsg** | Pointer to **float32** |  | [optional] 
**TotalEsgPercentile** | Pointer to **float32** |  | [optional] 
**EnvironmentScore** | Pointer to **float32** |  | [optional] 
**EnvironmentScorePercentile** | Pointer to **float32** |  | [optional] 
**SocialScore** | Pointer to **float32** |  | [optional] 
**SocialScorePercentile** | Pointer to **float32** |  | [optional] 
**GovernanceScore** | Pointer to **float32** |  | [optional] 
**GovernanceScorePercentile** | Pointer to **float32** |  | [optional] 
**ControversyLevel** | Pointer to **float32** |  | [optional] 
**ActivitiesInvolvement** | Pointer to [**[]FundamentalsESGScoresActivitiesInvolvement**](FundamentalsESGScoresActivitiesInvolvement.md) |  | [optional] 

## Methods

### NewFundamentalsESGScores

`func NewFundamentalsESGScores() *FundamentalsESGScores`

NewFundamentalsESGScores instantiates a new FundamentalsESGScores object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFundamentalsESGScoresWithDefaults

`func NewFundamentalsESGScoresWithDefaults() *FundamentalsESGScores`

NewFundamentalsESGScoresWithDefaults instantiates a new FundamentalsESGScores object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDisclaimer

`func (o *FundamentalsESGScores) GetDisclaimer() string`

GetDisclaimer returns the Disclaimer field if non-nil, zero value otherwise.

### GetDisclaimerOk

`func (o *FundamentalsESGScores) GetDisclaimerOk() (*string, bool)`

GetDisclaimerOk returns a tuple with the Disclaimer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisclaimer

`func (o *FundamentalsESGScores) SetDisclaimer(v string)`

SetDisclaimer sets Disclaimer field to given value.

### HasDisclaimer

`func (o *FundamentalsESGScores) HasDisclaimer() bool`

HasDisclaimer returns a boolean if a field has been set.

### GetRatingDate

`func (o *FundamentalsESGScores) GetRatingDate() string`

GetRatingDate returns the RatingDate field if non-nil, zero value otherwise.

### GetRatingDateOk

`func (o *FundamentalsESGScores) GetRatingDateOk() (*string, bool)`

GetRatingDateOk returns a tuple with the RatingDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRatingDate

`func (o *FundamentalsESGScores) SetRatingDate(v string)`

SetRatingDate sets RatingDate field to given value.

### HasRatingDate

`func (o *FundamentalsESGScores) HasRatingDate() bool`

HasRatingDate returns a boolean if a field has been set.

### GetTotalEsg

`func (o *FundamentalsESGScores) GetTotalEsg() float32`

GetTotalEsg returns the TotalEsg field if non-nil, zero value otherwise.

### GetTotalEsgOk

`func (o *FundamentalsESGScores) GetTotalEsgOk() (*float32, bool)`

GetTotalEsgOk returns a tuple with the TotalEsg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalEsg

`func (o *FundamentalsESGScores) SetTotalEsg(v float32)`

SetTotalEsg sets TotalEsg field to given value.

### HasTotalEsg

`func (o *FundamentalsESGScores) HasTotalEsg() bool`

HasTotalEsg returns a boolean if a field has been set.

### GetTotalEsgPercentile

`func (o *FundamentalsESGScores) GetTotalEsgPercentile() float32`

GetTotalEsgPercentile returns the TotalEsgPercentile field if non-nil, zero value otherwise.

### GetTotalEsgPercentileOk

`func (o *FundamentalsESGScores) GetTotalEsgPercentileOk() (*float32, bool)`

GetTotalEsgPercentileOk returns a tuple with the TotalEsgPercentile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalEsgPercentile

`func (o *FundamentalsESGScores) SetTotalEsgPercentile(v float32)`

SetTotalEsgPercentile sets TotalEsgPercentile field to given value.

### HasTotalEsgPercentile

`func (o *FundamentalsESGScores) HasTotalEsgPercentile() bool`

HasTotalEsgPercentile returns a boolean if a field has been set.

### GetEnvironmentScore

`func (o *FundamentalsESGScores) GetEnvironmentScore() float32`

GetEnvironmentScore returns the EnvironmentScore field if non-nil, zero value otherwise.

### GetEnvironmentScoreOk

`func (o *FundamentalsESGScores) GetEnvironmentScoreOk() (*float32, bool)`

GetEnvironmentScoreOk returns a tuple with the EnvironmentScore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentScore

`func (o *FundamentalsESGScores) SetEnvironmentScore(v float32)`

SetEnvironmentScore sets EnvironmentScore field to given value.

### HasEnvironmentScore

`func (o *FundamentalsESGScores) HasEnvironmentScore() bool`

HasEnvironmentScore returns a boolean if a field has been set.

### GetEnvironmentScorePercentile

`func (o *FundamentalsESGScores) GetEnvironmentScorePercentile() float32`

GetEnvironmentScorePercentile returns the EnvironmentScorePercentile field if non-nil, zero value otherwise.

### GetEnvironmentScorePercentileOk

`func (o *FundamentalsESGScores) GetEnvironmentScorePercentileOk() (*float32, bool)`

GetEnvironmentScorePercentileOk returns a tuple with the EnvironmentScorePercentile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentScorePercentile

`func (o *FundamentalsESGScores) SetEnvironmentScorePercentile(v float32)`

SetEnvironmentScorePercentile sets EnvironmentScorePercentile field to given value.

### HasEnvironmentScorePercentile

`func (o *FundamentalsESGScores) HasEnvironmentScorePercentile() bool`

HasEnvironmentScorePercentile returns a boolean if a field has been set.

### GetSocialScore

`func (o *FundamentalsESGScores) GetSocialScore() float32`

GetSocialScore returns the SocialScore field if non-nil, zero value otherwise.

### GetSocialScoreOk

`func (o *FundamentalsESGScores) GetSocialScoreOk() (*float32, bool)`

GetSocialScoreOk returns a tuple with the SocialScore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSocialScore

`func (o *FundamentalsESGScores) SetSocialScore(v float32)`

SetSocialScore sets SocialScore field to given value.

### HasSocialScore

`func (o *FundamentalsESGScores) HasSocialScore() bool`

HasSocialScore returns a boolean if a field has been set.

### GetSocialScorePercentile

`func (o *FundamentalsESGScores) GetSocialScorePercentile() float32`

GetSocialScorePercentile returns the SocialScorePercentile field if non-nil, zero value otherwise.

### GetSocialScorePercentileOk

`func (o *FundamentalsESGScores) GetSocialScorePercentileOk() (*float32, bool)`

GetSocialScorePercentileOk returns a tuple with the SocialScorePercentile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSocialScorePercentile

`func (o *FundamentalsESGScores) SetSocialScorePercentile(v float32)`

SetSocialScorePercentile sets SocialScorePercentile field to given value.

### HasSocialScorePercentile

`func (o *FundamentalsESGScores) HasSocialScorePercentile() bool`

HasSocialScorePercentile returns a boolean if a field has been set.

### GetGovernanceScore

`func (o *FundamentalsESGScores) GetGovernanceScore() float32`

GetGovernanceScore returns the GovernanceScore field if non-nil, zero value otherwise.

### GetGovernanceScoreOk

`func (o *FundamentalsESGScores) GetGovernanceScoreOk() (*float32, bool)`

GetGovernanceScoreOk returns a tuple with the GovernanceScore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGovernanceScore

`func (o *FundamentalsESGScores) SetGovernanceScore(v float32)`

SetGovernanceScore sets GovernanceScore field to given value.

### HasGovernanceScore

`func (o *FundamentalsESGScores) HasGovernanceScore() bool`

HasGovernanceScore returns a boolean if a field has been set.

### GetGovernanceScorePercentile

`func (o *FundamentalsESGScores) GetGovernanceScorePercentile() float32`

GetGovernanceScorePercentile returns the GovernanceScorePercentile field if non-nil, zero value otherwise.

### GetGovernanceScorePercentileOk

`func (o *FundamentalsESGScores) GetGovernanceScorePercentileOk() (*float32, bool)`

GetGovernanceScorePercentileOk returns a tuple with the GovernanceScorePercentile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGovernanceScorePercentile

`func (o *FundamentalsESGScores) SetGovernanceScorePercentile(v float32)`

SetGovernanceScorePercentile sets GovernanceScorePercentile field to given value.

### HasGovernanceScorePercentile

`func (o *FundamentalsESGScores) HasGovernanceScorePercentile() bool`

HasGovernanceScorePercentile returns a boolean if a field has been set.

### GetControversyLevel

`func (o *FundamentalsESGScores) GetControversyLevel() float32`

GetControversyLevel returns the ControversyLevel field if non-nil, zero value otherwise.

### GetControversyLevelOk

`func (o *FundamentalsESGScores) GetControversyLevelOk() (*float32, bool)`

GetControversyLevelOk returns a tuple with the ControversyLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControversyLevel

`func (o *FundamentalsESGScores) SetControversyLevel(v float32)`

SetControversyLevel sets ControversyLevel field to given value.

### HasControversyLevel

`func (o *FundamentalsESGScores) HasControversyLevel() bool`

HasControversyLevel returns a boolean if a field has been set.

### GetActivitiesInvolvement

`func (o *FundamentalsESGScores) GetActivitiesInvolvement() []FundamentalsESGScoresActivitiesInvolvement`

GetActivitiesInvolvement returns the ActivitiesInvolvement field if non-nil, zero value otherwise.

### GetActivitiesInvolvementOk

`func (o *FundamentalsESGScores) GetActivitiesInvolvementOk() (*[]FundamentalsESGScoresActivitiesInvolvement, bool)`

GetActivitiesInvolvementOk returns a tuple with the ActivitiesInvolvement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivitiesInvolvement

`func (o *FundamentalsESGScores) SetActivitiesInvolvement(v []FundamentalsESGScoresActivitiesInvolvement)`

SetActivitiesInvolvement sets ActivitiesInvolvement field to given value.

### HasActivitiesInvolvement

`func (o *FundamentalsESGScores) HasActivitiesInvolvement() bool`

HasActivitiesInvolvement returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


