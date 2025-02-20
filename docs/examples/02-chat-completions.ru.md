## Работа с чатом

* [Создание запроса к чату](#Создание-запроса-к-чату)
* [Потоковая передача содержимого](#Потоковая-передача-содержимого)

---

### Создание запроса к чату
[Создание запроса к чату в документации](https://api-docs.deepseek.com/api/create-chat-completion)

Чтобы создать запрос к чату, необходимо создать объект `ChatCompletionsRequest` и передать его в метод `CreateChatCompletions`.  
В ответ на запрос возвращается объект `ChatCompletionsResponse`.

```go
import (
    "context"

    "github.com/rvinnie/deepseek-sdk-go/deepseek"
)

func main() {
    // Создание клиента
    client, err := deepseek.NewClient("API_TOKEN")
    if err != nil {
        // Обработка ошибки
    }

    // Создание объекта запроса
    req := &deepseek.ChatCompletionsRequest{
        Messages: []deepseek.ChatCompletionsMessage{
            {
                Role:    deepseek.ChatMessageRoleSystem,
                Content: "You are a helpful assistant",
            },
            {
                Role:    deepseek.ChatMessageRoleUser,
                Content: "Hi",
            },
        },
        Model:            deepseek.DeepSeekChat,
        FrequencyPenalty: 0,
        PresencePenalty:  0,
        ResponseFormat: &deepseek.ResponseFormat{
            Type: "text",
        },
        Stop:          nil,
        Stream:        false,
        StreamOptions: nil,
        Temperature:   1,
        TopP:          1,
        Tools:         nil,
        ToolChoice: deepseek.ToolChoice{
            Type: "none",
        },
        LogProbs:    false,
        TopLogProbs: nil,
    }

    // Получение ответа от чата
    resp, err := client.CreateChatCompletions(context.Background(), req)
    if err != nil {
        // Обработка ошибки
    }
}
```

---

### Потоковая передача содержимого

[Потоковая передача содержимого в документации](https://api-docs.deepseek.com/api/create-chat-completion)

Чтобы создать запрос к чату на потоковую передачу данных, необходимо создать объект `ChatCompletionsRequest` и передать его в метод `CreateChatCompletionsWithStream`.   
В ответ на запрос возвращается объект `http.Response`.  

Далее создаем объект `*bufio.Reader` и передаем его в функцию `ReadStream`, которая читает поток данных из ответа.

```go
import (
    "bufio"
    "context"
    "errors"
    "io"
    
    "github.com/rvinnie/deepseek-sdk-go/deepseek"
)

func main() {
    // Создание клиента
    client, err := deepseek.NewClient("<API_TOKEN>")
    if err != nil {
        // Обработка ошибки
    }

    // Создание объекта запроса
    req := &deepseek.ChatCompletionsRequest{
        Messages: []deepseek.ChatCompletionsMessage{
            {
                Role:    deepseek.ChatMessageRoleSystem,
                Content: "You are a helpful assistant",
            },
            {
                Role:    deepseek.ChatMessageRoleUser,
                Content: "Hi",
            },
        },
        Model:            deepseek.DeepSeekChat,
        FrequencyPenalty: 0,
        PresencePenalty:  0,
        ResponseFormat: &deepseek.ResponseFormat{
            Type: "text",
        },
        Stop:          nil,
        Stream:        true,
        StreamOptions: nil,
        Temperature:   1,
        TopP:          1,
        Tools:         nil,
        ToolChoice: deepseek.ToolChoice{
            Type: "none",
        },
        LogProbs:    false,
        TopLogProbs: nil,
    }

    // Получение объекта ответа http.Response
    resp, err := client.CreateChatCompletionsWithStream(context.Background(), req)
    if err != nil {
        // Обработка ошибки
    }
    defer func() {
        _ = resp.Body.Close()
    }()

    // Создание объекта для чтение потока данных из ответа
    reader := bufio.NewReader(resp.Body)

    // Создание цикла для чтение потока данных
    var full string
    for {
        // Вызов функции чтения потока данных
		part, err := deepseek.ReadStream(reader)
        if errors.Is(err, io.EOF) {
            // Конец потока данных
            break
        }
        if err != nil {
            // Обработка ошибки
            break
        }
        for _, choice := range part.Choices {
            // Агрегация всех чанков в одно сообщение
            full += choice.Delta.Content
        }
    }
}

```
