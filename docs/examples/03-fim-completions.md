## Chat Operations (Fill-In-the-Middle)

* [Create Chat Request](#create-chat-request)

---

### Create Chat Request
[API Documentation](https://api-docs.deepseek.com/api/create-completion)

Create a `BetaCompletionsRequest` object and pass it to `CreateBetaCompletions`.  
Returns a `BetaCompletionsResponse`.  
```go
import (
    "context"
    
    "github.com/rvinnie/deepseek-sdk-go/deepseek"
)

func main() {
    // Create client
    client, err := deepseek.NewClient("<API_TOKEN>")
    if err != nil {
        // Handle error
    }

    // Create request object
    req := &deepseek.BetaCompletionsRequest{
        Model:            deepseek.DeepSeekChat,
        Prompt:           "Once upon a time, ",
        Echo:             false,
        Logprobs:         0,
        MaxTokens:        1024,
        FrequencyPenalty: 0,
        PresencePenalty:  0,
        Stop:             nil,
        Stream:           false,
        StreamOptions:    nil,
        Temperature:      1,
        TopP:             1,
    }

    // Get response
    resp, err := client.CreateBetaCompletions(context.Background(), req)
    if err != nil {
        // Handle error
    }
}

```
