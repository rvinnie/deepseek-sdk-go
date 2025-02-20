package deepseek

import (
	"fmt"
	"strings"
)

// ModelData represents model information.
type ModelData struct {
	// The model identifier, which can be referenced in the API endpoints.
	ID string `json:"id"`

	// The object type, which is always "model".
	Object string `json:"object"`

	// The organization that owns the model.
	OwnedBy string `json:"owned_by"`
}

// ModelsResponse represents the response of the currently available models list.
type ModelsResponse struct {
	// Possible values: [list]
	Object string `json:"object"`

	// List of ModelData
	Data []ModelData `json:"data"`
}

func (r ModelsResponse) String() string {
	var sb strings.Builder

	sb.WriteString("{")
	sb.WriteString(fmt.Sprintf("Object:%q ", r.Object))

	sb.WriteString("Data:[")
	if len(r.Data) > 0 {
		for i, model := range r.Data {
			if i > 0 {
				sb.WriteString(", ")
			}
			sb.WriteString(fmt.Sprintf(
				"{ID:%q, Object:%q, OwnedBy:%q}",
				model.ID,
				model.Object,
				model.OwnedBy,
			))
		}
	}
	sb.WriteString("]}")

	return sb.String()
}
