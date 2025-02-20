## 聊天操作

* [创建聊天请求](#创建聊天请求)
* [流式传输内容](#流式传输内容)

---

### 创建聊天请求
[API 文档](https://api-docs.deepseek.com/api/create-chat-completion)

创建 `ChatCompletionsRequest` 对象并传递给 `CreateChatCompletions` 方法。  
返回 `ChatCompletionsResponse` 对象。  

```go
import (
    "context"

    "github.com/rvinnie/deepseek-sdk-go/deepseek"
)

func main() {
    // 创建客户端
    client, err := deepseek.NewClient("API_TOKEN")
    if err != nil {
        // 处理错误
    }

	// 创建请求对象
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

    // 获取聊天响应
    resp, err := client.CreateChatCompletions(context.Background(), req)
    if err != nil {
        // 处理错误
    }
}
```

---

### 流式传输内容

[API 文档](https://api-docs.deepseek.com/api/create-chat-completion)

使用 `CreateChatCompletionsWithStream` 方法创建流式请求。返回 `http.Response` 对象。  
使用 `bufio.Reader` 和 `ReadStream` 处理数据流。  
```go
import (
    "bufio"
    "context"
    "errors"
    "io"
    
    "github.com/rvinnie/deepseek-sdk-go/deepseek"
)

func main() {
    // 创建客户端
    client, err := deepseek.NewClient("<API_TOKEN>")
    if err != nil {
        // 处理错误
    }

    // 创建请求对象
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

    // 获取 HTTP 响应
    resp, err := client.CreateChatCompletionsWithStream(context.Background(), req)
    if err != nil {
        // 处理错误 
    }
    defer func() {
        _ = resp.Body.Close()
    }()

    // 创建流读取器
    reader := bufio.NewReader(resp.Body)

    // 处理流数据
    var full string
    for {
        line, err := deepseek.ReadStream(reader)
        if errors.Is(err, io.EOF) {
            // 流结束
			break
        }
        if err != nil {
            // 处理错误
            break
        }
        for _, choice := range line.Choices {
            // 聚合分块内容
            full += choice.Delta.Content
        }
    }
}

```
