package deepseek

type StreamOptions struct {
	// If set, an additional chunk will be streamed before the data: [DONE] message.
	// The usage field on this chunk shows the token usage statistics for the entire request,
	// and the choices field will always be an empty array. All other chunks will also include a usage field,
	// but with a null value.
	IncludeUsage bool `json:"include_usage"`
}

// CompletionTokensDetails represents breakdown of tokens used in a completion.
type CompletionTokensDetails struct {
	// Tokens generated by the model for reasoning.
	ReasoningTokens int `json:"reasoning_tokens"`
}

// Usage statistics for the completion request.
type Usage struct {
	// Number of tokens in the generated completion.
	CompletionTokens int `json:"completion_tokens"`

	// Number of tokens in the prompt. It equals prompt_cache_hit_tokens + prompt_cache_miss_tokens.
	PromptTokens int `json:"prompt_tokens"`

	// Number of tokens in the prompt that hits the context cache.
	PromptCacheHitTokens int `json:"prompt_cache_hit_tokens"`

	// Number of tokens in the prompt that misses the context cache.
	PromptCacheMissTokens int `json:"prompt_cache_miss_tokens"`

	// Total number of tokens used in the request (prompt + completion).
	TotalTokens int `json:"total_tokens"`

	// Breakdown of tokens used in a completion.
	CompletionTokensDetails *CompletionTokensDetails `json:"completion_tokens_details,omitempty"`
}

// TopLogprobs represents a single token within the top log probabilities with its log probability and byte information.
type TopLogprobs struct {
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
}
