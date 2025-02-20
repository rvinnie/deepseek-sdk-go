package deepseek

import (
	"fmt"
	"strings"
)

// BalanceInfo represents the balance information.
type BalanceInfo struct {
	// The currency of the balance.
	Currency string `json:"currency"`

	// The total available balance, including the granted balance and the topped-up balance.
	TotalBalance string `json:"total_balance"`

	// The total not expired granted balance.
	GrantedBalance string `json:"granted_balance"`

	// The total topped-up balance.
	ToppedUpBalance string `json:"topped_up_balance"`
}

// BalanceResponse represents the response from the API endpoint.
type BalanceResponse struct {
	// Whether the user's balance is sufficient for API calls.
	IsAvailable bool `json:"is_available"`

	// List of Balance information
	BalanceInfos []BalanceInfo `json:"balance_infos"`
}

func (r BalanceResponse) String() string {
	var sb strings.Builder

	sb.WriteString("{")
	sb.WriteString(fmt.Sprintf("IsAvailable:%v ", r.IsAvailable))

	sb.WriteString("BalanceInfos:[")
	if len(r.BalanceInfos) > 0 {
		for i, info := range r.BalanceInfos {
			if i > 0 {
				sb.WriteString(", ")
			}
			sb.WriteString(fmt.Sprintf(
				"{Currency:%q, TotalBalance:%q, GrantedBalance:%q, ToppedUpBalance:%q}",
				info.Currency,
				info.TotalBalance,
				info.GrantedBalance,
				info.ToppedUpBalance,
			))
		}
	}
	sb.WriteString("]}")

	return sb.String()
}
