package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/zireael26/nameservice/x/nameservice/types"
)

func (k msgServer) BuyName(goCtx context.Context, msg *types.MsgBuyName) (*types.MsgBuyNameResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// try to get a name from the store
	whois, isFound := k.GetWhois(ctx, msg.Name)

	// decide min price to set in case the name has no previous owner
	minPrice := sdk.Coins{sdk.NewInt64Coin("token", 10)}

	// convert price and bid string to sdk.Coins
	price, _ := sdk.ParseCoinsNormalized(whois.Price)
	bid, _ := sdk.ParseCoinsNormalized(msg.Bid)

	// Convert whois and buyer addresses to sdk.AccAddress
	owner, _ := sdk.AccAddressFromBech32(whois.Value)
	buyer, _ := sdk.AccAddressFromBech32(msg.Creator)

	// if a name is found in store
	if isFound {
		// if the current price is higher than bid
		if price.IsAllGT(bid) {
			// Throw an error
			return nil, sdkerrors.Wrap(sdkerrors.ErrInsufficientFunds, "Bid is not high enough")
		}

		// Otherwise, if the bid is higher, send tokens from buyer's account to owner's account
		k.bankKeeper.SendCoins(ctx, buyer, owner, bid)
	} else { // If the name is not in the store
		// if the min price is higher than bid
		if minPrice.IsAllGT(bid) {
			// Throw an error
			return nil, sdkerrors.Wrap(sdkerrors.ErrInsufficientFunds, "Bid is less than min amount")
		}
		// Otherwise (when bid is higher), send tokens from buyer's account to the module's account (as a payment for the name)
		k.bankKeeper.SendCoinsFromAccountToModule(ctx, buyer, types.ModuleName, bid)
	}

	// Create an updated whois record
	newWhois := types.Whois{
		Index: msg.Name,
		Name:  msg.Name,
		Value: whois.Value,
		Price: bid.String(),
		Owner: buyer.String(),
	}

	// Write whois information to the store
	k.SetWhois(ctx, newWhois)

	return &types.MsgBuyNameResponse{}, nil
}
