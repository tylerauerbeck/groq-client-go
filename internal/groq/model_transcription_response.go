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

// checks if the TranscriptionResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TranscriptionResponse{}

// TranscriptionResponse struct for TranscriptionResponse
type TranscriptionResponse struct {
	Text *string `json:"text,omitempty"`
}

// NewTranscriptionResponse instantiates a new TranscriptionResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTranscriptionResponse() *TranscriptionResponse {
	this := TranscriptionResponse{}
	return &this
}

// NewTranscriptionResponseWithDefaults instantiates a new TranscriptionResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTranscriptionResponseWithDefaults() *TranscriptionResponse {
	this := TranscriptionResponse{}
	return &this
}

// GetText returns the Text field value if set, zero value otherwise.
func (o *TranscriptionResponse) GetText() string {
	if o == nil || IsNil(o.Text) {
		var ret string
		return ret
	}
	return *o.Text
}

// GetTextOk returns a tuple with the Text field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TranscriptionResponse) GetTextOk() (*string, bool) {
	if o == nil || IsNil(o.Text) {
		return nil, false
	}
	return o.Text, true
}

// HasText returns a boolean if a field has been set.
func (o *TranscriptionResponse) HasText() bool {
	if o != nil && !IsNil(o.Text) {
		return true
	}

	return false
}

// SetText gets a reference to the given string and assigns it to the Text field.
func (o *TranscriptionResponse) SetText(v string) {
	o.Text = &v
}

func (o TranscriptionResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TranscriptionResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Text) {
		toSerialize["text"] = o.Text
	}
	return toSerialize, nil
}

type NullableTranscriptionResponse struct {
	value *TranscriptionResponse
	isSet bool
}

func (v NullableTranscriptionResponse) Get() *TranscriptionResponse {
	return v.value
}

func (v *NullableTranscriptionResponse) Set(val *TranscriptionResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableTranscriptionResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableTranscriptionResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTranscriptionResponse(val *TranscriptionResponse) *NullableTranscriptionResponse {
	return &NullableTranscriptionResponse{value: val, isSet: true}
}

func (v NullableTranscriptionResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTranscriptionResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


