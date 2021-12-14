package keeper

import (
	"context"
	"strings"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/aiax-network/aiax-node/x/aiax/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ethereum/go-ethereum/common"
)

var (
  _ types.QueryServer = Keeper{}
)

func (k Keeper) ERC20Address(ctx context.Context, req *types.QueryERC20AddressRequest) (*types.QueryERC20AddressResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "empty request")
	}
	if req.Address == "" {
		return nil, status.Error(codes.InvalidArgument, "address cannot be empty")
	}
	address := common.HexToAddress(strings.ToLower(req.Address))
	exists, localAddress := k.ExternalERC20LocalLookup(sdk.UnwrapSDKContext(ctx), address)

	if !exists {
		return &types.QueryERC20AddressResponse{Address: ""}, nil
	} else {
		return &types.QueryERC20AddressResponse{Address: localAddress.String()}, nil
	}
}
