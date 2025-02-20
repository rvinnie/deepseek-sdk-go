## Model Operations

* [List Available Models](#list-available-models)

---

### List Available Models
[API Documentation](https://api-docs.deepseek.com/api/list-models)

Call `GetModels` method to retrieve available models. Returns a `ModelsResponse`.

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
    
    // Get models list 
    resp, err := client.GetModels(context.Background())
    if err != nil {
        // Handle error
    }
}

```
