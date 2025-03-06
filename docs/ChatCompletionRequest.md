# ChatCompletionRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Messages** | [**[]ChatMessage**](ChatMessage.md) |  | 
**Model** | **string** |  | 
**FrequencyPenalty** | Pointer to **NullableFloat32** |  | [optional] [default to 0]
**FunctionCall** | Pointer to [**NullableChatCompletionRequestFunctionCall**](ChatCompletionRequestFunctionCall.md) |  | [optional] 
**Functions** | Pointer to **[]map[string]interface{}** |  | [optional] 
**LogitBias** | Pointer to **map[string]interface{}** |  | [optional] 
**MaxTokens** | Pointer to **NullableInt32** |  | [optional] 
**N** | Pointer to **NullableInt32** |  | [optional] [default to 1]
**PresencePenalty** | Pointer to **NullableFloat32** |  | [optional] [default to 0]
**Stop** | Pointer to [**NullableChatCompletionRequestStop**](ChatCompletionRequestStop.md) |  | [optional] 
**Stream** | Pointer to **NullableBool** |  | [optional] [default to false]
**Temperature** | Pointer to **NullableFloat32** |  | [optional] [default to 1]
**TopP** | Pointer to **NullableFloat32** |  | [optional] [default to 1]
**User** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewChatCompletionRequest

`func NewChatCompletionRequest(messages []ChatMessage, model string, ) *ChatCompletionRequest`

NewChatCompletionRequest instantiates a new ChatCompletionRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChatCompletionRequestWithDefaults

`func NewChatCompletionRequestWithDefaults() *ChatCompletionRequest`

NewChatCompletionRequestWithDefaults instantiates a new ChatCompletionRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessages

`func (o *ChatCompletionRequest) GetMessages() []ChatMessage`

GetMessages returns the Messages field if non-nil, zero value otherwise.

### GetMessagesOk

`func (o *ChatCompletionRequest) GetMessagesOk() (*[]ChatMessage, bool)`

GetMessagesOk returns a tuple with the Messages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessages

`func (o *ChatCompletionRequest) SetMessages(v []ChatMessage)`

SetMessages sets Messages field to given value.


### GetModel

`func (o *ChatCompletionRequest) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *ChatCompletionRequest) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *ChatCompletionRequest) SetModel(v string)`

SetModel sets Model field to given value.


### GetFrequencyPenalty

`func (o *ChatCompletionRequest) GetFrequencyPenalty() float32`

GetFrequencyPenalty returns the FrequencyPenalty field if non-nil, zero value otherwise.

### GetFrequencyPenaltyOk

`func (o *ChatCompletionRequest) GetFrequencyPenaltyOk() (*float32, bool)`

GetFrequencyPenaltyOk returns a tuple with the FrequencyPenalty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrequencyPenalty

`func (o *ChatCompletionRequest) SetFrequencyPenalty(v float32)`

SetFrequencyPenalty sets FrequencyPenalty field to given value.

### HasFrequencyPenalty

`func (o *ChatCompletionRequest) HasFrequencyPenalty() bool`

HasFrequencyPenalty returns a boolean if a field has been set.

### SetFrequencyPenaltyNil

`func (o *ChatCompletionRequest) SetFrequencyPenaltyNil(b bool)`

 SetFrequencyPenaltyNil sets the value for FrequencyPenalty to be an explicit nil

### UnsetFrequencyPenalty
`func (o *ChatCompletionRequest) UnsetFrequencyPenalty()`

UnsetFrequencyPenalty ensures that no value is present for FrequencyPenalty, not even an explicit nil
### GetFunctionCall

`func (o *ChatCompletionRequest) GetFunctionCall() ChatCompletionRequestFunctionCall`

GetFunctionCall returns the FunctionCall field if non-nil, zero value otherwise.

### GetFunctionCallOk

`func (o *ChatCompletionRequest) GetFunctionCallOk() (*ChatCompletionRequestFunctionCall, bool)`

GetFunctionCallOk returns a tuple with the FunctionCall field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunctionCall

`func (o *ChatCompletionRequest) SetFunctionCall(v ChatCompletionRequestFunctionCall)`

SetFunctionCall sets FunctionCall field to given value.

### HasFunctionCall

`func (o *ChatCompletionRequest) HasFunctionCall() bool`

HasFunctionCall returns a boolean if a field has been set.

### SetFunctionCallNil

`func (o *ChatCompletionRequest) SetFunctionCallNil(b bool)`

 SetFunctionCallNil sets the value for FunctionCall to be an explicit nil

### UnsetFunctionCall
`func (o *ChatCompletionRequest) UnsetFunctionCall()`

UnsetFunctionCall ensures that no value is present for FunctionCall, not even an explicit nil
### GetFunctions

`func (o *ChatCompletionRequest) GetFunctions() []map[string]interface{}`

GetFunctions returns the Functions field if non-nil, zero value otherwise.

### GetFunctionsOk

`func (o *ChatCompletionRequest) GetFunctionsOk() (*[]map[string]interface{}, bool)`

GetFunctionsOk returns a tuple with the Functions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunctions

`func (o *ChatCompletionRequest) SetFunctions(v []map[string]interface{})`

SetFunctions sets Functions field to given value.

### HasFunctions

`func (o *ChatCompletionRequest) HasFunctions() bool`

HasFunctions returns a boolean if a field has been set.

### SetFunctionsNil

`func (o *ChatCompletionRequest) SetFunctionsNil(b bool)`

 SetFunctionsNil sets the value for Functions to be an explicit nil

### UnsetFunctions
`func (o *ChatCompletionRequest) UnsetFunctions()`

UnsetFunctions ensures that no value is present for Functions, not even an explicit nil
### GetLogitBias

`func (o *ChatCompletionRequest) GetLogitBias() map[string]interface{}`

GetLogitBias returns the LogitBias field if non-nil, zero value otherwise.

### GetLogitBiasOk

`func (o *ChatCompletionRequest) GetLogitBiasOk() (*map[string]interface{}, bool)`

GetLogitBiasOk returns a tuple with the LogitBias field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogitBias

`func (o *ChatCompletionRequest) SetLogitBias(v map[string]interface{})`

SetLogitBias sets LogitBias field to given value.

### HasLogitBias

`func (o *ChatCompletionRequest) HasLogitBias() bool`

HasLogitBias returns a boolean if a field has been set.

### SetLogitBiasNil

`func (o *ChatCompletionRequest) SetLogitBiasNil(b bool)`

 SetLogitBiasNil sets the value for LogitBias to be an explicit nil

### UnsetLogitBias
`func (o *ChatCompletionRequest) UnsetLogitBias()`

UnsetLogitBias ensures that no value is present for LogitBias, not even an explicit nil
### GetMaxTokens

`func (o *ChatCompletionRequest) GetMaxTokens() int32`

GetMaxTokens returns the MaxTokens field if non-nil, zero value otherwise.

### GetMaxTokensOk

`func (o *ChatCompletionRequest) GetMaxTokensOk() (*int32, bool)`

GetMaxTokensOk returns a tuple with the MaxTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxTokens

`func (o *ChatCompletionRequest) SetMaxTokens(v int32)`

SetMaxTokens sets MaxTokens field to given value.

### HasMaxTokens

`func (o *ChatCompletionRequest) HasMaxTokens() bool`

HasMaxTokens returns a boolean if a field has been set.

### SetMaxTokensNil

`func (o *ChatCompletionRequest) SetMaxTokensNil(b bool)`

 SetMaxTokensNil sets the value for MaxTokens to be an explicit nil

### UnsetMaxTokens
`func (o *ChatCompletionRequest) UnsetMaxTokens()`

UnsetMaxTokens ensures that no value is present for MaxTokens, not even an explicit nil
### GetN

`func (o *ChatCompletionRequest) GetN() int32`

GetN returns the N field if non-nil, zero value otherwise.

### GetNOk

`func (o *ChatCompletionRequest) GetNOk() (*int32, bool)`

GetNOk returns a tuple with the N field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetN

`func (o *ChatCompletionRequest) SetN(v int32)`

SetN sets N field to given value.

### HasN

`func (o *ChatCompletionRequest) HasN() bool`

HasN returns a boolean if a field has been set.

### SetNNil

`func (o *ChatCompletionRequest) SetNNil(b bool)`

 SetNNil sets the value for N to be an explicit nil

### UnsetN
`func (o *ChatCompletionRequest) UnsetN()`

UnsetN ensures that no value is present for N, not even an explicit nil
### GetPresencePenalty

`func (o *ChatCompletionRequest) GetPresencePenalty() float32`

GetPresencePenalty returns the PresencePenalty field if non-nil, zero value otherwise.

### GetPresencePenaltyOk

`func (o *ChatCompletionRequest) GetPresencePenaltyOk() (*float32, bool)`

GetPresencePenaltyOk returns a tuple with the PresencePenalty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPresencePenalty

`func (o *ChatCompletionRequest) SetPresencePenalty(v float32)`

SetPresencePenalty sets PresencePenalty field to given value.

### HasPresencePenalty

`func (o *ChatCompletionRequest) HasPresencePenalty() bool`

HasPresencePenalty returns a boolean if a field has been set.

### SetPresencePenaltyNil

`func (o *ChatCompletionRequest) SetPresencePenaltyNil(b bool)`

 SetPresencePenaltyNil sets the value for PresencePenalty to be an explicit nil

### UnsetPresencePenalty
`func (o *ChatCompletionRequest) UnsetPresencePenalty()`

UnsetPresencePenalty ensures that no value is present for PresencePenalty, not even an explicit nil
### GetStop

`func (o *ChatCompletionRequest) GetStop() ChatCompletionRequestStop`

GetStop returns the Stop field if non-nil, zero value otherwise.

### GetStopOk

`func (o *ChatCompletionRequest) GetStopOk() (*ChatCompletionRequestStop, bool)`

GetStopOk returns a tuple with the Stop field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStop

`func (o *ChatCompletionRequest) SetStop(v ChatCompletionRequestStop)`

SetStop sets Stop field to given value.

### HasStop

`func (o *ChatCompletionRequest) HasStop() bool`

HasStop returns a boolean if a field has been set.

### SetStopNil

`func (o *ChatCompletionRequest) SetStopNil(b bool)`

 SetStopNil sets the value for Stop to be an explicit nil

### UnsetStop
`func (o *ChatCompletionRequest) UnsetStop()`

UnsetStop ensures that no value is present for Stop, not even an explicit nil
### GetStream

`func (o *ChatCompletionRequest) GetStream() bool`

GetStream returns the Stream field if non-nil, zero value otherwise.

### GetStreamOk

`func (o *ChatCompletionRequest) GetStreamOk() (*bool, bool)`

GetStreamOk returns a tuple with the Stream field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStream

`func (o *ChatCompletionRequest) SetStream(v bool)`

SetStream sets Stream field to given value.

### HasStream

`func (o *ChatCompletionRequest) HasStream() bool`

HasStream returns a boolean if a field has been set.

### SetStreamNil

`func (o *ChatCompletionRequest) SetStreamNil(b bool)`

 SetStreamNil sets the value for Stream to be an explicit nil

### UnsetStream
`func (o *ChatCompletionRequest) UnsetStream()`

UnsetStream ensures that no value is present for Stream, not even an explicit nil
### GetTemperature

`func (o *ChatCompletionRequest) GetTemperature() float32`

GetTemperature returns the Temperature field if non-nil, zero value otherwise.

### GetTemperatureOk

`func (o *ChatCompletionRequest) GetTemperatureOk() (*float32, bool)`

GetTemperatureOk returns a tuple with the Temperature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemperature

`func (o *ChatCompletionRequest) SetTemperature(v float32)`

SetTemperature sets Temperature field to given value.

### HasTemperature

`func (o *ChatCompletionRequest) HasTemperature() bool`

HasTemperature returns a boolean if a field has been set.

### SetTemperatureNil

`func (o *ChatCompletionRequest) SetTemperatureNil(b bool)`

 SetTemperatureNil sets the value for Temperature to be an explicit nil

### UnsetTemperature
`func (o *ChatCompletionRequest) UnsetTemperature()`

UnsetTemperature ensures that no value is present for Temperature, not even an explicit nil
### GetTopP

`func (o *ChatCompletionRequest) GetTopP() float32`

GetTopP returns the TopP field if non-nil, zero value otherwise.

### GetTopPOk

`func (o *ChatCompletionRequest) GetTopPOk() (*float32, bool)`

GetTopPOk returns a tuple with the TopP field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopP

`func (o *ChatCompletionRequest) SetTopP(v float32)`

SetTopP sets TopP field to given value.

### HasTopP

`func (o *ChatCompletionRequest) HasTopP() bool`

HasTopP returns a boolean if a field has been set.

### SetTopPNil

`func (o *ChatCompletionRequest) SetTopPNil(b bool)`

 SetTopPNil sets the value for TopP to be an explicit nil

### UnsetTopP
`func (o *ChatCompletionRequest) UnsetTopP()`

UnsetTopP ensures that no value is present for TopP, not even an explicit nil
### GetUser

`func (o *ChatCompletionRequest) GetUser() string`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *ChatCompletionRequest) GetUserOk() (*string, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *ChatCompletionRequest) SetUser(v string)`

SetUser sets User field to given value.

### HasUser

`func (o *ChatCompletionRequest) HasUser() bool`

HasUser returns a boolean if a field has been set.

### SetUserNil

`func (o *ChatCompletionRequest) SetUserNil(b bool)`

 SetUserNil sets the value for User to be an explicit nil

### UnsetUser
`func (o *ChatCompletionRequest) UnsetUser()`

UnsetUser ensures that no value is present for User, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


