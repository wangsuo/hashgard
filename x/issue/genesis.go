package issue

import (
	"bytes"

	"github.com/hashgard/hashgard/x/issue/types"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/hashgard/hashgard/x/issue/keeper"
)

// GenesisState - all issue state that must be provided at genesis
type GenesisState struct {
	StartingIssueId uint64          `json:"starting_issue_id"`
	Issues          []CoinIssueInfo `json:"issues"`
}

// NewGenesisState creates a new genesis state.
func NewGenesisState(startingIssueId uint64) GenesisState {
	return GenesisState{StartingIssueId: startingIssueId}
}

// DefaultGenesisState returns a default genesis state
func DefaultGenesisState() GenesisState {
	return NewGenesisState(types.CoinIssueMinId)
}

// Returns if a GenesisState is empty or has data in it
func (data GenesisState) IsEmpty() bool {
	emptyGenState := GenesisState{}
	return data.Equal(emptyGenState)
}

// Checks whether 2 GenesisState structs are equivalent.
func (data GenesisState) Equal(data2 GenesisState) bool {
	b1 := MsgCdc.MustMarshalBinaryBare(data)
	b2 := MsgCdc.MustMarshalBinaryBare(data2)
	return bytes.Equal(b1, b2)
}

// InitGenesis sets distribution information for genesis.
func InitGenesis(ctx sdk.Context, keeper keeper.Keeper, data GenesisState) {
	err := keeper.SetInitialIssueStartingIssueId(ctx, data.StartingIssueId)
	if err != nil {
		panic(err)
	}

	for _, issue := range data.Issues {
		keeper.AddIssue(ctx, &issue)
	}
}

// ExportGenesis returns a GenesisState for a given context and keeper.
func ExportGenesis(ctx sdk.Context, keeper keeper.Keeper) GenesisState {
	genesisState := GenesisState{}

	startingIssueId, _ := keeper.PeekCurrentIssueID(ctx)
	genesisState.StartingIssueId = startingIssueId

	genesisState.Issues = keeper.ListAll(ctx)

	return genesisState

}

// ValidateGenesis performs basic validation of bank genesis data returning an
// error for any failed validation criteria.
func ValidateGenesis(data GenesisState) error { return nil }
