# \ChatAPI

All URIs are relative to *https://api.groq.com/openai/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateChatCompletion**](ChatAPI.md#CreateChatCompletion) | **Post** /chat/completions | Create chat completion



## CreateChatCompletion

> ChatCompletionResponse CreateChatCompletion(ctx).ChatCompletionRequest(chatCompletionRequest).Execute()

Create chat completion



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	chatCompletionRequest := *openapiclient.NewChatCompletionRequest([]openapiclient.ChatMessage{*openapiclient.NewChatMessage("Role_example", "Content_example")}, "Model_example") // ChatCompletionRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ChatAPI.CreateChatCompletion(context.Background()).ChatCompletionRequest(chatCompletionRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChatAPI.CreateChatCompletion``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateChatCompletion`: ChatCompletionResponse
	fmt.Fprintf(os.Stdout, "Response from `ChatAPI.CreateChatCompletion`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateChatCompletionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **chatCompletionRequest** | [**ChatCompletionRequest**](ChatCompletionRequest.md) |  | 

### Return type

[**ChatCompletionResponse**](ChatCompletionResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

