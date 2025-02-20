[![MIT License](https://img.shields.io/badge/license-MIT-blue.svg)](LICENSE)
[![Golang](https://img.shields.io/badge/Go-v1.21-EEEEEE?logo=go&logoColor=white&labelColor=00ADD8)](https://go.dev/)

<div align="center">
    <h1 align="center">Deepseek API Golang Client Library
    </h1>
    <h3 align="center">Клиент для работы с <a href="https://api-docs.deepseek.com/">Deepseek API</a>
    </h3>
    <p align="center">
        <a href="README.md">English</a> | <a href="README.zh.md">中文</a> | Русский
    </p>
</div>

### Установка
`go get github.com/rvinnie/deepseek-sdk-go`

### Начало работы
1. Импортируйте модуль
```golang
import "github.com/rvinnie/deepseek-sdk-go/deepseek"
```
2. Создайте клиента для работы с Deepseek API
```golang
import "github.com/rvinnie/deepseek-sdk-go/deepseek"

func main() {
    client, err := deepseek.NewClient("API_TOKEN")
    if err != nil {
        // Обработка ошибки
    }
}
```
3. Вызовите нужный метод API. [Подробнее в документации к API Deepseek](https://api-docs.deepseek.com/api/deepseek-api)

## Примеры использования SDK
#### [Настройки SDK API Deepseek](https://github.com/rvinnie/deepseek-sdk-go/blob/main/docs/examples/01-configuration.ru.md)
* [Аутентификация](https://github.com/rvinnie/deepseek-sdk-go/blob/main/docs/examples/01-configuration.ru.md#Аутентификация)
#### [Работа с чатом](https://github.com/rvinnie/deepseek-sdk-go/blob/main/docs/examples/02-chat-completions.ru.md)
* [Создание запроса к чату](https://github.com/rvinnie/deepseek-sdk-go/blob/main/docs/examples/02-chat-completions.ru.md#Создание-запроса-к-чату)
* [Потоковая передача содержимого](https://github.com/rvinnie/deepseek-sdk-go/blob/main/docs/examples/02-chat-completions.ru.md#Потоковая-передача-содержимого)
#### [Работа с чатом (Fill-In-the-Middle)](https://github.com/rvinnie/deepseek-sdk-go/blob/main/docs/examples/03-fim-completions.ru.md)
* [Создание запроса к чату](https://github.com/rvinnie/deepseek-sdk-go/blob/main/docs/examples/03-fim-completions.ru.md#Создание-запроса-к-чату)
#### [Работа с моделями](https://github.com/rvinnie/deepseek-sdk-go/blob/main/docs/examples/04-models.ru.md)
* [Получение списка моделей](https://github.com/rvinnie/deepseek-sdk-go/blob/main/docs/examples/04-models.ru.md#Получение-списка-моделей)
#### [Работа с аккаунтом](https://github.com/rvinnie/deepseek-sdk-go/blob/main/docs/examples/05-user.ru.md)
* [Получение баланса пользователя](https://github.com/rvinnie/deepseek-sdk-go/blob/main/docs/examples/05-user.ru.md#Получение-баланса-пользователя)