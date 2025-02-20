## 模型操作

* [获取可用模型列表](#获取可用模型列表)

---

### 获取可用模型列表
[API 文档](https://api-docs.deepseek.com/api/list-models)

调用 `GetModels` 方法获取可用模型列表。返回 `ModelsResponse` 对象。

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
    
    // 获取模型列表 
    resp, err := client.GetModels(context.Background())
    if err != nil {
        // 处理错误
    }
}

```
