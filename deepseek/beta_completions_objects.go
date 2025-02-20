package deepseek

import (
	"fmt"
	"strings"
	"time"
)

// BetaCompletions represents the request body for a Fill-In-the-Middle (FIM) completion.
type BetaCompletionsRequest struct {
	// ID of the model to use.
	Model string `json:"model"`

	// The prompt to generate completions for.
	Prompt string `json:"prompt"`

	// Echo back the prompt in addition to the completion
	Echo bool `json:"echo,omitempty"`

	// Number between -2.0 and 2.0.
	// Positive values penalize new tokens based on their existing frequency in the text so far,
	// decreasing the model's likelihood to repeat the same line verbatim.
	FrequencyPenalty float64 `json:"frequency_penalty,omitempty"`

	// Include the log probabilities on the logprobs most likely output tokens, as well the chosen tokens.
	// For example, if logprobs is 20, the API will return a list of the 20 most likely tokens.
	// The API will always return the logprob of the sampled token, so there may be up to logprobs+1 elements in the response.
	// The maximum value for logprobs is 20.
	Logprobs int `json:"logprobs,omitempty"`

	// The maximum number of tokens that can be generated in the completion.
	MaxTokens int `json:"max_tokens,omitempty"`

	// Number between -2.0 and 2.0.
	// Positive values penalize new tokens based on whether they appear in the text so far,
	// increasing the model's likelihood to talk about new topics.
	PresencePenalty float64 `json:"presence_penalty,omitempty"`

	// Up to 16 sequences where the API will stop generating further tokens.
	// The returned text will not contain the stop sequence.
	Stop []string `json:"stop,omitempty"`

	// If set, partial message deltas will be sent.
	// Tokens will be sent as data-only server-sent events (SSE) as they become available,
	// with the stream terminated by a data: [DONE] message.
	Stream bool `json:"stream,omitempty"`

	// Options for streaming response. Only set this when you set stream: true.
	StreamOptions *StreamOptions `json:"stream_options,omitempty"`

	// The suffix that comes after a completion of inserted text.
	Suffix string `json:"suffix,omitempty"`

	// What sampling temperature to use, between 0 and 2.
	// Higher values like 0.8 will make the output more random,
	// while lower values like 0.2 will make it more focused and deterministic.
	// We generally recommend altering this or top_p but not both.
	Temperature float64 `json:"temperature,omitempty"`

	// An alternative to sampling with temperature, called nucleus sampling,
	// where the model considers the results of the tokens with top_p probability mass.
	// So 0.1 means only the tokens comprising the top 10% probability mass are considered.
	// We generally recommend altering this or temperature but not both
	TopP float64 `json:"top_p,omitempty"`
}

// LogprobsBetaCompletions represents log probability information for a choice or token.
type LogprobsBetaCompletions struct {
	TextOffsets   []int         `json:"text_offset"`
	TokenLogprobs []float64     `json:"token_logprobs"`
	Tokens        []string      `json:"tokens"`
	TopLogprobs   []TopLogprobs `json:"top_logprobs"`
}

// Choice represents choice the model generated for the input prompt.
type ChoiceBetaCompletions struct {
	// The reason the model stopped generating tokens.
	// This will be stop if the model hit a natural stop point or a provided stop sequence,
	// length if the maximum number of tokens specified in the request was reached,
	// content_filter if content was omitted due to a flag from our content filters,
	// or insufficient_system_resource if the request is interrupted due to insufficient resource of the inference system.
	FinishReason string `json:"finish_reason"`

	// Index of the choice.
	Index int `json:"index"`

	// Log probabilities of the generated tokens.
	Logprobs *LogprobsBetaCompletions `json:"logprobs,omitempty"`

	// The generated completion text.
	Text string `json:"text"`
}

// BetaCompletionsResponse represents the response body for a Fill-In-the-Middle (FIM) completion.
type BetaCompletionsResponse struct {
	// A unique identifier for the completion.
	ID string `json:"id"`

	// The list of completion choices the model generated for the input prompt.
	Choices []ChoiceBetaCompletions `json:"choices"`

	// The Unix timestamp (in seconds) of when the completion was created.
	Created int `json:"created"`

	// The model used for completion.
	Model string `json:"model"`

	// This fingerprint represents the backend configuration that the model runs with.
	SystemFingerprint string `json:"system_fingerprint,omitempty"`

	// The object type, which is always "text_completion"
	Object string `json:"object"`

	// Usage statistics for the completion request.
	Usage Usage `json:"usage"`
}

func (r BetaCompletionsResponse) String() string {
	var sb strings.Builder
	sb.WriteString("{")

	// Basic fields
	sb.WriteString(fmt.Sprintf("ID: %q", r.ID))
	sb.WriteString(fmt.Sprintf(", Created: %s", time.Unix(int64(r.Created), 0).Format(time.RFC3339)))
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
		sb.WriteString(fmt.Sprintf(", CompletionTokensDetails: {ReasoningTokens: %d}",
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
		sb.WriteString(fmt.Sprintf(", Text: %q", choice.Text))

		// Logprobs
		if choice.Logprobs != nil {
			sb.WriteString(", Logprobs: {")
			// TextOffsets
			sb.WriteString("TextOffsets: [")
			for i, offset := range choice.Logprobs.TextOffsets {
				if i > 0 {
					sb.WriteString(", ")
				}
				sb.WriteString(fmt.Sprintf("%d", offset))
			}
			sb.WriteString("]")

			// TokenLogprobs
			sb.WriteString(", TokenLogprobs: [")
			for i, logprob := range choice.Logprobs.TokenLogprobs {
				if i > 0 {
					sb.WriteString(", ")
				}
				sb.WriteString(fmt.Sprintf("%.2f", logprob))
			}
			sb.WriteString("]")

			// Tokens
			sb.WriteString(", Tokens: [")
			for i, token := range choice.Logprobs.Tokens {
				if i > 0 {
					sb.WriteString(", ")
				}
				sb.WriteString(fmt.Sprintf("%q", token))
			}
			sb.WriteString("]")

			// TopLogprobs
			sb.WriteString(", TopLogprobs: [")
			for i, tlp := range choice.Logprobs.TopLogprobs {
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
			sb.WriteString("}")
		}
		sb.WriteString("}")
	}
	sb.WriteString("]")
	sb.WriteString("}")

	return sb.String()
}
