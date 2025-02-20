## Chat Operations

* [Create Chat Request](#create-chat-request)
* [Streaming Content](#streaming-content)

---

### Create Chat Request
[API Documentation](https://api-docs.deepseek.com/api/create-chat-completion)

Create a `ChatCompletionsRequest` object and pass it to the `CreateChatCompletions` method.  
Returns a `ChatCompletionsResponse` object.

```go
import (
    "context"

    "github.com/rvinnie/deepseek-sdk-go/deepseek"
)

func main() {
    // Create client
    client, err := deepseek.NewClient("API_TOKEN")
    if err != nil {
        // Handle error
    }

    // Create request object
    req := &deepseek.ChatCompletionsRequest{
        Messages: []deepseek.ChatCompletionsMessage{
            {
                Role:    deepseek.ChatMessageRoleSystem,
                Content: "You are a helpful assistant",
            },
            {
                Role:    deepseek.ChatMessageRoleUser,
                Content: "Hi",
            },
        },
        Model:            deepseek.DeepSeekChat,
        FrequencyPenalty: 0,
        PresencePenalty:  0,
        ResponseFormat: &deepseek.ResponseFormat{
            Type: "text",
        },
        Stop:          nil,
        Stream:        false,
        StreamOptions: nil,
        Temperature:   1,
        TopP:          1,
        Tools:         nil,
        ToolChoice: deepseek.ToolChoice{
            Type: "none",
        },
        LogProbs:    false,
        TopLogProbs: nil,
    }

    // Get chat response
    resp, err := client.CreateChatCompletions(context.Background(), req)
    if err != nil {
        // Handle error
    }
}
```

---

### Streaming Content

[API Documentation](https://api-docs.deepseek.com/api/create-chat-completion)

For streaming requests, use `CreateChatCompletionsWithStream` method. Returns an `http.Response` object.  
Use `bufio.Reader` with `ReadStream` to process the data stream.  

```go
import (
    "bufio"
    "context"
    "errors"
    "io"
    
    "github.com/rvinnie/deepseek-sdk-go/deepseek"
)

func main() {
    // Create client
    client, err := deepseek.NewClient("<API_TOKEN>")
    if err != nil {
        // Handle error
    }

    // Create request object
    req := &deepseek.ChatCompletionsRequest{
        Messages: []deepseek.ChatCompletionsMessage{
            {
                Role:    deepseek.ChatMessageRoleSystem,
                Content: "You are a helpful assistant",
            },
            {
                Role:    deepseek.ChatMessageRoleUser,
                Content: "Hi",
            },
        },
        Model:            deepseek.DeepSeekChat,
        FrequencyPenalty: 0,
        PresencePenalty:  0,
        ResponseFormat: &deepseek.ResponseFormat{
            Type: "text",
        },
        Stop:          nil,
        Stream:        true,
        StreamOptions: nil,
        Temperature:   1,
        TopP:          1,
        Tools:         nil,
        ToolChoice: deepseek.ToolChoice{
            Type: "none",
        },
        LogProbs:    false,
        TopLogProbs: nil,
    }

    // Get HTTP response
    resp, err := client.CreateChatCompletionsWithStream(context.Background(), req)
    if err != nil {
        // Handle error
    }
    defer func() {
        _ = resp.Body.Close()
    }()

    // Create stream reader
    reader := bufio.NewReader(resp.Body)

    // Process stream
    var full string
    for {
        part, err := deepseek.ReadStream(reader)
        if errors.Is(err, io.EOF) {
            // End of stream
            break
        }
        if err != nil {
            // Handle error
            break
        }
        for _, choice := range part.Choices {
            // Aggregate chunks
            full += choice.Delta.Content
        }
    }
}

```
