package rest

import (
	"net/http"
  "strings"

	"github.com/gorilla/mux"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/ethereum/go-ethereum/common"
)

func RegisterAiaxRESTRoutes(ctx client.Context, r *mux.Router) {
	r.HandleFunc("/aiax/erc20/{address}", HandleERC20Fn(ctx)).Methods("GET")
}

func HandleERC20Fn(ctx client.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		address := common.HexToAddress(strings.ToLower(vars["address"]))
    
    // TODO:

	}
}
