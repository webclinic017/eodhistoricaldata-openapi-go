/*
 * eodhistoricaldata
 *
 * eodhistoricaldata API
 *
 * API version: 1.0.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// Quote struct for Quote
type Quote struct {
	Code          *string  `json:"code,omitempty"`
	Timestamp     *int64   `json:"timestamp,omitempty"`
	Gmtoffset     *int32   `json:"gmtoffset,omitempty"`
	Date          *string  `json:"date,omitempty"`
	Open          *float64 `json:"open,omitempty"`
	High          *float64 `json:"high,omitempty"`
	Low           *float64 `json:"low,omitempty"`
	Close         *float64 `json:"close,omitempty"`
	AdjustedClose *float64 `json:"adjusted_close,omitempty"`
	PreviousClose *float64 `json:"previousClose,omitempty"`
	Volume        *float64 `json:"volume,omitempty"`
	Change        *float32 `json:"change,omitempty"`
}

// NewQuote instantiates a new Quote object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQuote() *Quote {
	this := Quote{}
	return &this
}

// NewQuoteWithDefaults instantiates a new Quote object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQuoteWithDefaults() *Quote {
	this := Quote{}
	return &this
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *Quote) GetCode() string {
	if o == nil || o.Code == nil {
		var ret string
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Quote) GetCodeOk() (*string, bool) {
	if o == nil || o.Code == nil {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *Quote) HasCode() bool {
	if o != nil && o.Code != nil {
		return true
	}

	return false
}

// SetCode gets a reference to the given string and assigns it to the Code field.
func (o *Quote) SetCode(v string) {
	o.Code = &v
}

// GetTimestamp returns the Timestamp field value if set, zero value otherwise.
func (o *Quote) GetTimestamp() int64 {
	if o == nil || o.Timestamp == nil {
		var ret int64
		return ret
	}
	return *o.Timestamp
}

// GetTimestampOk returns a tuple with the Timestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Quote) GetTimestampOk() (*int64, bool) {
	if o == nil || o.Timestamp == nil {
		return nil, false
	}
	return o.Timestamp, true
}

// HasTimestamp returns a boolean if a field has been set.
func (o *Quote) HasTimestamp() bool {
	if o != nil && o.Timestamp != nil {
		return true
	}

	return false
}

// SetTimestamp gets a reference to the given int64 and assigns it to the Timestamp field.
func (o *Quote) SetTimestamp(v int64) {
	o.Timestamp = &v
}

// GetGmtoffset returns the Gmtoffset field value if set, zero value otherwise.
func (o *Quote) GetGmtoffset() int32 {
	if o == nil || o.Gmtoffset == nil {
		var ret int32
		return ret
	}
	return *o.Gmtoffset
}

// GetGmtoffsetOk returns a tuple with the Gmtoffset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Quote) GetGmtoffsetOk() (*int32, bool) {
	if o == nil || o.Gmtoffset == nil {
		return nil, false
	}
	return o.Gmtoffset, true
}

// HasGmtoffset returns a boolean if a field has been set.
func (o *Quote) HasGmtoffset() bool {
	if o != nil && o.Gmtoffset != nil {
		return true
	}

	return false
}

// SetGmtoffset gets a reference to the given int32 and assigns it to the Gmtoffset field.
func (o *Quote) SetGmtoffset(v int32) {
	o.Gmtoffset = &v
}

// GetDate returns the Date field value if set, zero value otherwise.
func (o *Quote) GetDate() string {
	if o == nil || o.Date == nil {
		var ret string
		return ret
	}
	return *o.Date
}

// GetDateOk returns a tuple with the Date field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Quote) GetDateOk() (*string, bool) {
	if o == nil || o.Date == nil {
		return nil, false
	}
	return o.Date, true
}

// HasDate returns a boolean if a field has been set.
func (o *Quote) HasDate() bool {
	if o != nil && o.Date != nil {
		return true
	}

	return false
}

// SetDate gets a reference to the given string and assigns it to the Date field.
func (o *Quote) SetDate(v string) {
	o.Date = &v
}

// GetOpen returns the Open field value if set, zero value otherwise.
func (o *Quote) GetOpen() float64 {
	if o == nil || o.Open == nil {
		var ret float64
		return ret
	}
	return *o.Open
}

// GetOpenOk returns a tuple with the Open field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Quote) GetOpenOk() (*float64, bool) {
	if o == nil || o.Open == nil {
		return nil, false
	}
	return o.Open, true
}

// HasOpen returns a boolean if a field has been set.
func (o *Quote) HasOpen() bool {
	if o != nil && o.Open != nil {
		return true
	}

	return false
}

// SetOpen gets a reference to the given float64 and assigns it to the Open field.
func (o *Quote) SetOpen(v float64) {
	o.Open = &v
}

// GetHigh returns the High field value if set, zero value otherwise.
func (o *Quote) GetHigh() float64 {
	if o == nil || o.High == nil {
		var ret float64
		return ret
	}
	return *o.High
}

// GetHighOk returns a tuple with the High field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Quote) GetHighOk() (*float64, bool) {
	if o == nil || o.High == nil {
		return nil, false
	}
	return o.High, true
}

// HasHigh returns a boolean if a field has been set.
func (o *Quote) HasHigh() bool {
	if o != nil && o.High != nil {
		return true
	}

	return false
}

// SetHigh gets a reference to the given float64 and assigns it to the High field.
func (o *Quote) SetHigh(v float64) {
	o.High = &v
}

// GetLow returns the Low field value if set, zero value otherwise.
func (o *Quote) GetLow() float64 {
	if o == nil || o.Low == nil {
		var ret float64
		return ret
	}
	return *o.Low
}

// GetLowOk returns a tuple with the Low field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Quote) GetLowOk() (*float64, bool) {
	if o == nil || o.Low == nil {
		return nil, false
	}
	return o.Low, true
}

// HasLow returns a boolean if a field has been set.
func (o *Quote) HasLow() bool {
	if o != nil && o.Low != nil {
		return true
	}

	return false
}

// SetLow gets a reference to the given float64 and assigns it to the Low field.
func (o *Quote) SetLow(v float64) {
	o.Low = &v
}

// GetClose returns the Close field value if set, zero value otherwise.
func (o *Quote) GetClose() float64 {
	if o == nil || o.Close == nil {
		var ret float64
		return ret
	}
	return *o.Close
}

// GetCloseOk returns a tuple with the Close field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Quote) GetCloseOk() (*float64, bool) {
	if o == nil || o.Close == nil {
		return nil, false
	}
	return o.Close, true
}

// HasClose returns a boolean if a field has been set.
func (o *Quote) HasClose() bool {
	if o != nil && o.Close != nil {
		return true
	}

	return false
}

// SetClose gets a reference to the given float64 and assigns it to the Close field.
func (o *Quote) SetClose(v float64) {
	o.Close = &v
}

// GetAdjustedClose returns the AdjustedClose field value if set, zero value otherwise.
func (o *Quote) GetAdjustedClose() float64 {
	if o == nil || o.AdjustedClose == nil {
		var ret float64
		return ret
	}
	return *o.AdjustedClose
}

// GetAdjustedCloseOk returns a tuple with the AdjustedClose field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Quote) GetAdjustedCloseOk() (*float64, bool) {
	if o == nil || o.AdjustedClose == nil {
		return nil, false
	}
	return o.AdjustedClose, true
}

// HasAdjustedClose returns a boolean if a field has been set.
func (o *Quote) HasAdjustedClose() bool {
	if o != nil && o.AdjustedClose != nil {
		return true
	}

	return false
}

// SetAdjustedClose gets a reference to the given float64 and assigns it to the AdjustedClose field.
func (o *Quote) SetAdjustedClose(v float64) {
	o.AdjustedClose = &v
}

// GetPreviousClose returns the PreviousClose field value if set, zero value otherwise.
func (o *Quote) GetPreviousClose() float64 {
	if o == nil || o.PreviousClose == nil {
		var ret float64
		return ret
	}
	return *o.PreviousClose
}

// GetPreviousCloseOk returns a tuple with the PreviousClose field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Quote) GetPreviousCloseOk() (*float64, bool) {
	if o == nil || o.PreviousClose == nil {
		return nil, false
	}
	return o.PreviousClose, true
}

// HasPreviousClose returns a boolean if a field has been set.
func (o *Quote) HasPreviousClose() bool {
	if o != nil && o.PreviousClose != nil {
		return true
	}

	return false
}

// SetPreviousClose gets a reference to the given float64 and assigns it to the PreviousClose field.
func (o *Quote) SetPreviousClose(v float64) {
	o.PreviousClose = &v
}

// GetVolume returns the Volume field value if set, zero value otherwise.
func (o *Quote) GetVolume() float64 {
	if o == nil || o.Volume == nil {
		var ret float64
		return ret
	}
	return *o.Volume
}

// GetVolumeOk returns a tuple with the Volume field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Quote) GetVolumeOk() (*float64, bool) {
	if o == nil || o.Volume == nil {
		return nil, false
	}
	return o.Volume, true
}

// HasVolume returns a boolean if a field has been set.
func (o *Quote) HasVolume() bool {
	if o != nil && o.Volume != nil {
		return true
	}

	return false
}

// SetVolume gets a reference to the given float64 and assigns it to the Volume field.
func (o *Quote) SetVolume(v float64) {
	o.Volume = &v
}

// GetChange returns the Change field value if set, zero value otherwise.
func (o *Quote) GetChange() float32 {
	if o == nil || o.Change == nil {
		var ret float32
		return ret
	}
	return *o.Change
}

// GetChangeOk returns a tuple with the Change field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Quote) GetChangeOk() (*float32, bool) {
	if o == nil || o.Change == nil {
		return nil, false
	}
	return o.Change, true
}

// HasChange returns a boolean if a field has been set.
func (o *Quote) HasChange() bool {
	if o != nil && o.Change != nil {
		return true
	}

	return false
}

// SetChange gets a reference to the given float32 and assigns it to the Change field.
func (o *Quote) SetChange(v float32) {
	o.Change = &v
}

func (o Quote) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Code != nil {
		toSerialize["code"] = o.Code
	}
	if o.Timestamp != nil {
		toSerialize["timestamp"] = o.Timestamp
	}
	if o.Gmtoffset != nil {
		toSerialize["gmtoffset"] = o.Gmtoffset
	}
	if o.Date != nil {
		toSerialize["date"] = o.Date
	}
	if o.Open != nil {
		toSerialize["open"] = o.Open
	}
	if o.High != nil {
		toSerialize["high"] = o.High
	}
	if o.Low != nil {
		toSerialize["low"] = o.Low
	}
	if o.Close != nil {
		toSerialize["close"] = o.Close
	}
	if o.AdjustedClose != nil {
		toSerialize["adjusted_close"] = o.AdjustedClose
	}
	if o.PreviousClose != nil {
		toSerialize["previousClose"] = o.PreviousClose
	}
	if o.Volume != nil {
		toSerialize["volume"] = o.Volume
	}
	if o.Change != nil {
		toSerialize["change"] = o.Change
	}
	return json.Marshal(toSerialize)
}

type NullableQuote struct {
	value *Quote
	isSet bool
}

func (v NullableQuote) Get() *Quote {
	return v.value
}

func (v *NullableQuote) Set(val *Quote) {
	v.value = val
	v.isSet = true
}

func (v NullableQuote) IsSet() bool {
	return v.isSet
}

func (v *NullableQuote) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQuote(val *Quote) *NullableQuote {
	return &NullableQuote{value: val, isSet: true}
}

func (v NullableQuote) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQuote) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
