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

// New struct for New
type New struct {
	Date    *string   `json:"date,omitempty"`
	Title   *string   `json:"title,omitempty"`
	Content *string   `json:"content,omitempty"`
	Link    *string   `json:"link,omitempty"`
	Symbols *[]string `json:"symbols,omitempty"`
	Tags    *[]string `json:"tags,omitempty"`
}

// NewNew instantiates a new New object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNew() *New {
	this := New{}
	return &this
}

// NewNewWithDefaults instantiates a new New object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNewWithDefaults() *New {
	this := New{}
	return &this
}

// GetDate returns the Date field value if set, zero value otherwise.
func (o *New) GetDate() string {
	if o == nil || o.Date == nil {
		var ret string
		return ret
	}
	return *o.Date
}

// GetDateOk returns a tuple with the Date field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *New) GetDateOk() (*string, bool) {
	if o == nil || o.Date == nil {
		return nil, false
	}
	return o.Date, true
}

// HasDate returns a boolean if a field has been set.
func (o *New) HasDate() bool {
	if o != nil && o.Date != nil {
		return true
	}

	return false
}

// SetDate gets a reference to the given string and assigns it to the Date field.
func (o *New) SetDate(v string) {
	o.Date = &v
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *New) GetTitle() string {
	if o == nil || o.Title == nil {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *New) GetTitleOk() (*string, bool) {
	if o == nil || o.Title == nil {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *New) HasTitle() bool {
	if o != nil && o.Title != nil {
		return true
	}

	return false
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *New) SetTitle(v string) {
	o.Title = &v
}

// GetContent returns the Content field value if set, zero value otherwise.
func (o *New) GetContent() string {
	if o == nil || o.Content == nil {
		var ret string
		return ret
	}
	return *o.Content
}

// GetContentOk returns a tuple with the Content field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *New) GetContentOk() (*string, bool) {
	if o == nil || o.Content == nil {
		return nil, false
	}
	return o.Content, true
}

// HasContent returns a boolean if a field has been set.
func (o *New) HasContent() bool {
	if o != nil && o.Content != nil {
		return true
	}

	return false
}

// SetContent gets a reference to the given string and assigns it to the Content field.
func (o *New) SetContent(v string) {
	o.Content = &v
}

// GetLink returns the Link field value if set, zero value otherwise.
func (o *New) GetLink() string {
	if o == nil || o.Link == nil {
		var ret string
		return ret
	}
	return *o.Link
}

// GetLinkOk returns a tuple with the Link field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *New) GetLinkOk() (*string, bool) {
	if o == nil || o.Link == nil {
		return nil, false
	}
	return o.Link, true
}

// HasLink returns a boolean if a field has been set.
func (o *New) HasLink() bool {
	if o != nil && o.Link != nil {
		return true
	}

	return false
}

// SetLink gets a reference to the given string and assigns it to the Link field.
func (o *New) SetLink(v string) {
	o.Link = &v
}

// GetSymbols returns the Symbols field value if set, zero value otherwise.
func (o *New) GetSymbols() []string {
	if o == nil || o.Symbols == nil {
		var ret []string
		return ret
	}
	return *o.Symbols
}

// GetSymbolsOk returns a tuple with the Symbols field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *New) GetSymbolsOk() (*[]string, bool) {
	if o == nil || o.Symbols == nil {
		return nil, false
	}
	return o.Symbols, true
}

// HasSymbols returns a boolean if a field has been set.
func (o *New) HasSymbols() bool {
	if o != nil && o.Symbols != nil {
		return true
	}

	return false
}

// SetSymbols gets a reference to the given []string and assigns it to the Symbols field.
func (o *New) SetSymbols(v []string) {
	o.Symbols = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *New) GetTags() []string {
	if o == nil || o.Tags == nil {
		var ret []string
		return ret
	}
	return *o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *New) GetTagsOk() (*[]string, bool) {
	if o == nil || o.Tags == nil {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *New) HasTags() bool {
	if o != nil && o.Tags != nil {
		return true
	}

	return false
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *New) SetTags(v []string) {
	o.Tags = &v
}

func (o New) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Date != nil {
		toSerialize["date"] = o.Date
	}
	if o.Title != nil {
		toSerialize["title"] = o.Title
	}
	if o.Content != nil {
		toSerialize["content"] = o.Content
	}
	if o.Link != nil {
		toSerialize["link"] = o.Link
	}
	if o.Symbols != nil {
		toSerialize["symbols"] = o.Symbols
	}
	if o.Tags != nil {
		toSerialize["tags"] = o.Tags
	}
	return json.Marshal(toSerialize)
}

type NullableNew struct {
	value *New
	isSet bool
}

func (v NullableNew) Get() *New {
	return v.value
}

func (v *NullableNew) Set(val *New) {
	v.value = val
	v.isSet = true
}

func (v NullableNew) IsSet() bool {
	return v.isSet
}

func (v *NullableNew) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNew(val *New) *NullableNew {
	return &NullableNew{value: val, isSet: true}
}

func (v NullableNew) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNew) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}