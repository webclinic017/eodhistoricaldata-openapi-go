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

// BondFundamentalsRating struct for BondFundamentalsRating
type BondFundamentalsRating struct {
	MoodyRating           *string `json:"MoodyRating,omitempty"`
	MoodyRatingUpdateDate *string `json:"MoodyRatingUpdateDate,omitempty"`
	SPRating              *string `json:"SPRating,omitempty"`
	SPRatingUpdateDate    *string `json:"SPRatingUpdateDate,omitempty"`
}

// NewBondFundamentalsRating instantiates a new BondFundamentalsRating object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBondFundamentalsRating() *BondFundamentalsRating {
	this := BondFundamentalsRating{}
	return &this
}

// NewBondFundamentalsRatingWithDefaults instantiates a new BondFundamentalsRating object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBondFundamentalsRatingWithDefaults() *BondFundamentalsRating {
	this := BondFundamentalsRating{}
	return &this
}

// GetMoodyRating returns the MoodyRating field value if set, zero value otherwise.
func (o *BondFundamentalsRating) GetMoodyRating() string {
	if o == nil || o.MoodyRating == nil {
		var ret string
		return ret
	}
	return *o.MoodyRating
}

// GetMoodyRatingOk returns a tuple with the MoodyRating field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BondFundamentalsRating) GetMoodyRatingOk() (*string, bool) {
	if o == nil || o.MoodyRating == nil {
		return nil, false
	}
	return o.MoodyRating, true
}

// HasMoodyRating returns a boolean if a field has been set.
func (o *BondFundamentalsRating) HasMoodyRating() bool {
	if o != nil && o.MoodyRating != nil {
		return true
	}

	return false
}

// SetMoodyRating gets a reference to the given string and assigns it to the MoodyRating field.
func (o *BondFundamentalsRating) SetMoodyRating(v string) {
	o.MoodyRating = &v
}

// GetMoodyRatingUpdateDate returns the MoodyRatingUpdateDate field value if set, zero value otherwise.
func (o *BondFundamentalsRating) GetMoodyRatingUpdateDate() string {
	if o == nil || o.MoodyRatingUpdateDate == nil {
		var ret string
		return ret
	}
	return *o.MoodyRatingUpdateDate
}

// GetMoodyRatingUpdateDateOk returns a tuple with the MoodyRatingUpdateDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BondFundamentalsRating) GetMoodyRatingUpdateDateOk() (*string, bool) {
	if o == nil || o.MoodyRatingUpdateDate == nil {
		return nil, false
	}
	return o.MoodyRatingUpdateDate, true
}

// HasMoodyRatingUpdateDate returns a boolean if a field has been set.
func (o *BondFundamentalsRating) HasMoodyRatingUpdateDate() bool {
	if o != nil && o.MoodyRatingUpdateDate != nil {
		return true
	}

	return false
}

// SetMoodyRatingUpdateDate gets a reference to the given string and assigns it to the MoodyRatingUpdateDate field.
func (o *BondFundamentalsRating) SetMoodyRatingUpdateDate(v string) {
	o.MoodyRatingUpdateDate = &v
}

// GetSPRating returns the SPRating field value if set, zero value otherwise.
func (o *BondFundamentalsRating) GetSPRating() string {
	if o == nil || o.SPRating == nil {
		var ret string
		return ret
	}
	return *o.SPRating
}

// GetSPRatingOk returns a tuple with the SPRating field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BondFundamentalsRating) GetSPRatingOk() (*string, bool) {
	if o == nil || o.SPRating == nil {
		return nil, false
	}
	return o.SPRating, true
}

// HasSPRating returns a boolean if a field has been set.
func (o *BondFundamentalsRating) HasSPRating() bool {
	if o != nil && o.SPRating != nil {
		return true
	}

	return false
}

// SetSPRating gets a reference to the given string and assigns it to the SPRating field.
func (o *BondFundamentalsRating) SetSPRating(v string) {
	o.SPRating = &v
}

// GetSPRatingUpdateDate returns the SPRatingUpdateDate field value if set, zero value otherwise.
func (o *BondFundamentalsRating) GetSPRatingUpdateDate() string {
	if o == nil || o.SPRatingUpdateDate == nil {
		var ret string
		return ret
	}
	return *o.SPRatingUpdateDate
}

// GetSPRatingUpdateDateOk returns a tuple with the SPRatingUpdateDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BondFundamentalsRating) GetSPRatingUpdateDateOk() (*string, bool) {
	if o == nil || o.SPRatingUpdateDate == nil {
		return nil, false
	}
	return o.SPRatingUpdateDate, true
}

// HasSPRatingUpdateDate returns a boolean if a field has been set.
func (o *BondFundamentalsRating) HasSPRatingUpdateDate() bool {
	if o != nil && o.SPRatingUpdateDate != nil {
		return true
	}

	return false
}

// SetSPRatingUpdateDate gets a reference to the given string and assigns it to the SPRatingUpdateDate field.
func (o *BondFundamentalsRating) SetSPRatingUpdateDate(v string) {
	o.SPRatingUpdateDate = &v
}

func (o BondFundamentalsRating) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.MoodyRating != nil {
		toSerialize["MoodyRating"] = o.MoodyRating
	}
	if o.MoodyRatingUpdateDate != nil {
		toSerialize["MoodyRatingUpdateDate"] = o.MoodyRatingUpdateDate
	}
	if o.SPRating != nil {
		toSerialize["SPRating"] = o.SPRating
	}
	if o.SPRatingUpdateDate != nil {
		toSerialize["SPRatingUpdateDate"] = o.SPRatingUpdateDate
	}
	return json.Marshal(toSerialize)
}

type NullableBondFundamentalsRating struct {
	value *BondFundamentalsRating
	isSet bool
}

func (v NullableBondFundamentalsRating) Get() *BondFundamentalsRating {
	return v.value
}

func (v *NullableBondFundamentalsRating) Set(val *BondFundamentalsRating) {
	v.value = val
	v.isSet = true
}

func (v NullableBondFundamentalsRating) IsSet() bool {
	return v.isSet
}

func (v *NullableBondFundamentalsRating) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBondFundamentalsRating(val *BondFundamentalsRating) *NullableBondFundamentalsRating {
	return &NullableBondFundamentalsRating{value: val, isSet: true}
}

func (v NullableBondFundamentalsRating) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBondFundamentalsRating) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
