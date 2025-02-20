package deepseek

const (
	// ChatMessageRoleSystem is the role of a system message
	ChatMessageRoleSystem = "system"
	// ChatMessageRoleUser is the role of a user message
	ChatMessageRoleUser = "user"
	// ChatMessageRoleAssistant is the role of an assistant message
	ChatMessageRoleAssistant = "assistant"
)

// Official DeepSeek Models
const (
	DeepSeekChat     = "deepseek-chat"     // DeepSeekChat is the official model for chat completions
	DeepSeekReasoner = "deepseek-reasoner" // DeepSeekReasoner is the official model for reasoning completions
)
