package tests

import (
	"testing"

	"github.com/hashgard/hashgard/x/issue"
	"github.com/hashgard/hashgard/x/issue/msgs"
	"github.com/stretchr/testify/require"
	abci "github.com/tendermint/tendermint/abci/types"
)

func TestHandlerNewMsgIssue(t *testing.T) {
	mapp, keeper, _, _, _, _ := getMockApp(t, 0, issue.GenesisState{}, nil)
	mapp.BeginBlock(abci.RequestBeginBlock{})
	ctx := mapp.NewContext(false, abci.Header{})
	mapp.InitChainer(ctx, abci.RequestInitChain{})

	//querier := issue.NewQuerier(keeper)
	handler := issue.NewHandler(keeper)

	res := handler(ctx, msgs.NewMsgIssue(&CoinIssueInfo))
	require.True(t, res.IsOK())

	var issueID string
	keeper.Getcdc().MustUnmarshalBinaryLengthPrefixed(res.Data, &issueID)
	require.NotNil(t, issueID)
}
