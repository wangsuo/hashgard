package rest

import (
	"fmt"
	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/types/rest"
	"github.com/gorilla/mux"
	"net/http"

	"github.com/hashgard/hashgard/x/issue/client/queriers"
	"github.com/hashgard/hashgard/x/issue/domain"
	issueutils "github.com/hashgard/hashgard/x/issue/utils"
)

// RegisterRoutes - Central function to define routes that get registered by the main application
func registerQueryRoutes(cliCtx context.CLIContext, r *mux.Router, cdc *codec.Codec, queryRoute string) {
	r.HandleFunc(fmt.Sprintf("/issue/issues/{%s}", IssueID), queryProposalHandlerFn(cdc, cliCtx)).Methods("GET")
}
func queryProposalHandlerFn(cdc *codec.Codec, cliCtx context.CLIContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		issueID := vars[IssueID]
		if error := issueutils.CheckIssueId(issueID); error != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, error.Error())
			return
		}
		res, err := queriers.QueryIssueByID(issueID, cliCtx, cdc, domain.StoreKey)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusInternalServerError, err.Error())
			return
		}
		rest.PostProcessResponse(w, cdc, res, cliCtx.Indent)
	}
}