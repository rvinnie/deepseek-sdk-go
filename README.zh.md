[![MIT License](https://img.shields.io/badge/license-MIT-blue.svg)](LICENSE)
[![Golang](https://img.shields.io/badge/Go-v1.21-EEEEEE?logo=go&logoColor=white&labelColor=00ADD8)](https://go.dev/)

<div align="center">
    <h1 align="center">Deepseek API Golang 客户端库
    </h1>
    <h3 align="center">用于 <a href="https://api-docs.deepseek.com/">Deepseek API</a> 的客户端
    </h3>
    <p align="center">
        <a href="README.md">English</a> | 中文 | <a href="README.ru.md">Русский</a>
    </p>
</div>

### 安装
`go get github.com/rvinnie/deepseek-sdk-go`

### 快速入门
1. 导入模块

```golang
import "github.com/rvinnie/deepseek-sdk-go/deepseek"
```
2. 创建 Deepseek API 客户端
```golang
import "github.com/rvinnie/deepseek-sdk-go/deepseek"

func main() {
    client, err := deepseek.NewClient("API_TOKEN")
    if err != nil {
        // 处理错误
    }
}
```
3. 调用所需 API 方法。 [查看 Deepseek API 文档](https://api-docs.deepseek.com/api/deepseek-api)

## SDK 使用示例
#### [Deepseek API SDK 配置](https://github.com/rvinnie/deepseek-sdk-go/blob/main/docs/examples/01-configuration.zh.md)
* [认证](https://github.com/rvinnie/deepseek-sdk-go/blob/main/docs/examples/01-configuration.zh.md#认证)
#### [聊天操作](https://github.com/rvinnie/deepseek-sdk-go/blob/main/docs/examples/02-chat-completions.zh.md)
* [创建聊天请求](https://github.com/rvinnie/deepseek-sdk-go/blob/main/docs/examples/02-chat-completions.zh.md#创建聊天请求)
* [流式传输内容](https://github.com/rvinnie/deepseek-sdk-go/blob/main/docs/examples/02-chat-completions.zh.md#流式传输内容)
#### [聊天操作（Fill-In-the-Middle）](https://github.com/rvinnie/deepseek-sdk-go/blob/main/docs/examples/03-fim-completions.zh.md)
* [创建聊天请求](https://github.com/rvinnie/deepseek-sdk-go/blob/main/docs/examples/03-fim-completions.zh.md#创建聊天请求)
#### [模型操作](https://github.com/rvinnie/deepseek-sdk-go/blob/main/docs/examples/04-models.zh.md)
* [获取可用模型列表](https://github.com/rvinnie/deepseek-sdk-go/blob/main/docs/examples/04-models.zh.md#获取可用模型列表)
#### [账户操作](https://github.com/rvinnie/deepseek-sdk-go/blob/main/docs/examples/05-user.zh.md)
* [获取用户余额](https://github.com/rvinnie/deepseek-sdk-go/blob/main/docs/examples/05-user.zh.md#获取用户余额)