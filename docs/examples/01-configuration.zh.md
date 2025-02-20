## Deepseek API SDK 配置

[API 参考文档](https://api-docs.deepseek.com/api/deepseek-api)
* [认证](#认证)
---

### 认证

使用 API 需要创建包含密钥的客户端。密钥可在 [文档](https://platform.deepseek.com/api_keys)中生成。

```go
import "github.com/rvinnie/deepseek-sdk-go/deepseek"

func main() {
    client, err := deepseek.NewClient("<API_TOKEN>")
    if err != nil {
        // 处理错误
    }
}
```
