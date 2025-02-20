## Работа с чатом (Fill-In-the-Middle)

* [Создание запроса к чату](#Создание-запроса-к-чату)

---

### Создание запроса к чату
[Создание запроса к чату в документации](https://api-docs.deepseek.com/api/create-completion)

Чтобы создать запрос к чату, необходимо создать объект `BetaCompletionsRequest` и передать его в метод `CreateBetaCompletions`.   
В ответ на запрос возвращается объект `BetaCompletionsResponse`.

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

    // Создание объекта запроса
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

    // Получение ответа от чата
    resp, err := client.CreateBetaCompletions(context.Background(), req)
    if err != nil {
        // Обработка ошибки
    }
}

```
