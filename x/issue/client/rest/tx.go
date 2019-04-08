package rest

import (
	"fmt"
	"github.com/cosmos/cosmos-sdk/client/context"
	clientrest "github.com/cosmos/cosmos-sdk/client/rest"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/rest"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"

	"github.com/hashgard/hashgard/x/issue/domain"
	"github.com/hashgard/hashgard/x/issue/msgs"
	"github.com/hashgard/hashgard/x/issue/params"
	issueutils "github.com/hashgard/hashgard/x/issue/utils"
)

type PostIssueReq struct {
	BaseReq            rest.BaseReq `json:"base_req"`
	params.IssueParams `json:"issue"`
}
type PostIssueBaseReq struct {
	BaseReq rest.BaseReq `json:"base_req"`
}

// RegisterRoutes - Central function to define routes that get registered by the main application
func registerTxRoutes(cliCtx context.CLIContext, r *mux.Router, cdc *codec.Codec) {
	r.HandleFunc("/issue/issues", postIssueHandlerFn(cdc, cliCtx)).Methods("POST")
	r.HandleFunc(fmt.Sprintf("/issue/mint/{%s}/{%s}/{%s}", IssueID, Amount, To), postMintHandlerFn(cdc, cliCtx)).Methods("POST")
	r.HandleFunc(fmt.Sprintf("/issue/burn/{%s}/{%s}", IssueID, Amount), postBurnHandlerFn(cdc, cliCtx)).Methods("POST")
	r.HandleFunc(fmt.Sprintf("/issue/finish-minting/{%s}", IssueID), postFinishMintingHandlerFn(cdc, cliCtx)).Methods("POST")
}
func postIssueHandlerFn(cdc *codec.Codec, cliCtx context.CLIContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req PostIssueReq
		if !rest.ReadRESTReq(w, r, cdc, &req) {
			return
		}

		req.BaseReq = req.BaseReq.Sanitize()
		if !req.BaseReq.ValidateBasic(w) {
			return
		}
		fromAddress, _, err := context.GetFromFields(req.BaseReq.From)
		if err != nil {
			return
		}
		coinIssueInfo := domain.CoinIssueInfo{
			Issuer:          fromAddress,
			Name:            req.Name,
			TotalSupply:     req.TotalSupply,
			Decimals:        domain.DefaultDecimals,
			MintingFinished: req.MintingFinished,
		}
		// create the message
		msg := msgs.NewMsgIssue(&coinIssueInfo)
		if err := msg.ValidateBasic(); err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		if req.BaseReq.GenerateOnly {
			clientrest.WriteGenerateStdTxResponse(w, cdc, cliCtx, req.BaseReq, []sdk.Msg{msg})
			return
		}
		clientrest.CompleteAndBroadcastTxREST(w, cliCtx, req.BaseReq, []sdk.Msg{msg}, cdc)
	}
}

func postMintHandlerFn(cdc *codec.Codec, cliCtx context.CLIContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		issueID := vars[IssueID]
		if error := issueutils.CheckIssueId(issueID); error != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, error.Error())
			return
		}
		num, error := strconv.ParseInt(vars[Amount], 10, 64)
		if error != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, error.Error())
			return
		}
		amount := sdk.NewInt(num)
		to, error := sdk.AccAddressFromBech32(vars[To])
		if error != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, error.Error())
			return
		}
		var req PostIssueBaseReq
		if !rest.ReadRESTReq(w, r, cdc, &req) {
			return
		}
		req.BaseReq = req.BaseReq.Sanitize()
		if !req.BaseReq.ValidateBasic(w) {
			return
		}
		fromAddress, _, error := context.GetFromFields(req.BaseReq.From)
		if error != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, error.Error())
			return
		}

		msg := msgs.NewMsgIssueMint(issueID, fromAddress, amount, to)
		if err := msg.ValidateBasic(); err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		if req.BaseReq.GenerateOnly {
			clientrest.WriteGenerateStdTxResponse(w, cdc, cliCtx, req.BaseReq, []sdk.Msg{msg})
			return
		}
		clientrest.CompleteAndBroadcastTxREST(w, cliCtx, req.BaseReq, []sdk.Msg{msg}, cdc)
	}
}
func postBurnHandlerFn(cdc *codec.Codec, cliCtx context.CLIContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		issueID := vars[IssueID]
		if error := issueutils.CheckIssueId(issueID); error != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, error.Error())
			return
		}
		num, error := strconv.ParseInt(vars[Amount], 10, 64)
		if error != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, error.Error())
			return
		}
		amount := sdk.NewInt(num)

		var req PostIssueBaseReq
		if !rest.ReadRESTReq(w, r, cdc, &req) {
			return
		}

		req.BaseReq = req.BaseReq.Sanitize()
		if !req.BaseReq.ValidateBasic(w) {
			return
		}
		fromAddress, _, error := context.GetFromFields(req.BaseReq.From)
		if error != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, error.Error())
			return
		}

		msg := msgs.NewMsgIssueBurn(issueID, fromAddress, amount)
		if err := msg.ValidateBasic(); err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		if req.BaseReq.GenerateOnly {
			clientrest.WriteGenerateStdTxResponse(w, cdc, cliCtx, req.BaseReq, []sdk.Msg{msg})
			return
		}
		clientrest.CompleteAndBroadcastTxREST(w, cliCtx, req.BaseReq, []sdk.Msg{msg}, cdc)
	}
}

func postFinishMintingHandlerFn(cdc *codec.Codec, cliCtx context.CLIContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		issueID := vars[IssueID]
		if error := issueutils.CheckIssueId(issueID); error != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, error.Error())
			return
		}
		var req PostIssueBaseReq
		if !rest.ReadRESTReq(w, r, cdc, &req) {
			return
		}

		req.BaseReq = req.BaseReq.Sanitize()
		if !req.BaseReq.ValidateBasic(w) {
			return
		}
		fromAddress, _, error := context.GetFromFields(req.BaseReq.From)
		if error != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, error.Error())
			return
		}
		msg := msgs.NewMsgIssueFinishMinting(issueID, fromAddress)
		if err := msg.ValidateBasic(); err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		if req.BaseReq.GenerateOnly {
			clientrest.WriteGenerateStdTxResponse(w, cdc, cliCtx, req.BaseReq, []sdk.Msg{msg})
			return
		}
		clientrest.CompleteAndBroadcastTxREST(w, cliCtx, req.BaseReq, []sdk.Msg{msg}, cdc)
	}
}
