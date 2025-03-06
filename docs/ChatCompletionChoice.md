# ChatCompletionChoice

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Index** | Pointer to **int32** |  | [optional] 
**Message** | Pointer to [**ChatMessage**](ChatMessage.md) |  | [optional] 
**FinishReason** | Pointer to **string** |  | [optional] 

## Methods

### NewChatCompletionChoice

`func NewChatCompletionChoice() *ChatCompletionChoice`

NewChatCompletionChoice instantiates a new ChatCompletionChoice object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChatCompletionChoiceWithDefaults

`func NewChatCompletionChoiceWithDefaults() *ChatCompletionChoice`

NewChatCompletionChoiceWithDefaults instantiates a new ChatCompletionChoice object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIndex

`func (o *ChatCompletionChoice) GetIndex() int32`

GetIndex returns the Index field if non-nil, zero value otherwise.

### GetIndexOk

`func (o *ChatCompletionChoice) GetIndexOk() (*int32, bool)`

GetIndexOk returns a tuple with the Index field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndex

`func (o *ChatCompletionChoice) SetIndex(v int32)`

SetIndex sets Index field to given value.

### HasIndex

`func (o *ChatCompletionChoice) HasIndex() bool`

HasIndex returns a boolean if a field has been set.

### GetMessage

`func (o *ChatCompletionChoice) GetMessage() ChatMessage`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ChatCompletionChoice) GetMessageOk() (*ChatMessage, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ChatCompletionChoice) SetMessage(v ChatMessage)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *ChatCompletionChoice) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetFinishReason

`func (o *ChatCompletionChoice) GetFinishReason() string`

GetFinishReason returns the FinishReason field if non-nil, zero value otherwise.

### GetFinishReasonOk

`func (o *ChatCompletionChoice) GetFinishReasonOk() (*string, bool)`

GetFinishReasonOk returns a tuple with the FinishReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinishReason

`func (o *ChatCompletionChoice) SetFinishReason(v string)`

SetFinishReason sets FinishReason field to given value.

### HasFinishReason

`func (o *ChatCompletionChoice) HasFinishReason() bool`

HasFinishReason returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


