package keeper_test

import (
	"testing"

	testkeeper "github.com/cosmos/ibc-apps/modules/async-icq/v8/interchain-query-demo/testutil/keeper"
	"github.com/cosmos/ibc-apps/modules/async-icq/v8/interchain-query-demo/x/interquery/types"
	"github.com/stretchr/testify/require"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func TestParamsQuery(t *testing.T) {
	keeper, ctx := testkeeper.InterqueryKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	params := types.DefaultParams()
	keeper.SetParams(ctx, params)

	response, err := keeper.Params(wctx, &types.QueryParamsRequest{})
	require.NoError(t, err)
	require.Equal(t, &types.QueryParamsResponse{Params: params}, response)
}
