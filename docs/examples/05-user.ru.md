## Работа с аккаунтом

* [Получение баланса пользователя](#Получение-баланса-пользователя)

---

### Получение баланса пользователя
[Получение баланса пользователя в документации](https://api-docs.deepseek.com/api/get-user-balance)

Чтобы создать запрос на получение баланс пользователя, необходимо вызвать метод `GetUserBalance`.  
В ответ на запрос возвращается объект `BalanceResponse`.

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

    // Получение баланса пользователя
    resp, err := client.GetUserBalance(context.Background())
    if err != nil {
        // Обработка ошибки
    }
}

```
