package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	testkeeper "github.com/tianqiyu119/qytmos/testutil/keeper"
	"github.com/tianqiyu119/qytmos/x/qytmos/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.QytmosKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
