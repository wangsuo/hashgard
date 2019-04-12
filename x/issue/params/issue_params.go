package params

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// Param issue for issue
type IssueParams struct {
	Name            string  `json:"name"`
	Symbol          string  `json:"symbol"`
	TotalSupply     sdk.Int `json:"total_supply"`
	Decimals        uint    `json:"decimals"`
	Description     string  `json:"description"`
	BurningFinished bool    `json:"burning_finished"`
	MintingFinished bool    `json:"minting_finished"`
}
