[![MIT License](https://img.shields.io/badge/license-MIT-blue.svg)](LICENSE)
[![Golang](https://img.shields.io/badge/Go-v1.21-EEEEEE?logo=go&logoColor=white&labelColor=00ADD8)](https://go.dev/)

<div align="center">
    <h1 align="center">Deepseek API Golang Client Library
    </h1>
    <h3 align="center">Client for <a href="https://api-docs.deepseek.com/">Deepseek API</a>
    </h3>
    <p align="center">
        English | <a href="README.zh.md">中文</a> | <a href="README.ru.md">Русский</a>
    </p>
</div>

### Installation
`go get github.com/rvinnie/deepseek-sdk-go`

### Getting Started
1. Import the module
```golang
import "github.com/rvinnie/deepseek-sdk-go/deepseek"
```
2. Create a Deepseek API client
```golang
import "github.com/rvinnie/deepseek-sdk-go/deepseek"

func main() {
    client, err := deepseek.NewClient("API_TOKEN")
    if err != nil { 
       // Handle error
    }
}
```
3. Call the desired API method. [See Deepseek API documentation](https://api-docs.deepseek.com/api/deepseek-api)

## SDK Usage Examples
#### [Deepseek API SDK Configuration](https://github.com/rvinnie/deepseek-sdk-go/blob/main/docs/examples/01-configuration.md)
* [Authentication](https://github.com/rvinnie/deepseek-sdk-go/blob/main/docs/examples/01-configuration.md#authentication)
#### [Chat Operations](https://github.com/rvinnie/deepseek-sdk-go/blob/main/docs/examples/02-chat-completions.md)
* [Create Chat Request](https://github.com/rvinnie/deepseek-sdk-go/blob/main/docs/examples/02-chat-completions.md#create-chat-request)
* [Streaming Content](https://github.com/rvinnie/deepseek-sdk-go/blob/main/docs/examples/02-chat-completions.md#streaming-content)
#### [Chat Operations (Fill-In-the-Middle)](https://github.com/rvinnie/deepseek-sdk-go/blob/main/docs/examples/03-fim-completions.md)
* [Create Chat Request](https://github.com/rvinnie/deepseek-sdk-go/blob/main/docs/examples/03-fim-completions.md#create-chat-request)
#### [Model Operations](https://github.com/rvinnie/deepseek-sdk-go/blob/main/docs/examples/04-models.md)
* [List Available Models](https://github.com/rvinnie/deepseek-sdk-go/blob/main/docs/examples/04-models.md#list-available-models)
#### [Account Operations](https://github.com/rvinnie/deepseek-sdk-go/blob/main/docs/examples/05-user.md)
* [Get User Balance](https://github.com/rvinnie/deepseek-sdk-go/blob/main/docs/examples/05-user.md#get-user-balance)