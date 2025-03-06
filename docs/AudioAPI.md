# \AudioAPI

All URIs are relative to *https://api.groq.com/openai/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateTranscription**](AudioAPI.md#CreateTranscription) | **Post** /audio/transcriptions | Create transcription
[**CreateTranslation**](AudioAPI.md#CreateTranslation) | **Post** /audio/translations | Create translation



## CreateTranscription

> TranscriptionResponse CreateTranscription(ctx).File(file).Model(model).Language(language).Prompt(prompt).ResponseFormat(responseFormat).Temperature(temperature).Execute()

Create transcription



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/tylerauerbeck/groq-client-go"
)

func main() {
	file := os.NewFile(1234, "some_file") // *os.File | 
	model := "model_example" // string | 
	language := "language_example" // string |  (optional)
	prompt := "prompt_example" // string |  (optional)
	responseFormat := "responseFormat_example" // string |  (optional) (default to "json")
	temperature := float32(8.14) // float32 |  (optional) (default to 0)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AudioAPI.CreateTranscription(context.Background()).File(file).Model(model).Language(language).Prompt(prompt).ResponseFormat(responseFormat).Temperature(temperature).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AudioAPI.CreateTranscription``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateTranscription`: TranscriptionResponse
	fmt.Fprintf(os.Stdout, "Response from `AudioAPI.CreateTranscription`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateTranscriptionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **file** | ***os.File** |  | 
 **model** | **string** |  | 
 **language** | **string** |  | 
 **prompt** | **string** |  | 
 **responseFormat** | **string** |  | [default to &quot;json&quot;]
 **temperature** | **float32** |  | [default to 0]

### Return type

[**TranscriptionResponse**](TranscriptionResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateTranslation

> TranslationResponse CreateTranslation(ctx).File(file).Model(model).Prompt(prompt).ResponseFormat(responseFormat).Temperature(temperature).Execute()

Create translation



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/tylerauerbeck/groq-client-go"
)

func main() {
	file := os.NewFile(1234, "some_file") // *os.File | 
	model := "model_example" // string | 
	prompt := "prompt_example" // string |  (optional)
	responseFormat := "responseFormat_example" // string |  (optional) (default to "json")
	temperature := float32(8.14) // float32 |  (optional) (default to 0)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AudioAPI.CreateTranslation(context.Background()).File(file).Model(model).Prompt(prompt).ResponseFormat(responseFormat).Temperature(temperature).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AudioAPI.CreateTranslation``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateTranslation`: TranslationResponse
	fmt.Fprintf(os.Stdout, "Response from `AudioAPI.CreateTranslation`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateTranslationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **file** | ***os.File** |  | 
 **model** | **string** |  | 
 **prompt** | **string** |  | 
 **responseFormat** | **string** |  | [default to &quot;json&quot;]
 **temperature** | **float32** |  | [default to 0]

### Return type

[**TranslationResponse**](TranslationResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

