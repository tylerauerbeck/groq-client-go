/*
Groq API

API for interacting with Groq's language models and audio transcription/translation services

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the ChatCompletionRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ChatCompletionRequest{}

// ChatCompletionRequest struct for ChatCompletionRequest
type ChatCompletionRequest struct {
	Messages []ChatMessage `json:"messages"`
	Model string `json:"model"`
	FrequencyPenalty NullableFloat32 `json:"frequency_penalty,omitempty"`
	FunctionCall NullableChatCompletionRequestFunctionCall `json:"function_call,omitempty"`
	Functions []map[string]interface{} `json:"functions,omitempty"`
	LogitBias map[string]interface{} `json:"logit_bias,omitempty"`
	MaxTokens NullableInt32 `json:"max_tokens,omitempty"`
	N NullableInt32 `json:"n,omitempty"`
	PresencePenalty NullableFloat32 `json:"presence_penalty,omitempty"`
	Stop NullableChatCompletionRequestStop `json:"stop,omitempty"`
	Stream NullableBool `json:"stream,omitempty"`
	Temperature NullableFloat32 `json:"temperature,omitempty"`
	TopP NullableFloat32 `json:"top_p,omitempty"`
	User NullableString `json:"user,omitempty"`
}

type _ChatCompletionRequest ChatCompletionRequest

// NewChatCompletionRequest instantiates a new ChatCompletionRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewChatCompletionRequest(messages []ChatMessage, model string) *ChatCompletionRequest {
	this := ChatCompletionRequest{}
	this.Messages = messages
	this.Model = model
	var frequencyPenalty float32 = 0
	this.FrequencyPenalty = *NewNullableFloat32(&frequencyPenalty)
	var n int32 = 1
	this.N = *NewNullableInt32(&n)
	var presencePenalty float32 = 0
	this.PresencePenalty = *NewNullableFloat32(&presencePenalty)
	var stream bool = false
	this.Stream = *NewNullableBool(&stream)
	var temperature float32 = 1
	this.Temperature = *NewNullableFloat32(&temperature)
	var topP float32 = 1
	this.TopP = *NewNullableFloat32(&topP)
	return &this
}

// NewChatCompletionRequestWithDefaults instantiates a new ChatCompletionRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewChatCompletionRequestWithDefaults() *ChatCompletionRequest {
	this := ChatCompletionRequest{}
	var frequencyPenalty float32 = 0
	this.FrequencyPenalty = *NewNullableFloat32(&frequencyPenalty)
	var n int32 = 1
	this.N = *NewNullableInt32(&n)
	var presencePenalty float32 = 0
	this.PresencePenalty = *NewNullableFloat32(&presencePenalty)
	var stream bool = false
	this.Stream = *NewNullableBool(&stream)
	var temperature float32 = 1
	this.Temperature = *NewNullableFloat32(&temperature)
	var topP float32 = 1
	this.TopP = *NewNullableFloat32(&topP)
	return &this
}

// GetMessages returns the Messages field value
func (o *ChatCompletionRequest) GetMessages() []ChatMessage {
	if o == nil {
		var ret []ChatMessage
		return ret
	}

	return o.Messages
}

// GetMessagesOk returns a tuple with the Messages field value
// and a boolean to check if the value has been set.
func (o *ChatCompletionRequest) GetMessagesOk() ([]ChatMessage, bool) {
	if o == nil {
		return nil, false
	}
	return o.Messages, true
}

// SetMessages sets field value
func (o *ChatCompletionRequest) SetMessages(v []ChatMessage) {
	o.Messages = v
}

// GetModel returns the Model field value
func (o *ChatCompletionRequest) GetModel() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Model
}

// GetModelOk returns a tuple with the Model field value
// and a boolean to check if the value has been set.
func (o *ChatCompletionRequest) GetModelOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Model, true
}

// SetModel sets field value
func (o *ChatCompletionRequest) SetModel(v string) {
	o.Model = v
}

// GetFrequencyPenalty returns the FrequencyPenalty field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ChatCompletionRequest) GetFrequencyPenalty() float32 {
	if o == nil || IsNil(o.FrequencyPenalty.Get()) {
		var ret float32
		return ret
	}
	return *o.FrequencyPenalty.Get()
}

// GetFrequencyPenaltyOk returns a tuple with the FrequencyPenalty field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ChatCompletionRequest) GetFrequencyPenaltyOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return o.FrequencyPenalty.Get(), o.FrequencyPenalty.IsSet()
}

// HasFrequencyPenalty returns a boolean if a field has been set.
func (o *ChatCompletionRequest) HasFrequencyPenalty() bool {
	if o != nil && o.FrequencyPenalty.IsSet() {
		return true
	}

	return false
}

// SetFrequencyPenalty gets a reference to the given NullableFloat32 and assigns it to the FrequencyPenalty field.
func (o *ChatCompletionRequest) SetFrequencyPenalty(v float32) {
	o.FrequencyPenalty.Set(&v)
}
// SetFrequencyPenaltyNil sets the value for FrequencyPenalty to be an explicit nil
func (o *ChatCompletionRequest) SetFrequencyPenaltyNil() {
	o.FrequencyPenalty.Set(nil)
}

// UnsetFrequencyPenalty ensures that no value is present for FrequencyPenalty, not even an explicit nil
func (o *ChatCompletionRequest) UnsetFrequencyPenalty() {
	o.FrequencyPenalty.Unset()
}

// GetFunctionCall returns the FunctionCall field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ChatCompletionRequest) GetFunctionCall() ChatCompletionRequestFunctionCall {
	if o == nil || IsNil(o.FunctionCall.Get()) {
		var ret ChatCompletionRequestFunctionCall
		return ret
	}
	return *o.FunctionCall.Get()
}

// GetFunctionCallOk returns a tuple with the FunctionCall field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ChatCompletionRequest) GetFunctionCallOk() (*ChatCompletionRequestFunctionCall, bool) {
	if o == nil {
		return nil, false
	}
	return o.FunctionCall.Get(), o.FunctionCall.IsSet()
}

// HasFunctionCall returns a boolean if a field has been set.
func (o *ChatCompletionRequest) HasFunctionCall() bool {
	if o != nil && o.FunctionCall.IsSet() {
		return true
	}

	return false
}

// SetFunctionCall gets a reference to the given NullableChatCompletionRequestFunctionCall and assigns it to the FunctionCall field.
func (o *ChatCompletionRequest) SetFunctionCall(v ChatCompletionRequestFunctionCall) {
	o.FunctionCall.Set(&v)
}
// SetFunctionCallNil sets the value for FunctionCall to be an explicit nil
func (o *ChatCompletionRequest) SetFunctionCallNil() {
	o.FunctionCall.Set(nil)
}

// UnsetFunctionCall ensures that no value is present for FunctionCall, not even an explicit nil
func (o *ChatCompletionRequest) UnsetFunctionCall() {
	o.FunctionCall.Unset()
}

// GetFunctions returns the Functions field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ChatCompletionRequest) GetFunctions() []map[string]interface{} {
	if o == nil {
		var ret []map[string]interface{}
		return ret
	}
	return o.Functions
}

// GetFunctionsOk returns a tuple with the Functions field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ChatCompletionRequest) GetFunctionsOk() ([]map[string]interface{}, bool) {
	if o == nil || IsNil(o.Functions) {
		return nil, false
	}
	return o.Functions, true
}

// HasFunctions returns a boolean if a field has been set.
func (o *ChatCompletionRequest) HasFunctions() bool {
	if o != nil && !IsNil(o.Functions) {
		return true
	}

	return false
}

// SetFunctions gets a reference to the given []map[string]interface{} and assigns it to the Functions field.
func (o *ChatCompletionRequest) SetFunctions(v []map[string]interface{}) {
	o.Functions = v
}

// GetLogitBias returns the LogitBias field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ChatCompletionRequest) GetLogitBias() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.LogitBias
}

// GetLogitBiasOk returns a tuple with the LogitBias field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ChatCompletionRequest) GetLogitBiasOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.LogitBias) {
		return map[string]interface{}{}, false
	}
	return o.LogitBias, true
}

// HasLogitBias returns a boolean if a field has been set.
func (o *ChatCompletionRequest) HasLogitBias() bool {
	if o != nil && !IsNil(o.LogitBias) {
		return true
	}

	return false
}

// SetLogitBias gets a reference to the given map[string]interface{} and assigns it to the LogitBias field.
func (o *ChatCompletionRequest) SetLogitBias(v map[string]interface{}) {
	o.LogitBias = v
}

// GetMaxTokens returns the MaxTokens field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ChatCompletionRequest) GetMaxTokens() int32 {
	if o == nil || IsNil(o.MaxTokens.Get()) {
		var ret int32
		return ret
	}
	return *o.MaxTokens.Get()
}

// GetMaxTokensOk returns a tuple with the MaxTokens field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ChatCompletionRequest) GetMaxTokensOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.MaxTokens.Get(), o.MaxTokens.IsSet()
}

// HasMaxTokens returns a boolean if a field has been set.
func (o *ChatCompletionRequest) HasMaxTokens() bool {
	if o != nil && o.MaxTokens.IsSet() {
		return true
	}

	return false
}

// SetMaxTokens gets a reference to the given NullableInt32 and assigns it to the MaxTokens field.
func (o *ChatCompletionRequest) SetMaxTokens(v int32) {
	o.MaxTokens.Set(&v)
}
// SetMaxTokensNil sets the value for MaxTokens to be an explicit nil
func (o *ChatCompletionRequest) SetMaxTokensNil() {
	o.MaxTokens.Set(nil)
}

// UnsetMaxTokens ensures that no value is present for MaxTokens, not even an explicit nil
func (o *ChatCompletionRequest) UnsetMaxTokens() {
	o.MaxTokens.Unset()
}

// GetN returns the N field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ChatCompletionRequest) GetN() int32 {
	if o == nil || IsNil(o.N.Get()) {
		var ret int32
		return ret
	}
	return *o.N.Get()
}

// GetNOk returns a tuple with the N field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ChatCompletionRequest) GetNOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.N.Get(), o.N.IsSet()
}

// HasN returns a boolean if a field has been set.
func (o *ChatCompletionRequest) HasN() bool {
	if o != nil && o.N.IsSet() {
		return true
	}

	return false
}

// SetN gets a reference to the given NullableInt32 and assigns it to the N field.
func (o *ChatCompletionRequest) SetN(v int32) {
	o.N.Set(&v)
}
// SetNNil sets the value for N to be an explicit nil
func (o *ChatCompletionRequest) SetNNil() {
	o.N.Set(nil)
}

// UnsetN ensures that no value is present for N, not even an explicit nil
func (o *ChatCompletionRequest) UnsetN() {
	o.N.Unset()
}

// GetPresencePenalty returns the PresencePenalty field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ChatCompletionRequest) GetPresencePenalty() float32 {
	if o == nil || IsNil(o.PresencePenalty.Get()) {
		var ret float32
		return ret
	}
	return *o.PresencePenalty.Get()
}

// GetPresencePenaltyOk returns a tuple with the PresencePenalty field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ChatCompletionRequest) GetPresencePenaltyOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return o.PresencePenalty.Get(), o.PresencePenalty.IsSet()
}

// HasPresencePenalty returns a boolean if a field has been set.
func (o *ChatCompletionRequest) HasPresencePenalty() bool {
	if o != nil && o.PresencePenalty.IsSet() {
		return true
	}

	return false
}

// SetPresencePenalty gets a reference to the given NullableFloat32 and assigns it to the PresencePenalty field.
func (o *ChatCompletionRequest) SetPresencePenalty(v float32) {
	o.PresencePenalty.Set(&v)
}
// SetPresencePenaltyNil sets the value for PresencePenalty to be an explicit nil
func (o *ChatCompletionRequest) SetPresencePenaltyNil() {
	o.PresencePenalty.Set(nil)
}

// UnsetPresencePenalty ensures that no value is present for PresencePenalty, not even an explicit nil
func (o *ChatCompletionRequest) UnsetPresencePenalty() {
	o.PresencePenalty.Unset()
}

// GetStop returns the Stop field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ChatCompletionRequest) GetStop() ChatCompletionRequestStop {
	if o == nil || IsNil(o.Stop.Get()) {
		var ret ChatCompletionRequestStop
		return ret
	}
	return *o.Stop.Get()
}

// GetStopOk returns a tuple with the Stop field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ChatCompletionRequest) GetStopOk() (*ChatCompletionRequestStop, bool) {
	if o == nil {
		return nil, false
	}
	return o.Stop.Get(), o.Stop.IsSet()
}

// HasStop returns a boolean if a field has been set.
func (o *ChatCompletionRequest) HasStop() bool {
	if o != nil && o.Stop.IsSet() {
		return true
	}

	return false
}

// SetStop gets a reference to the given NullableChatCompletionRequestStop and assigns it to the Stop field.
func (o *ChatCompletionRequest) SetStop(v ChatCompletionRequestStop) {
	o.Stop.Set(&v)
}
// SetStopNil sets the value for Stop to be an explicit nil
func (o *ChatCompletionRequest) SetStopNil() {
	o.Stop.Set(nil)
}

// UnsetStop ensures that no value is present for Stop, not even an explicit nil
func (o *ChatCompletionRequest) UnsetStop() {
	o.Stop.Unset()
}

// GetStream returns the Stream field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ChatCompletionRequest) GetStream() bool {
	if o == nil || IsNil(o.Stream.Get()) {
		var ret bool
		return ret
	}
	return *o.Stream.Get()
}

// GetStreamOk returns a tuple with the Stream field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ChatCompletionRequest) GetStreamOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.Stream.Get(), o.Stream.IsSet()
}

// HasStream returns a boolean if a field has been set.
func (o *ChatCompletionRequest) HasStream() bool {
	if o != nil && o.Stream.IsSet() {
		return true
	}

	return false
}

// SetStream gets a reference to the given NullableBool and assigns it to the Stream field.
func (o *ChatCompletionRequest) SetStream(v bool) {
	o.Stream.Set(&v)
}
// SetStreamNil sets the value for Stream to be an explicit nil
func (o *ChatCompletionRequest) SetStreamNil() {
	o.Stream.Set(nil)
}

// UnsetStream ensures that no value is present for Stream, not even an explicit nil
func (o *ChatCompletionRequest) UnsetStream() {
	o.Stream.Unset()
}

// GetTemperature returns the Temperature field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ChatCompletionRequest) GetTemperature() float32 {
	if o == nil || IsNil(o.Temperature.Get()) {
		var ret float32
		return ret
	}
	return *o.Temperature.Get()
}

// GetTemperatureOk returns a tuple with the Temperature field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ChatCompletionRequest) GetTemperatureOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Temperature.Get(), o.Temperature.IsSet()
}

// HasTemperature returns a boolean if a field has been set.
func (o *ChatCompletionRequest) HasTemperature() bool {
	if o != nil && o.Temperature.IsSet() {
		return true
	}

	return false
}

// SetTemperature gets a reference to the given NullableFloat32 and assigns it to the Temperature field.
func (o *ChatCompletionRequest) SetTemperature(v float32) {
	o.Temperature.Set(&v)
}
// SetTemperatureNil sets the value for Temperature to be an explicit nil
func (o *ChatCompletionRequest) SetTemperatureNil() {
	o.Temperature.Set(nil)
}

// UnsetTemperature ensures that no value is present for Temperature, not even an explicit nil
func (o *ChatCompletionRequest) UnsetTemperature() {
	o.Temperature.Unset()
}

// GetTopP returns the TopP field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ChatCompletionRequest) GetTopP() float32 {
	if o == nil || IsNil(o.TopP.Get()) {
		var ret float32
		return ret
	}
	return *o.TopP.Get()
}

// GetTopPOk returns a tuple with the TopP field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ChatCompletionRequest) GetTopPOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return o.TopP.Get(), o.TopP.IsSet()
}

// HasTopP returns a boolean if a field has been set.
func (o *ChatCompletionRequest) HasTopP() bool {
	if o != nil && o.TopP.IsSet() {
		return true
	}

	return false
}

// SetTopP gets a reference to the given NullableFloat32 and assigns it to the TopP field.
func (o *ChatCompletionRequest) SetTopP(v float32) {
	o.TopP.Set(&v)
}
// SetTopPNil sets the value for TopP to be an explicit nil
func (o *ChatCompletionRequest) SetTopPNil() {
	o.TopP.Set(nil)
}

// UnsetTopP ensures that no value is present for TopP, not even an explicit nil
func (o *ChatCompletionRequest) UnsetTopP() {
	o.TopP.Unset()
}

// GetUser returns the User field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ChatCompletionRequest) GetUser() string {
	if o == nil || IsNil(o.User.Get()) {
		var ret string
		return ret
	}
	return *o.User.Get()
}

// GetUserOk returns a tuple with the User field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ChatCompletionRequest) GetUserOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.User.Get(), o.User.IsSet()
}

// HasUser returns a boolean if a field has been set.
func (o *ChatCompletionRequest) HasUser() bool {
	if o != nil && o.User.IsSet() {
		return true
	}

	return false
}

// SetUser gets a reference to the given NullableString and assigns it to the User field.
func (o *ChatCompletionRequest) SetUser(v string) {
	o.User.Set(&v)
}
// SetUserNil sets the value for User to be an explicit nil
func (o *ChatCompletionRequest) SetUserNil() {
	o.User.Set(nil)
}

// UnsetUser ensures that no value is present for User, not even an explicit nil
func (o *ChatCompletionRequest) UnsetUser() {
	o.User.Unset()
}

func (o ChatCompletionRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ChatCompletionRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["messages"] = o.Messages
	toSerialize["model"] = o.Model
	if o.FrequencyPenalty.IsSet() {
		toSerialize["frequency_penalty"] = o.FrequencyPenalty.Get()
	}
	if o.FunctionCall.IsSet() {
		toSerialize["function_call"] = o.FunctionCall.Get()
	}
	if o.Functions != nil {
		toSerialize["functions"] = o.Functions
	}
	if o.LogitBias != nil {
		toSerialize["logit_bias"] = o.LogitBias
	}
	if o.MaxTokens.IsSet() {
		toSerialize["max_tokens"] = o.MaxTokens.Get()
	}
	if o.N.IsSet() {
		toSerialize["n"] = o.N.Get()
	}
	if o.PresencePenalty.IsSet() {
		toSerialize["presence_penalty"] = o.PresencePenalty.Get()
	}
	if o.Stop.IsSet() {
		toSerialize["stop"] = o.Stop.Get()
	}
	if o.Stream.IsSet() {
		toSerialize["stream"] = o.Stream.Get()
	}
	if o.Temperature.IsSet() {
		toSerialize["temperature"] = o.Temperature.Get()
	}
	if o.TopP.IsSet() {
		toSerialize["top_p"] = o.TopP.Get()
	}
	if o.User.IsSet() {
		toSerialize["user"] = o.User.Get()
	}
	return toSerialize, nil
}

func (o *ChatCompletionRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"messages",
		"model",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varChatCompletionRequest := _ChatCompletionRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varChatCompletionRequest)

	if err != nil {
		return err
	}

	*o = ChatCompletionRequest(varChatCompletionRequest)

	return err
}

type NullableChatCompletionRequest struct {
	value *ChatCompletionRequest
	isSet bool
}

func (v NullableChatCompletionRequest) Get() *ChatCompletionRequest {
	return v.value
}

func (v *NullableChatCompletionRequest) Set(val *ChatCompletionRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableChatCompletionRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableChatCompletionRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableChatCompletionRequest(val *ChatCompletionRequest) *NullableChatCompletionRequest {
	return &NullableChatCompletionRequest{value: val, isSet: true}
}

func (v NullableChatCompletionRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableChatCompletionRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


