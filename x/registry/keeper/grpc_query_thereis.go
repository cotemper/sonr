package keeper

import (
	"context"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/sonr-io/sonr/x/registry/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) ThereisAll(c context.Context, req *types.QueryAllThereisRequest) (*types.QueryAllThereisResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var thereiss []types.Thereis
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	thereisStore := prefix.NewStore(store, types.KeyPrefix(types.ThereisKeyPrefix))

	pageRes, err := query.Paginate(thereisStore, req.Pagination, func(key []byte, value []byte) error {
		var thereis types.Thereis
		if err := k.cdc.Unmarshal(value, &thereis); err != nil {
			return err
		}

		thereiss = append(thereiss, thereis)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllThereisResponse{Thereis: thereiss, Pagination: pageRes}, nil
}

func (k Keeper) Thereis(c context.Context, req *types.QueryGetThereisRequest) (*types.QueryGetThereisResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	val, found := k.GetThereis(
		ctx,
		req.Index,
	)
	if !found {
		return nil, status.Error(codes.InvalidArgument, "not found")
	}

	return &types.QueryGetThereisResponse{Thereis: val}, nil
}
