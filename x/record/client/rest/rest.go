package rest

import (
	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/gorilla/mux"
)

const (
	StartRecordID	= "start_record_id"
	Limit			= "limit"
	Hash   			= "hash"
	Sender	    	= "sender"
)

// RegisterRoutes register distribution REST routes.
func RegisterRoutes(cliCtx context.CLIContext, r *mux.Router, cdc *codec.Codec) {
	registerQueryRoutes(cliCtx, r, cdc)
}
