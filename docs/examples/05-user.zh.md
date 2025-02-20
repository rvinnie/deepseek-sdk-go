## 账户操作

* [获取用户余额](#获取用户余额)

---

### 获取用户余额
[API 文档](https://api-docs.deepseek.com/api/get-user-balance)

调用 `GetUserBalance` 方法获取用户余额。返回 `BalanceResponse` 对象。  

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

    // 获取用户余额
    resp, err := client.GetUserBalance(context.Background())
    if err != nil {
        // 处理错误 
    }
}

```
