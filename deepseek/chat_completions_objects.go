package deepseek

import (
	"fmt"
	"strings"
	"time"
)

// ChatCompletionsMessage represents a single message in a chat completion conversation.
type ChatCompletionsMessage struct {
	// The role of the messages author, e.g., "system", "user", "assistant", "tool".
	Role string `json:"role"`

	// The contents of the message.
	Content string `json:"content"`

	// An optional name for the participant. Provides the model information to differentiate between participants of the same role.
	Name string `json:"name,omitempty"`

	// The prefix of the message [Beta Feature].
	Prefix bool `json:"prefix,omitempty"`

	// (Beta) Used for the deepseek-reasoner model in the Chat Prefix Completion feature as the input for the CoT in the last assistant message. When using this feature, the prefix parameter must be set to true.
	ReasoningContent string `json:"reasoning_content,omitempty"`

	// Tool call that this message is responding to.
	ToolCallID string `json:"tool_call_id,omitempty"`
}

// ResponseFormat defines the structure for the response format.
type ResponseFormat struct {
	// Must be one of text or json_object
	Type string `json:"type"`
}

// Tool defines the structure for a tool.
type Tool struct {
	// The type of the tool. Currently, only function is supported.
	Type string `json:"type"`

	// The function details.
	Function Function `json:"function"`
}

// Function defines the structure of a function tool.
type Function struct {
	// The name of the function to be called.
	// Must be a-z, A-Z, 0-9, or contain underscores and dashes, with a maximum length of 64.
	Name string `json:"name"`

	// A description of what the function does, used by the model to choose when and how to call the function.
	Description string `json:"description"`

	// The parameters the functions accepts, described as a JSON Schema object.
	// See the Function Calling Guide for examples, and the JSON Schema reference for documentation about the format.
	// Omitting parameters defines a function with an empty parameter list.
	Parameters *FunctionParameters `json:"parameters,omitempty"`
}

// FunctionParameters defines the parameters for a function.
type FunctionParameters struct {
	// The type of the parameters, e.g., "object".
	Type string `json:"type"`

	// The properties of the parameters.
	Properties map[string]interface{} `json:"properties,omitempty"`

	// A list of required parameter names.
	Required []string `json:"required,omitempty"`
}

// ToolChoice defines the structure for a tool choice.
type ToolChoice struct {
	// The type of the tool, e.g., "function".
	Type string `json:"type"`
	// The function details.
	Function ToolChoiceFunction `json:"function,omitempty"`
}

// ToolChoiceFunction defines the function details within ToolChoice.
type ToolChoiceFunction struct {
	// The name of the function to call (required).
	Name string `json:"name"`
}

// ChatCompletionsRequest defines the structure for a chat completion request.
type ChatCompletionsRequest struct {
	// A list of messages comprising the conversation so far.
	Messages []ChatCompletionsMessage `json:"messages"`

	// The ID of the model to use: deepseek-chat, deepseek-reasoner.
	Model string `json:"model"`

	// Number between -2.0 and 2.0.
	// Positive values penalize new tokens based on their existing frequency in the text so far,
	// decreasing the model's likelihood to repeat the same line verbatim.
	FrequencyPenalty float32 `json:"frequency_penalty,omitempty"`

	// Integer between 1 and 8192.
	// The maximum number of tokens that can be generated in the chat completion.
	// The total length of input tokens and generated tokens is limited by the model's context length.
	// If max_tokens is not specified, the default value 4096 is used.
	MaxTokens *int `json:"max_tokens,omitempty"`

	// Number between -2.0 and 2.0.
	// Positive values penalize new tokens based on whether they appear in the text so far,
	// increasing the model's likelihood to talk about new topics.
	PresencePenalty float32 `json:"presence_penalty,omitempty"`

	// An object specifying the format that the model must output.
	// Setting to { "type": "json_object" } enables JSON Output, which guarantees the message the model generates is valid JSON.
	// Important: When using JSON Output, you must also instruct the model to produce JSON yourself via a system or user message.
	// Without this, the model may generate an unending stream of whitespace until the generation reaches the token limit,
	// resulting in a long-running and seemingly "stuck" request. Also note that the message content may be partially cut off
	// if finish_reason="length", which indicates the generation exceeded max_tokens or the conversation exceeded the max context length.
	ResponseFormat *ResponseFormat `json:"response_format,omitempty"`

	// Up to 16 sequences where the API will stop generating further tokens.
	Stop []string `json:"stop,omitempty"`

	// If set, partial message deltas will be sent.
	// Tokens will be sent as data-only server-sent events (SSE) as they become available,
	// with the stream terminated by a data: [DONE] message.
	Stream bool `json:"stream,omitempty"`

	// Options for streaming response. Only set this when you set stream: true.
	StreamOptions *StreamOptions `json:"stream_options,omitempty"`

	// What sampling temperature to use, between 0 and 2.
	// Higher values like 0.8 will make the output more random,
	// while lower values like 0.2 will make it more focused and deterministic.
	// We generally recommend altering this or top_p but not both.
	Temperature float32 `json:"temperature,omitempty"`

	// An alternative to sampling with temperature, called nucleus sampling,
	// where the model considers the results of the tokens with top_p probability mass.
	// So 0.1 means only the tokens comprising the top 10% probability mass are considered.
	// We generally recommend altering this or temperature but not both.
	TopP float32 `json:"top_p,omitempty"`

	// A list of tools the model may call. Currently, only functions are supported as a tool.
	// Use this to provide a list of functions the model may generate JSON inputs for.
	// A max of 128 functions are supported.
	Tools []Tool `json:"tools,omitempty"`

	// Controls which (if any) tool is called by the model.
	// 'none' means the model will not call any tool and instead generates a message.
	// 'auto' means the model can pick between generating a message or calling one or more tools.
	// 'required' means the model must call one or more tools.
	// Specifying a particular tool via {"type": "function", "function": {"name": "my_function"}} forces the model to call that tool.
	// 'none' is the default when no tools are present. auto is the default if tools are present.
	ToolChoice interface{} `json:"tool_choice,omitempty"`

	// Whether to return log probabilities of the output tokens or not.
	// If true, returns the log probabilities of each output token returned in the content of message.
	LogProbs bool `json:"logprobs,omitempty"`

	// An integer between 0 and 20 specifying the number of most likely tokens to return at each token position,
	// each with an associated log probability.
	// logprobs must be set to true if this parameter is used.
	TopLogProbs *int `json:"top_logprobs,omitempty"` // The number of top most likely tokens to return log probabilities for.
}

// Choice represents a completion choice generated by the model.
type ChoiceChatCompletions struct {
	// The reason the model stopped generating tokens.
	// This will be 'stop' if the model hit a natural stop point or a provided stop sequence,
	// length if the maximum number of tokens specified in the request was reached,
	// content_filter if content was omitted due to a flag from our content filters,
	// tool_calls if the model called a tool, or insufficient_system_resource if the request
	// is interrupted due to insufficient resource of the inference system.
	FinishReason string `json:"finish_reason"`

	// The index of the choice in the list of choices.
	Index int `json:"index"`

	// A chat completion message generated by the model.
	Message *Message `json:"message,omitempty"`

	// A chat completion delta generated by streamed model responses
	Delta *Delta `json:"delta,omitempty"`

	// Log probability information for the choice.
	Logprobs *LogprobsChatCompletions `json:"logprobs,omitempty"`
}

// LogprobsChatCompletions represents log probability information for a choice or token.
type LogprobsChatCompletions struct {
	// A list of message content tokens with log probability information.
	Content []ContentToken `json:"content"`
}

// ContentToken represents a single token within the content with its log probability and byte information.
type ContentToken struct {
	// The token.
	Token string `json:"token"`

	// The log probability of this token, if it is within the top 20 most likely tokens.
	// Otherwise, the value -9999.0 is used to signify that the token is very unlikely.
	Logprob float64 `json:"logprob"`

	// A list of integers representing the UTF-8 bytes representation of the token.
	// Useful in instances where characters are represented by multiple tokens and their byte
	// representations must be combined to generate the correct text representation.
	// Can be null if there is no bytes representation for the token.
	Bytes []int `json:"bytes,omitempty"`

	// List of the most likely tokens and their log probability, at this token position.
	// In rare cases, there may be fewer than the number of requested top_logprobs returned.
	TopLogprobs []TopLogprobs `json:"top_logprobs"`
}

// TopLogprobToken represents a single token within the top log probabilities with its log probability and byte information.
type TopLogprobToken struct {
	Token   string  `json:"token"`           // The token string.
	Logprob float64 `json:"logprob"`         // The log probability of this token. -9999.0 if not in top 20.
	Bytes   []int   `json:"bytes,omitempty"` // UTF-8 byte representation of the token. Can be nil.
}

// Message represents a message generated by the model.
type Message struct {
	// The contents of the message.
	Content string `json:"content"`

	// For deepseek-reasoner model only. The reasoning contents of the assistant message, before the final answer.
	ReasoningContent string `json:"reasoning_content,omitempty"`

	// The tool calls generated by the model, such as function calls.
	ToolCalls []ToolCall `json:"tool_calls,omitempty"`

	// The role of the author of this message.
	Role string `json:"role"`
}

// Delta represents a chat completion delta generated by streamed model responses.
type Delta struct {
	// The contents of the message.
	Content string `json:"content,omitempty"`
	// For deepseek-reasoner model only. The reasoning contents of the assistant message, before the final answer.
	ReasoningContent string `json:"reasoning_content,omitempty"`
	// The role of the author of this message.
	Role string `json:"role,omitempty"`
}

// ToolCall represents a tool call in the completion.
type ToolCall struct {
	// The ID of the tool call.
	ID string `json:"id"`

	// The type of the tool. Currently, only function is supported.
	Type string `json:"type"`

	// The function that the model called.
	Function ToolCallFunction `json:"function"`
}

// ToolCallFunction represents a function call in the tool.
type ToolCallFunction struct {
	// The name of the function to call.
	Name string `json:"name"`

	// The arguments to call the function with, as generated by the model in JSON format.
	// Note that the model does not always generate valid JSON,
	// and may hallucinate parameters not defined by your function schema.
	// Validate the arguments in your code before calling your function.
	Arguments string `json:"arguments"`
}

// ChatCompletionsResponse defines the structure for a chat completion response.
type ChatCompletionsResponse struct {
	// A unique identifier for the chat completion.
	ID string `json:"id"`

	// A list of chat completion choices.
	Choices []ChoiceChatCompletions `json:"choices"`

	// The Unix timestamp (in seconds) of when the chat completion was created.
	Created int64 `json:"created"`

	// The model used for the chat completion.
	Model string `json:"model"`

	// This fingerprint represents the backend configuration that the model runs with.
	SystemFingerprint string `json:"system_fingerprint"`

	// The object type, which is always chat.completion.
	Object string `json:"object"`

	// Usage statistics for the completion request.
	Usage Usage `json:"usage"`
}

func (r ChatCompletionsResponse) String() string {
	var sb strings.Builder
	sb.WriteString("{")

	// Basic fields
	sb.WriteString(fmt.Sprintf("ID: %q", r.ID))
	sb.WriteString(fmt.Sprintf(", Created: %s", time.Unix(r.Created, 0).Format(time.RFC3339)))
	sb.WriteString(fmt.Sprintf(", Model: %q", r.Model))
	sb.WriteString(fmt.Sprintf(", SystemFingerprint: %q", r.SystemFingerprint))
	sb.WriteString(fmt.Sprintf(", Object: %q", r.Object))

	// Usage
	sb.WriteString(", Usage: {")
	sb.WriteString(fmt.Sprintf("CompletionTokens: %d", r.Usage.CompletionTokens))
	sb.WriteString(fmt.Sprintf(", PromptTokens: %d", r.Usage.PromptTokens))
	sb.WriteString(fmt.Sprintf(", PromptCacheHitTokens: %d", r.Usage.PromptCacheHitTokens))
	sb.WriteString(fmt.Sprintf(", PromptCacheMissTokens: %d", r.Usage.PromptCacheMissTokens))
	sb.WriteString(fmt.Sprintf(", TotalTokens: %d", r.Usage.TotalTokens))
	if r.Usage.CompletionTokensDetails != nil {
		sb.WriteString(fmt.Sprintf(", CompletionTokensDetails: {ReasoningTokens: %q}",
			r.Usage.CompletionTokensDetails.ReasoningTokens))
	}
	sb.WriteString("}")

	// Choices
	sb.WriteString(", Choices: [")
	for i, choice := range r.Choices {
		if i > 0 {
			sb.WriteString(", ")
		}
		sb.WriteString("{")
		sb.WriteString(fmt.Sprintf("Index: %d", choice.Index))
		sb.WriteString(fmt.Sprintf(", FinishReason: %q", choice.FinishReason))

		// Delta
		if choice.Delta != nil {
			delta := choice.Delta
			sb.WriteString(", Delta: {")
			sb.WriteString(fmt.Sprintf("Role: %q", delta.Role))
			sb.WriteString(fmt.Sprintf(", Content: %q", delta.Content))
			if delta.ReasoningContent != "" {
				sb.WriteString(fmt.Sprintf(", ReasoningContent: %q", delta.ReasoningContent))
			}
		}

		// Message
		msg := choice.Message
		sb.WriteString(", Message: {")
		sb.WriteString(fmt.Sprintf("Role: %q", msg.Role))
		sb.WriteString(fmt.Sprintf(", Content: %q", msg.Content))
		if msg.ReasoningContent != "" {
			sb.WriteString(fmt.Sprintf(", ReasoningContent: %q", msg.ReasoningContent))
		}

		// ToolCalls
		if len(msg.ToolCalls) > 0 {
			sb.WriteString(", ToolCalls: [")
			for i, tc := range msg.ToolCalls {
				if i > 0 {
					sb.WriteString(", ")
				}
				sb.WriteString("{")
				sb.WriteString(fmt.Sprintf("ID: %q", tc.ID))
				sb.WriteString(fmt.Sprintf(", Type: %q", tc.Type))
				sb.WriteString("Function: {")
				sb.WriteString(fmt.Sprintf("Name: %q", tc.Function.Name))
				sb.WriteString(fmt.Sprintf(", Arguments: %q", tc.Function.Arguments))
				sb.WriteString("}")
				sb.WriteString("}")
			}
			sb.WriteString("]")
		}
		sb.WriteString("}")

		// Logprobs
		if choice.Logprobs != nil {
			sb.WriteString(", Logprobs: {")
			sb.WriteString("Content: [")
			for i, ct := range choice.Logprobs.Content {
				if i > 0 {
					sb.WriteString(", ")
				}
				sb.WriteString("{")
				sb.WriteString(fmt.Sprintf("Token: %q", ct.Token))
				sb.WriteString(fmt.Sprintf(", Logprob: %.2f", ct.Logprob))
				if len(ct.Bytes) > 0 {
					sb.WriteString(fmt.Sprintf(", Bytes: %v", ct.Bytes))
				}
				if len(ct.TopLogprobs) > 0 {
					sb.WriteString(", TopLogprobs: [")
					for _, tlp := range ct.TopLogprobs {
						if i > 0 {
							sb.WriteString(", ")
						}
						sb.WriteString("{")
						sb.WriteString(fmt.Sprintf("Token: %q", tlp.Token))
						sb.WriteString(fmt.Sprintf(", Logprob: %.2f", tlp.Logprob))
						if len(tlp.Bytes) > 0 {
							sb.WriteString(fmt.Sprintf(", Bytes: %v", tlp.Bytes))
						}
						sb.WriteString("}")
					}
					sb.WriteString("]")
				}
				sb.WriteString("}")
			}
			sb.WriteString("]")
			sb.WriteString("}")
		}
		sb.WriteString("}")
	}
	sb.WriteString("]")
	sb.WriteString("}")

	return sb.String()
}
