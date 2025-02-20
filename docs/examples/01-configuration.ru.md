## Настройки SDK API Deepseek

[Справочник API Deepseek](https://api-docs.deepseek.com/api/deepseek-api)
* [Аутентификация](#Аутентификация)
---

### Аутентификация

Для работы с API необходимо создать клиента, указав секретный ключ. Создать ключ можно в [документации](https://platform.deepseek.com/api_keys)

```go
import "github.com/rvinnie/deepseek-sdk-go/deepseek"

func main() {
    client, err := deepseek.NewClient("<API_TOKEN>")
    if err != nil {
        // Обработка ошибки
    }
}
```
