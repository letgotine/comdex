package keeper

import (
	"context"
	"fmt"
	"github.com/tendermint/tendermint/libs/log"

	"github.com/comdex-official/comdex/x/auctionsV2/expected"
	"github.com/comdex-official/comdex/x/auctionsV2/types"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
)

type (
	Keeper struct {
		cdc            codec.BinaryCodec
		storeKey       sdk.StoreKey
		memKey         sdk.StoreKey
		paramstore     paramtypes.Subspace
		bankKeeper     types.BankKeeper
		liquidationsV2 expected.LiquidationsV2Keeper
	}
)

func (k Keeper) Params(ctx context.Context, request *types.QueryParamsRequest) (*types.QueryParamsResponse, error) {
	//TODO implement me
	panic("implement me")
}

func NewKeeper(
	cdc codec.BinaryCodec,
	storeKey,
	memKey sdk.StoreKey,
	ps paramtypes.Subspace,
	bankKeeper types.BankKeeper,
	liquidationsV2Keeper expected.LiquidationsV2Keeper,
) Keeper {
	// set KeyTable if it has not already been set
	if !ps.HasKeyTable() {
		ps = ps.WithKeyTable(types.ParamKeyTable())
	}

	return Keeper{
		cdc:            cdc,
		storeKey:       storeKey,
		memKey:         memKey,
		paramstore:     ps,
		bankKeeper:     bankKeeper,
		liquidationsV2: liquidationsV2Keeper,
	}
}

func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}
