package deepseek

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"strings"
)

const (
	LastStreamMessage   = "data: [DONE]"
	PrefixStreamMessage = "data: "
)

func ReadStream(reader *bufio.Reader) (*ChatCompletionsResponse, error) {
	for {
		l, err := reader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				return nil, io.EOF
			}
			return nil, createError("unknown", "unexpected_error", fmt.Sprintf("Unable to read stream: %v", err))
		}

		l = strings.TrimSpace(l)
		if l == LastStreamMessage {
			return nil, io.EOF
		}
		if isValidLine(l) {
			lWithoutPrefix := l[len(PrefixStreamMessage):]
			var response ChatCompletionsResponse
			if err := json.Unmarshal([]byte(lWithoutPrefix), &response); err != nil {
				return nil, createError("unknown", "unexpected_error", fmt.Sprintf("Unmarshal error occurs: %v", err))
			}
			return &response, nil
		}
	}
}

func isValidLine(l string) bool {
	return len(l) > len(PrefixStreamMessage) && l[:len(PrefixStreamMessage)] == PrefixStreamMessage
}
