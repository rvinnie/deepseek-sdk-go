## 聊天操作（Fill-In-the-Middle）

* [创建聊天请求](#创建聊天请求)

---

### 创建聊天请求
[API 文档](https://api-docs.deepseek.com/api/create-completion)

创建 `BetaCompletionsRequest` 对象并传递给 `CreateBetaCompletions` 方法。  
返回 `BetaCompletionsResponse` 对象。  

```go
import (
    "context"
    
    "github.com/rvinnie/deepseek-sdk-go/deepseek"
)

func main() {
    // 创建客户端  
    client, err := deepseek.NewClient("<API_TOKEN>")
    if err != nil {
        // 处理错误
    }

    // 创建请求对象
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

    // 获取响应
    resp, err := client.CreateBetaCompletions(context.Background(), req)
    if err != nil {
        // 处理错误
    }
}

```
