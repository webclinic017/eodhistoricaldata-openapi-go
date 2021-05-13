# Quote

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | Pointer to **string** |  | [optional] 
**Timestamp** | Pointer to **float64** |  | [optional] 
**Gmtoffset** | Pointer to **int32** |  | [optional] 
**Date** | Pointer to **string** |  | [optional] 
**Open** | Pointer to **float64** |  | [optional] 
**High** | Pointer to **float64** |  | [optional] 
**Low** | Pointer to **float64** |  | [optional] 
**Close** | Pointer to **float64** |  | [optional] 
**AdjustedClose** | Pointer to **float64** |  | [optional] 
**PreviousClose** | Pointer to **float64** |  | [optional] 
**Volume** | Pointer to **float64** |  | [optional] 
**Change** | Pointer to **float32** |  | [optional] 

## Methods

### NewQuote

`func NewQuote() *Quote`

NewQuote instantiates a new Quote object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQuoteWithDefaults

`func NewQuoteWithDefaults() *Quote`

NewQuoteWithDefaults instantiates a new Quote object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *Quote) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *Quote) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *Quote) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *Quote) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetTimestamp

`func (o *Quote) GetTimestamp() float64`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *Quote) GetTimestampOk() (*float64, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *Quote) SetTimestamp(v float64)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *Quote) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetGmtoffset

`func (o *Quote) GetGmtoffset() int32`

GetGmtoffset returns the Gmtoffset field if non-nil, zero value otherwise.

### GetGmtoffsetOk

`func (o *Quote) GetGmtoffsetOk() (*int32, bool)`

GetGmtoffsetOk returns a tuple with the Gmtoffset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGmtoffset

`func (o *Quote) SetGmtoffset(v int32)`

SetGmtoffset sets Gmtoffset field to given value.

### HasGmtoffset

`func (o *Quote) HasGmtoffset() bool`

HasGmtoffset returns a boolean if a field has been set.

### GetDate

`func (o *Quote) GetDate() string`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *Quote) GetDateOk() (*string, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *Quote) SetDate(v string)`

SetDate sets Date field to given value.

### HasDate

`func (o *Quote) HasDate() bool`

HasDate returns a boolean if a field has been set.

### GetOpen

`func (o *Quote) GetOpen() float64`

GetOpen returns the Open field if non-nil, zero value otherwise.

### GetOpenOk

`func (o *Quote) GetOpenOk() (*float64, bool)`

GetOpenOk returns a tuple with the Open field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpen

`func (o *Quote) SetOpen(v float64)`

SetOpen sets Open field to given value.

### HasOpen

`func (o *Quote) HasOpen() bool`

HasOpen returns a boolean if a field has been set.

### GetHigh

`func (o *Quote) GetHigh() float64`

GetHigh returns the High field if non-nil, zero value otherwise.

### GetHighOk

`func (o *Quote) GetHighOk() (*float64, bool)`

GetHighOk returns a tuple with the High field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHigh

`func (o *Quote) SetHigh(v float64)`

SetHigh sets High field to given value.

### HasHigh

`func (o *Quote) HasHigh() bool`

HasHigh returns a boolean if a field has been set.

### GetLow

`func (o *Quote) GetLow() float64`

GetLow returns the Low field if non-nil, zero value otherwise.

### GetLowOk

`func (o *Quote) GetLowOk() (*float64, bool)`

GetLowOk returns a tuple with the Low field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLow

`func (o *Quote) SetLow(v float64)`

SetLow sets Low field to given value.

### HasLow

`func (o *Quote) HasLow() bool`

HasLow returns a boolean if a field has been set.

### GetClose

`func (o *Quote) GetClose() float64`

GetClose returns the Close field if non-nil, zero value otherwise.

### GetCloseOk

`func (o *Quote) GetCloseOk() (*float64, bool)`

GetCloseOk returns a tuple with the Close field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClose

`func (o *Quote) SetClose(v float64)`

SetClose sets Close field to given value.

### HasClose

`func (o *Quote) HasClose() bool`

HasClose returns a boolean if a field has been set.

### GetAdjustedClose

`func (o *Quote) GetAdjustedClose() float64`

GetAdjustedClose returns the AdjustedClose field if non-nil, zero value otherwise.

### GetAdjustedCloseOk

`func (o *Quote) GetAdjustedCloseOk() (*float64, bool)`

GetAdjustedCloseOk returns a tuple with the AdjustedClose field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdjustedClose

`func (o *Quote) SetAdjustedClose(v float64)`

SetAdjustedClose sets AdjustedClose field to given value.

### HasAdjustedClose

`func (o *Quote) HasAdjustedClose() bool`

HasAdjustedClose returns a boolean if a field has been set.

### GetPreviousClose

`func (o *Quote) GetPreviousClose() float64`

GetPreviousClose returns the PreviousClose field if non-nil, zero value otherwise.

### GetPreviousCloseOk

`func (o *Quote) GetPreviousCloseOk() (*float64, bool)`

GetPreviousCloseOk returns a tuple with the PreviousClose field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreviousClose

`func (o *Quote) SetPreviousClose(v float64)`

SetPreviousClose sets PreviousClose field to given value.

### HasPreviousClose

`func (o *Quote) HasPreviousClose() bool`

HasPreviousClose returns a boolean if a field has been set.

### GetVolume

`func (o *Quote) GetVolume() float64`

GetVolume returns the Volume field if non-nil, zero value otherwise.

### GetVolumeOk

`func (o *Quote) GetVolumeOk() (*float64, bool)`

GetVolumeOk returns a tuple with the Volume field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolume

`func (o *Quote) SetVolume(v float64)`

SetVolume sets Volume field to given value.

### HasVolume

`func (o *Quote) HasVolume() bool`

HasVolume returns a boolean if a field has been set.

### GetChange

`func (o *Quote) GetChange() float32`

GetChange returns the Change field if non-nil, zero value otherwise.

### GetChangeOk

`func (o *Quote) GetChangeOk() (*float32, bool)`

GetChangeOk returns a tuple with the Change field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChange

`func (o *Quote) SetChange(v float32)`

SetChange sets Change field to given value.

### HasChange

`func (o *Quote) HasChange() bool`

HasChange returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


