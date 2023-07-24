package keeper_test

import (
	"context"
	"testing"

	keepertest "github.com/cooldev900/deal/testutil/keeper"
	"github.com/cooldev900/deal/x/deal/keeper"
	"github.com/cooldev900/deal/x/deal/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.DealKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}

func TestMsgServer(t *testing.T) {
	ms, ctx := setupMsgServer(t)
	require.NotNil(t, ms)
	require.NotNil(t, ctx)
}
