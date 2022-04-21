package qytmos_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	keepertest "github.com/tianqiyu119/qytmos/testutil/keeper"
	"github.com/tianqiyu119/qytmos/testutil/nullify"
	"github.com/tianqiyu119/qytmos/x/qytmos"
	"github.com/tianqiyu119/qytmos/x/qytmos/types"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.QytmosKeeper(t)
	qytmos.InitGenesis(ctx, *k, genesisState)
	got := qytmos.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
