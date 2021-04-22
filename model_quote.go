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
	Date          *string  `json:"date,omitempty"`
	Open          *float64 `json:"open,omitempty"`
	High          *float64 `json:"high,omitempty"`
	Low           *float64 `json:"low,omitempty"`
	Close         *float64 `json:"close,omitempty"`
	AdjustedClose *float64 `json:"adjusted_close,omitempty"`
	Volume        *int64   `json:"volume,omitempty"`
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

// GetVolume returns the Volume field value if set, zero value otherwise.
func (o *Quote) GetVolume() int64 {
	if o == nil || o.Volume == nil {
		var ret int64
		return ret
	}
	return *o.Volume
}

// GetVolumeOk returns a tuple with the Volume field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Quote) GetVolumeOk() (*int64, bool) {
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

// SetVolume gets a reference to the given int64 and assigns it to the Volume field.
func (o *Quote) SetVolume(v int64) {
	o.Volume = &v
}

func (o Quote) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
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
	if o.Volume != nil {
		toSerialize["volume"] = o.Volume
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
