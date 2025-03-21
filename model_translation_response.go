/*
Groq API

API for interacting with Groq's language models and audio transcription/translation services

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the TranslationResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TranslationResponse{}

// TranslationResponse struct for TranslationResponse
type TranslationResponse struct {
	Text *string `json:"text,omitempty"`
}

// NewTranslationResponse instantiates a new TranslationResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTranslationResponse() *TranslationResponse {
	this := TranslationResponse{}
	return &this
}

// NewTranslationResponseWithDefaults instantiates a new TranslationResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTranslationResponseWithDefaults() *TranslationResponse {
	this := TranslationResponse{}
	return &this
}

// GetText returns the Text field value if set, zero value otherwise.
func (o *TranslationResponse) GetText() string {
	if o == nil || IsNil(o.Text) {
		var ret string
		return ret
	}
	return *o.Text
}

// GetTextOk returns a tuple with the Text field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TranslationResponse) GetTextOk() (*string, bool) {
	if o == nil || IsNil(o.Text) {
		return nil, false
	}
	return o.Text, true
}

// HasText returns a boolean if a field has been set.
func (o *TranslationResponse) HasText() bool {
	if o != nil && !IsNil(o.Text) {
		return true
	}

	return false
}

// SetText gets a reference to the given string and assigns it to the Text field.
func (o *TranslationResponse) SetText(v string) {
	o.Text = &v
}

func (o TranslationResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TranslationResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Text) {
		toSerialize["text"] = o.Text
	}
	return toSerialize, nil
}

type NullableTranslationResponse struct {
	value *TranslationResponse
	isSet bool
}

func (v NullableTranslationResponse) Get() *TranslationResponse {
	return v.value
}

func (v *NullableTranslationResponse) Set(val *TranslationResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableTranslationResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableTranslationResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTranslationResponse(val *TranslationResponse) *NullableTranslationResponse {
	return &NullableTranslationResponse{value: val, isSet: true}
}

func (v NullableTranslationResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTranslationResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


