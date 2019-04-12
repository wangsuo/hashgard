package msgs

import (
	"github.com/cosmos/cosmos-sdk/codec"

	"github.com/hashgard/hashgard/x/issue/types"
)

var MsgCdc = codec.New()

// Register concrete types on codec codec
func RegisterCodec(cdc *codec.Codec) {
	cdc.RegisterConcrete(MsgIssue{}, "issue/MsgIssue", nil)
	cdc.RegisterConcrete(MsgIssueDescription{}, "issue/MsgIssueDescription", nil)
	cdc.RegisterConcrete(MsgIssueMint{}, "issue/MsgIssueMint", nil)
	cdc.RegisterConcrete(MsgIssueBurn{}, "issue/MsgIssueBurn", nil)
	cdc.RegisterConcrete(MsgIssueFinishMinting{}, "issue/MsgIssueFinishMinting", nil)

	cdc.RegisterInterface((*types.Issue)(nil), nil)
	cdc.RegisterConcrete(&types.CoinIssueInfo{}, "issue/CoinIssueInfo", nil)
}

//nolint
func init() {
	RegisterCodec(MsgCdc)
}