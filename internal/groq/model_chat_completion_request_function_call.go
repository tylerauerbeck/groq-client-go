/*
Groq API

API for interacting with Groq's language models and audio transcription/translation services

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"fmt"
	"gopkg.in/validator.v2"
)

// ChatCompletionRequestFunctionCall - struct for ChatCompletionRequestFunctionCall
type ChatCompletionRequestFunctionCall struct {
	MapmapOfStringAny *map[string]interface{}
	String *string
}

// map[string]interface{}AsChatCompletionRequestFunctionCall is a convenience function that returns map[string]interface{} wrapped in ChatCompletionRequestFunctionCall
func MapmapOfStringAnyAsChatCompletionRequestFunctionCall(v *map[string]interface{}) ChatCompletionRequestFunctionCall {
	return ChatCompletionRequestFunctionCall{
		MapmapOfStringAny: v,
	}
}

// stringAsChatCompletionRequestFunctionCall is a convenience function that returns string wrapped in ChatCompletionRequestFunctionCall
func StringAsChatCompletionRequestFunctionCall(v *string) ChatCompletionRequestFunctionCall {
	return ChatCompletionRequestFunctionCall{
		String: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *ChatCompletionRequestFunctionCall) UnmarshalJSON(data []byte) error {
	var err error
	// this object is nullable so check if the payload is null or empty string
	if string(data) == "" || string(data) == "{}" {
		return nil
	}

	match := 0
	// try to unmarshal data into MapmapOfStringAny
	err = newStrictDecoder(data).Decode(&dst.MapmapOfStringAny)
	if err == nil {
		jsonMapmapOfStringAny, _ := json.Marshal(dst.MapmapOfStringAny)
		if string(jsonMapmapOfStringAny) == "{}" { // empty struct
			dst.MapmapOfStringAny = nil
		} else {
			if err = validator.Validate(dst.MapmapOfStringAny); err != nil {
				dst.MapmapOfStringAny = nil
			} else {
				match++
			}
		}
	} else {
		dst.MapmapOfStringAny = nil
	}

	// try to unmarshal data into String
	err = newStrictDecoder(data).Decode(&dst.String)
	if err == nil {
		jsonString, _ := json.Marshal(dst.String)
		if string(jsonString) == "{}" { // empty struct
			dst.String = nil
		} else {
			if err = validator.Validate(dst.String); err != nil {
				dst.String = nil
			} else {
				match++
			}
		}
	} else {
		dst.String = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.MapmapOfStringAny = nil
		dst.String = nil

		return fmt.Errorf("data matches more than one schema in oneOf(ChatCompletionRequestFunctionCall)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(ChatCompletionRequestFunctionCall)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src ChatCompletionRequestFunctionCall) MarshalJSON() ([]byte, error) {
	if src.MapmapOfStringAny != nil {
		return json.Marshal(&src.MapmapOfStringAny)
	}

	if src.String != nil {
		return json.Marshal(&src.String)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *ChatCompletionRequestFunctionCall) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.MapmapOfStringAny != nil {
		return obj.MapmapOfStringAny
	}

	if obj.String != nil {
		return obj.String
	}

	// all schemas are nil
	return nil
}

// Get the actual instance value
func (obj ChatCompletionRequestFunctionCall) GetActualInstanceValue() (interface{}) {
	if obj.MapmapOfStringAny != nil {
		return *obj.MapmapOfStringAny
	}

	if obj.String != nil {
		return *obj.String
	}

	// all schemas are nil
	return nil
}

type NullableChatCompletionRequestFunctionCall struct {
	value *ChatCompletionRequestFunctionCall
	isSet bool
}

func (v NullableChatCompletionRequestFunctionCall) Get() *ChatCompletionRequestFunctionCall {
	return v.value
}

func (v *NullableChatCompletionRequestFunctionCall) Set(val *ChatCompletionRequestFunctionCall) {
	v.value = val
	v.isSet = true
}

func (v NullableChatCompletionRequestFunctionCall) IsSet() bool {
	return v.isSet
}

func (v *NullableChatCompletionRequestFunctionCall) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableChatCompletionRequestFunctionCall(val *ChatCompletionRequestFunctionCall) *NullableChatCompletionRequestFunctionCall {
	return &NullableChatCompletionRequestFunctionCall{value: val, isSet: true}
}

func (v NullableChatCompletionRequestFunctionCall) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableChatCompletionRequestFunctionCall) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


