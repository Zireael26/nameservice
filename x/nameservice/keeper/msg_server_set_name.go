package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/zireael26/nameservice/x/nameservice/types"
)

func (k msgServer) SetName(goCtx context.Context, msg *types.MsgSetName) (*types.MsgSetNameResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Fetch existing whois record from store
	whois, _ := k.GetWhois(ctx, msg.Name)

	// check if message sender is the owner of the name
	if !(msg.Creator == whois.Owner) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "You do not own the name")
	}

	// if owner turns out to be the message creator, let them set the value
	newWhois := types.Whois{
		Index: msg.Name,
		Name:  msg.Name,
		Value: msg.Value,
		Owner: whois.Owner,
		Price: whois.Price,
	}

	k.SetWhois(ctx, newWhois)

	return &types.MsgSetNameResponse{}, nil
}
