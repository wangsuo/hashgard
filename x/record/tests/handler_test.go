package tests

import (
	"testing"

	"github.com/hashgard/hashgard/x/record"
	"github.com/hashgard/hashgard/x/record/msgs"
	"github.com/stretchr/testify/require"
	abci "github.com/tendermint/tendermint/abci/types"
)

func TestHandlerNewMsgRecord(t *testing.T) {
	mapp, keeper, _, _, _ := getMockApp(t, record.GenesisState{}, nil)
	header := abci.Header{Height: mapp.LastBlockHeight() + 1}
	mapp.BeginBlock(abci.RequestBeginBlock{Header: header})
	ctx := mapp.NewContext(false, abci.Header{})
	mapp.InitChainer(ctx, abci.RequestInitChain{})

	handler := record.NewHandler(keeper)

	msg := msgs.NewMsgRecord(SenderAccAddr, &RecordParams)
	err := msg.ValidateBasic()
	require.Nil(t, err)

	res := handler(ctx, msg)
	require.True(t, res.IsOK())

	recordHash := string(res.Tags[3].Value)
	require.NotNil(t, recordHash)
}
