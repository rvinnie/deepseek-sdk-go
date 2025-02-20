## Account Operations

* [Get User Balance](#get-user-balance)

---

### Get User Balance
[API Documentation](https://api-docs.deepseek.com/api/get-user-balance)

Call `GetUserBalance` method to retrieve user balance. Returns a `BalanceResponse`.

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

    // Get user balance
    resp, err := client.GetUserBalance(context.Background())
    if err != nil {
        // Handle error
    }
}

```
