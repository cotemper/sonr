package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/sonr-io/sonr/x/registry/types"
)

func (k msgServer) CreateThereis(goCtx context.Context, msg *types.MsgCreateThereis) (*types.MsgCreateThereisResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value already exists
	_, isFound := k.GetThereis(
		ctx,
		msg.Index,
	)
	if isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "index already set")
	}

	var thereis = types.Thereis{
		Creator: msg.Creator,
		Index:   msg.Index,
		Did:     msg.Did,
		Value:   msg.Value,
	}

	k.SetThereis(
		ctx,
		thereis,
	)
	return &types.MsgCreateThereisResponse{}, nil
}

func (k msgServer) UpdateThereis(goCtx context.Context, msg *types.MsgUpdateThereis) (*types.MsgUpdateThereisResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value exists
	valFound, isFound := k.GetThereis(
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

	var thereis = types.Thereis{
		Creator: msg.Creator,
		Index:   msg.Index,
		Did:     msg.Did,
		Value:   msg.Value,
	}

	k.SetThereis(ctx, thereis)

	return &types.MsgUpdateThereisResponse{}, nil
}

func (k msgServer) DeleteThereis(goCtx context.Context, msg *types.MsgDeleteThereis) (*types.MsgDeleteThereisResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value exists
	valFound, isFound := k.GetThereis(
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

	k.RemoveThereis(
		ctx,
		msg.Index,
	)

	return &types.MsgDeleteThereisResponse{}, nil
}
