## Deepseek API SDK Configuration

[API Reference](https://api-docs.deepseek.com/api/deepseek-api)
* [Authentication](#authentication)
---

### Authentication

To use the API, create a client with your secret key. Generate keys in the [documentation](https://platform.deepseek.com/api_keys).

```go
import "github.com/rvinnie/deepseek-sdk-go/deepseek"

func main() {
    client, err := deepseek.NewClient("<API_TOKEN>")
    if err != nil {
        // Handle error
    }
}
```
