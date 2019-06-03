package handlers

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/hashgard/hashgard/x/issue/keeper"
	"github.com/hashgard/hashgard/x/issue/msgs"
	"github.com/hashgard/hashgard/x/issue/utils"
)

//Handle MsgIssue
func HandleMsgIssue(ctx sdk.Context, keeper keeper.Keeper, msg msgs.MsgIssue) sdk.Result {
	fee := keeper.GetParams(ctx).IssueFee
	if err := keeper.Fee(ctx, msg.Owner, fee); err != nil {
		return err.Result()
	}
	coinIssueInfo := msg.CoinIssueInfo
	_, err := keeper.CreateIssue(ctx, coinIssueInfo)
	if err != nil {
		return err.Result()
	}

	return sdk.Result{
		Data: keeper.Getcdc().MustMarshalBinaryLengthPrefixed(coinIssueInfo.IssueId),
		Tags: utils.GetIssueTags(coinIssueInfo.IssueId, coinIssueInfo.Issuer),
	}
}
