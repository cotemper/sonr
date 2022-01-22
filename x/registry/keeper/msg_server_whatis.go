package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/sonr-io/sonr/x/registry/types"
)

func (k msgServer) CreateWhatis(goCtx context.Context, msg *types.MsgCreateWhatis) (*types.MsgCreateWhatisResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value already exists
	_, isFound := k.GetWhatis(
		ctx,
		msg.Index,
	)
	if isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "index already set")
	}

	var whatis = types.Whatis{
		Creator: msg.Creator,
		Index:   msg.Index,
		Did:     msg.Did,
		Value:   msg.Value,
	}

	k.SetWhatis(
		ctx,
		whatis,
	)
	return &types.MsgCreateWhatisResponse{}, nil
}

func (k msgServer) UpdateWhatis(goCtx context.Context, msg *types.MsgUpdateWhatis) (*types.MsgUpdateWhatisResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value exists
	valFound, isFound := k.GetWhatis(
		ctx,
		msg.Index,
	)
	if !isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, "index not set")
	}

	// Checks if the the msg creator is the same as the current owner
	if msg.Creator != valFound.Creator {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	var whatis = types.Whatis{
		Creator: msg.Creator,
		Index:   msg.Index,
		Did:     msg.Did,
		Value:   msg.Value,
	}

	k.SetWhatis(ctx, whatis)

	return &types.MsgUpdateWhatisResponse{}, nil
}

func (k msgServer) DeleteWhatis(goCtx context.Context, msg *types.MsgDeleteWhatis) (*types.MsgDeleteWhatisResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value exists
	valFound, isFound := k.GetWhatis(
		ctx,
		msg.Index,
	)
	if !isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, "index not set")
	}

	// Checks if the the msg creator is the same as the current owner
	if msg.Creator != valFound.Creator {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	k.RemoveWhatis(
		ctx,
		msg.Index,
	)

	return &types.MsgDeleteWhatisResponse{}, nil
}
