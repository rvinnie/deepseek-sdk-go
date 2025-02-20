## Работа с моделями

* [Получение списка моделей](#Получение-списка-моделей)

---

### Получение списка моделей
[Получение списка моделей в документации](https://api-docs.deepseek.com/api/list-models)

Чтобы создать запрос на получение списка моделей, необходимо вызвать метод `GetModels`.  
В ответ на запрос возвращается объект `ModelsResponse`.

```go
import (
    "context"
    
    "github.com/rvinnie/deepseek-sdk-go/deepseek"
)

func main() {
    // Создание клиента
    client, err := deepseek.NewClient("<API_TOKEN>")
    if err != nil {
        // Обработка ошибки
    }
    
    // Получение списка моделей 
    resp, err := client.GetModels(context.Background())
    if err != nil {
        // Обработка ошибки
    }
}

```
