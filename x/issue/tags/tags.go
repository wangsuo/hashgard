package tags

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// Issue tags
var (
	Action          = sdk.TagAction
	Issuer          = "issuer"
	IssueID         = "issue-id"
	Name            = "name"
	Symbol          = "symbol"
	TotalSupply     = "total-supply"
	MintingFinished = "minting-finished"
)